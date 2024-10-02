package httpxtest

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/core/coretest"
	"github.com/advanced-go/stdlib/httpx"
	"io"
	"testing"
)

func Deserialize[E coretest.ErrorHandler, T any](gotBody, wantBody io.Reader, t *testing.T) (gotT, wantT T, success bool) {
	var e E

	gotStatus := core.StatusOK()
	gotT, gotStatus = httpx.Content[T](gotBody)
	if !gotStatus.OK() && !gotStatus.NoContent() {
		//t.Errorf("Deserialize() %v err = %v", "got", gotStatus.Err)
		e.Handle(gotStatus, t, "got")
		return
	}

	wantStatus := core.StatusOK()
	wantT, wantStatus = httpx.Content[T](wantBody)
	if !wantStatus.OK() && !wantStatus.NoContent() {
		//t.Errorf("Deserialize() %v err = %v", "want", wantStatus.Err)
		e.Handle(wantStatus, t, "want")
		return
	}

	if gotStatus.Code != wantStatus.Code {
		t.Errorf("Deserialize() got status code = %v, want status code = %v", gotStatus.Code, wantStatus.Code)
		return
	}
	return gotT, wantT, true
}
