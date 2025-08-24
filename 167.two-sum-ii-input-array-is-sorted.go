package main

// @leet start
func twoSum(numbers []int, target int) []int {
	minLen := 0
	maxLen := len(numbers) - 1
	for {
		currentVal := numbers[minLen] + numbers[maxLen]
		if currentVal == target {
			return []int{minLen + 1, maxLen + 1}
		} else if currentVal > target {
			maxLen--
		} else if currentVal < target {
			minLen++
		}
	}
}

// @leet end
