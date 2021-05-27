package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pc_break() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(time.Now().UnixNano())
	//n := rand.Intn(100) + 1
	//fmt.Println(n)
	//	break 语句的快速入门案例
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		n1 := rand.Intn(100) + 1
		//fmt.Printf("%d\t", n1)
		count++
		if n1 == 99 {
			break
		}
		fmt.Println("使用了 ", count)
	}
	// break 高级用法 使用标签跳出指定的循环
label1:
	for i := 0; i < 3; i++ {
		//label2:
		for j := 0; j < 3; j++ {
			if j == 2 {
				fmt.Println("end")
				fmt.Println(i)
				break label1
			}
		}
	}
	// 课堂练习1
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println(sum)
			fmt.Println(i)
			break
		}
	}
	// 课堂练习2
	var username string
	var password string
	ChanceNum := 3

	for ChanceNum > 0 {
		fmt.Println("输入username")
		fmt.Scanln(&username)
		fmt.Println("输入password")
		fmt.Scanln(&password)
		if username == "张无忌" && password == "888" {
			fmt.Println("success")
			break
		} else {
			fmt.Println("input error")
			ChanceNum--
			fmt.Printf("还有 %d 次\n", ChanceNum)
		}
	}
	if ChanceNum == 0 {
		fmt.Println("次数用完了")
	}

}
