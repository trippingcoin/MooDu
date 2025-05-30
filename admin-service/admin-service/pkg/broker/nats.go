package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type NATSPublisher struct {
	conn *nats.Conn
}

func NewNATSPublisher(url string) (*NATSPublisher, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NATSPublisher{conn: nc}, nil
}

func (p *NATSPublisher) Publish(subject string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if err := p.conn.Publish(subject, data); err != nil {
		return err
	}
	log.Printf("Published event to subject %s", subject)
	return nil
}
