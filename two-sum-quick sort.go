
func partition(nums []int, lo int, hi int) (p int) {

	v := nums[lo]
	i := lo
	j := hi
	for i < j {

		for nums[i] <= v {
			if i == hi {
				break
			} else {
				i++
			}
		}

		for nums[j] >= v {
			if j == lo {
				break
			} else {
				j--
			}
		}

		if i < j {
			swap(&nums[i], &nums[j])
		}
	}

	swap(&nums[lo], &nums[j])
	return j
}

func sort(nums []int, lo int, hi int) {

	if lo >= hi {
		return
	}

	p := partition(nums, lo, hi)
	sort(nums, lo, p-1)
	sort(nums, p+1, hi)
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func binarysearch(nums []int, lo int, hi int, v int) int {

	if lo > hi {
		return -1
	}

	mid := (lo + hi) / 2

	if nums[mid] == v {
		return mid
	}

	if nums[mid] > v {
		return binarysearch(nums, lo, mid-1, v)
	}

	if nums[mid] < v {
		return binarysearch(nums, mid+1, hi, v)
	}

	return -1
}


func twoSum(nums []int, target int) []int {
    
	var nums_copy = make([]int, len(nums))
	copy(nums_copy, nums[0:])

	sort(nums, 0, len(nums)-1)

	var p, q int
	for i := 0; i < len(nums); i++ {

		ret := binarysearch(nums, i+1, len(nums)-1, target-nums[i])
		if ret != -1 {
			p = nums[i]
			q = nums[ret]
			break
		}
	}

	fmt.Println([]int{p, q})
	fmt.Println(nums_copy)

	index1 := 0
	index2 := 0
	for i := 0; i < len(nums_copy); i++ {
		if p == nums_copy[i] && index1 == 0 {
			index1 = i
		} else if q == nums_copy[i] && index2 == 0 {
			index2 = i
		}
	}

	return []int{index1, index2}
    
}

