package handler

import (
	"context"
	"runtime"
	"time"

	proto "github.com/micro/go-micro/debug/proto"
)

var (
	// DefaultHandler is default debug handler
	DefaultHandler = newDebug()
)

type Debug struct {
	started int64
}

func newDebug() *Debug {
	return &Debug{
		started: time.Now().Unix(),
	}
}

func (d *Debug) Health(ctx context.Context, req *proto.HealthRequest, rsp *proto.HealthResponse) error {
	rsp.Status = "ok"
	return nil
}

func (d *Debug) Stats(ctx context.Context, req *proto.StatsRequest, rsp *proto.StatsResponse) error {
	var mstat runtime.MemStats
	runtime.ReadMemStats(&mstat)

	rsp.Started = uint64(d.started)
	rsp.Uptime = uint64(time.Now().Unix() - d.started)
	rsp.Memory = mstat.Alloc
	rsp.Gc = mstat.PauseTotalNs
	rsp.Threads = uint64(runtime.NumGoroutine())
	return nil
}

func (d *Debug) Log(ctx context.Context, req *proto.LogRequest, rsp proto.Debug_LogStream) error {
	return nil
}
