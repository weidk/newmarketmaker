package route

import (
	"NewMarkerMaker/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
)

var ROUTE = g.Map{
	"/onemmquote":            controller.SendTestMMQuote,       // 发送做市报价
	"/cancelmmquote":         controller.CancelMMQuote,         // 撤销报价
	"/BatchCancelMMQuoteApi": controller.BatchCancelMMQuoteApi, // 批量撤销报价
	"/InitMMQuote":           controller.InitMMQuote,           // 初始化做市报价
	"/BatchSendMMQuote":      controller.BatchSendMMQuote,      // 批量发送做市报价
	"/SetAutoMM":             controller.SetAutoMM,             // 设置是否自动启动自动做市
	"/SetMMCanSendApi":       controller.SetMMCanSendApi,       // 设置报价是否能发送
}
