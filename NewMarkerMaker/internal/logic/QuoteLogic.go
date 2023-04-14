package logic

import (
	"NewMarkerMaker/internal/consts"
	"NewMarkerMaker/internal/model"
	"encoding/json"
	"github.com/duke-git/lancet/v2/mathutil"
	_ "github.com/gogf/gf/contrib/drivers/oracle/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
	"math"
)

var MMQuoteSetDict = map[string]model.MMQUOTESET{}
var MMQuotingDict = map[string]model.MMQuoteAdd{}
var IsAutoSendMMQuote = false

func init() {
	//InitMMQuoteSet()
	go UpdateQuoteStatus()

}

func SendMMQuote(quote *model.MMQuoteAdd) error {
	if quote.BUYPRICE != 0. && quote.SELLPRICE != 0. {
		quoteBody, err := json.Marshal(quote)
		if err != nil {
			g.Log().Error(ctx, err)
		} else {
			SendMsg(quoteBody)
			consts.MMAddCh <- *quote
		}
		return nil
	} else {
		g.Log().Error(ctx, "净价计算失败，取消发送该笔报价"+gconv.String(quote))
		return gerror.New("净价计算失败，取消发送该笔报价")
	}
}

func SendXBondQuote(quote model.XBondQuoteAdd) {
	quoteBody, err := json.Marshal(quote)
	if err != nil {
		g.Log().Error(ctx, err)
	} else {
		SendMsg(quoteBody)
	}
}

func CancelQuote(quoteid string) error {
	cancel := model.QuoteCancel{
		TYPE:      "C",
		CLORDID:   GenerateClordID(),
		ORDERTIME: gtime.Datetime(),
		QUOTEID:   quoteid,
	}

	cancelBody, err := json.Marshal(cancel)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	} else {
		SendMsg(cancelBody)
		return nil
	}
}

// 回调函数，接收下行反馈，并更新报价状态
func UpdateQuoteStatus() {
	for {
		select {
		case callback := <-consts.CallBackCh:
			if m, ok := MMQuotingDict[callback.BONDCODE]; ok {
				if m.CLORDID == callback.CLORDID || m.QUOTEID == callback.QUOTEID {
					if callback.TABLE != "DELEGATIONFEADBACK" {
						m.QUOTESTATUS = callback.QUOTESTATUS
						m.QUOTEID = callback.QUOTEID
						if callback.EXECID != "" {
							if callback.SIDE == "1" {
								m.BUYDEALQTY += gconv.Float64(callback.DEALQTY)
							} else {
								m.SELLDEALQTY += gconv.Float64(callback.DEALQTY)
							}
							//如果成交，则暂停该债券报价继续发送
							if v, ok := MMQuoteSetDict[callback.BONDCODE]; ok {
								v.CANSEND = false
								MMQuoteSetDict[callback.BONDCODE] = v
							}
							// 撤销报价
							CancelQuote(callback.QUOTEID)
							// 更新数据库里参数设置为不可发送报价
							g.Model("MM_QUOTE_SET").Data("CANSEND", 0).Where("BONDCODE", callback.BONDCODE).Update()
						}
					}
					m.TRANSACTTIME = callback.TRANSACTTIME
					m.INFO = callback.INFO
					g.Dump(m)
					MMQuotingDict[callback.BONDCODE] = m
				}
			}
		case mm := <-consts.MMAddCh:
			MMQuotingDict[mm.BONDCODE] = mm
			//g.Dump(mm)
		}
	}
}

