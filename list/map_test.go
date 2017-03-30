package list

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	l := []string{"1", "2", "3"}
	f := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i * i
	}

	want := []int{1, 4, 9}
	got, ok := Map(l, f).([]int)
	if !ok {
		t.Fatal("Cant convert want type")
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
