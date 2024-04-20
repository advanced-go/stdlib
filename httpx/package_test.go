package httpx

import (
	"fmt"
	"reflect"
	"strings"
)

func Example_PkgPath() {
	dotCom := ".com"
	pkgPath := reflect.TypeOf(any(pkg{})).PkgPath()
	fmt.Printf("test: PkgPath = \"%v\"\n", strings.Replace(pkgPath, dotCom, "", len(dotCom)))

	//Output:
	//test: PkgPath = "github/advanced-go/stdlib/httpx"

}