// 从数据库中读取做市报价设置及初始化最优次优报价
func InitMMQuoteSet() []model.MMQUOTESET {
	lo.Synchronize().Do(func() { IsAutoSendMMQuote = false })

	mmQuoteSets := []model.MMQUOTESET{}
	g.Model("MM_QUOTE_SET").Scan(&mmQuoteSets)
	hqdb := g.DB("newrmb")

	mmQuotings := []model.CBMARKETQUOTE{}
	g.DB("FITRADE").Model("CBMARKETQUOTE").Where("QUOTESTATUS", 16).Scan(&mmQuotings)

	for _, quote := range mmQuotings {
		q := model.MMQuoteAdd{
			TYPE:         "M",
			CLORDID:      quote.CLORDID,
			QUOTEID:      quote.QUOTEID,
			QUOTESTATUS:  quote.QUOTESTATUS,
			BONDCODE:     quote.SECURITYID,
			MAXFLOOR:     quote.MAXFLOOR,
			BUYORDERQTY:  quote.BUYORDERQTY,
			BUYPRICE:     quote.BUYPRICE,
			BUYYTMYILED:  quote.BUYYTMYILED,
			BUYDEALQTY:   quote.BUYLASTQTY,
			SELLORDERQTY: quote.SELLORDERQTY,
			SELLPRICE:    quote.SELLPRICE,
			SELLYTMYILED: quote.SELLYTMYILED,
			SELLDEALQTY:  quote.SELLLASTQTY,
			SETTLTYPE:    quote.BUYSETTLTYPE,
			TRANSACTTIME: quote.TRANSACTTIME,
		}
		MMQuotingDict[quote.SECURITYID] = q
	}

	for _, set := range mmQuoteSets {
		//cbYield, _ := g.DB().GetOne(ctx, "select valu_yield_cnbd from dm.inf_bondvalue  where substr(bond_code,1,length(bond_code) - 3) = ?", set.BONDCODE)
		//set.BASEYIELD = gconv.Float32(cbYield["VALU_YIELD_CNBD"])
		mmQuote := []model.CBMARKETQUOTESW{}
		hqdb.Model("CBMARKETQUOTESW").
			Where("MDPRICELEVEL<=", 2).
			Where("PARTYID101_803_266!=", consts.PartyID101_803_266).
			Where("SETTLTYPE", set.SETTLETYPE).
			Where("securityid", set.BONDCODE).
			Where(" inserttime>trunc(sysdate)").
			OrderDesc("ID").
			Where("rownum<", 5).
			Scan(&mmQuote)
		BESTBID, ok := lo.Find(mmQuote, func(x model.CBMARKETQUOTESW) bool { return x.MDPRICELEVEL == "1" && x.MDENTRYTYPE == "0" })
		if ok {
			set.BESTBID = BESTBID.MATURITYYIELD
		}
		SECONDBID, ok := lo.Find(mmQuote, func(x model.CBMARKETQUOTESW) bool { return x.MDPRICELEVEL == "2" && x.MDENTRYTYPE == "0" })
		if ok {
			set.SECONDBID = SECONDBID.MATURITYYIELD
		}
		BESTOFR, ok := lo.Find(mmQuote, func(x model.CBMARKETQUOTESW) bool { return x.MDPRICELEVEL == "1" && x.MDENTRYTYPE == "1" })
		if ok {
			set.BESTOFR = BESTOFR.MATURITYYIELD
		}
		SECONDOFR, ok := lo.Find(mmQuote, func(x model.CBMARKETQUOTESW) bool { return x.MDPRICELEVEL == "2" && x.MDENTRYTYPE == "1" })
		if ok {
			set.SECONDOFR = SECONDOFR.MATURITYYIELD
		}
		if set.BESTBID > 0 && set.BESTOFR > 0 {
			set.BESTMIDYIELD = math.Round(10000*(set.BESTBID+set.BESTOFR)/2) / 10000
		}
		if set.SECONDBID > 0 && set.SECONDOFR > 0 {
			set.SECONDMIDYIELD = math.Round(10000*(set.SECONDBID+set.SECONDOFR)/2) / 10000
		}

		if v, ok := MMQuotingDict[set.BONDCODE]; ok {
			set.LASTBID = v.BUYYTMYILED
			set.LASTOFR = v.SELLYTMYILED
		}
		UpdateSelfQuote(&set)

		MMQuoteSetDict[set.BONDCODE] = set
	}
	return lo.Values[string, model.MMQUOTESET](MMQuoteSetDict)
}

// 更新除本方外的最优报价，并发送
func UpdateBestQuote(mmhq model.MarketDataSnapshotFullRefreshVo) {
	if mmhq.PartyID101_803_266 != consts.PartyID101_803_266 {
		if val, ok := MMQuoteSetDict[mmhq.SecurityID]; ok {
			if mmhq.SettlType == val.SETTLETYPE {
				if mmhq.MDPriceLevel == "1" || mmhq.MDPriceLevel == "2" {
					if mmhq.MDPriceLevel == "1" {
						if mmhq.MDEntryType == "0" {
							val.BESTBID = gconv.Float64(mmhq.MaturityYield)
						} else {
							val.BESTOFR = gconv.Float64(mmhq.MaturityYield)
						}
						val.BESTMIDYIELD = math.Round(10000*(val.BESTBID+val.BESTOFR)/2) / 10000

					} else if mmhq.MDPriceLevel == "2" {
						if mmhq.MDEntryType == "0" {
							val.SECONDBID = gconv.Float64(mmhq.MaturityYield)
						} else {
							val.SECONDOFR = gconv.Float64(mmhq.MaturityYield)
						}

						val.SECONDMIDYIELD = math.Round(10000*(val.SECONDBID+val.SECONDOFR)/2) / 10000
					}
					UpdateSelfQuote(&val)
					if IsAutoSendMMQuote {
						err := SendOneMMQuoteBag(&val)
						if err == nil {
							//g.Log().Error(ctx, "自动发送报价错误："+gconv.String(err))
							g.Dump("发送做市报价：" + val.BONDCODE)
						}
					}
					MMQuoteSetDict[mmhq.SecurityID] = val
					//g.Dump(val)
				}
			}

		}
	}
}

