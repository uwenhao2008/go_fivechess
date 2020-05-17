package main

import (
	"fmt"
)

func initChessMap(chessmap [10][10]string) [10][10]string {
	fmt.Println("初始化棋盘如下：")
	for i := 0; i < len(chessmap); i++ {
		for j := 0; j < len(chessmap[i]); j++ {
			chessmap[i][j] = "_"
		}
		// fmt.Print(chessmap[i], "\n")
	}
	return chessmap
}

// func changeGameMap(chessmap [10][10]string, playColor string, corXY [2]int) func(string) [2]int {

// }

/*  闭包函数
func adder() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x
		return sum
	}
}
*/

// 获取用户输入坐标
func changeGameMap(chessmap [10][10]string, n int, corXY [2]int) {
	i := n%2
	// 黑棋用X  白棋用O
	if i == 0{
		chessmap[corXY[0]][corXY[1]] = "X"
	}else{
		chessmap[corXY[0]][corXY[1]] = "O"
	}
	mapview(chessmap)
}

// 从控制台获取用户输入的坐标
func getPlayerXY(player string) [2]int {
	var arr [2]int
	fmt.Printf("请%s选手输入落子坐标x,y \n", player)

	// 命令行调试专用
/*	for i:=0;i<len(arr);i++{
		arr[i] = consoleStrToInt(os.Args[i+1])
	}*/

	fmt.Scanf("%d %d \n", &arr[0], &arr[1])
	return arr
}

//控制台输入的字符串转数字
/*func consoleStrToInt(s string) int{
	n,_ := strconv.Atoi(s)
	return n
}*/

// 黑白双方轮流下棋(默认黑棋先走 X)返回下一步下棋的人
func startChess(chessmap [10][10]string) {
	var player string
	n := 0
	for n < 10 {
		changeGameMap(chessmap, n, getPlayerXY(player))
		n++
	}
}

func mapview(chessmap [10][10]string) {
	for i := 0; i < len(chessmap); i++ {
		fmt.Println(chessmap[i])
	}
}

func main() {
	// 棋盘初始化
	var chessmap [10][10]string
	gameMap := initChessMap(chessmap)
	// 开始游戏
	startChess(gameMap)
}
