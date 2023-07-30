package main

func main() {
	// 1. defer执行顺序
	// 第二个先输出
	//defer func() {
	//	fmt.Println("第一个defer")
	//}()
	//defer func() {
	//	fmt.Println("第二个defer")
	//}()

	// 3. 输出问题
	// 输出2，因为永远是return前才执行
	//i := 0
	//defer func() {
	//	fmt.Println(i)
	//}()
	//i = 2
	// 输出0，因为先存下了i
	//i := 0
	//defer func(val int) {
	//	println(i)
	//}(i)
	//i = 2
}
