func longestPalindrome(s string) string {
    
	var arr [1000][1000]bool

	st := 0
	en := 0

	str_len := len(s)

	if str_len == 1 {
		return s
	}

	if 2 == str_len {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	for i := 0; i < str_len; i++ {
		arr[i][i] = true
	}
	for i := 0; i < str_len-1; i++ {
		if s[i] == s[i+1] {
			st = i
			en = i + 1
			arr[i][i+1] = true
		}
	}

	for i := 2; i < str_len; i++ {
		for j := 0; j+i < str_len; j++ {
			if s[j] == s[j+i] && arr[j+1][j+i-1] {
				st = j
				en = j + i
				arr[j][j+i] = true
			}
		}
	}

	return string(s[st : en+1])
}