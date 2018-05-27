/*
 https://projecteuler.net/problem=3

 The prime factors of 13195 are 5, 7, 13 and 29.

 What is the largest prime factor of the number 600851475143

 Answer: 6857
*/

package go_euler

func Solution003() int {
	answer := 1
	point := 3
	divisor := 600851475143

	for divisor%2 == 0 {
		answer = 2
		divisor = divisor / 2
	}

	for divisor != 1 {
		for divisor%point == 0 {
			answer = point
			divisor = divisor / point
		}
		point += 2
	}
	return answer
}
