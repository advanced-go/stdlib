package uri

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	testRespName = "file://[cwd]/timeseries1test/get-all-resp-v1.txt"
)

func ExampleBuildRsc() {
	ver := ""
	rsc := "access"
	r := BuildRsc(ver, rsc)

	fmt.Printf("test: BuildRsc(\"%v\",\"%v\") -> [%v]\n", ver, rsc, r)

	ver = "v1"
	r = BuildRsc(ver, rsc)
	fmt.Printf("test: BuildRsc(\"%v\",\"%v\") -> [%v]\n", ver, rsc, r)

	//Output:
	//test: BuildRsc("","access") -> [access]
	//test: BuildRsc("v1","access") -> [v1/access]

}

func ExampleBuildOrigin() {
	host := ""
	o := BuildOrigin(host)
	fmt.Printf("test: BuildOrigin(\"%v\") -> [origin:%v]\n", host, o)

	host = "www.google.com"
	o = BuildOrigin(host)
	fmt.Printf("test: BuildOrigin(\"%v\") -> [origin:%v]\n", host, o)

	host = "localhost:8080"
	o = BuildOrigin(host)
	fmt.Printf("test: BuildOrigin(\"%v\") -> [origin:%v]\n", host, o)

	host = "internalhost"
	o = BuildOrigin(host)
	fmt.Printf("test: BuildOrigin(\"%v\") -> [origin:%v]\n", host, o)

	//Output:
	//test: BuildOrigin("") -> [origin:]
	//test: BuildOrigin("www.google.com") -> [origin:https://www.google.com]
	//test: BuildOrigin("localhost:8080") -> [origin:http://localhost:8080]
	//test: BuildOrigin("internalhost") -> [origin:http://internalhost]

}

func ExampleBuildPath() {
	auth := "github/advanced-go/timeseries"
	rsc := "access"
	values := make(url.Values)
	p := BuildPath(auth, rsc, values)

	fmt.Printf("test: BuildPath(\"%v\",\"%v\") -> [%v]\n", auth, rsc, p)

	values.Add("region", "*")
	rsc = "v2/access"
	p = BuildPath(auth, rsc, values)
	fmt.Printf("test: BuildPath(\"%v\",\"%v\") -> [%v]\n", auth, rsc, p)

	//Output:
	//test: BuildPath("github/advanced-go/timeseries","access") -> [github/advanced-go/timeseries:access]
	//test: BuildPath("github/advanced-go/timeseries","v2/access") -> [github/advanced-go/timeseries:v2/access?region=%2A]

}

func ExampleResolve() {
	host := ""
	auth := "github/advanced-go/timeseries"
	rsc := "access"
	values := make(url.Values)

	url := Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	values.Add("region", "*")
	url = Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "www.google.com"
	url = Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "localhost:8080"
	rsc = "v2/" + rsc
	url = Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	h := make(http.Header)
	url = Resolve(host, auth, rsc, values, h)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	h.Add(BuildPath(auth, rsc, values), testRespName)
	url = Resolve(host, auth, rsc, values, h)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "www.google.com"
	rsc = "search"
	values.Del("region")
	values.Add("q", "golang")
	auth = ""
	url = Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	//Output:
	//test: Resolve("","github/advanced-go/timeseries","access") -> [github/advanced-go/timeseries:access]
	//test: Resolve("","github/advanced-go/timeseries","access") -> [github/advanced-go/timeseries:access?region=%2A]
	//test: Resolve("www.google.com","github/advanced-go/timeseries","access") -> [https://www.google.com/github/advanced-go/timeseries:access?region=%2A]
	//test: Resolve("localhost:8080","github/advanced-go/timeseries","v2/access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=%2A]
	//test: Resolve("localhost:8080","github/advanced-go/timeseries","v2/access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=%2A]
	//test: Resolve("localhost:8080","github/advanced-go/timeseries","v2/access") -> [file://[cwd]/timeseries1test/get-all-resp-v1.txt]
	//test: Resolve("www.google.com","","search") -> [https://www.google.com/search?q=golang]

}
