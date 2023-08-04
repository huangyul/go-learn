package main

import "fmt"

type List interface {
	Add(idx int, val any) error
	Delete(idx int) (any, error)
}

func main() {
	u1 := User{
		Name: "Jerry",
		Age:  18,
	}
	u1.ChangeName("Tom")
	u1.ChangeAge(30)
	fmt.Printf("%+v", u1)
}
