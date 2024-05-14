package uri

import (
	"fmt"
	"net/url"
)

func ExampleParseRaw() {
	u := "http://localhost:8080/github/advanced-go/stdlib/uri.Uproot?q=golang"
	uri, err := url.Parse(u)
	fmt.Printf("test: ParseRaw(\"%v\") -> [scheme:%v] [host:%v] [path:%v] [frag:%v] [query:%v] [err:%v]\n", u, uri.Scheme, uri.Host, uri.Path, uri.Fragment, uri.RawQuery, err)

	u = "http://localhost:8080/github/advanced-go/stdlib/uri:Uproot?q=golang"
	uri, err = url.Parse(u)
	fmt.Printf("test: ParseRaw(\"%v\") -> [scheme:%v] [path:%v] [frag:%v] [query:%v] [err:%v]\n", u, uri.Scheme, uri.Path, uri.Fragment, uri.RawQuery, err)

	u = "http://localhost:8080/github/advanced-go/stdlib/uri?q=golang#Uproot"
	uri, err = url.Parse(u)
	fmt.Printf("test: ParseRaw(\"%v\") -> [scheme:%v] [path:%v] [frag:%v] [query:%v] [err:%v]\n", u, uri.Scheme, uri.Path, uri.Fragment, uri.RawQuery, err)

	u = uri.Path
	uri, err = url.Parse(u)
	fmt.Printf("test: ParseRaw(\"%v\") -> [scheme:%v] [path:%v] [frag:%v] [query:%v] [err:%v]\n", u, uri.Scheme, uri.Path, uri.Fragment, uri.RawQuery, err)

	//Output:
	//test: ParseRaw("http://localhost:8080/github/advanced-go/stdlib/uri.Uproot?q=golang") -> [scheme:http] [host:localhost:8080] [path:/github/advanced-go/stdlib/uri.Uproot] [frag:] [query:q=golang] [err:<nil>]
	//test: ParseRaw("http://localhost:8080/github/advanced-go/stdlib/uri:Uproot?q=golang") -> [scheme:http] [path:/github/advanced-go/stdlib/uri:Uproot] [frag:] [query:q=golang] [err:<nil>]
	//test: ParseRaw("http://localhost:8080/github/advanced-go/stdlib/uri?q=golang#Uproot") -> [scheme:http] [path:/github/advanced-go/stdlib/uri] [frag:Uproot] [query:q=golang] [err:<nil>]
	//test: ParseRaw("/github/advanced-go/stdlib/uri") -> [scheme:] [path:/github/advanced-go/stdlib/uri] [frag:] [query:] [err:<nil>]

}

func ExampleParseVersion() {
	p := new(Parsed)

	prev := p.Path
	parseVersion(p)
	fmt.Printf("test: parseVersion(\"%v\") -> [path:%v] [vers:%v]\n", prev, p.Path, p.Version)

	p.Path = "search"
	prev = p.Path
	parseVersion(p)
	fmt.Printf("test: parseVersion(\"%v\") -> [path:%v] [vers:%v]\n", prev, p.Path, p.Version)

	p.Path = "v1/search"
	prev = p.Path
	parseVersion(p)
	fmt.Printf("test: parseVersion(\"%v\") -> [path:%v] [vers:%v]\n", prev, p.Path, p.Version)

	//Output:
	//test: parseVersion("") -> [path:] [vers:]
	//test: parseVersion("search") -> [path:search] [vers:]
	//test: parseVersion("v1/search") -> [path:search] [vers:v1]

}

func ExampleUproot_Validate() {
	// Empty
	path := ""
	p := Uproot(path)
	fmt.Printf("test: Uproot-Empty(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	// Urn should not be changed
	path = "urn:github.resource"
	p = Uproot(path)
	fmt.Printf("test: Uproot-URN(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	// No URN separator, valid with authority only
	path = "http://localhost:8080/github/advanced-go/search/query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot-Authority-Only(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	// 1 URN separator
	path = "http://localhost:8080/github/advanced-go/search:query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot-Authority+Path(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	// More than 1 URN separator
	path = "http://localhost:8080/github/advanced-go/:search:/query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot->1 URN(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	//Output:
	//test: Uproot-Empty() -> [ok:false] [auth:] [vers:] [path:] [err:error: invalid input, URI is empty]
	//test: Uproot-URN(urn:github.resource) -> [ok:true] [auth:urn:github.resource] [vers:] [path:urn:github.resource] [err:<nil>]
	//test: Uproot-Authority-Only(http://localhost:8080/github/advanced-go/search/query?term=golang) -> [ok:true] [auth:github/advanced-go/search/query] [vers:] [path:] [err:<nil>]
	//test: Uproot-Authority+Path(http://localhost:8080/github/advanced-go/search:query?term=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:] [path:query] [err:<nil>]
	//test: Uproot->1 URN(http://localhost:8080/github/advanced-go/:search:/query?term=golang) -> [ok:false] [auth:] [vers:] [path:] [err:error: path has multiple URN separators [/github/advanced-go/:search:/query]]

}

func ExampleUproot() {
	path := "/github/advanced-go/search:query?term=golang"
	p := Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	path = "/github/advanced-go/search:v1/query"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	path = "http://localhost:8080/github/advanced-go/search:query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	path = "http://localhost:8080/github/advanced-go/search:v1/query"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Authority, p.Version, p.Path, p.Err)

	//Output:
	//test: Uproot(/github/advanced-go/search:query?term=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:] [path:query] [err:<nil>]
	//test: Uproot(/github/advanced-go/search:v1/query) -> [ok:true] [auth:github/advanced-go/search] [vers:v1] [path:query] [err:<nil>]
	//test: Uproot(http://localhost:8080/github/advanced-go/search:query?term=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:] [path:query] [err:<nil>]
	//test: Uproot(http://localhost:8080/github/advanced-go/search:v1/query) -> [ok:true] [auth:github/advanced-go/search] [vers:v1] [path:query] [err:<nil>]

}

/*

	// valid path only and an empty nss
	uri = "/valid-empty-nss?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	// valid embedded path only
	uri = "/github/valid-leading-slash/example-domain/activity:entry"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	// valid URN
	uri = "github.com/valid-no-leading-slash/example-domain/activity:entry"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	uri = "https://www.google.com/valid-uri?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	uri = "https://www.google.com/github.com/valid-uri-nss/search?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	uri = "https://www.google.com/github.com/valid-uri-with-nss:search?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)


*/
