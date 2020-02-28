package handler

import (
	"context"

	"github.com/micro/go-log"

	jasontest "jasontest/proto/jasontest"
)

type Jasontest struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Jasontest) Call(ctx context.Context, req *jasontest.Request, rsp *jasontest.Response) error {
	log.Log("Received Jasontest.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Jasontest) Stream(ctx context.Context, req *jasontest.StreamingRequest, stream jasontest.Jasontest_StreamStream) error {
	log.Logf("Received Jasontest.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&jasontest.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Jasontest) PingPong(ctx context.Context, stream jasontest.Jasontest_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&jasontest.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
