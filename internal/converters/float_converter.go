package converters

import (
	"fmt"
	"reflect"
	"strconv"
)

type FloatConverter struct {
	Bytes int
}

func (c FloatConverter) String(a interface{}) (string, error) {
	return fmt.Sprint(a), nil
}

func (c FloatConverter) Set(value *reflect.Value, s string) error {
	i, err := strconv.ParseFloat(s, c.Bytes)
	if err != nil {
		return err
	}
	value.SetFloat(i)
	return nil
}
