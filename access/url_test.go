package access

import (
	"fmt"
	"net/http"
	url2 "net/url"
)

func ExampleURLComponents() {
	values := make(url2.Values)
	values.Add("q", "*")
	uri := "http://localhost:8081/github/advanced-go/search:yahoo" + "?" + values.Encode()
	req, _ := http.NewRequest("", uri, nil)

	url, host, path, query := CreateURLComponents(req)
	fmt.Printf("test: CreateURLComponents(\"%v\") -> [url:%v] [host:%v] [path:%v] [query:%v]\n", uri, url, host, path, query)

	uri = "http://www.google.com/search/+all/usa" + "?" + values.Encode()
	req, _ = http.NewRequest("", uri, nil)

	url, host, path, query = CreateURLComponents(req)
	fmt.Printf("test: CreateURLComponents(\"%v\") -> [url:%v] [host:%v] [path:%v] [query:%v]\n", uri, url, host, path, query)

	//Output:
	//fail

}
