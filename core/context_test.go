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
	urls := NewExchangeMap(url, "response-url", "status-url")
	ctx := NewExchangeMapContext(context.Background(), urls)
	fmt.Printf("test: NewExchangeMapContext(%v) -> %v\n", urls, ExchangeMapFromContext(ctx))

	//Output:
	//test: NewExchangeMapContext(&{map[request:https://google.com/search?q=test response:response-url status:status-url]}) -> &{map[request:https://google.com/search?q=test response:response-url status:status-url]}

}
