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

func _ExampleResolve() {
	host := ""
	auth := "github/advanced-go/timeseries"
	rsc := "access"
	values := make(url.Values)
	r := NewResolver(nil)

	url := r.Url(host, auth, rsc, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	values.Add("region", "*")
	url = r.Url(host, auth, rsc, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "www.google.com"
	url = r.Url(host, auth, rsc, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "localhost:8080"
	rsc = "v2/" + rsc
	url = r.Url(host, auth, rsc, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	h := make(http.Header)
	url = r.Url(host, auth, rsc, values, h)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	h.Add(BuildPath(auth, rsc, values), testRespName)
	url = r.Url(host, auth, rsc, values, h)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "www.google.com"
	rsc = "search"
	values.Del("region")
	values.Add("q", "golang")
	auth = ""
	url = r.Url(host, auth, rsc, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	//Output:
	//test: Url("","github/advanced-go/timeseries","access") -> [github/advanced-go/timeseries:access]
	//test: Url("","github/advanced-go/timeseries","access") -> [github/advanced-go/timeseries:access?region=%2A]
	//test: Url("www.google.com","github/advanced-go/timeseries","access") -> [https://www.google.com/github/advanced-go/timeseries:access?region=%2A]
	//test: Url("localhost:8080","github/advanced-go/timeseries","v2/access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=%2A]
	//test: Url("localhost:8080","github/advanced-go/timeseries","v2/access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=%2A]
	//test: Url("localhost:8080","github/advanced-go/timeseries","v2/access") -> [file://[cwd]/timeseries1test/get-all-resp-v1.txt]
	//test: Url("www.google.com","","search") -> [https://www.google.com/search?q=golang]

}

func resolverWithProxy() *Resolver {
	return NewResolver([]HostEntry{
		{Key: proxyKey, Host: "localhost:8081", Proxy: false},
		{Key: defaultKey, Host: "www.google.com", Proxy: false},
		{Key: yahooKey, Host: "www.search.yahoo.com", Proxy: true},
		{Key: bingKey, Host: "www.bing.com", Proxy: false},
	},
	)
}

func resolverWithoutProxy() *Resolver {
	return NewResolver([]HostEntry{
		{Key: defaultKey, Host: "www.google.com", Proxy: false},
		{Key: yahooKey, Host: "www.search.yahoo.com", Proxy: false},
		{Key: bingKey, Host: "www.bing.com", Proxy: false},
	},
	)
}

func ExampleResolver() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithProxy()

	url1 := r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = "duckduckgo.com"
	url1 = r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = defaultKey
	url1 = r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Url("","github/advanced-go/search","access") -> [github/advanced-go/search:access]
	//test: Url("duckduckgo.com","github/advanced-go/search","access") -> [https://duckduckgo.com/github/advanced-go/search:access]
	//test: Url("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Url("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]
	//test: Url("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_Empty() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithProxy()

	host = "duckduckgo.com"
	url1 := r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override(nil)
	host = defaultKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Url("duckduckgo.com","github/advanced-go/search","access") -> [https://duckduckgo.com/github/advanced-go/search:access]
	//test: Url2("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Url2("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]
	//test: Url2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_No_Proxy() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithProxy()

	host = defaultKey
	url1 := r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override([]HostEntry{{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false}})
	host = defaultKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Url("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Url2("default","github/advanced-go/search","access") -> [https://www.duckduckgo.com/github/advanced-go/search:access]
	//test: Url2("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]
	//test: Url2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_Proxy() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithProxy()

	host = defaultKey
	url1 := r.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override([]HostEntry{
		{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false},
		{Key: proxyKey, Host: "localhost:8888", Proxy: false},
		{Key: bingKey, Host: "www.bing.com", Proxy: true},
	})
	host = defaultKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.Url(host, auth, rsc, nil, nil)
	fmt.Printf("test: Url2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Url("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Url2("default","github/advanced-go/search","access") -> [https://www.duckduckgo.com/github/advanced-go/search:access]
	//test: Url2("bing","github/advanced-go/search","access") -> [http://localhost:8888/github/advanced-go/search:access]
	//test: Url2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Host() {
	host := ""
	r := resolverWithProxy()

	//host = defaultKey
	host2 := r.Host(host)
	fmt.Printf("test: Host(\"%v\") -> [%v]\n", host, host2)

	host = "www.duckduckgo.com"
	host2 = r.Host(host)
	fmt.Printf("test: Host(\"%v\") -> [%v]\n", host, host2)

	host = defaultKey
	host2 = r.Host(host)
	fmt.Printf("test: Host(\"%v\") -> [%v]\n", host, host2)

	r2 := r.Override([]HostEntry{
		{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false},
		{Key: proxyKey, Host: "localhost:8888", Proxy: false},
		{Key: bingKey, Host: "www.bing.com", Proxy: true},
	})
	host = defaultKey
	host2 = r2.Host(host)
	fmt.Printf("test: Host2(\"%v\") -> [%v]\n", host, host2)

	host = bingKey
	host2 = r2.Host(host)
	fmt.Printf("test: Host2(\"%v\") -> [%v]\n", host, host2)

	host = yahooKey
	host2 = r2.Host(host)
	fmt.Printf("test: Host2(\"%v\") -> [%v]\n", host, host2)

	//Output:
	//test: Host("") -> [error: host is empty]
	//test: Host("www.duckduckgo.com") -> [www.duckduckgo.com]
	//test: Host("default") -> [www.google.com]
	//test: Host2("default") -> [www.duckduckgo.com]
	//test: Host2("bing") -> [localhost:8888]
	//test: Host2("yahoo") -> [localhost:8081]

}
