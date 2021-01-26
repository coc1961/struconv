package struconv

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	UInt       uint
	UInt8      uint8
	UInt16     uint16
	UInt32     uint32
	UInt64     uint64
	UIntptr    uintptr
	Float32    float32
	Float64    float64
	String     string
	Time       time.Time
	Bool       bool
	Complex64  complex64
	Complex128 complex128
}

type MyPointerStruct struct {
	Int        *int
	Int8       *int8
	Int16      *int16
	Int32      *int32
	Int64      *int64
	UInt       *uint
	UInt8      *uint8
	UInt16     *uint16
	UInt32     *uint32
	UInt64     *uint64
	UIntptr    *uintptr
	Float32    *float32
	Float64    *float64
	String     *string
	Time       *time.Time
	Bool       *bool
	Complex64  *complex64
	Complex128 *complex128
}

func TestScanner_Scan_StingValues(t *testing.T) {
	s := New()

	var st MyStruct
	err := s.Scan(&st, map[string]interface{}{
		"String":     "Hola Juan",
		"Int":        "1000",
		"Int8":       "100",
		"Int16":      "1000",
		"Int32":      "1000",
		"Int64":      "1000",
		"UInt":       "1000",
		"UInt8":      "1000",
		"UInt16":     "1000",
		"UInt32":     "1000",
		"UInt64":     "1000",
		"UIntptr":    "1000",
		"Float32":    "10.32",
		"Float64":    "99.54",
		"Time":       "2021-01-22T16:08:09.000Z",
		"Bool":       "true",
		"Complex64":  "1234.555",
		"Complex128": "5678.998",
	})

	assert.Nil(t, err)
}

func TestScanner_Scan_NaturalValues(t *testing.T) {
	s := New()

	var st MyStruct
	err := s.Scan(&st, map[string]interface{}{
		"String":     "Hola Juan",
		"Int":        int(1000),
		"Int8":       int8(100),
		"Int16":      int16(1000),
		"Int32":      int32(1000),
		"Int64":      int64(1000),
		"UInt":       uint(1000),
		"UInt8":      uint8(100),
		"UInt16":     uint16(1000),
		"UInt32":     uint32(1000),
		"UInt64":     uint64(1000),
		"UIntptr":    uintptr(1000),
		"Float32":    float32(10.32),
		"Float64":    float64(99.54),
		"Time":       time.Now(),
		"Bool":       true,
		"Complex64":  complex64(1234.555),
		"Complex128": complex128(5678.998),
	})

	assert.Nil(t, err)
}

func TestScanner_Scan_InvalidData(t *testing.T) {
	s := New()

	var st MyStruct
	err := s.Scan(&st, map[string]interface{}{
		"String":     "Hola Juan",
		"Int":        "aa",
		"Int8":       "aa",
		"Int16":      "aa",
		"Int32":      "aa",
		"Int64":      "aa",
		"UInt":       "aa",
		"UInt8":      "aa",
		"UInt16":     "aa",
		"UInt32":     "aa",
		"UInt64":     "aa",
		"UIntptr":    "aa",
		"Float32":    "aa",
		"Float64":    "aa",
		"Time":       "aa",
		"Bool":       "aa",
		"Complex64":  "aa",
		"Complex128": "aa",
	})

	fmt.Println(err)
	assert.NotNil(t, err)
}

func TestScanner_Scan_NilData(t *testing.T) {
	s := New()

	var st MyStruct
	err := s.Scan(&st, map[string]interface{}{
		"String":     nil,
		"Int":        nil,
		"Int8":       nil,
		"Int16":      nil,
		"Int32":      nil,
		"Int64":      nil,
		"UInt":       nil,
		"UInt8":      nil,
		"UInt16":     nil,
		"UInt32":     nil,
		"UInt64":     nil,
		"UIntptr":    nil,
		"Float32":    nil,
		"Float64":    nil,
		"Time":       nil,
		"Bool":       nil,
		"Complex64":  nil,
		"Complex128": nil,
	})

	fmt.Println(err)
	assert.Nil(t, err)
}

func TestScanner_Scan_InvalidDestType(t *testing.T) {
	s := New()

	var st MyStruct
	err := s.Scan(st, map[string]interface{}{})

	fmt.Println(err)
	assert.NotNil(t, err)

	s1 := "Hello"
	err = s.Scan(&s1, map[string]interface{}{})

	fmt.Println(err)
	assert.NotNil(t, err)
}

func TestScanner_Scan_NoConverter(t *testing.T) {
	s := New()

	var st InvalidStruct
	err := s.Scan(&st, map[string]interface{}{
		"NoConverterField": InvalidStruct{},
	})

	fmt.Println(err)
	assert.NotNil(t, err)

}

type InvalidStruct struct {
	NoConverterField MyStruct
}

func TestScanner_Scan_Complex64PointerValues(t *testing.T) {
	s := New()

	complex64 := complex64(200.2)
	complex128 := complex128(100.2)
	var st MyPointerStruct
	err := s.Scan(&st, map[string]interface{}{
		"Complex64":  &complex64,
		"Complex128": &complex128,
	})

	assert.NotNil(t, st.Complex64)
	assert.NotNil(t, st.Complex128)
	assert.Nil(t, err)

	err = s.Scan(&st, map[string]interface{}{
		"Complex64":  "200.2",
		"Complex128": "100.2",
	})

	assert.NotNil(t, st.Complex64)
	assert.NotNil(t, st.Complex128)
	assert.Nil(t, err)
}
