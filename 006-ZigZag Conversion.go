func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	sc := make([]byte, 0)
	for i := 1; i <= numRows; i++ {

		step_down := (numRows - i) * 2
		step_up := (i - 1) * 2

		flag := false
		step := 0
		for j := i; j <= len(s); j += step {

			sc = append(sc, s[j-1])
			step = (numRows - 1) * 2
			if flag {

				if step_up != 0 {
					step = step_up
				}
				flag = false
			} else {

				if 0 != step_down {
					step = step_down
				}
				flag = true
			}
		}
	}

	return string(sc)
}