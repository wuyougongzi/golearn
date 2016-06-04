package main

import (
	"errors"
	"fmt"
)

func normalError(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can not work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prop string
}

//implement the error's interface
func (ar argError) Error() string {
	return fmt.Sprintf("%d - %s", ar.arg, ar.prop)
}

func definsErrors(arg int) (int, error) {
	if arg == 42 {
		//return -1, &argError{arg, "can not work it"}
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := normalError(i); e != nil {
			fmt.Println("normal failed", e)
		} else {
			fmt.Println("normal success", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := definsErrors(i); e != nil {
			fmt.Println("defineErrors failed", e)
		} else {
			fmt.Println("defineErrors success", r)
		}
	}

	_, e := definsErrors(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prop)
	}
}
