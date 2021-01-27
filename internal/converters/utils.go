package converters

import "reflect"

func Pointer(a interface{}) reflect.Type {
	pt := reflect.TypeOf(a)
	return reflect.PtrTo(pt)
}
