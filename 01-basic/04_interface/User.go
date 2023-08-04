package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func newUser() {
	u := User{
		Name: "user",
		Age:  10,
	}
	fmt.Printf("%+v", u)
}

func (u User) ChangeName(name string) {
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}
