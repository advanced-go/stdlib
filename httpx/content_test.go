package httpx

import (
	"fmt"
	"io"
	"net/url"
)

type Activity struct {
	ActivityID   string `json:"ActivityID"`
	ActivityType string `json:"ActivityType"`
	Agent        string `json:"Agent"`
	AgentUri     string `json:"AgentUri"`
	Assignment   string `json:"Assignment"`
	Controller   string `json:"Controller"`
	Behavior     string `json:"Behavior"`
	Description  string `json:"Description"`
}

func ExampleContent_Error() {
	s := "file://[cwd]/httpxtest/resource/activity-error-response.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	a, status := Content[[]Activity](resp.Body)
	fmt.Printf("test: Content -> %v [status:%v]\n", a, status)
	if status.Err != nil && status.Err.Error() == "EOF" {
		buf, status1 := io.ReadAll(resp.Body)
		fmt.Printf("test: Content -> [buf:%v] [status:%v]\n", buf, status1)

	}

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/resource/activity-error-response.txt) -> [error:[<nil>]] [statusCode:200]
	//test: Content -> [] [status:Json Decode Failure [unexpected EOF]]

}

func ExampleContent_Empty() {
	s := "file://[cwd]/httpxtest/resource/activity-empty-response.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	a, status := Content[[]Activity](resp.Body)
	fmt.Printf("test: Content -> %v [status:%v]\n", a, status)
	if status.Err != nil && status.Err.Error() == "EOF" {
		buf, status1 := io.ReadAll(resp.Body)
		fmt.Printf("test: Content -> [buf:%v] [status:%v]\n", buf, status1)

	}

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/resource/activity-empty-response.txt) -> [error:[<nil>]] [statusCode:200]
	//test: Content -> [] [status:Not Found]

}
