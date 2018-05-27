/*
https://projecteuler.net/problem=4

A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.
Find the largest palindrome made from the product of two 3-digit numbers.

Answer: 906609
*/

package go_euler

func Solution004() int {
	max_pali := 0
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			t := i * j
			if palindrome_test(t) {
				if t > max_pali {
					max_pali = t
				}
			}
		}
	}
	return max_pali
}

func palindrome_test(test_me int) bool {
	reversed := 0
	original := test_me

	for 0 < original {
		reversed = reversed*10 + (original % 10)
		original /= 10
	}

	return (test_me == reversed)
}
