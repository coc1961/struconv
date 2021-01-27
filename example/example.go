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

	c := struconv.New()
	c.AddScanner(struconv.Type(Person{}), PersonConverter{})

	var u User

	data := map[string]interface{}{
		"UserID":   "fgomez",
		"Password": "x3f5h7j89x997",
		"Person":   "Fernado Gomez,35",
	}

	err := c.Scan(&u, data)

	b, _ := json.MarshalIndent(u, "", "  ")

	fmt.Println(err)
	fmt.Println(string(b))
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

// Set Person to Value
func (c PersonConverter) Set(value *reflect.Value, s string) error {

	// Convert string to Person
	fields := strings.Split(s, ",")

	age, _ := strconv.Atoi(fields[1])
	person := Person{
		Name: fields[0],
		Age:  age,
	}

	// Set Person to User Field
	value.Set(reflect.ValueOf(person))
	return nil
}
