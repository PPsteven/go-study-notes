package main

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"testing"
)

var birdString = `[{"body": {"height": 163.3,"weight": 32},"name": "Lily"}]`

type BodyInfo struct {
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}
type Bird struct {
	Body BodyInfo `json:"body"`
	Name string `json:"name"`
}

// stuctured
func TestReadHeightByStructuredData(t *testing.T) {
	var bird []Bird
	_ = json.Unmarshal([]byte(birdString), &bird)
	t.Logf("height: %0.1f", bird[0].Body.Height)
}

// unstructured data like maps
func TestReadHeight(t *testing.T) {
	var people []interface{}
	_ = json.Unmarshal([]byte(birdString), &people)
	if len(people) <= 0 { return }

	person := people[0] // interface
	for k, v := range person.(map[string]interface{}) {
		if k == "body" {
			for kk, vv := range v.(map[string]interface{}) {
				if kk == "height" {
					t.Logf("height: %0.1f", vv.(float64))
				}
			}
		}
	}
}

// use gjson
func TestReadHeightByGjson(t *testing.T) {
	v := gjson.Get(birdString, "0").Get("body.height").Float()
	t.Logf("height: %0.1f", v)
}