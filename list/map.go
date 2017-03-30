package list

import (
	"reflect"
)

func Map(list interface{}, f interface{}) interface{} {
	lv, fv, ft := reflect.ValueOf(list), reflect.ValueOf(f), reflect.TypeOf(f)
	retType := ft.Out(0)
	retSlice := reflect.SliceOf(retType)
	retVal := reflect.MakeSlice(retSlice, lv.Len(), lv.Cap())
	for i := 0; i < lv.Len(); i++ {
		v := lv.Index(i)
		out := fv.Call([]reflect.Value{v})[0]
		retVal.Index(i).Set(out)
	}
	return retVal.Interface()
}
