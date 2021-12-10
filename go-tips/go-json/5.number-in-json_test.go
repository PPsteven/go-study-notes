package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

// handle numbers in json: json.Decoder and json.Decoder.UseNumber

func TestNumberInJson(t *testing.T) {
	s := `{"float":3.12,"int":3.0}`
	var res interface{}
	_ = json.Unmarshal([]byte(s), &res)
	t.Logf("%v", res)

	// map[float:3.12 int:3]
	// 3.0 was transfer to 3

	d := json.NewDecoder(bytes.NewReader([]byte(s)))
	d.UseNumber()
	_ = d.Decode(&res)
	t.Logf("%v", res)
	// map[float:3.12 int:3.0]
}