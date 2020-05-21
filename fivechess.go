package main

import (
	"fmt"
)

func initChessMap(chessmap [10][10]string) [10][10]string {
	fmt.Println("初始化棋盘如下： \n")
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

// 获取用户输入坐标    用户输入的坐标 1 3  在这里才被转换为棋盘上对应的  因为棋盘是从0开始计数
func changeGameMap(chessmap [10][10]string, corXY [2]int, s string) [10][10]string {
	chessmap[corXY[0]][corXY[1]] = s
	mapview(chessmap)
	return chessmap
}

// 从控制台获取用户输入的坐标  多个返回值的时候，就因为少了一个逗号浪费了半个小时~~~
func getPlayerXY(n int) ([2]int, string) {
	var arr, temp [2]int
	var play string
	if n%2 == 0 {
		play = "X" // 黑棋
		fmt.Printf("请%s选手输入落子坐标x,y \n", play)
	} else {
		play = "O" // 白棋
		fmt.Printf("请%s选手输入落子坐标x,y \n", play)
	}
	// 命令行调试专用
	/*
			for i:=0;i<len(arr);i++{
			arr[i] = consoleStrToInt(os.Args[i+1])
		}
	*/
	fmt.Scanf("%d %d \n", &temp[0], &temp[1])
	// 用户的输入习惯是  >=1  不会小于1
	if temp[0] < 1 || temp[1] < 1 || temp[1] > 10 || temp[1] > 10 {
		fmt.Printf("输入的落子坐标不合法,请重新输入 \n")
		getPlayerXY(n)
	} else {
		arr[0], arr[1] = temp[0]-1, temp[1]-1
	}
	return arr, play
}

//控制台输入的字符串转数字
/*
func consoleStrToInt(s string) int{
	n,_ := strconv.Atoi(s)
	return n
}
*/

// 落子的坐标有其他棋子
func isChessValid(chessmap [10][10]string, ordXY [2]int) bool {
	if chessmap[ordXY[0]][ordXY[1]] != "_" {
		fmt.Print("棋手落子位置非法 \n")
		return false
	}
	if ordXY[0] > 9 || ordXY[0] < 0 || ordXY[1] > 9 || ordXY[1] < 0 {
		fmt.Print("棋手坐标违法，超出棋盘边界 \n")
		return false
	}
	return true
}

// 使用结构体
type FiveChess struct {
	color   string //  棋子颜色
	lineNum int
}

//  扫描五子棋是否形成  分为 X  Y  斜方向的情况  游戏结束 返回 false  没有结束返回 true
func scanWalker(chessmap [10][10]string) bool {
	lenMap := len(chessmap)
	s := FiveChess{}
	for i := 0; i < lenMap; i++ {
		for j := 0; j < lenMap; j++ {
			s.color = chessmap[i][j]
			if s.color == "_" {
				continue
			}
			s.lineNum = 0
			//不能超出棋盘边界  主要是判断i+1>lenMap的情况
			if i <= lenMap-1 || j <= lenMap-1 {
				// 判断 Y 轴方向五子棋
				if chessmap[i+1][j] == s.color {
					s.lineNum++
					continue
				} else {
					s.lineNum = 0
				}
				// 判断 X 轴五子棋
				if chessmap[i][j+1] == s.color {
					s.lineNum++
					continue
				} else {
					s.lineNum = 0
				}
				// 判断斜线方向
				if chessmap[i+1][j+1] == s.color {
					s.lineNum++
					continue
				} else {
					s.lineNum = 0
				}
				if s.lineNum == 5 {
					fmt.Print("GAME is over \n")
					return false
				}
			} else {
				continue
			}
		}
	}
	return true //  这里出问题了  例如棋盘只有两个棋子，最后也会进入这条语句
}

func isGameOver(chessmap [10][10]string) bool {
	// 棋盘满足获胜条件 返回的false 没有结束返回true
	state := scanWalker(chessmap)
	if state {
		return true // 游戏没有结束
	} else {
		return false
	}
}

func chessRule(chessmap [10][10]string) bool {
	if isGameOver(chessmap) {
		return true //  游戏没有结束 返回true
	} else {
		return false
	}
}

// 黑白双方轮流下棋(默认黑棋先走 X)返回下一步下棋的人
func startChess(chessmap [10][10]string) {
	m := chessmap
	for i := 1; i < 10; i++ { // 先循环10次防止死机
		// 选手走棋函数  解决 判断黑白子选手，以及用户输入的坐标问题
		arr, str := getPlayerXY(i)
		if isChessValid(m, arr) { // 落子的位置合法
			m = changeGameMap(m, arr, str)
		}
		state := chessRule(m) //  游戏没有结束返回true
		if state {
			continue
		} else {
			fmt.Printf("棋手%s胜利了，祝贺他/她 \n", str)
		}
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
