package core

import (
	"context"
	"fmt"
)

/*
func ExampleUrlContext() {
	url := "https://google.com/search?q=test"
	ctx := NewUrlContext(context.Background(), url)
	fmt.Printf("test: NewUrlContext(\"%v\") -> %v\n", url, UrlFromContext(ctx))

	//Output:
	//test: NewUrlContext("https://google.com/search?q=test") -> https://google.com/search?q=test

}


*/

func ExampleExchangeContext() {
	url := "https://google.com/search?q=test"
	urls := NewExchangeOverride(url, "response-url", "status-url")
	ctx := NewExchangeOverrideContext(context.Background(), urls)
	fmt.Printf("test: NewExchangeContext(%v) -> %v\n", urls, ExchangeOverrideFromContext(ctx))

	//Output:
	//test: NewExchangeContext(&{map[request:https://google.com/search?q=test response:response-url status:status-url]}) -> &{map[request:https://google.com/search?q=test response:response-url status:status-url]}

}
