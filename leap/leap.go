package leap

const testVersion = 3

func IsLeapYear(i int) (r bool) {
	if i%4 == 0 {
		r = true
		if i%100 == 0 {
			r = false
			if i%400 == 0 {
				r = true
			}
		}
	}
	return r
}
