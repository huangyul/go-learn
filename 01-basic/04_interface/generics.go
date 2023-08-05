package main

type list[T any] interface {
	Add(a int, b T)
}

func Sum[T Number](n T) T {
	return n + 1
}

type Number interface {
	int | int64
}
