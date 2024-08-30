package httpxtest

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/http"
	"os"
)

func ReadRequest(uri any) (*http.Request, *core.Status) {
	if uri == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return nil, errors.New(fmt.Sprintf("error: invalid URL scheme : %v", u.Scheme))
	//}
	buf, err := os.ReadFile(io2.FileName(uri))
	if err != nil {
		return nil, core.NewStatusError(core.StatusIOError, err)
	}
	byteReader := bytes.NewReader(buf)
	reader := bufio.NewReader(byteReader)
	req, err1 := http.ReadRequest(reader)
	if err1 != nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, err1)
	}
	bytes1, err2 := ReadContent(buf)
	if err2 != nil {
		return req, core.NewStatusError(core.StatusIOError, err2)
	}
	if bytes1 != nil {
		req.Body = io.NopCloser(bytes1)
	}
	location := req.Header.Get(httpx.ContentLocation)
	if location != "" {
		ctx := core.NewUrlContext(nil, location)
		req2, err3 := http.NewRequestWithContext(ctx, req.Method, req.URL.String(), req.Body)
		if err3 != nil {
			return nil, core.NewStatusError(core.StatusInvalidArgument, err3)
		}
		req2.Header = req.Header
		return req2, core.StatusOK()
	}
	return req, core.StatusOK()
}
