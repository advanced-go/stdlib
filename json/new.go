package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

// NewConstraints - constraints
//type NewConstraints interface {
//	string | *url.URL | []byte | io.Reader | io.ReadCloser
//}

const (
	eofError = "EOF"
)

func decodeStatus(err error) *core.Status {
	if err == nil || err.Error() == "" {
		return core.StatusOK()
	}
	// If the error is "EOF", then the body was empty. If the error is "unexpected EOF", then the body has content
	// but the EOF was reached when more JSON content was expected.
	if err.Error() == eofError {
		return core.StatusNoContent()
	}
	return core.NewStatusError(core.StatusJsonDecodeError, err)
}

// New - create a new type from JSON content, supporting: string, *url.URL, []byte, io.Reader, io.ReadCloser
func New[T any](v any, h http.Header) (t T, status *core.Status) {
	var buf []byte

	if v == nil {
		return t, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: value parameter is nil"))
	}
	switch ptr := v.(type) {
	case string:
		if isStatusURL(ptr) {
			return t, NewStatusFrom(ptr)
		}
		buf, status = io2.ReadFileWithEncoding(ptr, h)
		if !status.OK() {
			return
		}
		err := json.Unmarshal(buf, &t)
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case *url.URL:
		if isStatusURL(ptr.String()) {
			return t, NewStatusFrom(ptr.String())
		}
		buf, status = io2.ReadFileWithEncoding(ptr.String(), h)
		if !status.OK() {
			return
		}
		err := json.Unmarshal(buf, &t)
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case []byte:
		buf, status = io2.Decode(ptr, h)
		if !status.OK() {
			return
		}
		err := json.Unmarshal(buf, &t)
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case io.Reader:
		reader, status0 := io2.NewEncodingReader(ptr, h)
		if !status0.OK() {
			return t, status0.AddLocation()
		}
		err := json.NewDecoder(reader).Decode(&t)
		_ = reader.Close()
		//if err != nil {
		//	return t, decodeStatus(err)
		//	}
		return t, decodeStatus(err)
	case io.ReadCloser:
		reader, status0 := io2.NewEncodingReader(ptr, h)
		if !status0.OK() {
			return t, status0.AddLocation()
		}
		err := json.NewDecoder(reader).Decode(&t)
		_ = reader.Close()
		_ = ptr.Close()
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case *http.Request:
		return New[T](ptr.Body, h)
	case *http.Response:
		return New[T](ptr.Body, h)
	default:
		return t, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: invalid type [%v]", reflect.TypeOf(v))))
	}
}

/*
	case *http.Response:
		if ptr1, ok := any(&t).(*[]byte); ok {
			buf, status = ReadAll(ptr.Body,h)
			if !status.OK() {
				return
			}
			*ptr1 = buf
			return t, StatusOK()
		}
		err := json.NewDecoder(ptr.Body).Decode(&t)
		_ = ptr.Body.Close()
		if err != nil {
			return t, NewStatusError(StatusJsonDecodeError, newLoc, err)
		}
		return t, StatusOK()
	case *http.Request:
		if ptr1, ok := any(&t).(*[]byte); ok {
			buf, status = ReadAll(ptr.Body)
			if !status.OK() {
				return
			}
			*ptr1 = buf
			return t, StatusOK()
		}
		err := json.NewDecoder(ptr.Body).Decode(&t)
		_ = ptr.Body.Close()
		if err != nil {
			return t, NewStatusError(StatusJsonDecodeError, newLoc, err)
		}
		return t, StatusOK()

*/
