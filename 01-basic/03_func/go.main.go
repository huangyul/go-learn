package main

import "fmt"

// defer
func main() {
	defer func() {
		fmt.Println("第一个defer")
	}()
	defer func() {
		fmt.Println("第二个defer")
	}()
}
