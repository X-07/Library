package tsUtils

func Max(values ...int) int {
	res := 0
	for _, value := range values {
		if value > res {
			res = value
		}
	}
	return res
}
