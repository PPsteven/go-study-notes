package main

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"testing"
)

func TestMap2Struct(t *testing.T) {
	var bird Bird
	birdMap := map[string]interface{}{
		"name": "Lily",
		"body": map[string]float64{
			"height": 132.09,
			"weight": 132.9,
		},
	}

	_ = mapstructure.Decode(birdMap, &bird)
	t.Logf("%+v", bird)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func TestStruct2Map(t *testing.T) {
	bird := &Bird{Body: BodyInfo{132.8,113.2}, Name: "Lily"}
	birdMap := Struct2Map(*bird)
	t.Logf("%v", birdMap)
}