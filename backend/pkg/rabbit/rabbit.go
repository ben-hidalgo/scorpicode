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

// Exchange rabbitmq exchange
type Exchange string

// RKey rabbitmq routing key
type RKey string

// Queue rabbitmq queue
type Queue string

// Rmq Hat Data Access Object
type Rmq interface {
	Send(ex Exchange, key RKey, msg string) error
	SendBlob(ex Exchange, key RKey, msg []byte) error
	SendJSON(ex Exchange, key RKey, msg interface{}) error
}

// enforces the interface is implemented
var _ Rmq = (*impl)(nil)

// New .
func New() Rmq {

	conn := Connect()

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

// Connect .
func Connect() *jazz.Connection {

	conn, err := jazz.Connect(AmqpDsn)
	if err != nil {
		logrus.Panicf("Could not connect to RabbitMQ: %v", err.Error())
	}
	return conn
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
