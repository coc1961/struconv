package converters

import "reflect"

func ComplexPointer(bytes int) reflect.Type {
	if bytes == 64 {
		r := complex64(1)
		return reflect.TypeOf(&r)
	}
	r := complex128(1)
	return reflect.TypeOf(&r)
}
