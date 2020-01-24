// Package trace provides an interface for distributed tracing
package trace

import (
	"context"
	"time"
)

// Trace is an interface for distributed tracing
type Trace interface {
	// Start a trace
	Start(ctx context.Context, name string) *Span
	// Finish the trace
	Finish(*Span) error
	// Read the traces
	Read(...ReadOption) ([]*Span, error)
}

// Span is used to record an entry
type Span struct {
	// Id of the trace
	Trace string
	// name of the span
	Name string
	// id of the span
	Id string
	// parent span id
	Parent string
	// Start time
	Started time.Time
	// Duration in nano seconds
	Duration time.Duration
	// associated data
	Metadata map[string]string
}

type spanKey struct{}

// FromContext returns a span from context
func FromContext(ctx context.Context) (*Span, bool) {
	s, ok := ctx.Value(spanKey{}).(*Span)
	return s, ok
}

// NewContext creates a new context with the span
func NewContext(ctx context.Context, s *Span) context.Context {
	return context.WithValue(ctx, spanKey{}, s)
}
