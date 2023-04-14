package consts

import (
	"NewMarkerMaker/internal/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Ctx                    = context.Background()
	Cfg, _                 = gcfg.New()
	CallBackCh             = make(chan model.QuoteCallBack, 10) // 回调channel
	MMAddCh                = make(chan model.MMQuoteAdd, 10)    // 新增做市报价channel
	Rabbitmqaddr           string
	MarketmakerHq          string
	XbondHq                string
	Client2ServerQueueName string
	Server2ClientQueueName string
	ClordIDHeader          string
	CalculateAddress       string
	PartyID101_803_266     string
)

func init() {
	rabbitmqaddr, err := Cfg.Get(Ctx, "rabbitmq.address")
	Rabbitmqaddr = gconv.String(rabbitmqaddr)
	LogOnError(err)

	marketmakerHq, err := Cfg.Get(Ctx, "rabbitmq.marketmakerHq")
	MarketmakerHq = gconv.String(marketmakerHq)
	LogOnError(err)

	xbondHq, err := Cfg.Get(Ctx, "rabbitmq.xbondHq")
	XbondHq = gconv.String(xbondHq)
	LogOnError(err)

	client2Server, err := Cfg.Get(Ctx, "rabbitmq.Client2Server")
	Client2ServerQueueName = gconv.String(client2Server)
	LogOnError(err)

	server2Client, err := Cfg.Get(Ctx, "rabbitmq.Server2Client")
	Server2ClientQueueName = gconv.String(server2Client)
	LogOnError(err)

	clordIDHeader, _ := Cfg.Get(Ctx, "clordid.test")
	ClordIDHeader = gconv.String(clordIDHeader)

	calculateAddress, _ := Cfg.Get(Ctx, "caculate.address")
	CalculateAddress = gconv.String(calculateAddress)

	partyID101_803_266, _ := Cfg.Get(Ctx, "quoteparams.partyID101_803_266")
	PartyID101_803_266 = gconv.String(partyID101_803_266)
}

func LogOnError(err error) {
	if err != nil {
		g.Log().Error(Ctx, err)
	}
}
