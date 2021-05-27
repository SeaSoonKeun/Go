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

// 变量作用范围 赋值语句不能在函数体外 除了全局变量
// 全局变量
var name = "tome"

// name := "tome" 报错
var name1 = "jack"

// 编写函数，对给定二维数组转置
// 123   -->  321
// 456		  654
// 789		  987
//func reverseArray()  {
//	arr := [3][3]int[1, 2, 3], [4, 5, 6], [7, 8, 9]
//	for i := 0; i < len(arr); i++ {
//		for j := 0; j < len(arr); j++ {
//			var tmp int
//			tmp = arr[i][j]
//			arr[i][j] = arr[j][i]
//			arr[j][i] = tmp
//		}
//	}
//}

// 字符串常用函数
func scanString() {
	s1 := "hello 中国"
	// 1、len函数
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d ：%c\n", i, s1[i])
	}
	// 2、字符串遍历 为了避免乱码的出现，使用rune切片进行解决
	r := []rune("hello 中国")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d ：%c\n", i, r[i])
	}
	// 3、字符串转整型
	// 4、整型转字符串
	// 5、字符串转byte
	// 6、byte转字符串
	// 7、10进制转2、8、16进制
	// 8、查找字符串是否包含子串
	// 9、查找字符串包含几个指定的子串
	// 10、不区分大小写的字符串比较（==符号默认是区分大小写的）
	// 11、查找子串在字符串中第一个出现的index值
	// 12、查找子串在字符串中最后一个出现的index值
	// 13、替换字符串中指定的子串
	// 14、按照指定的字符进行分割，拆分字符串
	// 15、字符串字符进行大小写转换
	// 16、去掉字符串左右两侧的空格
	// 17、将字符串左右两边指定的字符串去掉
	// 18、去掉字符串左侧指定的字符串
	// 19、去掉字符串右侧指定的字符串
	// 20、判断字符串是否以指定的字符串开头
	// 21、判断字符串是否以指定的字符串结束
}
