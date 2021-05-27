package main

// 虽然语法上没有做要求 包名和用户名必须一致 当时为了区分 一般的设置是一致的
// 同一个包下 不能有相同的函数名
func Fbnc(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fbnc(n-1) + Fbnc(n-2)
	}
}

//func MonkeyEat(Num int) int  {
//	// 思路：每天吃 n/2 + 1 第十天 只剩下一个
//	for i:=1;i<=9;i++{
//		Num = 2 * (Num + 1)
//	}
//	return Num
//}
var num int

func pinkNum(Times int) int {
	if Times == 9 {
		return 1
	}
	if Times < 9 || Times > 0 {
		num = 2 * (pinkNum(Times+1) + 1)
	}
	return num
}
