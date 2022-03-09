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

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if chessBoard[i][j] != 0 {
				dataNode := DataNode{
					row: i,
					col: j,
					val: chessBoard[i][j],
				}
				count++
				dataNodes = append(dataNodes, dataNode)
			}
		}

	}
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
