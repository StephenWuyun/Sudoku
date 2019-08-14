package sudoku


import (
	"fmt"
	"testing"
	"time"
)

var srcRawData = [9][9]int{
	[9]int{0, 0, 0, 0, 0, 0, 0, 0, 6},
	[9]int{3, 0, 2, 0, 9, 1, 7, 0, 0},
	[9]int{0, 1, 8, 2, 0, 0, 0, 9, 0},
	[9]int{0, 0, 0, 0, 8, 0, 0, 0, 0},
	[9]int{0, 8, 0, 4, 3, 0, 0, 0, 0},
	[9]int{0, 7, 9, 1, 2, 0, 6, 3, 8},
	[9]int{0, 0, 1, 0, 7, 0, 0, 4, 2},
	[9]int{9, 2, 3, 0, 0, 0, 0, 6, 0},
	[9]int{8, 4, 7, 0, 0, 0, 0, 0, 9},
}

//var srcRawData = [9][9]int{
//	[9]int{0, 0, 0, 0, 0, 0, 0, 0, 6},
//	[9]int{3, 0, 2, 0, 0, 1, 7, 0, 0},
//	[9]int{0, 0, 8, 2, 0, 0, 0, 9, 0},
//	[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
//	[9]int{0, 0, 0, 4, 0, 0, 0, 0, 0},
//	[9]int{0, 0, 9, 1, 0, 0, 6, 3, 8},
//	[9]int{0, 0, 1, 0, 0, 0, 0, 4, 2},
//	[9]int{9, 0, 3, 0, 0, 0, 0, 6, 0},
//	[9]int{8, 0, 7, 0, 0, 0, 0, 0, 9},
//}

func printData(data [9][9]int) {
	for _, intArray := range data {
		for _, v := range intArray {
			if v == 0 {
				fmt.Printf(". ")
			}else {
				fmt.Printf("%d ", v)
				//fmt.Printf("\033[1;31m%d\033[0m ", v)
			}

		}
		fmt.Println()
	}

	fmt.Printf("\n------------------------------\n")
}

// 验证九宫格是否合法
func validateRect(sd [9][9]int, r, c int) bool {
	var t [10]int
	startX := r/3 * 3
	startY := c/3 * 3

	hasNotSetField := false
	for i:=startX; i<startX+3; i++ {
		for j:=startY; j<startY+3; j++ {
			if sd[i][j] == 0 {
				hasNotSetField = true
			}else {
				// 发现重复
				if t[ sd[i][j] ] != 0 {
					return false
				}

				t[ sd[i][j] ] = 1

			}
		}
	}

	// 九宫格的值还没设置完
	if hasNotSetField {
		return true
	}

	for k:=1; k<=9; k++ {
		if t[k] == 0 {
			return false
		}
	}

	return true

}

// 用于验证最终结果的正确性（非必须）
func validateResult(sd [9][9]int) bool {
	// 验证行
	for i:=0; i<9; i++ {
		var t [10]int
		for j:=0; j<9; j++ {
			t[ sd[i][j] ] = 1
		}

		for k:=1; k<=9; k++ {
			if t[k] == 0 {
				return false
			}

		}
	}

	// 验证列
	for i:=0; i<9; i++ {
		var t [10]int
		for j:=0; j<9; j++ {
			t[ sd[j][i] ] = 1
		}
		for k:=1; k<=9; k++ {
			if t[k] == 0 {
				return false
			}
		}
	}

	resutl := true
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			resutl = resutl && validateRect(sd, i*3, j*3)
		}
	}

	return resutl
}

func validateResultWithMap(sd [9][9]int) bool {
	// 验证行
	for i:=0; i<9; i++ {
		tmpMap := make(map[int]int)
		for j:=0; j<9; j++ {
			v := sd[i][j]
			if _, ok := tmpMap[v]; ok {
				return false
			}else {
				tmpMap[v] = 1
			}
		}

	}

	// 验证列
	for i:=0; i<9; i++ {
		var t [10]int
		for j:=0; j<9; j++ {
			t[ sd[j][i] ] = 1
		}
		for k:=1; k<=9; k++ {
			if t[k] == 0 {
				return false
			}
		}
	}

	resutl := true
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			resutl = resutl && validateRect(sd, i*3, j*3)
		}
	}

	return resutl
}

type Node struct {
	x int
	y int
	possibleValue [9]int
	count int
	curIndex int
}

