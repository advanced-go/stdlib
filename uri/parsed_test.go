package uri

import (
	"fmt"
	"net/url"
)

func ExampleURLParse_Raw() {
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

func ExampleParsed_Version() {
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

func ExampleParsed_PathURL() {
	url := "https://www.google.com/github/advanced-go/search:google"
	p := Uproot(url)
	u := p.PathURL()
	fmt.Printf("test: Parsed(\"%v\") -> [pathURL:%v] [query:%v]\n", url, u, u.Query().Encode())

	url = "https://www.google.com/github/advanced-go/search:v2/google"
	p = Uproot(url)
	u = p.PathURL()
	fmt.Printf("test: Parsed(\"%v\") -> [pathURL:%v] [query:%v]\n", url, u, u.Query().Encode())

	url = "https://www.google.com/github/advanced-go/search:v2/google?q=golang"
	p = Uproot(url)
	u = p.PathURL()
	fmt.Printf("test: Parsed(\"%v\") -> [pathURL:%v] [query:%v]\n", url, u, u.Query().Encode())

	//Output:
	//test: Parsed("https://www.google.com/github/advanced-go/search:google") -> [pathURL:google] [query:]
	//test: Parsed("https://www.google.com/github/advanced-go/search:v2/google") -> [pathURL:google] [query:]
	//test: Parsed("https://www.google.com/github/advanced-go/search:v2/google?q=golang") -> [pathURL:google?q=golang] [query:q=golang]

}
