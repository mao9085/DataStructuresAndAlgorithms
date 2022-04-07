package main

import (
	"fmt"
)

func main() {
	//  初始化二维数组棋盘
	chessBoard := TwoDimensionalArray()
	// 二维数组转稀疏数组
	dataNodes := ArrayToSparsearry(chessBoard)
	for key, val := range dataNodes {
		fmt.Println(key)
		fmt.Println(val)
	}

	arrays := SparsearryToArray(dataNodes)
	fmt.Println(arrays)

}

type DataNode struct {
	row int
	col int
	val int
}

//  初始化二维数组
func TwoDimensionalArray() [11][11]int {

	var chessBoard [11][11]int
	chessBoard[1][2] = 1
	chessBoard[2][3] = 1
	fmt.Println("↓原始数据↓")
	for _, v := range chessBoard {
		for _, v2 := range v {
			// %v 相应值的默认格式，格式化输出
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}
	return chessBoard
}

//  二维数组转稀疏数组
func ArrayToSparsearry(chessBoard [11][11]int) []DataNode {
	len1 := len(chessBoard)
	len2 := len(chessBoard[0])

	var dataNodes []DataNode

	dataNode := DataNode{
		row: len1,
		col: len2,
		val: 0,
	}

	var count int

	dataNodes = append(dataNodes, dataNode)
	dataNodes[0] = dataNode

	for i, v1 := range chessBoard {
		for j, v2 := range v1 {
			// 找出二维矩阵中不为 0 的数值，并将其添加到稀疏数组中
			if chessBoard[i][j] != 0 {
				// 结构体赋值
				dataNode := DataNode{i, j, v2}
				// 添加棋盘棋子信息到 稀疏数组中
				dataNodes = append(dataNodes, dataNode)
				// 为稀疏数组第一行的 value 值计数
				count++
			}
		}
	}

	//for i := 0; i < len1; i++ {
	//	for j := 0; j < len2; j++ {
	//		// 找出二维矩阵中不为 0 的数值，并将其添加到稀疏数组中
	//		if chessBoard[i][j] != 0 {
	//			// 构建稀疏数组某行的值 dataNode
	//			dataNode := DataNode{i,j,chessBoard[i][j]}
	//			// 添加棋盘棋子信息到 稀疏数组中
	//			dataNodes = append(dataNodes, dataNode)
	//			// 为稀疏数组第一行的 value 值计数
	//			count++
	//		}
	//	}
	//
	//}

	dataNodes[0].val = count

	return dataNodes
}

//  稀疏数组转二维数组
func SparsearryToArray(dataNodes []DataNode) [11][11]int {
	//dataNode := dataNodes[0]
	//row := dataNode.row
	//col :=dataNode.col

	var arrays [11][11]int
	//arrays2 := [11][11]int{}

	for key, value := range dataNodes {
		if key != 0 {
			arrays[value.row][value.col] = value.val
		}
	}

	//  二维切片初始化
	//array := make([][]int , row)
	//for i := range array{
	//	array[i] = make ([]int ,col)
	//}
	//  二维切片赋值不会
	//for _,value := range dataNodes{
	//	array[value.row][value.col] = value.val
	//}

	return arrays
}
