package main

import (
	"fmt"
	"strings"
)

//search for t index
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//the same to index func
func Include(vs []string, t string) int {
	return Index(vs, t)
}

//any one can make f func true will be return true or return false
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//All has to make f func true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !(f(v)) {
			return false
		}
	}
	return true
}

//return v in vs which can make f func be true
func Filter(vs []string, f func(string) bool) []string {
	vsr := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsr = append(vsr, v)
		}
	}
	return vsr
}

//which can makes f func be true one by one
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
