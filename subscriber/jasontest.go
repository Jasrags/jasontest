package subscriber

import (
	"context"
	"github.com/micro/go-log"

	jasontest "jasontest/proto/jasontest"
)

type Jasontest struct{}

func (e *Jasontest) Handle(ctx context.Context, msg *jasontest.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *jasontest.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
