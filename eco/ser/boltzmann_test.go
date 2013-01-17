package ser

import (
	"fmt"
	"testing"
)

func TestBoltzmann(t *testing.T) {
	var seed int64
	n := 10
	a := [][]int{
		// perfect Robinson matrix
		{81, 64, 49, 17, 16, 10, 9, 4, 3, 0},
		{64, 81, 64, 49, 17, 16, 10, 9, 4, 3},
		{49, 64, 81, 64, 49, 17, 16, 10, 9, 4},
		{17, 49, 64, 81, 64, 49, 17, 16, 10, 9},
		{16, 17, 49, 64, 81, 64, 49, 17, 16, 10},
		{10, 16, 17, 49, 64, 81, 64, 49, 17, 16},
		{9, 10, 16, 17, 49, 64, 81, 64, 49, 17},
		{4, 9, 10, 16, 17, 49, 64, 81, 64, 49},
		{3, 4, 9, 10, 16, 17, 49, 64, 81, 64},
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},
	}
	energyFn := "Psi3"
	iter := 100000
	permute := true
	rowPerm := make([]int, n)
	for i := 0; i < n; i++ {
		rowPerm[i] = i
	}
	colPerm := make([]int, n)
	for i := 0; i < n; i++ {
		colPerm[i] = i
	}
	seed = 1

	e := Boltzmann(n, n, a, rowPerm, colPerm, energyFn, iter, permute, seed)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", rowPerm[i]+1)
	}
	fmt.Println("Energy: ", e)
}

// should result:
// 2   1   3   4   6   5   8   7   9  10

/*
	// perfect Anti-Robinson matrix
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},
		{3, 0, 3, 4, 9, 10, 16, 17, 49, 64},
		{4, 3, 0, 3, 4, 9, 10, 16, 17, 49},
		{9, 4, 4, 0, 3, 4, 9, 10, 16, 17},
		{10, 9, 4, 3, 0, 3, 4, 9, 10, 16},
		{16, 10, 9, 4, 3, 0, 3, 4, 9, 10},
		{17, 16, 10, 9, 4, 3, 0, 3, 4, 9},
		{49, 17, 16, 10, 9, 4, 3, 0, 3, 4},
		{64, 49, 17, 16, 10, 9, 4, 3, 0, 3},
		{81, 64, 49, 17, 16, 10, 9, 4, 3, 0},

	// not-perfect Anti-Robinson matrix
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


	// perfect Robinson matrix
		{81,64, 49, 17, 16, 10, 9, 4, 3, 0},
		{64, 81,64, 49, 17, 16, 10, 9, 4, 3},
		{49, 64, 81,64, 49, 17, 16, 10, 9, 4},
		{17, 49, 64, 81,64, 49, 17, 16, 10, 9},
		{16, 17, 49, 64, 81,64, 49, 17, 16, 10},
		{10, 16, 17, 49, 64, 81,64, 49, 17, 16},
		{9, 10, 16, 17, 49, 64, 81,64, 49, 17},
		{4, 9, 10, 16, 17, 49, 64, 81,64, 49},
		{3, 4, 9, 10, 16, 17, 49, 64, 81,64},
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},

func TestPsi2(t *testing.T) {
	n := 10
	// perfect Robinson matrix
	a := [][]int{
		{81, 64, 49, 17, 16, 10, 9, 4, 3, 0},
		{64, 81, 64, 49, 17, 16, 10, 9, 4, 3},
		{49, 64, 81, 64, 49, 17, 16, 10, 9, 4},
		{17, 49, 64, 81, 64, 49, 17, 16, 10, 9},
		{16, 17, 49, 64, 81, 64, 49, 17, 16, 10},
		{10, 16, 17, 49, 64, 81, 64, 49, 17, 16},
		{9, 10, 16, 17, 49, 64, 81, 64, 49, 17},
		{4, 9, 10, 16, 17, 49, 64, 81, 64, 49},
		{3, 4, 9, 10, 16, 17, 49, 64, 81, 64},
		{0, 3, 4, 9, 10, 16, 17, 49, 64, 81},
	}

	rowPerm := make([]int, n)
	for i := 0; i < n; i++ {
		rowPerm[i] = i
	}
	colPerm := make([]int, n)
	for i := 0; i < n; i++ {
		colPerm[i] = i
	}

	e := Psi(a, n, n, rowPerm, colPerm)
	fmt.Println("Energy: ", e)
}
*/
