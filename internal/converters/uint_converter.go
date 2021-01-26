package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type UintConverter struct {
}

func (c UintConverter) String(a interface{}) (string, error) {
	return fmt.Sprint(a), nil
}

func (c UintConverter) Set(value *reflect.Value, s string) error {
	f, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil
	}
	value.SetUint(f)
	return nil

}
