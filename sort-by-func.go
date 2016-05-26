package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s, %d", p.name, p.age)
}

type PersonList []Person

func (list PersonList) Len() int {
	return len(list)
}

func (list PersonList) Less(i, j int) bool {
	return list[i].age < list[j].age
}

func (list PersonList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func main() {
	fruits := []string{"peace", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	pers := []Person{{"lbj", 10}, {"kobe", 13}, {"cp4", 4}}

	pList := PersonList(pers)

	for i, p := range pers {
		fmt.Println(i, p.String())
	}

	sort.Sort(pList)
	fmt.Println(pList)
}
