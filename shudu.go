package main

import (
	"fmt"
	"sort"
	"time"
)

const Row = 9
const Column = 9

type ShuDu [Row][Column]byte

func (s ShuDu) String() string {
	var str string
	for _, columns := range s {
		for column, value := range columns {
			str += fmt.Sprintf("%d", value)
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

// 获取数独某个位置可能的值
func (s ShuDu) getPossibleValueList(row, column int) (result []byte) {
	valueMap := make(map[byte]struct{})
	// 列
	for i := 0; i < 9; i++ {
		if s[row][i] != 0 {
			valueMap[s[row][i]] = struct{}{}
		}
	}
	// 行
	for i := 0; i < 0; i++ {
		if s[i][column] != 0 {
			valueMap[s[i][column]] = struct{}{}
		}
	}

	// 方块
	blockRow := row / 3
	blockColumn := column / 3

	for i := blockRow * 3; i < blockRow*3+3; i++ {
		for j := blockColumn * 3; j < blockColumn*3+3; j++ {
			if s[i][j] != 0 {
				valueMap[s[i][j]] = struct{}{}
			}
		}
	}

	result = make([]byte, 0, 9-len(valueMap))
	for value := 1; value <= 9; value++ {
		if _, ok := valueMap[byte(value)]; !ok {
			result = append(result, byte(value))
		}
	}

	return
}

// 获取数独各个未知位置可能的值
func (s ShuDu) getPositionPossibleValue() (result PossibleValueList) {
	for row, columns := range s {
		for column, value := range columns {
			if value == 0 {
				result = append(result, &PossibleValue{row: row, column: column, possibleValueList: s.getPossibleValueList(row, column)})
			}
		}
	}

	return
}

// 新填入一个值的时候检查是否合法
func (s ShuDu) CheckWhenAdd(row, column int, value byte) bool {
	for i := 0; i < 9; i++ {
		if i == column {
			continue
		}

		if s[row][i] == value {
			return false
		}
	}
	// 行
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		if s[i][column] == value {
			return false
		}
	}

	// 方块
	blockRow := row / 3
	blockColumn := column / 3

	for i := blockRow * 3; i < blockRow*3+3; i++ {
		for j := blockColumn * 3; j < blockColumn*3+3; j++ {
			if i == row && j == column {
				continue
			}
			if s[i][j] == value {
				return false
			}
		}
	}

	return true
}

func (s ShuDu) Clone() ShuDu {
	newShuDu := ShuDu{}
	for row, columns := range s {
		for column := range columns {
			newShuDu[row][column] = s[row][column]
		}
	}
	return newShuDu
}

func (s *ShuDu) Clear(p PossibleValueList, tryMap map[*PossibleValue]int, index int) {
	for i, possibleValue := range p {
		if i >= index {
			s[possibleValue.row][possibleValue.column] = 0
		}

		if i > index {
			tryMap[possibleValue] = -1
		}
	}
}

func (s *ShuDu) Calculate(onlyFindOneResult bool) (result []ShuDu) {
	// 首先获取数独初始待填充字可能值
	possibleValueList := s.getPositionPossibleValue()

	if len(possibleValueList) == 0 {
		return
	}

	tryPossibleValueMap := make(map[*PossibleValue]int, len(possibleValueList))
	for _, v := range possibleValueList {
		tryPossibleValueMap[v] = -1
	}
	// 排序 可能值数量小的在前面
	sort.Sort(possibleValueList)
	tryUnknownValueIndex := 0
try:
	unknownValue := possibleValueList[tryUnknownValueIndex]

	s.Clear(possibleValueList, tryPossibleValueMap, tryUnknownValueIndex)
	for i, value := range unknownValue.possibleValueList {
		if i <= tryPossibleValueMap[unknownValue] {
			// 已经搜索过
			continue
		}

		if s.CheckWhenAdd(unknownValue.row, unknownValue.column, value) {
			//backFlag = false
			tryPossibleValueMap[unknownValue] = i
			s[unknownValue.row][unknownValue.column] = value
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
	// 这种情况没有解 回退
	for tryUnknownValueIndex > 0 {
		tryUnknownValueIndex--
		// 让上一个格子尝试其他可能数字
		if tryPossibleValueMap[possibleValueList[tryUnknownValueIndex]] < len(possibleValueList[tryUnknownValueIndex].possibleValueList)-1 {
			goto try
		}
	}

	return
}

func main() {
	shudu := ShuDu{
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

	fmt.Print("初始化数独:\n", shudu)
	timeNow := time.Now()
	results := shudu.Calculate(true)
	t1 := time.Now().Sub(timeNow) / time.Millisecond
	fmt.Printf("耗时:%dms\n", t1)
	for _, result := range results {
		fmt.Print(result)
	}

}
