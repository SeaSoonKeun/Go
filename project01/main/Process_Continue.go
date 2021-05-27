package main

import "fmt"

func pc_continue() {
	// 课堂练习1
	//here:
	//	for i := 0; i < 2; i++ {
	//		for j := 0; j < 5; j++ {
	//			if j == 2 {
	//				continue here
	//			}
	//			fmt.Println("i = ", i, "j = ", j)
	//		}
	//	}
	// 课堂练习2
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
