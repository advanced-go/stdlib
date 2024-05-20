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
