package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

//1. struct to string
//2. string to string
//3. load json from file

type User struct {
	FirstName string `json:"first_name"` // key will be "first_name"
	BirthYear int `json:"birth_year"` // key will be "birth_year"
	Email string // key will be "Email"
}

var UserString = `{"first_name":"ppsteven","birth_year":2021,"Email":"ppsteven@outlook.com"}`

func TestStrtctToString(t *testing.T) {
	user := &User{"ppsteven", 2021, "ppsteven@outlook.com"}
	userByte, _ := json.Marshal(user)
	t.Logf("%s", string(userByte))
	// {"first_name":"ppsteven","birth_year":2021,"Email":"ppsteven@outlook.com"}
	assert.Equal(t, UserString, string(userByte))
}

func TestStringToStruct(t *testing.T) {
	var user User
	_ = json.Unmarshal([]byte(UserString), &user)
	t.Logf("%v", user)
}

func TestReadJsonFile(t *testing.T) {
	tempFile, _ := ioutil.TempFile("/tmp", "temp-*.json")
	defer os.Remove(tempFile.Name())
	_, _ = tempFile.WriteString(UserString)

	var user User
	userByte, _ := ioutil.ReadFile(tempFile.Name())
	_ = json.Unmarshal(userByte, &user)
	t.Logf("%v", user)

}