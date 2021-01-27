package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type IntConverter struct {
}

func (c IntConverter) String(a interface{}) (string, error) {
	if cp, ok := a.(*int); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*int8); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*int16); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*int32); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*int64); ok {
		return fmt.Sprint(*cp), nil
	}

	return fmt.Sprint(a), nil
}

func (c IntConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	switch value.Type().Kind() {
	case reflect.Ptr:
		switch value.Type().Elem().Kind() {
		case reflect.Int:
			i1 := int(i)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Int8:
			i1 := int8(i)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Int16:
			i1 := int16(i)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Int32:
			i1 := int32(i)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Int64:
			i1 := int64(i)
			value.Set(reflect.ValueOf(&i1))
		}
		return nil
	}

	value.SetInt(i)
	return nil
}
