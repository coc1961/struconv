package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/coc1961/struconv"
)

func main() {
	// Create Scanner, that supports basic data conversion
	c := struconv.New()

	// Add custom Converter for unknown Person Struct
	c.AddScanner(struconv.Type(Person{}), PersonConverter{})

	var u User

	// Map of data, Key = Struct Fields Name, Field data type or string
	data := map[string]interface{}{
		"UserID":   "fgomez",
		"Password": "x3f5h7j89x997",
		"Person":   "Fernado Gomez,35",
	}

	// Fill struct User
	err := c.Scan(&u, data)

	b, _ := json.MarshalIndent(u, "", "  ")

	fmt.Println(err)
	fmt.Println(string(b))

	// Print
	//  <nil>
	//	{
	//		"UserID": "fgomez",
	//		"Password": "x3f5h7j89x997",
	//		"Person": {
	//			"Name": "Fernado Gomez",
	//			"Age": 35
	//		}
	//	}

}

type User struct {
	UserID   string
	Password string
	Person   Person
}

type Person struct {
	Name string
	Age  int
}

// Custom Converter
type PersonConverter struct{}

// Convert Person to String
func (c PersonConverter) String(a interface{}) (string, error) {
	p := a.(Person)
	return fmt.Sprintf("%v,%v", p.Name, p.Age), nil
}

// Convert and Set Value to User Field
func (c PersonConverter) Set(value *reflect.Value, s string) error {

	// Convert string to Person
	fields := strings.Split(s, ",")

	age, _ := strconv.Atoi(fields[1])
	person := Person{
		Name: fields[0],
		Age:  age,
	}

	// Set Person to User Field
	value.Set(struconv.Value(person))
	return nil
}
