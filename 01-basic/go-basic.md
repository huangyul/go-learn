# go 基础语法

> 太基础的就不记录了，记录一些难点易错点

## package

- 同一个目录下的所有go文件必须是同一个package
- 如果是test文件是可以改package，一般不会改

## 基本类型和string

基本类型包括:

- int家族：正常用int即可
- uint家族：正常用uint即可
- float家族：优先使用float64

### string

###### 获取字符长度

```go
package main

import "unicode/utf8"

func main() {
	print("huang")                          // 5
	print("你好")                           // 6
	println(utf8.RuneCountInString("你好")) // 正确获取中文2

}

```

## func

#### defer

延迟调用：在函数返回的时候执行

注意点：执行顺序，最后的先执行，栈的特点：`后进先出`

```go
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
```

### 数组

语法是`[cap]type`

- 初始化要指定长度
- 可以直接初始化
- arr[i]的方式访问元素
- len和cap可以获取数组的长度和容量
- 使用for range循环