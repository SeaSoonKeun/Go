package main

import (
	"fmt"
)

func pc_circle() {
	// 循环 for用法
	//for i := 0; i < 100; i++ {
	//	fmt.Println("sucess")
	//}

	// for循环遍历string字符
	//var str = "abc)ik"
	//for index, value := range str {
	//	fmt.Printf("index=%d,val=%c\n", index,value)
	//}
	var str2 = "abc)ik！南理工"
	for index, value := range str2 {
		fmt.Printf("index=%d,val=%c\n", index, value)
	}

	//课堂练习1
	sum := 0
	count := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Println(count, sum)

	//课堂练习2

	for i := 0; i <= 6; i++ {
		fmt.Printf("%d\t*\t%d\t=\t6\n", i, 6-i)
	}

	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 打印金字塔
	/*
		  *    2空格 1星
		 ***   1空格 3星
		*****  0空格 5星
	*/
	var N int
	fmt.Scanln(&N)
	for i := 1; i <= N; i++ {
		for k := 1; k <= N-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 打印空心金字塔
	/*
			  *    2空格 1星
			 * *   1空格 3星
			*   *  0空格 5星
		   *******
	*/
	var M int
	fmt.Scanln(&M)
	for i := 1; i <= M; i++ {
		for k := 1; k <= M-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == M {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
