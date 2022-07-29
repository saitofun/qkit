package mqtt

import (
	"crypto/tls"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/x/misc/retry"
)

type Broker struct {
	Server  types.Endpoint
	Retry   retry.Retry
	Timeout types.Duration

	cid   string // cid client id
	topic string // topic registered topic
	qos   QOS    // qos should be 0, 1 or 2

	client  mqtt.Client         `env:"-"`
	handler mqtt.MessageHandler `env:"-"`
}

func (b *Broker) SetDefault() {
	b.Retry.SetDefault()
	if b.Timeout == 0 {
		b.Timeout = types.Duration(3 * time.Second)
	}
}

func (b *Broker) Init() {
	opt := mqtt.NewClientOptions()
	if !b.Server.IsZero() {
		opt = opt.AddBroker(b.Server.SchemeHost())
	}
	if b.cid != "" {
		opt.SetClientID(b.cid)
	}
	if b.Server.Username != "" {
		opt.SetUsername(b.Server.Username)
		if b.Server.Password != "" {
			opt.SetPassword(b.Server.Password.String())
		}
	}
	opt.SetKeepAlive(b.Timeout.Duration())

	client := mqtt.NewClient(opt)
	if err := b.Retry.Do(b.conn(client)); err != nil {
		panic(err)
	}
	b.client = client
}
func (b Broker) conn(client mqtt.Client) func() error {
	return func() error {
		tok := client.Connect()
		waited := false
		if b.Timeout == 0 {
			waited = tok.Wait()
		} else {
			waited = tok.WaitTimeout(b.Timeout.Duration())
		}
		if !waited {
			return errors.New("connect timeout")
		}
		if err := tok.Error(); err != nil {
			return errors.Wrap(err, "connect failed")
		}
		return nil
	}
}

// TODO mqtt tls
func (b *Broker) tls() *tls.Config { return nil }

func (b *Broker) RegisterMessageHandler(hdl mqtt.MessageHandler) {
	b.handler = hdl
}

func (b *Broker) Publish(topic string, pld []byte) mqtt.Token {
	return b.client.Publish(topic, byte(b.qos), false, pld)
}

func (b *Broker) Subscribe(topic string) mqtt.Token {
	return b.client.Subscribe(topic, byte(b.qos), b.handler)
}
