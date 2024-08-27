package core

import (
	"context"
	"fmt"
)

func ExampleUrlContext() {
	url := "https://google.com/search?q=test"
	ctx := NewUrlContext(context.Background(), url)
	fmt.Printf("test: NewUrlContext(\"%v\") -> %v\n", url, UrlFromContext(ctx))

	//ctxNew := NewRequestIdContext(ctx, "123-456-abc-xyz")
	//fmt.Printf("test: NewRequestIdContext(ctx,id) -> %v [newContext:%v]\n", RequestIdFromContext(ctx), ctxNew != ctx)

	//Output:
	//test: NewUrlContext("https://google.com/search?q=test") -> https://google.com/search?q=test

}
