package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// 变量操作
	/*	fmt.Println("hello,world")
		fmt.Println("hello,world")
		fmt.Println("hello,world")
		fmt.Println("hello,world")
		fmt.Println("姓名\t 年龄\t 籍贯\t 地址\njohn\t20\t 河北\t 北京")
		var a1 int = 900
		var b1 int = 1000
		var t1, t2, t3 = 123, "xcg", 345
		t4, t5, t6 := 123, "xcg", 234
		fmt.Println(t1, t2, t3)
		fmt.Println(t4, t5, t6)
		fmt.Println(a1)
		fmt.Println(b1)*/
	// 基本数据类型
	/*	var u uint = 1;
		var b byte = 255;
		fmt.Println(u, b)*/
	// 两种查看变量类型的方法
	// fmt.Printf
	// fmt.Sprintf
	// 整型
	var a = 12
	fmt.Printf("type:%T\n", a)
	// 64位系统默认int字节数为8
	fmt.Printf("字节数 %d\n", unsafe.Sizeof(a))
	var str string = fmt.Sprintf("%T", a)
	fmt.Println(str)
	// -3 10000011 11111110  补码：11111111
	// 3  00000011                00000011
	// 	  1000 = 8                11111110
	fmt.Println(-3 ^ 3)
	// 浮点型
	var f1 float32 = 123.00000901
	var f2 float32 = 123.00000901
	fmt.Println("sum = ", f1+f2)
	var f3 float64 = 123.00000901
	var f4 float64 = 123.00000901
	fmt.Println("sum = ", f3+f4)
	// 字符型
	var b1 byte = 'a'
	fmt.Println(b1)
	fmt.Printf("%c\n", b1)
	fmt.Printf("字节数：%d\n", unsafe.Sizeof(b1))
	fmt.Println(b1 + 10)
	// 布尔型
	var flag bool = true
	fmt.Println(flag)
	fmt.Println("bool占用空间：", unsafe.Sizeof(flag))
	// String类型
	var str1 string = "Hello World\t hhhhi"
	fmt.Println(str1)
	var str2 string = `hello World\t hhhhi`
	fmt.Println(str2)
	fmt.Println(str1 + str2)
	var str3 string = "hello " +
		"go " +
		"gopher"
	fmt.Println(str3)
	// 默认值
	var m1 int
	var m2 float64
	var m3 byte
	var m4 string
	var m5 bool
	fmt.Printf("%d\n%v\n%v\n%v\n%v\n", m1, m2, m3, m4, m5)
	// 类型转换
	var i2 int = 1234343
	var z1 float32 = float32(i2)
	var z2 float64 = float64(i2)
	var z3 int8 = int8(i2)
	fmt.Println(i2, z1, z2, z3)
	fmt.Printf("%v\t%v\t%v\t%v\t\n", i2, z1, z2, z3)
	fmt.Printf("%T\n", i2)
	// string类型和基本类型的转换
	var sz1 int = 123
	var sz2 float64 = 123.1
	//var sz3 byte = 'a'
	var sz4 bool = true
	//str = fmt.Sprintf("%d",sz1)
	//fmt.Printf("str Type %T, str = %q\n",str,str)
	//str = fmt.Sprintf("%f",sz2)
	//fmt.Printf("str Type %T, str = %q\n",str,str)
	//str = fmt.Sprintf("%c",sz3)
	//fmt.Printf("str Type %T, str = %q\n",str,str)
	//str = fmt.Sprintf("%t",sz4)
	//fmt.Printf("str Type %T, str = %q\n",str,str)
	str = strconv.FormatInt(int64(sz1), 10)
	fmt.Printf("str Type %T, str = %q\n", str, str)
	str = strconv.FormatFloat(sz2, 'f', 10, 64)
	fmt.Printf("str Type %T, str = %q\n", str, str)
	str = strconv.FormatBool(sz4)
	fmt.Printf("str Type %T, str = %q\n", str, str)
	//流程控制
	//pc_circle()
	pc_break()
}
