package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `mytag:"MyName"`
	Email string `mytag:"MyEmail"`
}

func main() {

	u := User{Name: "Bob", Email: "bob@mycompany.com"}
	t := reflect.TypeOf(u)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}
}
