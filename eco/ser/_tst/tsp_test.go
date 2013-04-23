package ser

import (
	//	"code.google.com/p/go-eco/eco/aux"
	"fmt"
	"testing"
)

func TestARBBw(t *testing.T) {
	n := 11
	a := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 5, 4, 9, 81, 16, 15, 49, 64, 20},
		{0, 5, 0, 3, 4, 9, 25, 90, 36, 49, 64},
		{0, 4, 3, 0, 4, 5, 9, 16, 12, 36, 49},
		{0, 9, 4, 4, 0, 6, 4, 9, 16, 25, 36},
		{0, 81, 9, 5, 6, 0, 4, 4, 9, 16, 25},
		{0, 16, 25, 9, 4, 4, 0, 5, 4, 9, 16},
		{0, 15, 90, 16, 9, 4, 5, 0, 3, 4, 9},
		{0, 49, 36, 12, 16, 9, 4, 3, 0, 5, 4},
		{0, 64, 49, 36, 25, 16, 9, 4, 5, 0, 4},
		{0, 20, 64, 49, 36, 25, 16, 9, 4, 4, 0},
	}
fmt.Println("###TEST###")

	p := MinimumTour(n, a)
	for i := 1; i < n; i++ {
		fmt.Printf("%d ", p[i]+1)
	}
}

/*
func TestARBBw(t *testing.T) {
	n := 10
	a := [][]int{
		{0, 5, 4, 9, 81, 16, 15, 49, 64, 20},
		{5, 0, 3, 4, 9, 25, 90, 36, 49, 64},
		{4, 3, 0, 4, 5, 9, 16, 12, 36, 49},
		{9, 4, 4, 0, 6, 4, 9, 16, 25, 36},
		{81, 9, 5, 6, 0, 4, 4, 9, 16, 25},
		{16, 25, 9, 4, 4, 0, 5, 4, 9, 16},
		{15, 90, 16, 9, 4, 5, 0, 3, 4, 9},
		{49, 36, 12, 16, 9, 4, 3, 0, 5, 4},
		{64, 49, 36, 25, 16, 9, 4, 5, 0, 4},
		{20, 64, 49, 36, 25, 16, 9, 4, 4, 0},
	}

	p := ARBBw(n, a, 1e-7)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", p[i]+1)
	}
}
*/

// should result:
// 2   1   3   4   6   5   8   7   9  10
