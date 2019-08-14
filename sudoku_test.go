package sudoku

import (
	"fmt"
	"testing"
	"time"
)



func Benchmark100(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		//
		//	fmt.Print("初始化数独:\n", sudoku)
		//	timeNow := time.Now()
		sudoku.Calculate(true)
		//	t1 := time.Now().Sub(timeNow) / time.Microsecond
		//	fmt.Printf("耗时:%dus\n", t1)
	}
}

func TestSudoku_Calculate(t *testing.T) {
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
	sudoku.Calculate(true)
	cost := time.Now().Sub(timeNow) / time.Microsecond
	fmt.Printf("耗时:%dus\n", cost)
	fmt.Println("counter:", counter)
	//fmt.Printf("CheckWhenAdd 耗时:%dus\n", CheckWhenAddCost/ time.Microsecond)
}
