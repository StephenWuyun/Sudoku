/*************************************************************************
    > File Name: sudoku.go
    > Author: moses
    > Mail: 354296290@qq.com 
    > Created Time: Wed Aug 14 09:55:44 2019
 ************************************************************************/
package main

import (
	"fmt"
	"sort"
	"time"
)

const Row = 9
const Column = 9

type Sudoku [Row][Column]byte

func (s Sudoku) String() string {
	var str string
	for _, columns := range s {
		for column, value := range columns {
			if value != 0 {
				str += fmt.Sprintf("%d", value)
			} else {
				str += "*"
			}
			if column == Column-1 {
				str += "\n"
			} else {
				str += "\t"
			}
		}
	}
	return str
}

type PossibleValue struct {
	row               int
	column            int
	possibleValueList []byte
	index             int
}

type PossibleValueList []*PossibleValue

func (p PossibleValueList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PossibleValueList) Len() int {
	return len(p)
}

func (p PossibleValueList) Less(i, j int) bool {
	return len(p[i].possibleValueList) < len(p[j].possibleValueList)
}

//  获取数独某个位置可能的值
func (s Sudoku) getPossibleValueList(row, column int) (result []byte) {
	var valueList [10]bool
	//  列
	for i := 0; i < 9; i++ {
		valueList[s[row][i]] = true
	}
	//  行
	for i := 0; i < 9; i++ {
		valueList[s[i][column]] = true

	}

	//  方块
	blockRow := row / 3
	blockColumn := column / 3

	for i := blockRow * 3; i < blockRow*3+3; i++ {
		for j := blockColumn * 3; j < blockColumn*3+3; j++ {
			valueList[s[i][j]] = true
		}
	}

	for value := 1; value <= 9; value++ {
		if !valueList[value] {
			result = append(result, byte(value))
		}
	}

	return
}

//  获取数独各个未知位置可能的值
func (s Sudoku) getPositionPossibleValue() (result PossibleValueList, rowMap [9][10]bool, columnMap [9][10]bool, blockMap [9][10]bool) {
	for row, columns := range s {
		for column, value := range columns {
			if value == 0 {
				result = append(result, &PossibleValue{index:-1, row: row, column: column, possibleValueList: s.getPossibleValueList(row, column) })
			} else {
				rowMap[row][value] = true
				columnMap[column][value] = true
				blockMap[row/3*3+column/3][value] = true
			}
		}
	}

	return
}

// 克隆数独
func (s Sudoku) Clone() Sudoku {
	newShuDu := Sudoku{}
	for row, columns := range s {
		for column := range columns {
			newShuDu[row][column] = s[row][column]
		}
	}
	return newShuDu
}

// 设置数独格子值 并处理标记位 --clear 清除值
func (s *Sudoku)SetFlag(row, column int, value byte, rowMap, columnMap, blockMap *[9][10]bool, clear bool) {
	if clear {
		s[row][column] = 0
	} else  {
		s[row][column] = value
	}
	flag := !clear
	rowMap[row][value] = flag
	columnMap[column][value] = flag
	blockMap[row/3*3+column/3][value] = flag
}

func (s *Sudoku) Calculate(onlyFindOneResult bool) (result []Sudoku) {
	// 首先获取数独初始待填充字可能值
	possibleValueList, rowMap, columnMap, blockMap := s.getPositionPossibleValue()

	if len(possibleValueList) == 0 {
		return []Sudoku{s.Clone()}
	}

	// 排序 可能值数量小的在前面
	sort.Sort(possibleValueList)
	//fmt.Println(possibleValueList)

	tryUnknownValueIndex := 0
try:
	// 尝试设置空格数字
	unknownValue := possibleValueList[tryUnknownValueIndex]
	// i从index+1开始 跳过之前已经尝试的值
	for i := unknownValue.index + 1; i < len(unknownValue.possibleValueList); i++ {
		value := unknownValue.possibleValueList[i]
		// 根据标记位 检查当前位置能否插入value
		if !rowMap[unknownValue.row][value] && !columnMap[unknownValue.column][value] && !blockMap[unknownValue.row/3*3+unknownValue.column/3][value] {
			unknownValue.index = i
			s.SetFlag(unknownValue.row, unknownValue.column, value, &rowMap, &columnMap, &blockMap, false)
			if tryUnknownValueIndex == len(possibleValueList)-1 {
				// 已经填充完所有格子 添加到解空间
				result = append(result, s.Clone())
				if onlyFindOneResult {
					return
				}
				// 回退 尝试找另外的解
				goto back
			} else {
				// 继续尝试填充下一个格子
				tryUnknownValueIndex++
				goto try
			}
		}
	}

back:
	// 回退
	for tryUnknownValueIndex > 0 {
		unknownValue := possibleValueList[tryUnknownValueIndex]
		s.SetFlag(unknownValue.row, unknownValue.column, s[unknownValue.row][unknownValue.column], &rowMap, &columnMap, &blockMap, true)
		unknownValue.index = -1
		tryUnknownValueIndex--
		// 回退格子 尝试其他可能数字
		unknownValue = possibleValueList[tryUnknownValueIndex]
		if unknownValue.index < len(unknownValue.possibleValueList)-1 {
			s.SetFlag(unknownValue.row, unknownValue.column, s[unknownValue.row][unknownValue.column], &rowMap, &columnMap, &blockMap, true)
			goto try
		}
	}

	return
}

func main() {
	sudoku := Sudoku{
		[Column]byte{0, 0, 0, 0, 0, 0, 0, 0, 6},
		[Column]byte{3, 0, 2, 0, 9, 1, 7, 0, 0},
		[Column]byte{0, 1, 8, 2, 0, 0, 0, 9, 0},
		[Column]byte{0, 0, 0, 0, 8, 0, 0, 0, 0},
		[Column]byte{0, 8, 0, 4, 3, 0, 0, 0, 0},
		[Column]byte{0, 7, 9, 1, 2, 0, 6, 3, 8},
		[Column]byte{0, 0, 1, 0, 7, 0, 0, 4, 2},
		[Column]byte{9, 2, 3, 0, 0, 0, 0, 6, 0},
		[Column]byte{8, 4, 7, 0, 0, 0, 0, 0, 9},
	}

	fmt.Print("初始化数独:\n", sudoku)
	timeNow := time.Now()
	results := sudoku.Calculate(true)
	t1 := time.Now().Sub(timeNow) / time.Microsecond
	fmt.Printf("耗时:%dus\n", t1)
	for _, result := range results {
		fmt.Print(result)
	}

}

