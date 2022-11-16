package two_sum

// Runtime: 119 ms
// Memory Usage: 3.6 MB
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}

	return nil
}
