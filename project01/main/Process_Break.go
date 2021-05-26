package main

import (
	"fmt"
	"math/rand"
	"time"
)

//	break 语句的快速入门案例
func pc_break() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(time.Now().UnixNano())
	//n := rand.Intn(100) + 1
	//fmt.Println(n)
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		n1 := rand.Intn(100) + 1
		fmt.Printf("%d\t", n1)
		count++
		if n1 == 99 {
			break
		}
		fmt.Println("使用了 ", count)
	}
}
