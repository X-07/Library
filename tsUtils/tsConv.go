package tsUtils

import (
	"fmt"
	"strconv"
	"strings"
)

func AtoI(saisie string) int {
	value, err := strconv.Atoi(saisie)
	if err != nil {
		value = 0
	}
	return value
}

func AtoI64(saisie string) int64 {
	value, err := strconv.ParseInt(saisie, 10, 64)
	if err != nil {
		value = 0
	}
	return value
}

func ItoA(value int) string {
	return strconv.Itoa(value)
}

func I64toA(value int64) string {
	return strconv.FormatInt(int64(value), 10)
}

func AtoF(saisie string) float32 {
	value, err := strconv.ParseFloat(saisie, 32)
	if err != nil {
		value = 0
	}
	return float32(value)
}

func FtoA(value float32) string {
	if value < 0.1 {
		return strconv.FormatFloat(float64(value), 'f', 2, 32)
	} else {
		return strconv.FormatFloat(float64(value), 'f', 1, 32)
	}

}

func AtoB(saisie string) bool {
	value, err := strconv.ParseBool(saisie)
	if err != nil {
		value = false
	}
	return value
}

func AtoCSV(value string) string {
	value = strings.ReplaceAll(value, "|", " ")
	value = strings.ReplaceAll(value, "\n", "\\n")
	value = strings.ReplaceAll(value, "\r", "\\r")
	return value
}

func CSVtoA(value string) string {
	value = strings.ReplaceAll(value, "\\n", "\n")
	value = strings.ReplaceAll(value, "\\r", "\r")
	return value
}

func BtoA(value bool) string {
	return strconv.FormatBool(value)
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d o", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %co", float64(b)/float64(div), "KMGTPE"[exp])
}

func SliceTrim(in []string) []string {
	var out []string
	for _, elmt := range in {
		out = append(out, strings.Trim(elmt, " "))
	}
	return out
}

func ConvertDate(date string) string {
	result := date
	dateElmt := strings.Split(date, " ")
	if len(dateElmt) == 3 {
		month := ""
		switch dateElmt[1] {
		case "janvier":
			month = "01"
		case "février":
			month = "02"
		case "mars":
			month = "03"
		case "avril":
			month = "04"
		case "mai":
			month = "05"
		case "juin":
			month = "06"
		case "juillet":
			month = "07"
		case "août":
			month = "08"
		case "septembre":
			month = "09"
		case "octobre":
			month = "10"
		case "novembre":
			month = "11"
		case "décembre":
			month = "12"
		}
		day := dateElmt[0]
		if len(day) == 1 {
			day = "0" + day
		}
		result = day + "/" + month + "/" + dateElmt[2]
	}
	return result
}
