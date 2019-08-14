package sudoku

//import (
//	"fmt"
//	"sort"
//	"time"
//)
//
//package  main
//
//import  (
//"fmt"
//"sort"
//"time"
//)
//
////var (
////	TogglePossibleValueStateCost time.Duration
////	ClearCost time.Duration
////)
//
//const  Row  =  9
//const  Column  =  9
//
//type  Sudoku  [Row][Column]byte
//
//func  (s  Sudoku)  String()  string  {
//	var  str  string
//	for  _,  columns  :=  range  s  {
//		for  column,  value  :=  range  columns  {
//			if value != 0 {
//				str  +=  fmt.Sprintf("%d",  value)
//			} else {
//				str += "*"
//			}
//			if  column  ==  Column-1  {
//				str  +=  "\n"
//			}  else  {
//				str  +=  "\t"
//			}
//		}
//	}
//	return  str
//}
//
//type  ValueItem  struct  {
//	value      byte
//	invalid  bool
//}
//
//type  PossibleValue  struct  {
//	row                              int
//	column                        int
//	possibleValueList  []*ValueItem
//}
//
//type  PossibleValueList  []*PossibleValue
//
//func  (p  PossibleValueList)  Swap(i,  j  int)  {
//	p[i],  p[j]  =  p[j],  p[i]
//}
//
//func  (p  PossibleValueList)  Len()  int  {
//	return  len(p)
//}
//
//func  (p  PossibleValueList)  Less(i,  j  int)  bool  {
//	return  len(p[i].possibleValueList)  <  len(p[j].possibleValueList)
//}
//
////  获取数独某个位置可能的值
//func  (s  Sudoku)  getPossibleValueList(row,  column  int)  (result  []*ValueItem)  {
//	valueMap  :=  make(map[byte]struct{})
//	//  列
//	for  i  :=  0;  i  <  9;  i++  {
//		if  s[row][i]  !=  0  {
//			valueMap[s[row][i]]  =  struct{}{}
//		}
//	}
//	//  行
//	for  i  :=  0;  i  <  9;  i++  {
//		if  s[i][column]  !=  0  {
//			valueMap[s[i][column]]  =  struct{}{}
//		}
//	}
//
//	//  方块
//	blockRow  :=  row  /  3
//	blockColumn  :=  column  /  3
//
//	for  i  :=  blockRow  *  3;  i  <  blockRow*3+3;  i++  {
//		for  j  :=  blockColumn  *  3;  j  <  blockColumn*3+3;  j++  {
//			if  s[i][j]  !=  0  {
//				valueMap[s[i][j]]  =  struct{}{}
//			}
//		}
//	}
//
//	result  =  make([]*ValueItem,  0,  9-len(valueMap))
//	for  value  :=  1;  value  <=  9;  value++  {
//		if  _,  ok  :=  valueMap[byte(value)];  !ok  {
//			result  =  append(result,  &ValueItem{value:  byte(value)})
//		}
//	}
//
//	return
//}
//
////  获取数独各个未知位置可能的值
//func  (s  Sudoku)  getPositionPossibleValue()  (result  PossibleValueList)  {
//	for  row,  columns  :=  range  s  {
//		for  column,  value  :=  range  columns  {
//			if  value  ==  0  {
//				result  =  append(result,  &PossibleValue{row:  row,   column:  column,  possibleValueList:  s.getPossibleValueList(row,  column)})
//			}
//		}
//	}
//
//	return
//}
//
////  在一行或者一列或者一个方块  已经有值  则设置后面的格子不再能填充这个数字
//func  (s  Sudoku)TogglePossibleValueState(nextTryUnknownValueIndex,  row,  column  int,   value  byte,  possibleValueList  PossibleValueList)  {
//	//beginTime := time.Now()
//	for  index  :=  nextTryUnknownValueIndex;  index  <  len(possibleValueList);   index++  {
//		possibleValue  :=  possibleValueList[index]
//		possibleValueRow  :=  possibleValue.row
//		possibleValueColumn  :=  possibleValue.column
//		if possibleValueRow == row || possibleValueColumn == column ||  (possibleValueRow/3 == row/3 && possibleValueColumn/3 == column/3) {
//			for _, internalValue := range possibleValue.possibleValueList {
//				if internalValue.value == value {
//					internalValue.invalid = true
//					break
//				}
//			}
//		}
//
//	}
//	//TogglePossibleValueStateCost += time.Now().Sub(beginTime)
//}
//
//func (s Sudoku) Clone() Sudoku {
//	newShuDu := Sudoku{}
//	for row, columns := range s {
//		for column := range columns {
//			newShuDu[row][column] = s[row][column]
//		}
//	}
//	return newShuDu
//}
//
//func (s *Sudoku) Clear(p PossibleValueList, tryMap map[*PossibleValue]int, index int, back  bool) {
//	//begin := time.Now()
//	for i := index; i < len(p); i++ {
//		possibleValue := p[i]
//		s[possibleValue.row][possibleValue.column] = 0
//		if i > index {
//			tryMap[possibleValue] = -1
//			if back {
//				// 回退
//				possibleValue.possibleValueList = s.getPossibleValueList (possibleValue.row, possibleValue.column)
//			}
//		}
//	}
//	//ClearCost += time.Now().Sub(begin)
//}
//
//func (s *Sudoku) Calculate(onlyFindOneResult bool) (result []Sudoku) {
//	// 首先获取数独初始待填充字可能值
//	possibleValueList := s.getPositionPossibleValue()
//
//	if len(possibleValueList) == 0 {
//		return []Sudoku{s.Clone()}
//	}
//
//	tryPossibleValueMap := make(map[*PossibleValue]int, len(possibleValueList))
//	for _, v := range possibleValueList {
//		tryPossibleValueMap[v] = -1
//	}
//	// 排序 可能值数量小的在前面
//	sort.Sort(possibleValueList)
//
//	tryUnknownValueIndex := 0
//	backward := false
//	counter := 0
//try:
//	counter++
//	unknownValue := possibleValueList[tryUnknownValueIndex]
//	s.Clear(possibleValueList, tryPossibleValueMap, tryUnknownValueIndex, backward)
//	backward = false
//	for i, value := range unknownValue.possibleValueList {
//		if i <= tryPossibleValueMap[unknownValue] {
//			// 已经搜索过
//			continue
//		}
//
//		if !value.invalid {
//			tryPossibleValueMap[unknownValue] = i
//			s[unknownValue.row][unknownValue.column] = value.value
//			if tryUnknownValueIndex == len(possibleValueList)-1 {
//				// 已经填充完所有格子 添加到解空间
//				result = append(result, s.Clone())
//				if onlyFindOneResult {
//					fmt.Println("counter:", counter)
//					return
//				}
//				// 回退 尝试找另外的解
//				goto back
//			} else {
//				// 继续尝试填充下一个格子
//				tryUnknownValueIndex++
//				s.TogglePossibleValueState(tryUnknownValueIndex,  unknownValue.row, unknownValue.column, value.value, possibleValueList)
//				goto try
//			}
//		}
//	}
//
//back:
//	backward = true
//	// 这种情况没有解 回退
//	for tryUnknownValueIndex > 0 {
//		tryUnknownValueIndex--
//		// 让上一个格子尝试其他可能数字
//		if tryPossibleValueMap[possibleValueList[tryUnknownValueIndex]] < len (possibleValueList[tryUnknownValueIndex].possibleValueList)-1 {
//			goto try
//		}
//	}
//
//	return
//}
//
//func main() {
//	sudoku := Sudoku{
//		[Column]byte{0, 0, 0, 0, 0, 0, 0, 0, 6},
//		[Column]byte{3, 0, 2, 0, 9, 1, 7, 0, 0},
//		[Column]byte{0, 1, 8, 2, 0, 0, 0, 9, 0},
//		[Column]byte{0, 0, 0, 0, 8, 0, 0, 0, 0},
//		[Column]byte{0, 8, 0, 4, 3, 0, 0, 0, 0},
//		[Column]byte{0, 7, 9, 1, 2, 0, 6, 3, 8},
//		[Column]byte{0, 0, 1, 0, 7, 0, 0, 4, 2},
//		[Column]byte{9, 2, 3, 0, 0, 0, 0, 6, 0},
//		[Column]byte{8, 4, 7, 0, 0, 0, 0, 0, 9},
//	}
//
//	fmt.Print("初始化数独:\n", sudoku)
//	timeNow := time.Now()
//	results := sudoku.Calculate(true)
//	t1 := time.Now().Sub(timeNow) / time.Microsecond
//	fmt.Printf("耗时:%dus\n", t1)
//	//fmt.Printf("TogglePossibleValueStateCost 耗时:%dus\n", TogglePossibleValueStateCost/ time.Microsecond)
//	//fmt.Printf("ClearCost 耗时:%dus\n", ClearCost/ time.Microsecond)
//
//	for _, result := range results {
//		fmt.Print(result)
//	}
//
//}
