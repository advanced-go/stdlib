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

	url := r.Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	values.Add("region", "*")
	url = r.Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "www.google.com"
	url = r.Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "localhost:8080"
	rsc = "v2/" + rsc
	url = r.Resolve(host, auth, rsc, values, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	h := make(http.Header)
	url = r.Resolve(host, auth, rsc, values, h)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	h.Add(BuildPath(auth, rsc, values), testRespName)
	url = r.Resolve(host, auth, rsc, values, h)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url)

	host = "www.google.com"
	rsc = "search"
	values.Del("region")
	values.Add("q", "golang")
	auth = ""
	url = r.Resolve(host, auth, rsc, values, nil)
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

func resolverWithEnvoy() *Resolver {
	return NewResolver([]HostEntry{
		{Key: proxyKey, Host: "localhost:8081", Proxy: false},
		{Key: defaultKey, Host: "www.google.com", Proxy: false},
		{Key: yahooKey, Host: "www.search.yahoo.com", Proxy: true},
		{Key: bingKey, Host: "www.bing.com", Proxy: false},
	},
	)
}

func resolverWithoutEnvoy() *Resolver {
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
	r := resolverWithEnvoy()

	url1 := r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = "duckduckgo.com"
	url1 = r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = defaultKey
	url1 = r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Resolve("","github/advanced-go/search","access") -> [github/advanced-go/search:access]
	//test: Resolve("duckduckgo.com","github/advanced-go/search","access") -> [https://duckduckgo.com/github/advanced-go/search:access]
	//test: Resolve("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Resolve("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]
	//test: Resolve("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_Empty() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithEnvoy()

	host = "duckduckgo.com"
	url1 := r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override(nil)
	host = defaultKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Resolve("duckduckgo.com","github/advanced-go/search","access") -> [https://duckduckgo.com/github/advanced-go/search:access]
	//test: Resolve2("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Resolve2("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]
	//test: Resolve2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_No_Proxy() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithEnvoy()

	host = defaultKey
	url1 := r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override([]HostEntry{{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false}})
	host = defaultKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Resolve("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Resolve2("default","github/advanced-go/search","access") -> [https://www.duckduckgo.com/github/advanced-go/search:access]
	//test: Resolve2("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]
	//test: Resolve2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_Proxy() {
	host := ""
	auth := "github/advanced-go/search"
	rsc := "access"
	r := resolverWithEnvoy()

	host = defaultKey
	url1 := r.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override([]HostEntry{
		{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false},
		{Key: proxyKey, Host: "localhost:8888", Proxy: false},
		{Key: bingKey, Host: "www.bing.com", Proxy: true},
	})
	host = defaultKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.Resolve(host, auth, rsc, nil, nil)
	fmt.Printf("test: Resolve2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: Resolve("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: Resolve2("default","github/advanced-go/search","access") -> [https://www.duckduckgo.com/github/advanced-go/search:access]
	//test: Resolve2("bing","github/advanced-go/search","access") -> [http://localhost:8888/github/advanced-go/search:access]
	//test: Resolve2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}
