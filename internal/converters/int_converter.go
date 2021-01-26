package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type IntConverter struct {
}

func (c IntConverter) String(a interface{}) (string, error) {
	return fmt.Sprint(a), nil
}

func (c IntConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	value.SetInt(i)
	return nil
}
