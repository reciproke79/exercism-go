package accumulate

const testVersion = 1

func Accumulate(strs []string, f func(string) string) (res []string) {
	for _, s := range strs {
		res = append(res, f(s))
	}
	return res
}
