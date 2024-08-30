package httpxtest

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/url"
)

// ParseRaw - parse a raw Uri without error
func ParseRaw(rawUri string) *url.URL {
	u, _ := url.Parse(rawUri)
	return u
}

func ReadContent(rawHttp []byte) (*bytes.Buffer, error) {
	var content = new(bytes.Buffer)
	var writeTo = false

	reader := bufio.NewReader(bytes.NewReader(rawHttp))
	for {
		line, err := reader.ReadString('\n')
		if len(line) <= 2 && !writeTo {
			writeTo = true
			continue
		}
		if err == io.EOF {
			break
		} else {
			if err != nil {
				return nil, err
			}
		}
		if writeTo {
			//fmt.Printf("%v", line)
			content.Write([]byte(line))
		}
	}
	return content, nil
}

func Content[T any](body io.Reader) (t T, status *core.Status) {
	if body == nil {
		return t, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: body is nil"))
	}
	var buf []byte
	buf, status = io2.ReadAll(body, nil)
	if !status.OK() || len(buf) == 0 {
		return
	}
	switch p := any(&t).(type) {
	case *[]byte:
		*p = buf
	case *string:
		*p = string(buf)
	default:
		err := json.NewDecoder(body).Decode(p)
		if err != nil {
			status = core.NewStatusError(core.StatusJsonDecodeError, err)
		}
	}
	return
}
