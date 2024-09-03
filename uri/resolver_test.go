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
	rsc := "access"
	ver := "v2"
	values := make(url.Values)

	p := BuildPath(rsc, values)
	fmt.Printf("test: BuildPath(\"%v\") -> [%v]\n", rsc, p)

	p = BuildPathWithAuthority(auth, "", rsc, values)
	fmt.Printf("test: BuildPathWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", auth, ver, rsc, p)

	values.Add("region", "*")
	p = BuildPath(rsc, values)
	fmt.Printf("test: BuildPath(\"%v\") -> [%v]\n", rsc, p)

	p = BuildPathWithAuthority(auth, ver, rsc, values)
	fmt.Printf("test: BuildPathWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", auth, ver, rsc, p)

	//Output:
	//test: BuildPath("access") -> [access]
	//test: BuildPathWithAuthority("github/advanced-go/timeseries","v2","access") -> [github/advanced-go/timeseries:access]
	//test: BuildPath("access") -> [access?region=*]
	//test: BuildPathWithAuthority("github/advanced-go/timeseries","v2","access") -> [github/advanced-go/timeseries:v2/access?region=*]

}

func ExampleResolve_Url() {
	errType := 123
	host := ""
	path := "/search"
	values := make(url.Values)
	r := NewResolver(nil)

	url1 := r.Url(host, path, errType, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	url1 = r.Url(host, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	values.Add("q", "golang")
	url1 = r.Url(host, path, values, nil)
	fmt.Printf("test: Url(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	url1 = r.Url(host, path, "q=golang", nil)
	fmt.Printf("test: Url_String(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	host = "www.google.com"
	url1 = r.Url(host, path, values, nil)
	fmt.Printf("test: Url_String(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	h := make(http.Header)
	h.Add(BuildPath(path, values), "https://www.search.yahoo.com?q=golang")
	host = "www.google.com"
	url1 = r.Url(host, path, values, h)
	fmt.Printf("test: Url_Override(\"%v\",\"%v\") -> [%v]\n", host, path, url1)

	//Output:
	//test: Url("","/search") -> [/searcherror: query type is invalid int]
	//test: Url("","/search") -> [/search]
	//test: Url("","/search") -> [/search?q=golang]
	//test: Url_String("","/search") -> [/search?q=golang]
	//test: Url_String("www.google.com","/search") -> [https://www.google.com/search?q=golang]
	//test: Url_Override("www.google.com","/search") -> [https://www.search.yahoo.com?q=golang]

}

func ExampleResolve_UrlWithAuthority() {
	host := ""
	auth := "github/advanced-go/timeseries"
	rsc := "access"
	ver := ""
	values := make(url.Values)
	r := NewResolver(nil)

	url1 := r.UrlWithAuthority(host, auth, "", rsc, values, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	values.Add("region", "*")
	url1 = r.UrlWithAuthority(host, auth, "", rsc, values, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	url1 = r.UrlWithAuthority(host, auth, "", rsc, "region=*", nil)
	fmt.Printf("test: UrlWithAuthority_String(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	host = "www.google.com"
	url1 = r.UrlWithAuthority(host, auth, "", rsc, values, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	host = "localhost:8080"
	ver = "v2"
	//rsc = "v2/" + rsc
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, values, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	h := make(http.Header)
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, values, h)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	h.Add(BuildPathWithAuthority(auth, ver, rsc, values), testRespName)
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, values, h)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	host = "www.google.com"
	rsc = "search"
	values.Del("region")
	values.Add("q", "golang")
	auth = ""
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, values, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, ver, rsc, url1)

	//Output:
	//test: UrlWithAuthority("","github/advanced-go/timeseries","","access") -> [github/advanced-go/timeseries:access]
	//test: UrlWithAuthority("","github/advanced-go/timeseries","","access") -> [github/advanced-go/timeseries:access?region=*]
	//test: UrlWithAuthority_String("","github/advanced-go/timeseries","","access") -> [github/advanced-go/timeseries:access?region=*]
	//test: UrlWithAuthority("www.google.com","github/advanced-go/timeseries","","access") -> [https://www.google.com/github/advanced-go/timeseries:access?region=*]
	//test: UrlWithAuthority("localhost:8080","github/advanced-go/timeseries","v2","access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=*]
	//test: UrlWithAuthority("localhost:8080","github/advanced-go/timeseries","v2","access") -> [http://localhost:8080/github/advanced-go/timeseries:v2/access?region=*]
	//test: UrlWithAuthority("localhost:8080","github/advanced-go/timeseries","v2","access") -> [file://[cwd]/timeseries1test/get-all-resp-v1.txt]
	//test: UrlWithAuthority("www.google.com","","v2","search") -> [https://www.google.com/search?q=golang]

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

func ExampleResolver_UrlWithAuthority() {
	host := ""
	auth := "github/advanced-go/search"
	ver := ""
	rsc := "access"
	r := resolverWithProxy()

	url1 := r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = "duckduckgo.com"
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = defaultKey
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: UrlWithAuthority("","github/advanced-go/search","access") -> [github/advanced-go/search:access]
	//test: UrlWithAuthority("duckduckgo.com","github/advanced-go/search","access") -> [https://duckduckgo.com/github/advanced-go/search:access]
	//test: UrlWithAuthority("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: UrlWithAuthority("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]
	//test: UrlWithAuthority("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_Empty() {
	host := ""
	auth := "github/advanced-go/search"
	ver := ""
	rsc := "access"
	r := resolverWithProxy()

	host = "duckduckgo.com"
	url1 := r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override(nil)
	host = defaultKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: UrlWithAuthority("duckduckgo.com","github/advanced-go/search","access") -> [https://duckduckgo.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_No_Proxy() {
	host := ""
	auth := "github/advanced-go/search"
	ver := ""
	rsc := "access"
	r := resolverWithProxy()

	host = defaultKey
	url1 := r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override([]HostEntry{{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false}})
	host = defaultKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: UrlWithAuthority("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("default","github/advanced-go/search","access") -> [https://www.duckduckgo.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("bing","github/advanced-go/search","access") -> [https://www.bing.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

}

func ExampleResolver_Overrides_Proxy() {
	host := ""
	auth := "github/advanced-go/search"
	ver := ""
	rsc := "access"
	r := resolverWithProxy()

	host = defaultKey
	url1 := r.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	r2 := r.Override([]HostEntry{
		{Key: defaultKey, Host: "www.duckduckgo.com", Proxy: false},
		{Key: proxyKey, Host: "localhost:8888", Proxy: false},
		{Key: bingKey, Host: "www.bing.com", Proxy: true},
	})
	host = defaultKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = bingKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	host = yahooKey
	url1 = r2.UrlWithAuthority(host, auth, ver, rsc, nil, nil)
	fmt.Printf("test: UrlWithAuthority2(\"%v\",\"%v\",\"%v\") -> [%v]\n", host, auth, rsc, url1)

	//Output:
	//test: UrlWithAuthority("default","github/advanced-go/search","access") -> [https://www.google.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("default","github/advanced-go/search","access") -> [https://www.duckduckgo.com/github/advanced-go/search:access]
	//test: UrlWithAuthority2("bing","github/advanced-go/search","access") -> [http://localhost:8888/github/advanced-go/search:access]
	//test: UrlWithAuthority2("yahoo","github/advanced-go/search","access") -> [http://localhost:8081/github/advanced-go/search:access]

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
