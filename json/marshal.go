package json

import (
	json2 "encoding/json"
	"errors"
	"github.com/advanced-go/stdlib/core"
)

func Marshal(v any) ([]byte, *core.Status) {
	if v == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: value is nil"))
	}
	buf, err := json2.Marshal(v)
	if err != nil {
		return nil, core.NewStatusError(core.StatusJsonEncodeError, err)
	}
	return buf, core.StatusOK()

}
