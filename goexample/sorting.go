package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println("str befored sorted:", sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println("strs: ", strs)
	fmt.Println("str after sorted:", sort.StringsAreSorted(strs))

	ins := []int{3, 4, 2, 1, 6, 9, 3}
	sort.Ints(ins)
	fmt.Println("ins: ", ins)

	s := sort.IntsAreSorted(ins)
	fmt.Println("Sorted: ", s)
}
