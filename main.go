package main

func main() {

}

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

/*
8
Destroying Asteroids

You are given an integer mass, which represents the original mass of a planet.
You are further given an integer array asteroids, where asteroids[i] is the mass of the ith asteroid.

You can arrange for the planet to collide with the asteroids in any arbitrary order.
If the mass of the planet is greater than or equal to the mass of the asteroid,
the asteroid is destroyed and the planet gains the mass of the asteroid. Otherwise, the planet is destroyed.

Return true if all asteroids can be destroyed. Otherwise, return false.
*/

func asteroidsDestroyed(mass int, asteroids []int) bool {
	if len(asteroids) == 0 {
		return true
	}
	stack := []int{}
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] <= mass {
			mass += asteroids[i]
		} else {
			stack = append(stack, asteroids[i])
		}
	}
	if len(asteroids) == len(stack) {
		return false
	}

	return asteroidsDestroyed(mass, stack)
}

/*
9
Count Asterisks

You are given a string s, where every two consecutive vertical bars '|' are grouped into a pair.
In other words, the 1st and 2nd '|' make a pair, the 3rd and 4th '|' make a pair, and so forth.

Return the number of '*' in s, excluding the '*' between each pair of '|'.

Note that each '|' will belong to exactly one pair.
*/

func countAsterisks(s string) int {
	var opened bool
	counter := 0
	for _, v := range s {
		if v == '*' && !opened {
			counter++
		} else if v == '|' && !opened {
			opened = true
		} else if v == '|' && opened {
			opened = false
		}
	}
	return counter
}

/*
10
Jump Game

You are given an integer array nums.
You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/

func canJump(nums []int) bool {
	best := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > best {
			best = i + nums[i]
		}

		if nums[i] == 0 && i < len(nums)-1 && i == best {
			return false
		}
	}

	return true
}

/*
11
Longest Subarray of 1's After Deleting One Element

Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.
*/

func longestSubarray(nums []int) int {
	var deleted bool
	end := 0
	start := 0
	maxCounter := 0
	counter := 0
	for end < len(nums) {
		if nums[end] == 0 && !deleted {
			start = end
			deleted = true
		} else if nums[end] == 0 && deleted {
			end = start
			counter = 0
			deleted = false
		} else {
			counter++
		}
		if counter > maxCounter {
			maxCounter = counter
		}
		end++
	}

	if maxCounter == len(nums) {
		maxCounter -= 1
	}

	return maxCounter
}

/*
12
Validate Binary Search Tree

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
 of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(r *TreeNode, min, max *int) bool {
	if r == nil {
		return true
	}

	if min != nil && r.Val <= *min {
		return false
	}
	if max != nil && r.Val >= *max {
		return false
	}
	return validate(r.Left, min, &r.Val) && validate(r.Right, &r.Val, max)
}
