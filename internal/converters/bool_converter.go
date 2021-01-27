package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type BoolConverter struct {
}

func (c BoolConverter) String(a interface{}) (string, error) {
	if cp, ok := a.(*bool); ok {
		return fmt.Sprint(*cp), nil
	}
	return fmt.Sprint(a), nil
}

func (c BoolConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	switch value.Type().Kind() {
	case reflect.Ptr:
		switch value.Type().Elem().Kind() {
		case reflect.Bool:
			value.Set(reflect.ValueOf(&i))
		}
		return nil
	}

	value.SetBool(i)
	return nil
}
