package uri

import "fmt"

func Example_IsStatusURL() {
	u := ""
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/iotest/resource/activity.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/iotest/resource/status/activity.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/iotest/resource/status-504.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	u = "file://[cwd]/iotest/resource/status/status-504.json"
	fmt.Printf("test: IsStatusURL(\"%v\") -> %v\n", u, IsStatusURL(u))

	//Output:
	//test: IsStatusURL("") -> false
	//test: IsStatusURL("file://[cwd]/iotest/resource/activity.json") -> false
	//test: IsStatusURL("file://[cwd]/iotest/resource/status/activity.json") -> false
	//test: IsStatusURL("file://[cwd]/iotest/resource/status-504.json") -> true
	//test: IsStatusURL("file://[cwd]/iotest/resource/status/status-504.json") -> true

}
