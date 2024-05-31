package controller

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func createTest(r *http.Request) (*http.Response, *core.Status) {
	return nil, nil
}

func ExampleCreate() {
	cfg := Config{
		CtrlName:     "ctrl-name",
		Host:         "localhost:8081",
		Authority:    "github/advanced-go/search",
		LivenessPath: core.HealthLivenessPath,
		Duration:     time.Second * 2,
		Handler:      nil,
	}

	ctrl := Create(cfg)
	fmt.Printf("test: Create() -> [name:%v] [prime:%v] [second:%v]\n", ctrl.Name, ctrl.Router.Primary, ctrl.Router.Secondary)

	cfg.Handler = createTest
	ctrl = Create(cfg)
	fmt.Printf("test: Create() -> [name:%v] [prime:%v] [second:%v]\n", ctrl.Name, ctrl.Router.Primary, ctrl.Router.Secondary)

	//Output:
	//test: Create() -> [name:ctrl-name] [prime:&{primary localhost:8081 github/advanced-go/search health/liveness 2s <nil>}] [second:<nil>]
	//test: Create() -> [name:ctrl-name] [prime:&{primary  github/advanced-go/search health/liveness 2s 0x426c60}] [second:&{secondary localhost:8081 github/advanced-go/search health/liveness 2s <nil>}]

}
