package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type ComplexConverter struct {
	Pointer bool
	Bytes   int
}

func (c ComplexConverter) String(a interface{}) (string, error) {
	if c.Pointer {
		if cp, ok := a.(*complex64); ok {
			return fmt.Sprint(*cp), nil
		}
		if cp, ok := a.(*complex128); ok {
			return fmt.Sprint(*cp), nil
		}
	}
	return fmt.Sprint(a), nil
}

func (c ComplexConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return err
	}
	if c.Pointer {
		if c.Bytes == 64 {
			i1 := complex64(i)
			value.Set(reflect.ValueOf(&i1))
		} else {
			value.Set(reflect.ValueOf(&i))
		}
	} else {
		value.SetComplex(i)
	}
	return nil
}
