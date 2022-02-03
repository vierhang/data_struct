package main

import "fmt"

type Node struct {
	Row int
	Col int
	val int
}

// 情景
//如 11 * 11 棋盘上，只有 1,2,1 / 2,3,2 (row,col,value)两个子
func main() {
	// 1. 创建一个原始数组 1 黑棋 2白棋
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	// 棋盘
	for _, values := range chessMap {
		for _, v2 := range values {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	// 2.转化成稀疏数组
	// 1. 遍历chessMap 如果有一个元素值不为0，创建一个node结构体
	// 2. 将其放入到对应的切片中
	var sparseArr []Node
	// 记录稀疏数组的规模
	sparseArr = append(sparseArr, Node{
		Row: 11,
		Col: 11,
		val: 0,
	})
	for row, v := range chessMap {
		for col, v2 := range v {
			if v2 != 0 {
				//创建节点
				sparseArr = append(sparseArr, Node{
					Row: row,
					Col: col,
					val: v2,
				})
			}
		}
	}
	// 输出稀疏数组
	for i, node := range sparseArr {
		fmt.Printf("%d:%d:%d:%d \n", i, node.Row, node.Col, node.val)
	}
}
