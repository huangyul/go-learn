package main

type Inner struct{}

func (i Inner) Name() string {
	return "inner"
}
func (i Inner) SayHello() {
	println("hello," + i.Name())
}

type Outer struct {
	Inner
}

func (o Outer) Name() string {
	return "outer"
}
