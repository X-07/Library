package tsFunction

import (
	"strconv"
)

// MiseEnFormeByte (nb octets)
func MiseEnFormeByte(bps int64) string {
	var unit string
	var speed float64
	var xSpeed string
	// strconv.FormatFloat(f, 'f', 5, 64) - 'f' format - 5 is the number of decimals - 64 is for float64 type

	if bps >= 1073741824 {
		unit = "G"
		speed = float64(bps) / 1073741824
		if bps >= 107374182400 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10737418240 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1048576 {
		unit = "M"
		speed = float64(bps) / 1048576
		if bps >= 104857600 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10485760 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1024 {
		unit = "K"
		speed = float64(bps) / 1024
		if bps >= 102400 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else if bps >= 10240 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else {
		unit = "o"
		xSpeed = strconv.FormatInt(bps, 10)
	}

	return xSpeed + unit
}

// MiseEnFormeBit (nb bits)
func MiseEnFormeBit(bps int64) string {
	var unit string
	var speed float64
	var xSpeed string
	// strconv.FormatFloat(f, 'f', 5, 64) - 'f' format - 5 is the number of decimals - 64 is for float64 type

	if bps >= 1000000000 {
		unit = "Gbit/s"
		speed = float64(bps) / 1000000000
		if bps >= 100000000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10000000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1000000 {
		unit = "Mbit/s"
		speed = float64(bps) / 1000000
		if bps >= 100000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1000 {
		unit = "Kbit/s"
		speed = float64(bps) / 1000
		if bps >= 100000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else {
		unit = "bit/s"
		xSpeed = strconv.FormatInt(bps, 10)
	}

	return xSpeed + unit
}

//MiseEnFormeGiga : function de mise en forme
func MiseEnFormeGiga(val int64) string {
	var xResult string
	result := float64(val) / 1048576 //1Go
	switch {
	case val == 0:
		xResult = "0.0"
	case val > 10485760: //10Go
		xResult = strconv.FormatFloat(result, 'f', 0, 64) // 0 decimale
	case val > 2097152: //2Go
		xResult = strconv.FormatFloat(result, 'f', 1, 64) // 1 decimale
	default:
		xResult = strconv.FormatFloat(result, 'f', 2, 64) // 2 decimale
	}
	return xResult
}

// SliceIntUniq removes duplicate values in given slice
func SliceIntUniq(s []int) []int {
	for i1 := 0; i1 < len(s); i1++ {
		for i2 := i1 + 1; i2 < len(s); i2++ {
			if s[i1] == s[i2] {
				// delete
				s = append(s[:i2], s[i2+1:]...)
				i2--
			}
		}
	}
	return s
}

// SliceStringUniq removes duplicate values in given slice
func SliceStringUniq(s []string) []string {
	for i1 := 0; i1 < len(s); i1++ {
		for i2 := i1 + 1; i2 < len(s); i2++ {
			if s[i1] == s[i2] {
				// delete
				s = append(s[:i2], s[i2+1:]...)
				i2--
			}
		}
	}
	return s
}
