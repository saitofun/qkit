package mqtt

import (
	"crypto/tls"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/x/mapx"
	"github.com/saitofun/qkit/x/misc/retry"
)

type Broker struct {
	Server        types.Endpoint
	Retry         retry.Retry
	Timeout       types.Duration
	Keepalive     types.Duration
	RetainPublish bool
	QoS           QOS

	agents *mapx.Map[string, *Client]
}

func (b *Broker) SetDefault() {
	b.Retry.SetDefault()
	if b.Timeout == 0 {
		b.Timeout = types.Duration(3 * time.Second)
	}
	if b.Keepalive == 0 {
		b.Keepalive = types.Duration(3 * time.Hour)
	}
	if b.Server.IsZero() {
		b.Server.Hostname, b.Server.Port = "127.0.0.1", 1883
	}
	b.Server.Scheme = "mqtt"
	if b.agents == nil {
		b.agents = mapx.New[string, *Client]()
	}
}

func (b *Broker) Init() {
	err := b.Retry.Do(func() error {
		_, err := b.Client("")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (b *Broker) options() *mqtt.ClientOptions {
	opt := mqtt.NewClientOptions()
	if !b.Server.IsZero() {
		opt = opt.AddBroker(b.Server.SchemeHost())
	}
	if b.Server.Username != "" {
		opt.SetUsername(b.Server.Username)
		if b.Server.Password != "" {
			opt.SetPassword(b.Server.Password.String())
		}
	}

	opt.SetKeepAlive(b.Keepalive.Duration())
	opt.SetWriteTimeout(b.Timeout.Duration())
	opt.SetConnectTimeout(b.Timeout.Duration())
	return opt
}

func (b *Broker) Client(cid string) (*Client, error) {
	opt := b.options()
	if cid != "" {
		opt.SetClientID(cid)
	}
	if b.Server.IsTLS() {
		opt.SetTLSConfig(&tls.Config{
			ClientAuth:         tls.NoClientCert,
			ClientCAs:          nil,
			InsecureSkipVerify: true,
		})
	}
	return b.ClientWithOptions(cid, opt)
}

func (b *Broker) ClientWithOptions(cid string, opt *mqtt.ClientOptions) (*Client, error) {
	client, err := b.agents.LoadOrStore(
		cid,
		func() (*Client, error) {
			c := &Client{
				cid:     cid,
				qos:     b.QoS,
				timeout: b.Timeout.Duration(),
				retain:  b.RetainPublish,
				cli:     mqtt.NewClient(opt),
			}
			if err := c.connect(); err != nil {
				return nil, err
			}
			return c, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !client.cli.IsConnectionOpen() && !client.cli.IsConnected() {
		b.agents.Remove(cid)
		return b.Client(cid)
	}
	return client, nil
}

func (b *Broker) Close(cid string) {
	if c, ok := b.agents.LoadAndRemove(cid); ok && c != nil {
		c.cli.Disconnect(500)
	}
}
