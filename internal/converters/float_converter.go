package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type FloatConverter struct {
}

func (c FloatConverter) String(a interface{}) (string, error) {
	if cp, ok := a.(*float32); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*float64); ok {
		return fmt.Sprint(*cp), nil
	}

	return fmt.Sprint(a), nil
}

func (c FloatConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	switch value.Type().Kind() {
	case reflect.Ptr:
		switch value.Type().Elem().Kind() {
		case reflect.Float32:
			i1 := float32(i)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Float64:
			i1 := float64(i)
			value.Set(reflect.ValueOf(&i1))
		}
		return nil
	}
	value.SetFloat(i)
	return nil
}
