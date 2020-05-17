package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	for i := 0; i < 4; i++ {
		//fmt.Printf("第%d次循环",i)
		fmt.Print(os.Args[i+1], reflect.TypeOf(os.Args[i+1]), "\n")
		//fmt.Printf("---->","%d",os.Args[1],"\n")
	}
	//v := 3
	fmt.Print(reflect.TypeOf(string(3)), 3, reflect.TypeOf(3),"\n")

	x, _ := strconv.Atoi(os.Args[1])
	fmt.Print(reflect.TypeOf(x))   // 字符串转数字
}
