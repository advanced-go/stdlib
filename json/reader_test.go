package json

import (
	"fmt"
	"io"
)

type testStruct struct {
	Status    string `json:"status"`
	CreatedTS string
	UpdatedTS string `json:"updated-ts"`

	CostFunction string `json:"cost-function"`

	// Routing
	PrimaryRoute   string `json:"primary-route"`
	SecondaryRoute string `json:"secondary-route"`
	SecondaryPct   string `json:"secondary-pct"`
}

func ExampleNewReader() {
	var length int64
	reader, _, status := NewReadCloser(nil)
	buf, _ := io.ReadAll(reader)
	fmt.Printf("test: NewReader(nil) -> [status:%v] [reader:%v]\n", status, string(buf))

	t := testStruct{
		Status:         "Status",
		CreatedTS:      "2024-05-24",
		UpdatedTS:      "2024-05-25",
		CostFunction:   "Testing",
		PrimaryRoute:   "www.google.com",
		SecondaryRoute: "www.search.yahoo.com",
		SecondaryPct:   "45",
	}

	reader, length, status = NewReadCloser(t)
	buf, _ = io.ReadAll(reader)
	fmt.Printf("test: NewReader(nil) -> [status:%v] [reader:%v] [length:%v]\n", status, string(buf), length)

	//Output:
	//test: NewReader(nil) -> [status:OK] [reader:]
	//test: NewReader(nil) -> [status:OK] [reader:{"status":"Status","CreatedTS":"2024-05-24","updated-ts":"2024-05-25","cost-function":"Testing","primary-route":"www.google.com","secondary-route":"www.search.yahoo.com","secondary-pct":"45"}] [length:191]

}
