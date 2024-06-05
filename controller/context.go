package controller

import (
	"context"
	"time"
)

//func WithDeadline(ctx context.Context, timeout time.Duration)

func ContextWithTimeout(ctx context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	} else {
		if _, ok := ctx.Deadline(); ok {
			return ctx, cancel
		}
	}
	if duration == 0 {
		return ctx, cancel
	}
	return context.WithTimeout(ctx, duration)
}

func cancel() {}
