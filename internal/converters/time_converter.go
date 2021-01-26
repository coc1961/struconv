package converters

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.000Z"
)

var parseTime = []string{
	timeFormat,
	"2006-01-02",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

type TimeConverter struct {
}

func (c TimeConverter) String(a interface{}) (string, error) {
	if t, ok := a.(time.Time); ok {
		return t.Format(timeFormat), nil
	}
	return fmt.Sprint(a), nil
}

func (c TimeConverter) Set(value *reflect.Value, s string) error {
	for n := 0; n < len(parseTime); n++ {
		i, err := time.Parse(timeFormat, s)
		if err == nil {
			value.Set(reflect.ValueOf(i))
			return nil
		}

	}
	return errors.New("invalid date format")
}
