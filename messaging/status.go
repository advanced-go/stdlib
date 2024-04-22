package messaging

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

func NewStatusDuration(code int, duration time.Duration) *core.Status {
	s := core.NewStatus(code)
	s.Duration = duration
	return s
}

func NewStatusDurationError(code int, duration time.Duration, err error) *core.Status {
	s := NewStatusDuration(code, duration)
	s.Err = err
	return s
}
