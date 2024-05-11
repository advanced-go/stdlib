package host

import "github.com/advanced-go/stdlib/messaging"

func Ping(uri any) {
	messaging.Ping(Exchange, uri)
}
