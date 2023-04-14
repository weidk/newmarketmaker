package logic

import (
	"NewMarkerMaker/internal/consts"
	"NewMarkerMaker/internal/model"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func GenerateClordID() string {
	t := gtime.Now().Format("Hisu")
	clordid := consts.ClordIDHeader + gstr.SubStr(t, 1)
	return clordid
	//consts.ClordIDHeader
}

func CalBondPrice(code, tradeDate, speed string, yield float64) float64 {
	bondInfo := model.CalReq{
		ClearSpeed:     speed,
		SecurityId:     code,
		TradeDate:      tradeDate,
		Yield:          yield,
		SettleCurrency: "CNY",
	}

	//if err := g.Validator().Data(bondInfo).Run(consts.Ctx); err != nil {
	//	fmt.Println(err)
	//}

	data, _ := json.Marshal(bondInfo)
	content := g.Client().PostBytes(consts.Ctx,
		consts.CalculateAddress,
		data,
	)
	var rsp model.CalRsp
	json.Unmarshal(content, &rsp)
	var price float64 = 0
	if rsp.Success == true {
		price = gconv.Float64(rsp.Price)
	} else {
		g.Log().Error(consts.Ctx, rsp.Message)
	}
	return price
}
