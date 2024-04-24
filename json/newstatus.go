package json

import (
	"encoding/json"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
	"os"
	"strings"
)

const (
	StatusOKUri       = "urn:status:ok"
	StatusNotFoundUri = "urn:status:notfound"
	StatusTimeoutUri  = "urn:status:timeout"
	statusToken       = "status"
)

type serializedStatusState struct {
	Code     int    `json:"code"`
	Location string `json:"location"`
	Err      string `json:"err"`
}

// isStatusURL - determine if the file name of the URL contains the text 'status'
func isStatusURL(url string) bool {
	if len(url) == 0 {
		return false
	}
	i := strings.LastIndex(url, statusToken)
	if i == -1 {
		return false
	}
	return strings.LastIndex(url, "/") < i
}

// NewStatusFrom - create a new Status from a URI
func NewStatusFrom(uri string) *core.Status {
	status := statusFromConst(uri)
	if status != nil {
		return status
	}
	//status = ValidateUri(uri)
	//if !status.OK() {
	//	return status
	//}
	buf, err1 := os.ReadFile(io.FileName(uri))
	if err1 != nil {
		return core.NewStatusError(core.StatusIOError, err1)
	}
	var status2 serializedStatusState
	err := json.Unmarshal(buf, &status2)
	if err != nil {
		return core.NewStatusError(core.StatusJsonDecodeError, err)
	}
	if len(status2.Err) > 0 {
		return core.NewStatusError(status2.Code, errors.New(status2.Err))
	}
	return core.NewStatus(status2.Code).AddLocation()
}

func statusFromConst(url string) *core.Status {
	if len(url) == 0 {
		return core.StatusOK()
	}
	switch url {
	case StatusOKUri:
		return core.StatusOK()
	case StatusNotFoundUri:
		return core.NewStatus(http.StatusNotFound)
	case StatusTimeoutUri:
		return core.NewStatus(http.StatusGatewayTimeout)
	}
	return nil
}
