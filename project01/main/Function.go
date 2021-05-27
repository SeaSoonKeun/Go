package main

import "fmt"

/*// swap
func swap(a *int,b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}*/

// 执行顺序： 全局变量 > init > main
/*var AGE = test()
func test() int {
	fmt.Println("test()")
	return 90
}
func init() {
	fmt.Println("init()")
}
func main() {
	fmt.Println("main")*/

// 匿名函数 只能调用一次
func test() {
	// 没有名字的函数
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)
}

//1) 编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
//2) 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg 后缀，则返回原文件名。
//3) 要求使用闭包的方式完成
//4) strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。
