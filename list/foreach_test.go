package list

import (
	"testing"
)

func TestForEach(t *testing.T) {
	l := []string{"a", "b", "c"}
	called := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
	}
	f := func(s string) {
		if _, ok := called[s]; !ok {
			t.Errorf("Unknown string received %v", s)
			return
		}
		called[s]++
	}
	ForEach(l, f)

	for k, v := range called {
		if v != 1 {
			t.Errorf("%v must call just one time. but called %v times", k, v)
		}
	}
}
