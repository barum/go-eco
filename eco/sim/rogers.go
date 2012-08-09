// Rogers - Tanimoto similarity matrix
// Rogers & Tanimoto (1960), Gower & Legendre (1986)

package sim

import (
	"code.google.com/p/go-eco/eco/aux"
)

// Rogers - Tanimoto similarity matrix
// Legendre & Legendre (1998): 255, eq. 7.2 (S1 index)
func RogersTanimotoBool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c, d float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, d = aux.GetABCD(data, i, j)
			v := (a + d) / (a + 2*(b+c) + d)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
