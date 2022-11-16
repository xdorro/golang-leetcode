package two_sum

// Runtime: 11 ms
// Memory Usage: 4.3 MB
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}

		m[v] = k
	}

	return nil
}
