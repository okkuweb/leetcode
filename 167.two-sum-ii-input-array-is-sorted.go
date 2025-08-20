package main
// @leet start
func twoSum(numbers []int, target int) []int {
	for x, v := range numbers {
		for y, w := range numbers {
			if v == w && x == y {
				continue
			}
			if v + w == target {
				return []int{x + 1, y + 1}
			}
		}
	}
	return nil
}
// @leet end
