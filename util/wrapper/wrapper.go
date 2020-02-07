package wrapper

import (
	"context"
	"strings"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/debug/stats"
	"github.com/micro/go-micro/v2/debug/trace"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)

type clientWrapper struct {
	client.Client
	headers metadata.Metadata
}

type traceWrapper struct {
	client.Client

	name  string
	trace trace.Tracer
}

var (
	HeaderPrefix = "Micro-"
)

func (c *clientWrapper) setHeaders(ctx context.Context) context.Context {
	// copy metadata
	mda, _ := metadata.FromContext(ctx)
	md := metadata.Copy(mda)

	// set headers
	for k, v := range c.headers {
		if _, ok := md[k]; !ok {
			md[k] = v
		}
	}

	return metadata.NewContext(ctx, md)
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	ctx = c.setHeaders(ctx)
	return c.Client.Call(ctx, req, rsp, opts...)
}

func (c *clientWrapper) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	ctx = c.setHeaders(ctx)
	return c.Client.Stream(ctx, req, opts...)
}

func (c *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	ctx = c.setHeaders(ctx)
	return c.Client.Publish(ctx, p, opts...)
}

func (c *traceWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	newCtx, s := c.trace.Start(ctx, req.Service()+"."+req.Endpoint())

	err := c.Client.Call(newCtx, req, rsp, opts...)
	if err != nil {
		s.Metadata["error"] = err.Error()
	}

	// finish the trace
	c.trace.Finish(s)

	return err
}

// FromService wraps a client to inject From-Service header into metadata
func FromService(name string, c client.Client) client.Client {
	return &clientWrapper{
		c,
		metadata.Metadata{
			HeaderPrefix + "From-Service": name,
		},
	}
}

// HandlerStats wraps a server handler to generate request/error stats
func HandlerStats(stats stats.Stats) server.HandlerWrapper {
	// return a handler wrapper
	return func(h server.HandlerFunc) server.HandlerFunc {
		// return a function that returns a function
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			// execute the handler
			err := h(ctx, req, rsp)
			// record the stats
			stats.Record(err)
			// return the error
			return err
		}
	}
}

// TraceCall is a call tracing wrapper
func TraceCall(name string, t trace.Tracer, c client.Client) client.Client {
	return &traceWrapper{
		name:   name,
		trace:  t,
		Client: c,
	}
}

// TraceHandler wraps a server handler to perform tracing
func TraceHandler(t trace.Tracer) server.HandlerWrapper {
	// return a handler wrapper
	return func(h server.HandlerFunc) server.HandlerFunc {
		// return a function that returns a function
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			// don't store traces for debug
			if strings.HasPrefix(req.Endpoint(), "Debug.") {
				return h(ctx, req, rsp)
			}

			// get the span
			newCtx, s := t.Start(ctx, req.Service()+"."+req.Endpoint())

			err := h(newCtx, req, rsp)
			if err != nil {
				s.Metadata["error"] = err.Error()
			}

			// finish
			t.Finish(s)

			return err
		}
	}
}
