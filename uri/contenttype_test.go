package uri

import "fmt"

func Example_IsStatusURL() {
	u := ""
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/io2test/resource/activity.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/io2test/resource/status/activity.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/io2test/resource/status-504.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/io2test/resource/status/status-504.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	//Output:
	//test: IsStatusURL("") -> false
	//test: IsStatusURL("file://[cwd]/io2test/resource/activity.json") -> false
	//test: IsStatusURL("file://[cwd]/io2test/resource/status/activity.json") -> false
	//test: IsStatusURL("file://[cwd]/io2test/resource/status-504.json") -> true
	//test: IsStatusURL("file://[cwd]/io2test/resource/status/status-504.json") -> true

}