// 将设置转换为报价，并发送
func SendOneMMQuoteBag(v *model.MMQUOTESET) error {
	if !v.CANSEND {
		return gerror.New("该报价目前处于不可发送状态")
	}
	if v.IsNew {
		quote := model.MMQuoteAdd{
			TYPE:         "M",
			CLORDID:      GenerateClordID(),
			BONDCODE:     v.BONDCODE,
			MAXFLOOR:     v.MAXFLOOR,
			BUYORDERQTY:  v.BUYQTY,
			BUYYTMYILED:  v.SELFBID,
			SELLORDERQTY: v.SELLQTY,
			SELLYTMYILED: v.SELFOFR,
			SETTLTYPE:    v.SETTLETYPE,
			BUYPRICE:     CalBondPrice(v.BONDCODE, gtime.Now().Format("Ymd"), v.SETTLETYPE, v.SELFBID),
			SELLPRICE:    CalBondPrice(v.BONDCODE, gtime.Now().Format("Ymd"), v.SETTLETYPE, v.SELFOFR),
			ORDERTIME:    gtime.Datetime(),
		}
		err := SendMMQuote(&quote)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		} else {
			v.LASTBID = quote.BUYYTMYILED
			v.LASTOFR = quote.SELLYTMYILED
			v.IsNew = false
		}
	} else {
		return gerror.New("该价格未变动，无序再次报价")
	}
	return nil
}

// 更新本方报价
func UpdateSelfQuote(val *model.MMQUOTESET) {
	val.IsNew = false
	// 次优价加减点
	if val.BESTYIELDTYPE == 2 {
		if val.SETYIELDBUY > 0 {
			val.SELFBID = val.SETYIELDBUY
		} else if val.SECONDBID < 0 || val.SECONDOFR < 0 || val.SECONDMIDYIELD < 0 {
			//最优/次优报价不全，则用估值+10bp
			val.SELFBID = val.BASEYIELD + 0.1
		} else {
			val.SELFBID = lo.Max([]float64{val.BASEYIELD, val.SECONDMIDYIELD + val.MIDYIELDADD, val.SECONDBID + val.BESTYIELDADD})
		}
		if val.SETYIELDSELL > 0 {
			val.SELFOFR = val.SETYIELDSELL
		} else if val.SECONDBID < 0 || val.SECONDOFR < 0 || val.SECONDMIDYIELD < 0 {
			val.SELFOFR = val.BASEYIELD - 0.1
		} else {
			val.SELFOFR = lo.Min([]float64{val.BASEYIELD, val.SECONDMIDYIELD - val.MIDYIELDSUB, val.SECONDOFR - val.BESTYIELDSUB})
		}
	} else { // 最优价加减点
		if val.SETYIELDBUY > 0 {
			val.SELFBID = val.SETYIELDBUY
		} else if val.BESTBID < 0 || val.BESTOFR < 0 || val.BESTMIDYIELD < 0 {
			val.SELFBID = val.BASEYIELD + 0.1
		} else {
			val.SELFBID = lo.Max([]float64{val.BASEYIELD, val.BESTMIDYIELD + val.MIDYIELDADD, val.BESTBID + val.BESTYIELDADD})
		}

		if val.SETYIELDSELL > 0 {
			val.SELFOFR = val.SETYIELDSELL
		} else if val.BESTBID < 0 || val.BESTOFR < 0 || val.BESTMIDYIELD < 0 {
			val.SELFOFR = val.BASEYIELD - 0.1
		} else {
			val.SELFOFR = lo.Min([]float64{val.BASEYIELD, val.BESTMIDYIELD - val.MIDYIELDSUB, val.BESTOFR - val.BESTYIELDSUB})
		}
	}
	val.SELFBID = mathutil.RoundToFloat(val.SELFBID, 4)
	val.SELFOFR = mathutil.RoundToFloat(val.SELFOFR, 4)
	if math.Abs(val.SELFBID-val.LASTBID) > 0.0005 || (math.Abs(val.SELFOFR-val.LASTOFR) > 0.0005) {
		val.IsNew = true
	} else {
		val.IsNew = false
	}
}

// 批量发送做市报价
func BatchSendMMQuote() g.Map {
	errMap := g.Map{}
	for k, v := range MMQuoteSetDict {
		err := SendOneMMQuoteBag(&v)
		if err != nil {
			errMap[k] = err
		} else {
			errMap[k] = "发送报价成功"
			MMQuoteSetDict[k] = v
		}
	}
	return errMap
}

// 批量撤销报价
func BatchCancelMMQuote() g.Map {
	errMap := g.Map{}
	for k, v := range MMQuotingDict {
		if v.QUOTESTATUS == "16" {
			err := CancelQuote(v.QUOTEID)
			if err != nil {
				errMap[k] = err
			} else {
				errMap[k] = "撤销报价成功"
			}
		}
	}
	return errMap
}
