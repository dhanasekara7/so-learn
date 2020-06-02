package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestObject struct {
	Kind  string `json:"kind"`
	Id    string `json:"id, omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestCreateSingleItemResponse(t *testing.T) {
	testObject := new(TestObject)
	testObject.Kind = "TestObject"
	testObject.Id = "f73h5jf8"
	testObject.Name = "Yuri Gagarin"
	testObject.Email = "Yuri.Gagarin@Vostok.com"

	fmt.Println(testObject)

	b, err := json.Marshal(testObject)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b[:]))
}

func main() {

}
