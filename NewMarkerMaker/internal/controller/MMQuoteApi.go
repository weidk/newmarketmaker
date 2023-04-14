package controller

import (
	"NewMarkerMaker/internal/logic"
	"NewMarkerMaker/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
)

// 发送做市报价（测试）
func SendTestMMQuote(r *ghttp.Request) {
	var quote model.MMQuoteAdd
	if err := r.Parse(&quote); err != nil {
		r.Response.WriteJson(g.Map{
			"status": 0,
			"msg":    err.Error(),
			"data":   "",
		})
	} else {
		quote.CLORDID = logic.GenerateClordID()
		quote.BUYPRICE = logic.CalBondPrice(quote.BONDCODE, gtime.Now().Format("Ymd"), "2", quote.BUYYTMYILED)
		quote.SELLPRICE = logic.CalBondPrice(quote.BONDCODE, gtime.Now().Format("Ymd"), "2", quote.SELLYTMYILED)
		quote.ORDERTIME = gtime.Datetime()
		//g.Dump(gconv.String(&quote))
		err := logic.SendMMQuote(&quote)
		if err != nil {
			r.Response.WriteJson(g.Map{
				"status": 0,
				"msg":    err.Error(),
				"data":   &quote,
			})
		} else {
			r.Response.WriteJson(g.Map{
				"status": 1,
				"msg":    "success",
				"data":   &quote,
			})
		}

	}
}

// 撤销做市报价
func CancelMMQuote(r *ghttp.Request) {
	QUOTEID := gconv.String(r.Get("data"))
	if QUOTEID == "" {
		r.Response.WriteJson(g.Map{
			"status": 0,
			"msg":    "撤销报价 quoteid为空",
			"data":   "",
		})
	} else {
		err := logic.CancelQuote(QUOTEID)
		if err != nil {
			r.Response.WriteJson(g.Map{
				"status": 0,
				"msg":    err.Error(),
				"data":   "",
			})
		} else {
			r.Response.WriteJson(g.Map{
				"status": 1,
				"msg":    "success",
				"data":   "",
			})
		}
	}
}

// 初始化报价
func InitMMQuote(r *ghttp.Request) {
	mmQuoteSets := logic.InitMMQuoteSet()
	r.Response.WriteJson(g.Map{
		"status": 0,
		"msg":    "success",
		"data":   &mmQuoteSets,
	})
}

// 批量发送报价
func BatchSendMMQuote(r *ghttp.Request) {
	rst := logic.BatchSendMMQuote()
	r.Response.WriteJson(&rst)
}

// 设置做市报价启动或停止
func SetAutoMM(r *ghttp.Request) {
	isAuto := r.Get("data")
	lo.Synchronize().Do(func() {
		logic.IsAutoSendMMQuote = gconv.Bool(isAuto)
		r.Response.WriteJson(g.Map{
			"msg":  "success",
			"data": &logic.IsAutoSendMMQuote,
		})
	})
}

// 批量撤销报价
func BatchCancelMMQuoteApi(r *ghttp.Request) {
	rst := logic.BatchCancelMMQuote()
	r.Response.WriteJson(&rst)
}

// 设置报价是否能发送
func SetMMCanSendApi(r *ghttp.Request) {
	canSend := gconv.Bool(r.Get("cansend"))
	bondcode := gconv.String(r.Get("bondcode"))
	g.Model("MM_QUOTE_SET").Data("CANSEND", gconv.Int8(canSend)).Where("BONDCODE", bondcode).Update()
	lo.Synchronize().Do(func() {
		if v, ok := logic.MMQuoteSetDict[bondcode]; ok {
			v.CANSEND = canSend
			logic.MMQuoteSetDict[bondcode] = v
			r.Response.WriteJson(g.Map{
				"msg":  "success",
				"data": &v,
			})
		}
	})
}
