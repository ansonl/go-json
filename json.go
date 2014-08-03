package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var input = `{  "array": [    1,    2,    3  ],  "boolean": true,  "null": null,  "number": 123,  "object": {    "a": "b",    "c": "d",    "e": "f"  },  "string": "Hello World"}`

//var input = `{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`

type UnknownMap map[interface{}]interface{}

type UnknownMapString map[string]interface{}

type UnknownArray []interface{}

func findType(unknown interface{}) {
	fmt.Println("Reflect says", unknown, "is a", reflect.TypeOf(unknown))
	switch vv := unknown.(type) {
	case map[interface{}]interface{}:
		fmt.Println(vv, "is of type", reflect.TypeOf(unknown))
		m := unknown.(map[interface{}]interface{})
		fmt.Println("- Type asserted to", reflect.TypeOf(m))
		mapThing := UnknownMap(m)
		fmt.Println("- Type convert/assigned to", reflect.TypeOf(mapThing))
		mapThing.FindMapContentType()
	case map[string]interface{}:
		fmt.Println(vv, "is of type", reflect.TypeOf(unknown))
		m := unknown.(map[string]interface{})
		fmt.Println("- Type asserted to", reflect.TypeOf(m))
		mapThing := UnknownMapString(m)
		fmt.Println("- Type convert/assigned to", reflect.TypeOf(mapThing))
		mapThing.FindMapStringContentType()
	case []interface{}:
		fmt.Println(vv, "is of type", reflect.TypeOf(unknown))
		m := unknown.([]interface{})
		fmt.Println("- Type asserted to", reflect.TypeOf(m))
		arrayThing := UnknownArray(m)
		fmt.Println("- Type convert/assigned to", reflect.TypeOf(arrayThing))
		arrayThing.FindArrayContentType()
	case string:
		fmt.Println(vv, "is string")
	case int:
		fmt.Println(vv, "is int")
	case float64:
		fmt.Println(vv, "is float64")
	case bool:
		fmt.Println(vv, "is bool")
	case nil:
		fmt.Println(vv, "is nil")
	default:
		fmt.Println(vv, "had NO cases matched")
	}
}

func (m UnknownMap) FindMapContentType() {
	fmt.Println("-> Parsing inside", m)
	for k, v := range m {
		fmt.Println("Reflect says KEY", k, "is a", reflect.TypeOf(k))
		fmt.Println("Reflect says VAL", v, "is a", reflect.TypeOf(v))
		fmt.Println("--> Passing", v, "to find type ")
		findType(v)
	}
}

func (m UnknownMapString) FindMapStringContentType() {
	fmt.Println("-> Parsing inside", m)
	for k, v := range m {
		fmt.Println("Reflect says KEY", k, "is a", reflect.TypeOf(k))
		fmt.Println("Reflect says VAL", v, "is a", reflect.TypeOf(v))
		fmt.Println("--> Passing", v, "to find type ")
		findType(v)
	}
}

func (array UnknownArray) FindArrayContentType() {
	for i, v := range array {
		fmt.Println("Reflect says VAL", v, " at index ", i, "is a", reflect.TypeOf(v))
		fmt.Println("--> Passing", v, "to find type ")
		findType(v)
	}
}

func main() {
	var val interface{}

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)

	//val = val.(UnknownMapString)

	findType(val)

	/*
		var f interface{}
		f = map[string]interface{}{
			"Name": "Wednesday",
			"Age":  6,
			"Parents": []interface{}{
				"Gomez",
				"Morticia",
			},
		}
		m := f.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	*/
}
