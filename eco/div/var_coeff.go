// Coefficient of variation

package div

import (
	"code.google.com/p/go-eco/eco/aux"
	"math"
)

// Coefficient of variation, for population
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq
func VarCoeff_D(data *aux.Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		// calculate mean and variance
		meanX := 0.0
		s := 0.0 // number of species
		m2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				delta := x - meanX
				meanX += delta / s
				if s > 1 {
					m2 += delta * (x - meanX)
				}
			}
		}
		varX := m2 / s
		v := math.Sqrt((s-1)*varX/s) / meanX
		out.Set(i, v)
	}
	return out
}

// Coefficient of variation, for sample
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// Marshall & Olkin: Inequalities: Theory of Majorization and Its Applications, New York 1979 (Academic Press).
// Algorithm inspired by R:ineq
func VarCoeffSmp_D(data *aux.Matrix) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		// calculate mean and variance
		meanX := 0.0
		s := 0.0 // number of species
		m2 := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				delta := x - meanX
				meanX += delta / s
				if s > 1 {
					m2 += delta * (x - meanX)
				}
			}
		}
		varX := m2 / (s - 1)
		v := math.Sqrt((s-1)*varX/s) / meanX
		out.Set(i, v)
	}
	return out
}

// Squared coefficient of variation, for population
func VarCoeffSq_D(data *aux.Matrix) *Vector {
	rows := data.R
	out := NewVector(rows)
	vc := VarCoeff_D(data)

	for i := 0; i < rows; i++ {
		v := vc.Get(i)
		v *= v
		out.Set(i, v)
	}
	return out
}

// Squared coefficient of variation, for saample
func VarCoeffSqSmp_D(data *aux.Matrix) *Vector {
	rows := data.R
	out := NewVector(rows)
	vc := VarCoeffSmp_D(data)

	for i := 0; i < rows; i++ {
		v := vc.Get(i)
		v *= v
		out.Set(i, v)
	}
	return out
}
