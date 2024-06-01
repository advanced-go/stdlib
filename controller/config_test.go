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

func ExampleNew() {
	cfg := Config{
		CtrlName:     "ctrl-name",
		Host:         "localhost:8081",
		Authority:    "github/advanced-go/search",
		LivenessPath: core.HealthLivenessPath,
		Duration:     time.Second * 2,
	}

	ctrl := New(cfg, nil)
	fmt.Printf("test: New() -> [name:%v] [prime:%v] [second:%v]\n", ctrl.Name, ctrl.Router.Primary, ctrl.Router.Secondary)

	ctrl = New(cfg, createTest)
	fmt.Printf("test: New() -> [name:%v] [prime:%v] [second:%v]\n", ctrl.Name, ctrl.Router.Primary, ctrl.Router.Secondary)

	//Output:
	//test: New() -> [name:ctrl-name] [prime:&{primary localhost:8081 github/advanced-go/search health/liveness 2s <nil>}] [second:<nil>]
	//test: New() -> [name:ctrl-name] [prime:&{primary  github/advanced-go/search health/liveness 2s 0xec6c80}] [second:&{secondary localhost:8081 github/advanced-go/search health/liveness 2s <nil>}]

}

func ExampleGetConfig() {
	list := []Config{
		{
			CtrlName:     "ctrl-1",
			Host:         "localhost:8081",
			Authority:    "github/advanced-go/search",
			LivenessPath: core.HealthLivenessPath,
			Duration:     time.Second * 2,
		},
		{
			CtrlName:     "ctrl-2",
			Host:         "localhost:8081",
			Authority:    "github/advanced-go/search",
			LivenessPath: core.HealthLivenessPath,
			Duration:     time.Second * 2,
		},
	}
	name := "invalid"
	cfg, ok := GetConfig(name, list)
	fmt.Printf("test: GetConfig(\"%v\") -> [ok:%v] [cfg:%v]\n", name, ok, cfg)

	name = "ctrl-2"
	cfg, ok = GetConfig(name, list)
	fmt.Printf("test: GetConfig(\"%v\") -> [ok:%v] [cfg:%v]\n", name, ok, cfg)

	//Output:
	//test: GetConfig("invalid") -> [ok:false] [cfg:{    0s}]
	//test: GetConfig("ctrl-2") -> [ok:true] [cfg:{ctrl-2 localhost:8081 github/advanced-go/search health/liveness 2s}]
	
}
