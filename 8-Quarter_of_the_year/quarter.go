package quarter

/*
EXERCISE:
Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.
For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November), is part of the fourth quarter.
Constraint:
1 <= month <= 12

REFERENCE: https://www.codewars.com/kata/5ce9c1000bab0b001134f5af
KYU: 8
*/

func GetQuarter(month int) int {
	if month < 1 || month > 12 {
		return -1
	}
	quarter := 0
	if month <= 3 {
		quarter = 1
	}
	if month > 3 && month <= 6 {
		quarter = 2
	}
	if month > 6 && month <= 9 {
		quarter = 3
	}
	if month > 9 && month <= 12 {
		quarter = 4
	}
	return quarter
}
