package rabbit

import (
	"backend/pkg/envconfig"
	"bytes"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
	"github.com/twitchtv/twirp"
)

// AmqpDsn .
var AmqpDsn = "amqp://rabbit:rabbit@localhost:5672/"

func init() {
	envconfig.SetString("AMQP_DSN", &AmqpDsn)
}

const (
	// ServiceMsgtypeTx topic exchange
	ServiceMsgtypeTx = "service_msgtype_tx"
)

// Rmq Hat Data Access Object
type Rmq interface {
	Send(ex, key, msg string) error
}

// enforces the interface is implemented
var _ Rmq = (*impl)(nil)

// New .
func New() Rmq {

	conn, err := jazz.Connect(AmqpDsn)
	if err != nil {
		logrus.Panicf("Could not connect to RabbitMQ: %v", err.Error())
	}

	reader := bytes.NewReader([]byte(schema))

	scheme, err := jazz.DecodeYaml(reader)
	if err != nil {
		logrus.Panicf("Could not read YAML: %v", err.Error())
	}

	err = conn.CreateScheme(scheme)
	if err != nil {
		logrus.Panicf("Could not create scheme: %v", err.Error())
	}

	return &impl{
		Conn: conn,
	}
}

// used to store the Repo in Context
type key int

// Key is the key for the repo in context; public for mock injection
const Key key = 0

// From returns the dao and panics if not found
func From(ctx context.Context) Rmq {

	switch v := ctx.Value(Key).(type) {
	case Rmq:
		return v
	default:
		panic("rabbit.From() no value found")
	}
}

// ServerHooks is a Twirp middleware
func ServerHooks() *twirp.ServerHooks {

	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {

			return context.WithValue(ctx, Key, New()), nil
		},
	}
}

const schema = `
exchanges:
  service_msgtype_tx:
    durable: true
    type: topic
queues:
  hats_q:
    durable: true
    bindings:
      - exchange: "service_msgtype_tx"
        key: "hats.*"
`
