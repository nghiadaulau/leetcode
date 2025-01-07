package two_sum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)

	for i := range nums {
		if index, found := hash[target-nums[i]]; found {
			return []int{index, i}
		}
		hash[nums[i]] = i
	}
	return []int{-1, -1}
}
