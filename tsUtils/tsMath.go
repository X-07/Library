package tsUtils

import "github.com/google/uuid"

func Max(values ...int) int {
	res := 0
	for _, value := range values {
		if value > res {
			res = value
		}
	}
	return res
}

// length: longueur du RANDOM généré
func RandomString(length int) string {
	return uuid.NewString()[:length]
}
