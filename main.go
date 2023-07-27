package main

/*
1
Search Insert Position
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {

		mid := min + (max-min)/2
		if nums[mid] < target {
			min = mid + 1
		} else if nums[mid] > target {
			max = mid - 1
		} else {
			return mid
		}
	}

	return max + 1
}

/*
2
Climbing Stairs
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

// Slow but small
func climbStairs1(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

// Optimal
func climbStairs2(n int) int {
	one := 1
	two := 2

	if n <= 2 {
		return n
	}

	prev := 0
	for i := 3; i < n; i++ {
		prev = one
		one = two
		two = prev + one
	}
	return one + two
}

/*
3
Combinations
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.
*/

func combine(n int, k int) [][]int {

	res := [][]int{}
	tmp := make([]int, k)
	i := 0

	for i >= 0 {
		tmp[i]++

		if tmp[i] > n {
			i--
		} else if i == k-1 {
			res = append(res, append([]int{}, tmp...))
		} else {
			i++
			tmp[i] = tmp[i-1]
		}
	}

	return res
}

/*
4
Generate Parentheses

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	tmpStr := make([]rune, n*2)
	if n == 0 {
		return result
	}
	delta, pos := 0, 0
	findPatterns(delta, pos, n*2, tmpStr, &result)
	return result
}

func findPatterns(delta, pos, numBrackets int, str []rune, res *[]string) {

	if delta <= numBrackets-pos-2 {
		str[pos] = '('
		findPatterns(delta+1, pos+1, numBrackets, str, res)
	}

	if delta > 0 {
		str[pos] = ')'
		findPatterns(delta-1, pos+1, numBrackets, str, res)
	}

	if pos == numBrackets {
		if delta == 0 {
			*res = append(*res, string(str))
		}
	}

}

/*
5
Count and Say

The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit. Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.

For example, the saying and conversion for digit string "3322251":

"23321511"

Given a positive integer n, return the nth term of the count-and-say sequence.
*/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prev := countAndSay(n - 1)
	var word string
	var count byte
	for i := range prev {
		count++
		if i == len(prev)-1 || prev[i] != prev[i+1] {
			word += string([]byte{count + 48}) + string(prev[i])
			count = 0
		}
	}
	return word
}

/*
6
Asteroid Collision

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction
(positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.
*/

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	if len(asteroids) < 2 {
		return asteroids
	}
	stack := []int{}

	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else if asteroids[i] < 0 && len(stack) == 0 {
			res = append(res, asteroids[i])
		} else if asteroids[i] < 0 {
			if stack[len(stack)-1]+asteroids[i] > 0 {
				continue
			} else if stack[len(stack)-1]+asteroids[i] < 0 {
				stack = stack[:len(stack)-1]
				i--
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	res = append(res, stack...)
	return res
}

/*
7
Maximum Subarray

Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.
*/

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	//start := 0 //somehow this works faster. idkw ???
	end := 0 // and this works faster. idkw???
	for ; end < len(nums); end++ {
		curSum += nums[end]
		if nums[end] > curSum {
			curSum = nums[end]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
