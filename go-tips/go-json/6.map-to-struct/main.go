package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

type BodyInfo struct {
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

type Bird struct {
	Body BodyInfo `json:"body"`
	Name string   `json:"name"`
}

// 1. map to struct: mapstructure
// 2. struct to map: reflect

func mapToStruct(inf interface{}) (bird *Bird) {
	_ = mapstructure.Decode(inf, &bird)
	return bird
}

func structToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func main() {
	birdMap := map[string]interface{}{
		"name": "Lily",
		"body": map[string]float64{
			"height": 132.09,
			"weight": 132.9,
		},
	}

	bird := mapToStruct(birdMap)
	fmt.Printf("struct: %+v\n", bird)

	bird = &Bird{Body: BodyInfo{132.8,113.2}, Name: "Lily"}
	birdMap = structToMap(*bird)
	fmt.Printf("map: %+v\n", birdMap)
}