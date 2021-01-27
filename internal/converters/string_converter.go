package converters

import (
	"fmt"
	"reflect"
)

type StringConverter struct {
}

func (c StringConverter) String(a interface{}) (string, error) {
	if cp, ok := a.(*string); ok {
		return *cp, nil
	}

	return fmt.Sprint(a), nil
}

func (c StringConverter) Set(value *reflect.Value, s string) error {
	switch value.Type().Kind() {
	case reflect.Ptr:
		switch value.Type().Elem().Kind() {
		case reflect.String:
			value.Set(reflect.ValueOf(&s))
		}
		return nil
	}
	value.SetString(s)
	return nil
}
