package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type ComplexConverter struct {
}

func (c ComplexConverter) String(a interface{}) (string, error) {
	if cp, ok := a.(*complex64); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*complex128); ok {
		return fmt.Sprint(*cp), nil
	}
	return fmt.Sprint(a), nil
}

func (c ComplexConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return err
	}
	switch value.Type().Kind() {
	case reflect.Ptr:
		switch value.Type().Elem().Kind() {
		case reflect.Complex64:
			i1 := complex64(i)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Complex128:
			value.Set(reflect.ValueOf(&i))
		}
		return nil
	}

	value.SetComplex(i)
	return nil
}
