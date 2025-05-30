package nats

import (
	"time"

	"github.com/aftosmiros/moodu/user-service/config"
	"github.com/nats-io/nats.go"
)

func NewNatsConn(cfg config.Config) (*nats.Conn, error) {
	opts := []nats.Option{
		nats.Name("User Service"),
		nats.Timeout(10 * time.Second),
	}

	conn, err := nats.Connect(cfg.Nats.Host, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
