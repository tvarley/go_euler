/*
https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from
1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?

Answer: 232792560
*/

package go_euler

func Solution005() int {
	answer := 0
	test := 20
	check := false
	for !check {
		check = true
		for i := 20; i > 0 && check; i-- {
			check = check && (0 == (test % i))
		}
		if !check {
			test += 20
		}
	}
	answer = test
	return answer
}
