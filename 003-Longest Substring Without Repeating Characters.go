func lengthOfLongestSubstring(s string) int {

    if len(s) == 0{
        return 0
    }

	max := 1
	max_pre := 1
	index := make(map[byte]int)
	index[s[0]] = 0

	for i := 1; i < len(s); i++ {

		value, present := index[s[i]]
		if present {
			max_pre = Min(max_pre+1, i-value)
		} else {
			max_pre = max_pre + 1
		}

		max = Max(max, max_pre)

		index[s[i]] = i
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
