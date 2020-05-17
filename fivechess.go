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

// 获取用户输入坐标
func changeGameMap(chessmap [10][10]string, player string, corXY [2]int) {
	fmt.Printf("获取到的用户坐标为：%v \n", corXY)
	// i := &chessmap[corXY[0]][corXY[1]]
	// *i = player
	chessmap[corXY[0]][corXY[1]] = player
	mapview(chessmap)
}

// 从控制台获取用户输入的坐标
func getPlayerXY(player string) [2]int {
	var arr [2]int
	fmt.Printf("请%s选手输入落子坐标x,y \n", player)
	fmt.Scanf("%d %d \n", &arr[0], &arr[1])
	return arr
}

func closurefunc() {

}

// 黑白双方轮流下棋(默认黑棋先走 X)返回下一步下棋的人
func startChess(chessmap [10][10]string) {
	var player string
	n := 0
	for n < 10 {
		// 黑棋用X  白棋用O
		if n%2 == 0 { // 黑棋
			player = "X"
			changeGameMap(chessmap, player, getPlayerXY(player))
		} else { // 白棋
			player = "O"
			changeGameMap(chessmap, player, getPlayerXY(player))
		}
		n++
		mapview(chessmap)
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
