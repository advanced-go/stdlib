package httpxtest

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	io2 "github.com/advanced-go/stdlib/io"
	"net/http"
	"os"
	"testing"
)

/*
const (
	fileExistsError = "The system cannot find the file specified"
	fileScheme      = "file"
)

// NewResponse - read an HTTP response given a URL
func NewResponse(uri any) (*http.Response, *core.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: "Internal Error"}

	if uri == nil {
		return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	buf, status := io.ReadFile(uri)
	if !status.OK() {
		if strings.Contains(status.Err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found"}, core.NewStatusError(core.StatusInvalidArgument, status.Err)
		}
		return serverErr, core.NewStatusError(core.StatusIOError, status.Err)
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, core.NewStatusError(core.StatusIOError, err2)
	}
	return resp1, core.StatusOK()
}

*/

// NewResponseTest - used from test packages
func NewResponseTest(uri any, t *testing.T) *http.Response {
	resp, status := httpx.NewResponseFromUri(uri)
	if status.OK() {
		return resp
	}
	t.Errorf("ReadResponse() err = %v", status.Err.Error())
	return &http.Response{StatusCode: http.StatusTeapot}
}

func writeValues(buf *bytes.Buffer, name string, values []string) {
	for _, value := range values {
		buf.WriteString(fmt.Sprintf("%v: %v\n", name, value))
	}
}

func writeHeader(buf *bytes.Buffer, resp *http.Response) {
	for name, values := range resp.Header {
		writeValues(buf, name, values)
	}
}

func WriteResponse(url string, resp *http.Response) *core.Status {
	if url == "" || resp == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: url is empty or response is nil"))
	}
	var buf bytes.Buffer

	// Write status line
	buf.WriteString(fmt.Sprintf("%v %v\n", resp.Proto, resp.Status))

	// Write header
	writeHeader(&buf, resp)
	buf.WriteString("\n")

	// Write content
	buf1, status := io2.ReadAll(resp.Body, nil)
	if !status.OK() {
		return status
	}
	count, err := buf.Write(buf1)
	if err != nil {
		return core.NewStatusError(core.StatusIOError, err)
	}
	if count != len(buf1) {
		return core.NewStatusError(core.StatusIOError, errors.New("error: writing bytes"))
	}

	// Create filename and write file
	fname := io2.FileName(url)
	// 0666 is read only
	err = os.WriteFile(fname, buf.Bytes(), 0777)
	if err != nil {
		return core.NewStatusError(core.StatusIOError, err)
	}
	return core.StatusOK()
}
