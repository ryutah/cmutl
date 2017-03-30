package list

import (
	"reflect"
)

func ForEach(list interface{}, f interface{}) {
	lv, fv := reflect.ValueOf(list), reflect.ValueOf(f)

	for i := 0; i < lv.Len(); i++ {
		v := lv.Index(i)
		fv.Call([]reflect.Value{v})
	}
}
