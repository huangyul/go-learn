package main

type List interface {
	Add(idx int, val any) error
	Delete(idx int) (any, error)
}

type ListA struct{}

func (ListA) Add(idx int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (ListA) Delete(idx int) (any, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	//u1 := User{
	//	Name: "Jerry",
	//	Age:  18,
	//}
	//u1.ChangeName("Tom")
	//u1.ChangeAge(30)
	//fmt.Printf("%+v", u1)

	var o Outer

	o.SayHello()
}
