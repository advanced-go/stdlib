package httpxtest

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"io"
	"testing"
)

func Deserialize[T any](gotBody, wantBody io.Reader, t *testing.T) (gotT, wantT T, success bool) {
	gotStatus := core.StatusOK()
	gotT, gotStatus = httpx.Content[T](gotBody)
	if !gotStatus.OK() && !gotStatus.NotFound() {
		t.Errorf("Deserialize() %v err = %v", "got", gotStatus.Err)
		return
	}

	wantStatus := core.StatusOK()
	wantT, wantStatus = httpx.Content[T](wantBody)
	if !wantStatus.OK() && !wantStatus.NotFound() {
		t.Errorf("Deserialize() %v err = %v", "want", wantStatus.Err)
		return
	}

	if gotStatus.Code != wantStatus.Code {
		t.Errorf("Deserialize() got status code = %v, want status code = %v", gotStatus.Code, wantStatus.Code)
		return
	}
	return gotT, wantT, true
}
