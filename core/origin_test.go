package core

import "fmt"

func ExampleNewValues() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	values := NewValues(o)
	fmt.Printf("test: NewValues() -> [%v]\n", values)

	//Output:
	//test: NewValues() -> [map[az:[zone] host:[host] reg:[region] sz:[sub-zone]]]

}

func ExampleNewOrigin() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	values := NewValues(o)
	o = NewOrigin(values)
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	//Output:
	//test: NewOrigin() -> [{region zone sub-zone host }]

}

func ExampleOriginMatch() {
	target := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	filter := Origin{
		Region:     "",
		Zone:       "",
		SubZone:    "",
		Host:       "",
		InstanceId: "",
	}

	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.Region = "region"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.Zone = "zone"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.SubZone = "sub-zone"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.Host = "host"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.SubZone = ""
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	filter.SubZone = "invalid"
	fmt.Printf("test: OriginMatch(%v,%v) -> [match:%v]\n", target, filter, OriginMatch(target, filter))

	//Output:
	//test: OriginMatch({region zone sub-zone host },{    }) -> [match:true]
	//test: OriginMatch({region zone sub-zone host },{region    }) -> [match:true]
	//test: OriginMatch({region zone sub-zone host },{region zone   }) -> [match:true]
	//test: OriginMatch({region zone sub-zone host },{region zone sub-zone  }) -> [match:true]
	//test: OriginMatch({region zone sub-zone host },{region zone sub-zone host }) -> [match:true]
	//test: OriginMatch({region zone sub-zone host },{region zone  host }) -> [match:true]
	//test: OriginMatch({region zone sub-zone host },{region zone invalid host }) -> [match:false]

}
