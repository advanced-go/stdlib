package messaging

import (
	"context"
	"errors"
	"fmt"
	uri2 "github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

const (
	timeout = time.Second * 3
)

const (
	pingLocation = PkgPath + ":Ping"
)

// Ping - function to "ping" a resource
func Ping(ctx context.Context, uri any) *Status {
	return ping(ctx, HostExchange, uri)
}

func ping(ctx context.Context, ex *Exchange, uri any) *Status {
	to, status := createTo(uri)
	if !status.OK() {
		return status
	}
	var response *Message

	result := make(chan *Status)
	reply := make(chan *Message, 16)
	msg := NewControlMessage(to, PkgPath, PingEvent)
	msg.ReplyTo = NewReceiverReplyTo(reply)
	err := ex.Send(msg)
	if err != nil {
		return NewStatusError(http.StatusInternalServerError, err)
	}
	go Receiver(timeout, reply, result, func(msg *Message) bool {
		response = msg
		return true
	})
	status = <-result
	status.AddLocation()
	if response != nil {
		status.Code = response.Status().Code
		status.Err = response.Status().Err
	}
	close(reply)
	close(result)
	return status
}

func createTo(uri any) (string, *Status) {
	if uri == nil {
		return "", NewStatusError(http.StatusBadRequest, errors.New("error: Ping() uri is nil"))
	}
	path := ""
	if u, ok := uri.(*url.URL); ok {
		path = u.Path
	} else {
		if u2, ok1 := uri.(string); ok1 {
			path = u2
		} else {
			return "", NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: Ping() uri is invalid type: %v", reflect.TypeOf(uri).String())))
		}
	}
	nid, _, ok := uri2.UprootUrn(path)
	if !ok {
		return "", NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: Ping() uri is not a valid URN %v", path)))
	}
	return nid, StatusOK()
}
