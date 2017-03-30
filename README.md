# cmutl
Go言語 Map Reduceユーティリティ

# Usage
### Example
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/ryutah/cmutl/list"
)

func main() {
	fmt.Println("Start ForEach")
	lforeach := []string{"a", "b", "c"}
	list.ForEach(lforeach, func(s string) {
		fmt.Println(s)
	})
	fmt.Println("End ForEach")

	fmt.Println("Start Map")
	lmap := []int{1, 2, 3}
	lmapped := list.Map(lmap, func(i int) string {
		return strconv.Itoa(i)
	}).([]string)
	fmt.Println(lmapped)
	fmt.Println("End Map")

	fmt.Println("Start Zip")
	lzipa := []int{1, 2, 3}
	lzipb := []string{"a", "b", "c"}
	lzipped := list.Zip(lzipa, lzipb)
	list.ForEach(lzipped, func(z *list.ZipElem) {
		fmt.Printf("First %v, Second %v\n", z.First(), z.Second())
	})
	fmt.Println("End Zip")
}
```

### Result
```console
$ Start ForEach
$ a
$ b
$ c
$ End ForEach
$
$ Start Map
$ [1 2 3]
$ End Map
$
$ Start Zip
$ First 1, Second a
$ First 2, Second b
$ First 3, Second c
$ End Zip
```
