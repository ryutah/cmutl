package list

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	want := []*ZipElem{
		&ZipElem{1, "a"},
		&ZipElem{2, "b"},
		&ZipElem{3, "c"},
	}
	if got := Zip(a, b); !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
