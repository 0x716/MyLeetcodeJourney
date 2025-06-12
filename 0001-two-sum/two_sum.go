// Solution for LeetCode problem #1 - Two Sum
// Author: Lin Jun (GitHub: https://github.com/0x716)
// This is part of my personal LeetCode journey project:
// https://github.com/0x716/MyLeetcodeJourney
//
// All solutions are written for educational and personal learning purposes.
// Do not copy directly for submission or academic use without understanding.

package leetcode

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := numToIndex[complement]; ok {
			return []int{j, i}
		}

		numToIndex[num] = i
	}

	return []int{}
}
