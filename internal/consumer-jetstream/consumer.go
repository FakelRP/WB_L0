package consumer

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
)

type natsConsumer interface {
	Consume(handler jetstream.MessageHandler, opts ...jetstream.PullConsumeOpt) (jetstream.ConsumeContext, error)
}

type handler interface {
	Handle(ctx context.Context, msg []byte) error
}

// Consumer ...
type Consumer struct {
	natsConsumer   natsConsumer
	handler        handler
	consumeContext jetstream.ConsumeContext
}

// New ...
func New(
	natsConsumer natsConsumer,
	handler handler,
) *Consumer {
	return &Consumer{
		natsConsumer: natsConsumer,
		handler:      handler,
	}
}

// Run ...
func (c *Consumer) Run() error {
	var err error
	c.consumeContext, err = c.natsConsumer.Consume(func(msg jetstream.Msg) {

		fmt.Println()
		fmt.Printf("message: %+v", msg)

		err := c.handler.Handle(context.Background(), msg.Data())
		if err != nil {
			nak(msg)
		}
		ack(msg)
	})
	if err != nil {
		return fmt.Errorf("cannot consume. Err: %w", err)
	}

	return nil
}

// Close ...
func (c *Consumer) Close() error {
	c.consumeContext.Stop()

	return nil
}

func nak(msg jetstream.Msg) {
	err := msg.Nak() // todo можно настроить задержку переотправки
	if err != nil {
		// todo log
	}
}

func ack(msg jetstream.Msg) {
	err := msg.Ack()
	if err != nil {
		// todo log
	}
}
