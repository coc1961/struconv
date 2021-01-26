package struconv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"errors"

	"github.com/coc1961/struconv/internal/converters"
)

type Converter interface {
	String(a interface{}) (string, error)
	Set(value *reflect.Value, s string) error
}

type Scanner struct {
	converters map[reflect.Type]Converter
}

func New() *Scanner {
	s := &Scanner{}
	s.converters = make(map[reflect.Type]Converter)

	// String
	s.AddScanner(reflect.TypeOf(""), converters.StringConverter{})

	// Int
	s.AddScanner(reflect.TypeOf(int(1)), converters.IntConverter{})
	s.AddScanner(reflect.TypeOf(int8(1)), converters.IntConverter{})
	s.AddScanner(reflect.TypeOf(int16(1)), converters.IntConverter{})
	s.AddScanner(reflect.TypeOf(int32(1)), converters.IntConverter{})
	s.AddScanner(reflect.TypeOf(int64(1)), converters.IntConverter{})

	// Uint
	s.AddScanner(reflect.TypeOf(uint(1)), converters.UintConverter{})
	s.AddScanner(reflect.TypeOf(uint8(1)), converters.UintConverter{})
	s.AddScanner(reflect.TypeOf(uint16(1)), converters.UintConverter{})
	s.AddScanner(reflect.TypeOf(uint32(1)), converters.UintConverter{})
	s.AddScanner(reflect.TypeOf(uint64(1)), converters.UintConverter{})
	s.AddScanner(reflect.TypeOf(uintptr(1)), converters.UintConverter{})

	// Float
	s.AddScanner(reflect.TypeOf(float32(1)), converters.FloatConverter{Bytes: 32})
	s.AddScanner(reflect.TypeOf(float64(1)), converters.FloatConverter{Bytes: 64})

	// Bool
	s.AddScanner(reflect.TypeOf(true), converters.BoolConverter{})

	// Complex
	s.AddScanner(reflect.TypeOf(complex64(1)), converters.ComplexConverter{Bytes: 64})
	s.AddScanner(reflect.TypeOf(complex128(1)), converters.ComplexConverter{Bytes: 128})
	s.AddScanner(converters.ComplexPointer(64), converters.ComplexConverter{Pointer: true, Bytes: 64})
	s.AddScanner(converters.ComplexPointer(128), converters.ComplexConverter{Pointer: true, Bytes: 128})

	// Date
	s.AddScanner(reflect.TypeOf(time.Now()), converters.TimeConverter{})

	return s
}
func (s *Scanner) AddScanner(dataType reflect.Type, sc Converter) {
	s.converters[dataType] = sc
}

func (s *Scanner) Scan(dest interface{}, data map[string]interface{}) (retErr error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				retErr = e
				return
			}
			retErr = errors.New(fmt.Sprint(r))
		}
	}()

	retErr = nil
	value := reflect.ValueOf(dest)
	if value.Kind() != reflect.Ptr {
		return errors.New("no pointer parameter")
	}
	if reflect.Indirect(value).Kind() != reflect.Struct {
		return errors.New("no struct parameter")
	}

	var aerros multiError
	for k, v := range data {
		fld := value.Elem().FieldByName(k)
		if fld.Kind() == reflect.Invalid {
			continue
		}

		if v == nil {
			continue
		}

		src := reflect.ValueOf(v)

		fldType := fld.Type()
		srcType := src.Type()

		convSrc, ok := s.converters[srcType]
		convDst, ok1 := s.converters[fldType]
		if !ok || !ok1 {
			aerros = append(aerros, errors.New("no converter for "+k))
			continue
		}
		str, err := convSrc.String(v)
		if err != nil {
			aerros = append(aerros, fmt.Errorf("convert error for "+k+" %w", err))
			continue
		}
		err = convDst.Set(&fld, str)
		if err != nil {
			aerros = append(aerros, fmt.Errorf("convert error for "+k+" %w", err))
			continue
		}
	}
	if len(aerros) > 0 {
		retErr = aerros
	}
	return retErr
}

type multiError []error

func (e multiError) Error() string {
	err := make([]string, 0)

	for _, er := range e {
		err = append(err, er.Error())
	}
	b, _ := json.Marshal(err)
	return string(b)
}
