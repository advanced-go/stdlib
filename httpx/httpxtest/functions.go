package httpxtest

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

const (
	contentType     = "Content-Type"
	contentTypeJson = "application/json"
)

type Args struct {
	Item string
	Got  string
	Want string
	Err  error
}

func ReadHttp(basePath, reqName, respName string) ([]Args, *http.Request, *http.Response) {
	path := basePath + reqName
	req, err := ReadRequest(ParseRaw(path))
	if err != nil {
		return []Args{{Item: fmt.Sprintf("ReadRequest(%v)", path), Got: "", Want: "", Err: err}}, nil, nil
	}
	path = basePath + respName
	resp, status1 := readResponse(ParseRaw(path))
	if !status1.OK() {
		return []Args{{Item: fmt.Sprintf("ReadResponse(%v)", path), Got: "", Want: "", Err: status1.Err}}, nil, nil
	}
	return nil, req, resp
}

func Headers(got *http.Response, want *http.Response, names ...string) (failures []Args) {
	if names == nil {
		for _, name := range want.Header {
			names = append(names, name[0])
		}
	}
	for _, name := range names {
		wantVal := want.Header.Get(name)
		if wantVal == "" {
			return []Args{{Item: name, Got: "", Want: "", Err: errors.New(fmt.Sprintf("want header [%v] is missing or empty", name))}}
		}
		gotVal := got.Header.Get(name)
		if wantVal != gotVal {
			failures = append(failures, Args{Item: name, Got: gotVal, Want: wantVal, Err: nil})
		}
	}
	return failures
}

func Content(got *http.Response, want *http.Response) (failures []Args, gotBuf []byte, wantBuf []byte) {
	// validate content type matches
	fails, _ := validateContentType(got, want)
	if fails != nil {
		failures = fails
		return
	}
	var status *core.Status

	// validate body IO
	wantBuf, status = io.ReadAll(want.Body, nil)
	if status.Err != nil {
		failures = []Args{{Item: "want.Body", Got: "", Want: "", Err: status.Err}}
		return
	}
	gotBuf, status = io.ReadAll(got.Body, nil)
	if status.Err != nil {
		failures = []Args{{Item: "got.Body", Got: "", Want: "", Err: status.Err}}
	}
	if len(gotBuf) != len(wantBuf) {
		failures = []Args{{Item: "Body", Got: fmt.Sprintf("%v", gotBuf), Want: fmt.Sprintf("%v", wantBuf), Err: errors.New("content length does not match")}}
	}
	return
}

func validateContentType(got *http.Response, want *http.Response) (failures []Args, ct string) {
	ct = want.Header.Get(contentType)
	if ct == "" {
		return []Args{{Item: contentType, Got: "", Want: "", Err: errors.New("want Response header Content-Type is empty")}}, ct
	}
	gotCt := got.Header.Get(contentType)
	if gotCt != ct {
		return []Args{{Item: contentType, Got: gotCt, Want: ct, Err: nil}}, ct
	}
	return nil, ct
}

/*
// optional
//if testBytes != nil {
//	failures = testBytes(got, gotBytes, want, wantBytes)
//	if failures != nil {
return
}
//	}


	// if no content is wanted, return
	if len(wantBytes) == 0 {
		return
	}

	// validate content length
	//if len(gotBytes) != len(wantBytes) {
	//	failures = []Args{{Item: "Content-Length", Got: fmt.Sprintf("%v", len(gotBytes)), Want: fmt.Sprintf("%v", len(wantBytes))}}
	//	return
	//}

	// validate content type is application/json
	if ct != contentTypeJson {
		failures = []Args{{Item: "Content-Type", Got: "", Want: "", Err: errors.New(fmt.Sprintf("invalid content type for serialization [%v]", ct))}}
		return
	}

	// unmarshal
	err := json.Unmarshal(wantBytes, &wantT)
	if err != nil {
		failures = []Args{{Item: "want.Unmarshal()", Got: "", Want: "", Err: err}}
		return
	}
	err = json.Unmarshal(gotBytes, &gotT)
	if err != nil {
		failures = []Args{{Item: "got.Unmarshal()", Got: "", Want: "", Err: err}}
	} else {
		content = true
	}

*/
