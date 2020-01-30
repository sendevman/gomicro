package client

import (
	"context"
	"math"
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	delta := time.Duration(0)

	c := NewClient()

	for i := 0; i < 5; i++ {
		d, err := exponentialBackoff(context.TODO(), c.NewRequest("test", "test", nil), i)
		if err != nil {
			t.Fatal(err)
		}

		if d < delta {
			t.Fatalf("Expected greater than %v, got %v", delta, d)
		}

		delta = time.Millisecond * 100 * time.Duration(math.Pow(math.E, float64(i)))
	}
}
