func twoSum(nums []int, target int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j]+nums[j-i] == target {
				return []int{j - i, j}
			}
		}
	}
	return nil // Return nil if no solution found
}

