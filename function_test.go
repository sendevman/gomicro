package micro

import (
	"context"
	"sync"
	"testing"

	"github.com/micro/go-micro/registry/mock"
	proto "github.com/micro/go-micro/server/debug/proto"
)

func TestFunction(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// create service
	fn := NewFunction(
		Name("test.function"),
		Context(ctx),
		Registry(mock.NewRegistry()),
		AfterStart(func() error {
			wg.Done()
			return nil
		}),
	)

	// we can't test fn.Init as it parses the command line
	// fn.Init()

	go func() {
		// wait for start
		wg.Wait()

		// test call debug
		req := fn.Client().NewRequest(
			"test.function",
			"Debug.Health",
			new(proto.HealthRequest),
		)

		rsp := new(proto.HealthResponse)

		err := fn.Client().Call(context.TODO(), req, rsp)
		if err != nil {
			t.Fatal(err)
		}

		if rsp.Status != "ok" {
			t.Fatalf("function response: %s", rsp.Status)
		}

		cancel()
	}()

	// run service
	if err := fn.Run(); err != nil {
		t.Fatal(err)
	}
}
