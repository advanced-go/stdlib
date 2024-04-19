package io2

import (
	"errors"
	"fmt"
	"github.com/advaced-go/stdlib/core"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	readFileLoc = PkgPath + ":ReadFile"
	readAllLoc  = PkgPath + ":ReadAll"
)

// ReadFile - read a file with a Status
func ReadFile(uri string) ([]byte, *core.Status) {
	status := ValidateUri(uri)
	if !status.OK() {
		return nil, status
	}
	buf, err := os.ReadFile(FileName(uri))
	if err != nil {
		return nil, core.NewStatusError(core.StatusIOError, err)
	}
	return buf, core.StatusOK()
}

// ReadAll - read the body with a Status
func ReadAll(body io.Reader, h http.Header) ([]byte, *core.Status) {
	if body == nil {
		return nil, core.StatusOK()
	}
	if rc, ok := any(body).(io.ReadCloser); ok {
		defer func() {
			err := rc.Close()
			if err != nil {
				fmt.Printf("error: io.ReadCloser.Close() [%v]", err)
			}
		}()
	}
	reader, status := NewEncodingReader(body, h)
	if !status.OK() {
		return nil, status.AddLocation()
	}
	buf, err := io.ReadAll(reader)
	_ = reader.Close()
	if err != nil {
		return nil, core.NewStatusError(core.StatusIOError, err)
	}
	return buf, core.StatusOK()
}

func ValidateUri(uri string) *core.Status {
	if len(uri) == 0 {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URI is empty"))
	}
	if !strings.HasPrefix(uri, fileScheme) {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: URI is not of scheme file: %v", uri)))
	}
	////if !isJsonURL(uri) {
	//	return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URI is not a JSON file"))
	//}
	return core.StatusOK()
}