func generateNodeList(srcData [9][9]int, rowMap [10][10]int, columnMap [10][10]int, gridMap [3][3][10]int) (nodeCount int, nodeList []*Node) {

	nodeList = make([]*Node, 9*9)
	nodeCount = 0
	for i:=0; i<9; i++ {
		for j := 0; j < 9; j++ {
			v := srcData[i][j]
			if v == 0 {
				n := &Node{
					x:i,
					y:j,
					curIndex:-1,
				}
				count := 0
				for k:=1; k<=9; k++ {
					if rowMap[i][k] == 0 && columnMap[j][k]==0 && gridMap[i/3][j/3][k]==0 {
						n.possibleValue[count] = k
						count++
					}
				}
				n.count = count
				nodeList[nodeCount] = n
				nodeCount++
			}
		}
	}



	return
}

//func isValidStep(rowMap [10][10]int, columnMap [10][10]int, gridMap [3][3][10]int, r,c,i,j,k int) bool {
//	if rowMap[r][k] == 0 && columnMap[c][k] == 0 && gridMap[i][j][k]==0 {
//		return true
//	}
//
//	return false
//}

/*
思路：
1)对现有的数据进行分析，生成映射关系
2)计算出每个待填位置的可能值
3)对可能值的数量进行排序，数量少的优先遍历(优化）
4)循环判断各种可能值的组合，选出满足规则的
 */
func Arrange(srcData [9][9]int)  [9][9]int {
	var selected [9][9]int
	var rowMap [10][10]int
	var columnMap [10][10]int
	var gridMap [3][3][10]int
	//execCount := 0

	// 初始化映射
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			v := srcData[i][j]
			rowMap[i][ v ] = 1
			columnMap[j][ v ] = 1
			gridMap[i/3][j/3][v] = 1
			selected[i][j] = v
		}
	}

	nodeCount, nodeList := generateNodeList(srcData, rowMap, columnMap, gridMap)

	// 按取值数量对nodelist排序，后续优先计算数量少的
	for i := 0; i < nodeCount-1; i++ {
		for j := 0; j < nodeCount- 1 - i; j++ {
			if nodeList[j].count > nodeList[j+1].count {
				nodeList[j], nodeList[j+1] = nodeList[j+1], nodeList[j]
			}
		}
	}

	index := 0

	for index < nodeCount {
		n := nodeList[index]
		//k := n.curIndex
		r := n.x
		c := n.y
		m:=n.curIndex + 1
		gridR := r/3
		gridC := c/3

		for ; m<n.count; m++ {
			//execCount++
			k := n.possibleValue[m]
			if rowMap[r][k] == 0 && columnMap[c][k] == 0 && gridMap[gridR][gridC][k]==0{
				//fmt.Printf("[%d, %d] set value:%d\n", r, c, k)
				selected[r][c] = k
				rowMap[r][k] = 1
				columnMap[c][k] = 1
				gridMap[gridR][gridC][k] = 1
				n.curIndex = m
				break

			}
		}

		// not found
		if m >= n.count {
			if index == 0 {
				fmt.Printf("退出\n")
				break
			}

			// 回退之前的选择
			selected[r][c] = 0
			if n.curIndex != -1 {
				cv := n.possibleValue[ n.curIndex ]
				rowMap[r][ cv ] = 0
				columnMap[c][ cv ] = 0
				gridMap[gridR][gridC][ cv ] = 0
				n.curIndex = -1
			}

			index--
			nPrev := nodeList[index]
			if nPrev.curIndex != -1 {
				cv := nPrev.possibleValue[ nPrev.curIndex ]
				rowMap[nPrev.x][ cv ] = 0
				columnMap[nPrev.y][ cv ] = 0
				gridMap[nPrev.x/3][nPrev.y/3][ cv ] = 0
			}

		}else {
			index++
		}
	}

	//fmt.Printf("count:%d\n", execCount)

	return selected
}


func Test100(t *testing.T) {

	beginTime := time.Now().UnixNano()

	selected := Arrange(srcRawData)

	endTime := time.Now().UnixNano()

	fmt.Printf("原数据:\n")
	printData( srcRawData )

	fmt.Printf("\n-------------result------------\n\n")

	//printData(selected)

	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			v := srcRawData[i][j]
			v2 := selected[i][j]
			if v == 0 {
				fmt.Printf("\033[1;31m%d\033[0m ", v2)
			}else {
				fmt.Printf("%d ", v)
			}
		}
		fmt.Println()
	}

	// 验证结果的正确性，不是必要步骤
	fmt.Printf("\n验证结果:%v\n", validateResult(selected))

	fmt.Printf("\n用时:%d us\n\n", (endTime-beginTime)/1000)

}

func verifyBat() {
	selected := Arrange(srcRawData)
	r := validateResult(selected)
	if !r {
		fmt.Printf("err\n")
	}
}

func Benchmark101(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Arrange(srcRawData)
		//verifyBat()
	}
}


