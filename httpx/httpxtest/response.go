package httpxtest

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	fileExistsError = "The system cannot find the file specified"
	fileScheme      = "file"
)

// readResponse - read a Http response given a URL
func readResponse(u *url.URL) (*http.Response, *core.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: "Internal Error"}

	if u == nil {
		return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	if u.Scheme != fileScheme {
		return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	}
	buf, err := os.ReadFile(io.FileName(u))
	if err != nil {
		if strings.Contains(err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found"}, core.NewStatusError(core.StatusInvalidArgument, err)
		}
		return serverErr, core.NewStatusError(core.StatusIOError, err)
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, core.NewStatusError(core.StatusIOError, err2)
	}
	return resp1, core.StatusOK()

}
