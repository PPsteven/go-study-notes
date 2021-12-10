package main

import (
	"encoding/json"
	gsimplejson "github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
)

// 1. decode to struct when structure is fixed
// 2. get field value when structure is unfiexd by map
// 3. package: go-simplejson
// 4. package: gjson

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

// use gjson "github.com/tidwall/gjson"
func TestReadHeightByGjson(t *testing.T) {
	v := gjson.Get(birdString, "0").Get("body.height").Float()
	t.Logf("height: %0.1f", v)
}

// use go-simple-json "github.com/bitly/go-simplejson"
func TestArr(t *testing.T) {
	person := `
{
    "UID":"183201273",
    "Lily":{
        "xt-high-school":{
            "start":"2010",
            "end":"2013"
        },
        "TJ-university":{
            "start":"2013",
            "end":"2017"
        }
    }
}
`
	js, _ := gsimplejson.NewJson([]byte(person))

	// In this Json, person name and education experience is uncertain, so it's hard to decode this json to stuctured data
	for _, name := range []string{"Lily", "Jack"} {
		p := js.GetPath(name).Interface()
		if p != nil {
			t.Logf("%s is found, detail info is %v", name, p)
		} else {
			t.Logf("%s is not found", name)
		}
	}

	// Lily is found, detail info is map[TJ-university:map[end:2017 start:2013] xt-high-school:map[end:2013 start:2010]]
	// Jack is not found

	lilyStartYear, _ := js.GetPath("Lily", "TJ-university", "start").String()
	assert.Equal(t, "2013", lilyStartYear)
}
