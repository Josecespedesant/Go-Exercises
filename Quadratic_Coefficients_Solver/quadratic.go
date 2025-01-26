package quadratic

const a = 1

func SolveQuadraticCoefficients(x1 int, x2 int) [3]int {
	// For an equation ax^2+bx+c=0, x1 and x2 are two roots, x1 + x2 = -b/a, and x1 * x2 = b/a
	b := -a * (x1 + x2)
	c := a * (x1 * x2)
	coefficients := [3]int{a, b, c}
	return coefficients
}
