package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type UintConverter struct {
}

func (c UintConverter) String(a interface{}) (string, error) {
	if cp, ok := a.(*uint); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*uint8); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*uint16); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*uint32); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*uint64); ok {
		return fmt.Sprint(*cp), nil
	}
	if cp, ok := a.(*uintptr); ok {
		return fmt.Sprint(*cp), nil
	}
	return fmt.Sprint(a), nil
}

func (c UintConverter) Set(value *reflect.Value, s string) error {
	f, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil
	}

	switch value.Type().Kind() {
	case reflect.Ptr:
		switch value.Type().Elem().Kind() {
		case reflect.Uint:
			i1 := uint(f)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Uint8:
			i1 := uint8(f)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Uint16:
			i1 := uint16(f)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Uint32:
			i1 := uint32(f)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Uint64:
			i1 := uint64(f)
			value.Set(reflect.ValueOf(&i1))
		case reflect.Uintptr:
			i1 := uintptr(f)
			value.Set(reflect.ValueOf(&i1))
		}
		return nil
	}
	value.SetUint(f)
	return nil
}
