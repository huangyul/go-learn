package main

import "unicode/utf8"

func main() {
	var a int = 456
	var b int = 123
	println(a + b)

	println(len("huang"))                 // 5
	println(len("你好"))                    // 6
	println(utf8.RuneCountInString("你好")) // 2
}
