package main

import "fmt"

// 函数返回两个值 一个数组 一个字符串
func reData() ([2]int, string) {
	var arr = [2]int{1, 2}
	var s = "wen"
	return arr, s
}
func main() {
	fmt.Print(reData())
}
