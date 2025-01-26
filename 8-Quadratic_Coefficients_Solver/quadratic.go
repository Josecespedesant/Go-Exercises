package quadratic

/*
EXERCISE:
In this Kata you are expected to find the coefficients of quadratic equation of the given two roots (x1 and x2).
Equation will be the form of ax^2 + bx + c = 0
Return type is a Vector (tuple in Rust, Array in Ruby) containing coefficients of the equations in the order (a, b, c).
Since there are infinitely many solutions to this problem, we fix a = 1.
Remember, the roots can be written like (x-x1) * (x-x2) = 0

REFERENCE: https://www.codewars.com/kata/5d59576768ba810001f1f8d6
KYU: 8
*/

const a = 1

func SolveQuadraticCoefficients(x1 int, x2 int) [3]int {
	// For an equation ax^2+bx+c=0, x1 and x2 are two roots, x1 + x2 = -b/a, and x1 * x2 = b/a
	b := -a * (x1 + x2)
	c := a * (x1 * x2)
	coefficients := [3]int{a, b, c}
	return coefficients
}
