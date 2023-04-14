package logic

import (
	"NewMarkerMaker/internal/consts"
	"NewMarkerMaker/internal/model"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"net"
	"time"
)
import amqp "github.com/rabbitmq/amqp091-go"

type mq struct {
	conn       *amqp.Connection
	mmhqCh     *amqp.Channel
	mmhqQ      amqp.Queue
	mmhqMsg    <-chan amqp.Delivery
	xbondhqCh  *amqp.Channel
	xbondhqQ   amqp.Queue
	xbondhqMsg <-chan amqp.Delivery
	sendCh     *amqp.Channel
	sendQ      amqp.Queue
	receiveCh  *amqp.Channel
	receiveQ   amqp.Queue
	receiveMsg <-chan amqp.Delivery

	notifyConnClose chan *amqp.Error
	notifyMMhqClose chan *amqp.Error
}

var ctx = context.Background()
var Mq = mq{}

func init() {
	Setup()
	go ReConn()
}

func Setup() {
	Mq.getConn()
	Mq.getMMhqObj()
	Mq.getXBondhqObj()
	Mq.client2ServerChannel()
	Mq.server2ClientQueueName()

	Mq.notifyConnClose = Mq.conn.NotifyClose(make(chan *amqp.Error))
	Mq.notifyMMhqClose = Mq.mmhqCh.NotifyClose(make(chan *amqp.Error))
	go ReceiveMMHQ()
	go ReceiveMsg()
	//ReceiveXBondHQ()
}

func ReConn() {
	for Mq.notifyConnClose != nil {
		select {
		case err := <-Mq.notifyConnClose:
			g.Log().Warning(ctx, "connection closed, error %s", err)
			Setup()
			time.Sleep(3 * time.Second)
		}
	}
}

func (m *mq) getConn() {
	conn, err := amqp.DialConfig(consts.Rabbitmqaddr, amqp.Config{
		//Heartbeat: 30 * time.Second,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 60*time.Second)
		},
	})
	if err != nil {
		//g.Log().Error(ctx, err)
		time.Sleep(3 * time.Second)
		m.getConn()
	} else {
		m.conn = conn

		g.Log().Info(ctx, "rabbitmq connection is established")
	}

}

func (m *mq) getMMhqObj() {
	m.mmhqCh, _ = m.conn.Channel()
	m.mmhqQ, _ = m.mmhqCh.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	m.mmhqCh.QueueBind(
		m.mmhqQ.Name,         // queue name
		"",                   // routing key
		consts.MarketmakerHq, // exchange
		false,
		nil,
	)
	m.mmhqMsg, _ = m.mmhqCh.Consume(
		m.mmhqQ.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
}

func (m *mq) getXBondhqObj() {
	m.xbondhqCh, _ = m.conn.Channel()
	m.xbondhqQ, _ = m.xbondhqCh.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	m.xbondhqCh.QueueBind(
		m.xbondhqQ.Name, // queue name
		"",              // routing key
		consts.XbondHq,  // exchange
		false,
		nil,
	)
	m.xbondhqMsg, _ = m.xbondhqCh.Consume(
		Mq.xbondhqQ.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
}

func ReceiveMMHQ() {
	for {
		for d := range Mq.mmhqMsg {
			//g.Dump(d.Body)
			quote := model.MarketDataSnapshotFullRefreshVo{}
			if err := json.Unmarshal(d.Body, &quote); err != nil {
				g.Log().Error(ctx, err)
			} else {
				UpdateBestQuote(quote)
			}
		}
		g.Log().Error(ctx, "mmhqmsg is closed . waiting to reconnect")
		time.Sleep(time.Second)
	}
}

func ReceiveXBondHQ() {
	for {
		for d := range Mq.xbondhqMsg {
			g.Dump(d.Body)
			quote := model.CBXbondQuoteMarketData{}
			if err := json.Unmarshal(d.Body, &quote); err != nil {
				g.Log().Error(ctx, err)

			}
		}
		g.Log().Error(ctx, "mmhqmsg is closed . waiting to reconnect")
		time.Sleep(time.Second)
	}
}

func (m *mq) client2ServerChannel() {
	m.sendCh, _ = m.conn.Channel()
	m.sendQ, _ = m.sendCh.QueueDeclare(
		consts.Client2ServerQueueName, // name
		true,                          // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
}

func SendMsg(body []byte) {
	err := Mq.sendCh.PublishWithContext(ctx,
		"",            // exchange
		Mq.sendQ.Name, // routing key
		false,         // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			//Body:         []byte(body),
			Body: body,
		})
	if err != nil {
		g.Log().Error(ctx, err)
	}
	//else {
	//	g.Log().Info(ctx, "send msg: "+body)
	//}
}

func (m *mq) server2ClientQueueName() {
	m.receiveCh, _ = m.conn.Channel()
	m.receiveQ, _ = m.receiveCh.QueueDeclare(
		consts.Server2ClientQueueName, // name
		true,                          // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	m.receiveCh.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	m.receiveMsg, _ = m.receiveCh.Consume(
		Mq.receiveQ.Name, // queue
		"",               // consumer
		false,            // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
}

func ReceiveMsg() {
	for {
		for d := range Mq.receiveMsg {
			var callback model.QuoteCallBack
			//gconv.Struct(d.Body, &callback)
			json.Unmarshal(d.Body, &callback)
			consts.CallBackCh <- callback
			if e := d.Ack(false); e != nil {
				g.Log().Error(ctx, "ack error: %+v", e)
			}
		}
		g.Log().Error(ctx, "receive queue (server2ClientQueue) is closed . waiting to reconnect")
		time.Sleep(time.Second)
	}
}
