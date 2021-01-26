package converters

import (
	"fmt"
	"reflect"
)

type StringConverter struct {
}

func (c StringConverter) String(a interface{}) (string, error) {
	return fmt.Sprint(a), nil
}

func (c StringConverter) Set(value *reflect.Value, s string) error {
	value.SetString(s)
	return nil
}
