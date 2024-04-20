package messaging

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"reflect"
	"time"
)

// Status - message status
type Status struct {
	core.Status
	Duration time.Duration
}

func StatusOK() *Status {
	return NewStatus(http.StatusOK)
}

func NewStatus(code int) *Status {
	s := new(Status)
	s.Code = code
	return s
}

func NewStatusError(code int, err error) *Status {
	s := new(Status)
	s.Code = code
	s.Err = err
	s.AddLocation()
	return s
}

func NewStatusDuration(code int, duration time.Duration) *Status {
	s := new(Status)
	s.Code = code
	s.Duration = duration
	return s
}

func NewStatusDurationError(code int, duration time.Duration, err error) *Status {
	s := NewStatusDuration(code, duration)
	s.Err = err
	return s
}

func (s *Status) Runtime() *core.Status {
	v := reflect.ValueOf(*s)
	f := v.Field(0)
	i := f.Interface()
	if rts, ok := i.(core.Status); ok {
		return &rts
	}
	return nil
}
