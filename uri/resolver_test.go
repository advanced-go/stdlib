package uri

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	testRespName = "file://[cwd]/timeseries1test/get-all-resp-v1.txt"
	defaultKey   = "default"
	googleKey    = "google"
	yahooKey     = "yahoo"
	bingKey      = "bing"
)

func ExampleBuildHostWithScheme() {
	host := ""
	o := BuildHostWithScheme(host)
	fmt.Printf("test: BuildHostWithScheme(\"%v\") -> [origin:%v]\n", host, o)

	host = "www.google.com"
	o = BuildHostWithScheme(host)
	fmt.Printf("test: BuildHostWithScheme(\"%v\") -> [origin:%v]\n", host, o)

	host = "localhost:8080"
	o = BuildHostWithScheme(host)
	fmt.Printf("test: BuildHostWithScheme(\"%v\") -> [origin:%v]\n", host, o)

	host = "internalhost"
	o = BuildHostWithScheme(host)
	fmt.Printf("test: BuildHostWithScheme(\"%v\") -> [origin:%v]\n", host, o)

	//Output:
	//test: BuildHostWithScheme("") -> [origin:]
	//test: BuildHostWithScheme("www.google.com") -> [origin:https://www.google.com]
	//test: BuildHostWithScheme("localhost:8080") -> [origin:http://localhost:8080]
	//test: BuildHostWithScheme("internalhost") -> [origin:http://internalhost]

}

func ExampleBuildPath() {
	auth := "github/advanced-go/timeseries"
	path := "v2/access"
	//ver := "v2"
	values := make(url.Values)

	p := BuildPath("", path, values)
	fmt.Printf("test: BuildPath(\"%v\") -> [%v]\n", path, p)

	p = BuildPath(auth, path, values)
	fmt.Printf("test: BuildPath(\"%v\",\"%v\") -> [%v]\n", auth, path, p)

	values.Add("region", "*")
	p = BuildPath("", path, values)
	fmt.Printf("test: BuildPath(\"%v\") -> [%v]\n", path, p)

	p = BuildPath(auth, path, values)
	fmt.Printf("test: BuildPath(\"%v\",\"%v\") -> [%v]\n", auth, path, p)

	//Output:
	//test: BuildPath("v2/access") -> [v2/access]
	//test: BuildPath("github/advanced-go/timeseries","v2/access") -> [github/advanced-go/timeseries:v2/access]
	//test: BuildPath("v2/access") -> [v2/access?region=*]
	//test: BuildPath("github/advanced-go/timeseries","v2/access") -> [github/advanced-go/timeseries:v2/access?region=*]

}

func ExampleResolve_Url() {
	errType := 123
	host := ""
	path := "/search"
	values := make(url.Values)
	r := NewResolver("localhost:8081")

	url1 := r.Url(host, "", path, errType, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	url1 = r.Url(host, "", path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	values.Add("q", "golang")
	url1 = r.Url(host, "", path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	url1 = r.Url(host, "", path, "q=golang", nil)
	fmt.Printf("test: Url_String(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	host = "www.google.com"
	url1 = r.Url(host, "", path, values, nil)
	fmt.Printf("test: Url_String(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	h := make(http.Header)
	h.Add(BuildPath("", path, values), "https://www.search.yahoo.com?q=golang")
	host = "www.google.com"
	url1 = r.Url(host, "", path, values, h)
	fmt.Printf("test: Url_Override(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	//Output:
	//test: Url("","/search") -> [http://localhost:8081/searcherror: query type is invalid int]
	//test: Url("","/search") -> [http://localhost:8081/search]
	//test: Url("","/search") -> [http://localhost:8081/search?q=golang]
	//test: Url_String("","/search") -> [http://localhost:8081/search?q=golang]
	//test: Url_String("www.google.com","/search") -> [https://www.google.com/search?q=golang]
	//test: Url_Override("www.google.com","/search") -> [https://www.google.com/search?q=golang]

}

func ExampleResolve_UrlWithAuthority() {
	host := ""
	auth := "github/advanced-go/timeseries"
	path := "access"
	values := make(url.Values)
	r := NewResolver("")

	url1 := r.Url(host, auth, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	values.Add("region", "*")
	url1 = r.Url(host, auth, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	url1 = r.Url(host, auth, path, "region=*", nil)
	fmt.Printf("test: Url_String(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	host = "www.google.com"
	url1 = r.Url(host, auth, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	host = "localhost:8080"
	path = "v2/" + path
	url1 = r.Url(host, auth, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	h := make(http.Header)
	url1 = r.Url(host, auth, path, values, h)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	//h.Add(BuildPath(auth, path, values), testRespName)
	//url1 = r.Url(host, auth, path, values, h)
	//fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	host = "www.google.com"
	path = "v2/search"
	values.Del("region")
	values.Add("q", "golang")
	auth = ""
	url1 = r.Url(host, auth, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, path, url1)

	//Output:
	//test: Url("","github/advanced-go/timeseries","access") -> [/github/advanced-go/timeseries:access]
	//test: Url("","github/advanced-go/timeseries","access") -> [/github/advanced-go/timeseries:access?region=*]
	//test: Url_String("","github/advanced-go/timeseries","access") -> [/github/advanced-go/timeseries:access?region=*]
	//test: Url("www.google.com","github/advanced-go/timeseries","access") -> [https://www.google.com/github/advanced-go/timeseries:access?region=*]
	//test: Url("localhost:8080","github/advanced-go/timeseries","v2/access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=*]
	//test: Url("localhost:8080","github/advanced-go/timeseries","v2/access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=*]
	//test: Url("www.google.com","","v2/search") -> [https://www.google.com/v2/search?q=golang]

}

func ExampleCreateUrl() {
	path1 := "advanced-go/observation:v1/timeseries/egress/entry?region=*"
	path2 := "advanced-go/observation:v1/timeseries/egress/entry?region=**"
	url1 := "file:///f:/resource/info.json"
	url2 := "file:///f:/resource/test.json"

	path := ""
	h := make(http.Header)
	h.Add(XContentResolver, path)
	uri := createUrl(h, "")
	fmt.Printf("test: createUrl(empty) -> [%v]\n", uri)

	h = AddResolverEntry(nil, path1, url1)
	AddResolverEntry(h, path2, url2)
	uri = createUrl(h, path1)
	fmt.Printf("test: createUrl(\"%v\") -> %v\n", path1, uri)

	uri = createUrl(h, path2)
	fmt.Printf("test: createUrl(\"%v\") -> %v\n", path2, uri)

	//Output:
	//test: createUrl(empty) -> []
	//test: createUrl("advanced-go/observation:v1/timeseries/egress/entry?region=*") -> file:///f:/resource/info.json
	//test: createUrl("advanced-go/observation:v1/timeseries/egress/entry?region=**") -> file:///f:/resource/test.json

}
