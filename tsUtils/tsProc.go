package tsUtils

import (
	"math"
	"strconv"
)

// NumberFormatting 		1047527424 octets ==> 999 Mo
//
//	kindOf = "o" pour Octet ou "b" pour bit par ex ou toute autre lettre
func NumberFormatting(val int64, kindOf string) (float64, string) {
	var round float64
	var roundOn float64 = 0.5
	var places int = 2
	var suffixes [5]string

	suffixes[0] = kindOf + " "
	suffixes[1] = "K" + kindOf
	suffixes[2] = "M" + kindOf
	suffixes[3] = "G" + kindOf
	suffixes[4] = "T" + kindOf

	if val == 0 {
		return 0, suffixes[0]
	}

	base := math.Log(float64(val)) / math.Log(1024)
	size := math.Pow(1024, base-math.Floor(base))

	pow := math.Pow(10, float64(places))
	digit := pow * size
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal := round / pow

	return newVal, suffixes[int(math.Floor(base))]
}

// Deprecated: should not be used - Use NumberFormatting(val, kindOf) instead.
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
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 décimale
		} else if bps >= 10737418240 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 décimale
		}
	} else if bps >= 1048576 {
		unit = "M"
		speed = float64(bps) / 1048576
		if bps >= 104857600 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 décimale
		} else if bps >= 10485760 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 décimale
		}
	} else if bps >= 1024 {
		unit = "K"
		speed = float64(bps) / 1024
		if bps >= 102400 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else if bps >= 10240 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 décimale
		}
	} else {
		unit = "o"
		xSpeed = strconv.FormatInt(bps, 10)
	}

	return xSpeed + unit
}

// Deprecated: should not be used - Use NumberFormatting(val, kindOf) instead.
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
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 décimale
		} else if bps >= 10000000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 décimale
		}
	} else if bps >= 1000000 {
		unit = "Mbit/s"
		speed = float64(bps) / 1000000
		if bps >= 100000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 décimale
		} else if bps >= 10000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 décimale
		}
	} else if bps >= 1000 {
		unit = "Kbit/s"
		speed = float64(bps) / 1000
		if bps >= 100000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 décimale
		} else if bps >= 10000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 décimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 décimale
		}
	} else {
		unit = "bit/s"
		xSpeed = strconv.FormatInt(bps, 10)
	}

	return xSpeed + unit
}

// Deprecated: should not be used - Use NumberFormatting(val, kindOf) instead.
// MiseEnFormeGiga : function de mise en forme
func MiseEnFormeGiga(val int64) string {
	var xResult string
	result := float64(val) / 1048576 //1Go
	switch {
	case val == 0:
		xResult = "0.0"
	case val > 10485760: //10Go
		xResult = strconv.FormatFloat(result, 'f', 0, 64) // 0 décimale
	case val > 2097152: //2Go
		xResult = strconv.FormatFloat(result, 'f', 1, 64) // 1 décimale
	default:
		xResult = strconv.FormatFloat(result, 'f', 2, 64) // 2 décimale
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

// AppendIfNotContains: retourne si le nouvel élément est présent dans la liste (ajouté ou déjà présent)
func AppendIfNotContains(tab *[]string, str string, max int) bool {
	find := false
	for _, elmt := range *tab {
		if elmt == str {
			find = true
		}
	}
	if !find {
		if max == -1 || len(*tab) < max+1 {
			*tab = append(*tab, str)
			return true
		} else {
			return false
		}
	}
	return true
}

// IsValueIntoSlice: retourne si la valeur ('value') est présente dans la liste ('slice') ou pas.
func IsValueIntoSlice(slice []int64, value int64) bool {
	for _, elmt := range slice {
		if elmt == value {
			return true
		}
	}
	return false
}

// MinimiseString retourne un s sous la forme xxx....xxx de la longueur len
func MinimiseString(s string, newLen int) string {
	if len(s) <= newLen {
		return s
	}
	var deb, fin int
	if newLen < 11 {
		newLen = 10
	}
	fin = newLen/2 - 2
	deb = newLen - 4 - fin
	len := len(s)
	return s[:deb] + "...." + s[len-fin:len]
}
