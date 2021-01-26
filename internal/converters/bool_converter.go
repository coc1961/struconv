package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type BoolConverter struct {
}

func (c BoolConverter) String(a interface{}) (string, error) {
	return fmt.Sprint(a), nil
}

func (c BoolConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	value.SetBool(i)
	return nil
}
