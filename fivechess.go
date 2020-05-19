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

/*  闭包函数
func adder() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			//  这里的i是因为 adder定义的返回参数是 func(int) 是这里的int
			pos(i),
			neg(-2*i),
		)
	}
}
*/

// 获取用户输入坐标
func changeGameMap(chessmap [10][10]string, corXY [2]int, s string) [10][10]string {
	chessmap[corXY[0]][corXY[1]] = s
	return chessmap
}

// 从控制台获取用户输入的坐标  多个返回值的时候，就因为少了一个逗号浪费了半个小时~~~
func getPlayerXY(n int) ([2]int, string) {
	var arr [2]int
	var play string
	if n%2 == 0 {
		play = "X" // 黑棋
		fmt.Printf("请%s选手输入落子坐标x,y \n", play)
	} else {
		play = "O" // 白棋
		fmt.Printf("请%s选手输入落子坐标x,y \n", play)
	}
	// 命令行调试专用
	/*	for i:=0;i<len(arr);i++{
		arr[i] = consoleStrToInt(os.Args[i+1])
	}*/

	fmt.Scanf("%d %d \n", &arr[0], &arr[1])
	return arr, play
}

//控制台输入的字符串转数字
/*func consoleStrToInt(s string) int{
	n,_ := strconv.Atoi(s)
	return n
}*/

func closureChangeMap(chessmap [10][10]string) func([2]int, string) [10][10]string {
	temap := chessmap
	return func(arr [2]int, s string) [10][10]string {
		temap[arr[0]][arr[1]] = s
		return temap
	}
}

// 黑白双方轮流下棋(默认黑棋先走 X)返回下一步下棋的人
func startChess(chessmap [10][10]string) {
	m := chessmap
	for i := 1; i < 10; i++ { // 先循环10次防止死机
		// 选手走棋函数  解决 判断黑白子选手，以及用户输入的坐标问题
		arr, str := getPlayerXY(i)
		// 根据选手操作改变游戏地图
		//  这里的闭包 到底怎么关联呢？奇怪了
		/*
			其实不需要使用到闭包，因为这里就是少了一个保存临时变量 m 用于保存用户地图信息
		*/
		m = changeGameMap(m, arr, str)
		mapview(m)
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
