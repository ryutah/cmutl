package list

import (
	"reflect"
)

type ZipElem struct {
	first  interface{}
	second interface{}
}

func (z *ZipElem) First() interface{} {
	return z.first
}

func (z *ZipElem) Second() interface{} {
	return z.second
}

func Zip(a, b interface{}) []*ZipElem {
	var zipped []*ZipElem
	av, bv := reflect.ValueOf(a), reflect.ValueOf(b)
	for i := 0; i < av.Len(); i++ {
		z := new(ZipElem)
		z.first = av.Index(i).Interface()
		if i < bv.Len() {
			z.second = bv.Index(i).Interface()
		}
		zipped = append(zipped, z)
	}
	return zipped
}
