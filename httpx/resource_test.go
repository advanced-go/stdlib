package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
	"time"
)

const (
	originAuthority = "github/advanced-go/origin-resource"
)

var (
	testOrigins = []core.Origin{
		{Region: "region1", Zone: "Zone1", Host: "www.host1.com"},
		{Region: "region1", Zone: "Zone2", Host: "www.host2.com"},
		{Region: "region2", Zone: "Zone1", Host: "www.google.com"},
	}
	originRsc = NewResource[core.Origin](originAuthority, originMatch, originPatch)
)

func originMatch(item any, req *http.Request) bool {
	filter := core.NewOrigin(req.URL.Query())
	if entry, ok := item.(*core.Origin); ok {
		if core.OriginMatch(*entry, filter) {
			return true
		}
	}
	return false
}

func originPatch(item any, patch *Patch) {
	if item == nil || patch == nil {
		return
	}
	if target, ok := item.(*core.Origin); ok {
		for _, op := range patch.Updates {
			switch op.Op {
			case OpReplace:
				if op.Path == core.HostKey {
					if s, ok1 := op.Value.(string); ok1 {
						target.Host = s
					}
				}
			default:
			}
		}
	}
}

func init() {
	ctrl := controller.NewController("origin-resource", controller.NewPrimaryResource("localhost", originAuthority, time.Second*2, "", originRsc.Do), nil)
	controller.RegisterController(ctrl)
}

func createOriginReadCloser(body any) (io.ReadCloser, int64, *core.Status) {
	switch ptr := body.(type) {
	case []core.Origin:
		return json2.NewReadCloser(body)
	case []byte:
		return io.NopCloser(bytes.NewReader(ptr)), int64(len(ptr)), core.StatusOK()
	default:
		return nil, 0, core.NewStatus(http.StatusBadRequest)
	}
}

func ExampleOriginResource() {
	url := originAuthority + ":resiliency"
	rc, _, status0 := createOriginReadCloser(testOrigins)
	fmt.Printf("test: createReaderCloser() -> [status:%v]\n", status0)

	// Get authority
	auth := core.Authority(originRsc.Do)
	fmt.Printf("test: originRsc.Do-AUTH() -> [auth:%v]\n", auth)

	// Put
	req, _ := http.NewRequest(http.MethodPut, url, rc)
	resp, status := originRsc.Do(req)
	fmt.Printf("test: originRsc.Do-PUT() -> [status:%v] [resp:%v] [count:%v]\n", status, resp != nil, originRsc.Count())

	// Get no filter
	req, _ = http.NewRequest(http.MethodGet, url, nil)
	items, status1 := getItems(req)
	fmt.Printf("test: originRsc.Do-GET(*) -> [status:%v] [count:%v]\n", status1, len(items))

	// Get zone=zone1
	req, _ = http.NewRequest(http.MethodGet, url+"?az=zone1", nil)
	items, status1 = getItems(req)
	fmt.Printf("test: originRsc.Do-GET(az=zone1) -> [status:%v] [count:%v]\n", status1, len(items))

	// Patch replace www.google.com -> www.search.yahoo.com
	p := Patch{Updates: []Operation{
		{Op: OpReplace, Path: core.HostKey, Value: "www.search.yahoo.com"},
	}}
	buf, _ := json.Marshal(p)
	req, _ = http.NewRequest(http.MethodPatch, url+"?host=www.google.com", io.NopCloser(bytes.NewReader(buf)))
	resp, status = originRsc.Do(req)
	fmt.Printf("test: originRsc.Do-PATCH() -> [status:%v]\n", status)

	// GET patched Origin
	req, _ = http.NewRequest(http.MethodGet, url+"?host=www.search.yahoo.com", nil)
	items, status = getItems(req)
	fmt.Printf("test: originRsc.Do-GET(host=www.search.yahoo.com) -> [status:%v] [count:%v]\n", status, len(items))

	// DELETE - patched item
	req, _ = http.NewRequest(http.MethodDelete, url+"?host=www.search.yahoo.com", nil)
	resp, status = originRsc.Do(req)
	fmt.Printf("test: originRsc.Do-DELETE(host=www.search.yahoo.com) -> [status:%v] [count:%v]\n", status, originRsc.Count())

	// Get *
	//req, _ = http.NewRequest(http.MethodGet, url, nil)
	//items, status = getItems(req)
	//fmt.Printf("test: originRsc.Do-GET(*) -> [status:%v] [count:%v]\n", status, len(items))

	originRsc.Empty()
	//req, _ = http.NewRequest(http.MethodGet, url, nil)
	//items, status = getItems(req)
	fmt.Printf("test: originRsc.Count() -> [status:%v] [count:%v]\n", core.StatusOK(), originRsc.Count())

	//Output:
	//test: createReaderCloser() -> [status:OK]
	//test: originRsc.Do-AUTH() -> [auth:github/advanced-go/origin-resource]
	//test: originRsc.Do-PUT() -> [status:OK] [resp:true] [count:3]
	//test: originRsc.Do-GET(*) -> [status:Not Found] [count:0]
	//test: originRsc.Do-GET(az=zone1) -> [status:OK] [count:2]
	//test: originRsc.Do-PATCH() -> [status:OK]
	//test: originRsc.Do-GET(host=www.search.yahoo.com) -> [status:OK] [count:1]
	//test: originRsc.Do-DELETE(host=www.search.yahoo.com) -> [status:OK] [count:2]
	//test: originRsc.Count() -> [status:OK] [count:0]

}

func getItems(req *http.Request) ([]core.Origin, *core.Status) {
	resp, status := originRsc.Do(req)
	if !status.OK() {
		return nil, status
	}
	if resp.Body == nil {
		return nil, core.StatusNotFound()
	}
	return json2.New[[]core.Origin](resp.Body, nil)
}
