package host

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
)

func Ping(uri any) *core.Status {
	return messaging.Ping(Exchange, uri)
}
