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
	var temp [2]int
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

	fmt.Scanf("%d %d \n", &temp[0], &temp[1])
	arr[0], arr[1] = temp[0]-1, temp[1]-1
	return arr, play
}

//控制台输入的字符串转数字
/*
func consoleStrToInt(s string) int{
	n,_ := strconv.Atoi(s)
	return n
}
*/

/*
func closureChangeMap(chessmap [10][10]string) func([2]int, string) [10][10]string {
	temap := chessmap
	return func(arr [2]int, s string) [10][10]string {
		temap[arr[0]][arr[1]] = s
		return temap
	}
}
*/

// 落子的坐标有其他棋子
func isChessValid(chessmap [10][10]string, ordXY [2]int) bool {
	if chessmap[ordXY[0]][ordXY[1]] != "_" {
		fmt.Print("棋手落子位置非法")
		return false
	}
	return true
}

func isValid(chessmap [10][10]string, ordXY [2]int) bool {
	// 超出边界
	if ordXY[0] > 10 || ordXY[0] < 1 || ordXY[1] > 10 || ordXY[1] < 1 {
		fmt.Print("棋手坐标违法，超出棋盘边界")
		return false
	}
	// 棋手顺序违法  暂时先不考虑这个点
	// 落子的坐标有其他棋子
	if isChessValid(chessmap, ordXY) {
		return true
	} else {
		return false
	}
}

// 使用结构体
type FiveChess struct {
	color   string //  棋子颜色
	lineNum int
}

//  扫描五子棋是否形成  分为 X  Y  斜方向的情况
func scanWalker(chessmap [10][10]string) bool {
	for i := 0; i < len(chessmap); i++ {
		for j := 0; j < len(chessmap[i]); j++ {
			s := FiveChess{}
			s.color = chessmap[i][j]
			n := 0
			// 判断 X轴方向五子棋
			if chessmap[i+1][j] == s.color {
				n++
			} else {
				n = 0
			}
			// 判断Y轴五子棋
			if chessmap[i][j+1] == s.color {
				n++
			} else {
				n = 0
			}
			// 判断斜线方向
			if chessmap[i+1][j+1] == s.color {
				n++
			} else {
				n = 0
			}
			if n == 5 {
				return true
			}
		}
	}
	return false //  无法继续游戏
}

func gameOver(chessmap [10][10]string) bool {
	// XY轴五子棋未超出棋盘
	//X 与 Y轴要分开扫描   否则会出错   写一个函数
	state = scanWalker(chessmap)
	if state {
		return true
	} else {
		return false
	}
}

func chessRule(chessmap [10][10]string, ordXY [2]int) bool {
	//	1.不能超出棋盘边界          2.棋盘满足获胜条件
	if isValid(chessmap, ordXY) && gameOver(chessmap) {
		return true
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
		// 根据选手操作改变游戏地图
		//  这里的闭包 到底怎么关联呢？奇怪了
		/*
			其实不需要使用到闭包，因为这里就是少了一个保存临时变量 m 用于保存用户地图信息
		*/
		m = changeGameMap(m, arr, str)
		state = chessRule(m, arr)
		if state {
			mapview(m)
		} else {
			break
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
