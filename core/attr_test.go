package core

import (
	"fmt"
	"net/http"
)

func ExampleFmtAttrs() {
	var attrs []any

	s := formatAttrs(attrs)
	fmt.Printf("test: FmtAttrs() -> [empty:%v]\n", len(s) == 0)

	attrs = append(attrs, StatusName)
	attrs = append(attrs, "Bad Request")

	attrs = append(attrs, CodeName)
	attrs = append(attrs, http.StatusBadRequest)

	attrs = append(attrs, "isError")
	attrs = append(attrs, false)

	attrs = append(attrs, "empty-string")
	attrs = append(attrs, "")

	attrs = append(attrs, TimestampName)
	attrs = append(attrs, testTS)

	s = formatAttrs(attrs)
	fmt.Printf("test: FmtAttrs-Even() -> %v\n", s)

	attrs = append(attrs, "name-only")
	s = formatAttrs(attrs)
	fmt.Printf("test: FmtAttrs-Odd() -> %v\n", s)

	//Output:
	//test: FmtAttrs() -> [empty:true]
	//test: FmtAttrs-Even() -> "status":"Bad Request", "code":400, "isError":false, "empty-string":null, "timestamp":"2024-03-01T18:23:50.205Z"
	//test: FmtAttrs-Odd() -> "status":"Bad Request", "code":400, "isError":false, "empty-string":null, "timestamp":"2024-03-01T18:23:50.205Z", "name-only":null

}
