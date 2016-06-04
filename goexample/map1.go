package main

import (
	"fmt"
)

type PersonInfo struct {
	ID   string
	Name string
	Age  int
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	//map insert
	personDB["12345"] = PersonInfo{"12345", "Tom", 5}
	personDB["1"] = PersonInfo{"1", "t-mac", 35}
	personDB["2"] = PersonInfo{"2", "lbj", 34}

	//map find
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person: ", person.ID, person.Name, person.Age)
	} else {
		fmt.Println("Found no person")
	}

	person1, ok1 := personDB["12345"]
	if ok1 {
		fmt.Println("Found person: ", person1.ID, person1.Name, person1.Age)
	} else {
		fmt.Println("Found no person")
	}

	fmt.Println("...................")

	//itertor map && map delete
	for key, value := range personDB {
		if key == "1" {
			delete(personDB, key)
			continue
		}
		fmt.Println(key, " ", value.Name, " ", value.Age)
	}

}
