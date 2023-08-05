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

### 切片

语法是`[]type`，跟数组的区别就是有没有初始化长度

- 直接初始化
- make初始化:make([]type, length, capacity)
- arr[i] 的形式访问元素
- append 追加元素
- len 获取元素数
- 推荐写法：s1 := make([]type, 0, capacity)
- 使用 for range 来遍历

###### 共享数组

因为切片的子层都是基于数组，所以子切片与原本的切片会共享数组，改了其中一个另外一个也会被改变，当任何一个发生扩容，共享的性质就会发生改变

### comparable

- 在switch里面，值必须是可比较的
- 在map里面，key必须是可比较的

基本概念：

- 基本类型和string都是可比较的
- 如果元素可比较，数组也是可比较的

## GO接口与类型定义

接口是一组行为的抽象 *当你怀疑要不要定义接口的时候，就加上去*

```go
package main

type List interface {
	Add(idx int, val any) error
	Append(val any)
	Delete(idx int) (any, error)
}
```

### 方法接收器

定义在结构体上，只有使用指针的时候才能改变本身，因为函数是按值传递的，在方法里面相当于复制一份，而指针都指向同一个内容

### 衍生类型和类型别名

衍生类型:`type Type1 Type2`

是个全新的类型，但是不会继承方法

类型别名:`type Type1 = Type2`

既有方法也有方法

## 鸭子类型

结构体实现了接口所定义的方法

## 组合

- 当A组合B之后：
   - 可以直接在A上调用B的方法
   - B实现的所有接口，都认为A已经实现了
- A组合B后，在初始化A的时候将B看作普通字段来初始化

```go
package main

type TypeA struct{}

func (a *TypeA) SayA() {
}

type TypeB struct {
  TypeA
}

func main() {
  a := TypeB{
    TypeA: TypeA{}, // 这样去初始化里面的结构体
  }
  a.SayA() // 可以使用A的方法
}
```