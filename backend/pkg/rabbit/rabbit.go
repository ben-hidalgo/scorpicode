package rabbit

import (
	"backend/pkg/envconfig"
	"bytes"
	"context"

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
func New(jc *jazz.Connection) Rmq {
	return &impl{
		Conn: jc,
	}
}

// Connect .
func Connect() (*jazz.Connection, error) {

	jc, err := jazz.Connect(AmqpDsn)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader([]byte(schema))

	scheme, err := jazz.DecodeYaml(reader)
	if err != nil {
		return nil, err
	}

	err = jc.CreateScheme(scheme)
	if err != nil {
		return nil, err
	}

	return jc, nil
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

// ServerHooks is a Twirp middleware which injects the Rabbit Rmq impl
func ServerHooks(jc *jazz.Connection) *twirp.ServerHooks {

	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {

			return context.WithValue(ctx, Key, New(jc)), nil
		},
	}
}
