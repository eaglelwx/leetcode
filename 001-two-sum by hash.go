func twoSum(nums []int, target int) []int {
    
	hashMap := make(map[int]int)
	for i := range nums {
		
		index,present := hashMap[target-nums[i]]
		if present{
			return []int{index, i}
		}
		hashMap[nums[i]] = i
	}

	return []int{0, 0}
    
}