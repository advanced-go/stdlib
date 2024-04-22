package messaging

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

func NewStatusDurationError2(code int, duration time.Duration, err error) *core.Status {
	s := core.NewStatusDuration(code, duration)
	s.Err = err
	return s
}
