package core

import (
	"context"
	"fmt"
)

func ExampleUrlContext() {
	url := "https://google.com/search?q=test"
	ctx := NewUrlContext(context.Background(), url)
	fmt.Printf("test: NewUrlContext(\"%v\") -> %v\n", url, UrlFromContext(ctx))

	//Output:
	//test: NewUrlContext("https://google.com/search?q=test") -> https://google.com/search?q=test

}

func ExampleUrlMapContext() {
	url := "https://google.com/search?q=test"
	urls := make(map[string]string)
	urls[ContextRequestKey] = url
	urls[ContextResponseKey] = "CommandTag"
	ctx := NewUrlMapContext(context.Background(), urls)
	fmt.Printf("test: NewUrlMapContext(%v) -> %v\n", urls, UrlMapFromContext(ctx))

	//Output:
	//test: NewUrlMapContext(map[request:https://google.com/search?q=test response:CommandTag]) -> map[request:https://google.com/search?q=test response:CommandTag]

}
