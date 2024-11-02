package tsUtils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	return float32(AtoF64(saisie))
}

func AtoF32(saisie string) float32 {
	return float32(AtoF64(saisie))
}

func AtoF64(saisie string) float64 {
	value, err := strconv.ParseFloat(saisie, 64)
	if err != nil {
		value = 0
	}
	return value
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
	result := ""
	dateElmt := strings.Split(date, " ")
	switch len(dateElmt) {
	case 1:
		if val, err := strconv.Atoi(date); err == nil { // donc numérique
			if val > 1900 {
				result = date
			}
		}
	case 3:
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
		if month != "" {
			if _, err := strconv.Atoi(dateElmt[0]); err == nil { // donc numérique
				day := dateElmt[0]
				if len(day) == 1 {
					day = "0" + day
				}

				if _, err := strconv.Atoi(dateElmt[2]); err == nil { // donc numérique
					result = day + "/" + month + "/" + dateElmt[2]
				}
			}
		}
	}
	return result
}

// removeAccents() remplace les caractères accentués par leurs équivalents non accentués
func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	if err != nil {
		output = s
	}
	return output
}

// ### TranscodeName() - normalise le nom pour faciliter les comparaisons
func TranscodeName(nom string) string {
	//#remplace les caractères accentués par des caractères sans accent
	nom = removeAccents(nom)
	//#passer en minuscule le titre (name)
	nom = strings.ToLower(nom)
	//#supprime un caractère :|(|)|[|]|!|
	nom = strings.Replace(nom, ":", "", -1)
	nom = strings.Replace(nom, "(", "", -1)
	nom = strings.Replace(nom, ")", "", -1)
	nom = strings.Replace(nom, "[", "", -1)
	nom = strings.Replace(nom, "]", "", -1)
	nom = strings.Replace(nom, "!", "", -1)
	//#remplace un caractère par blanc '|’|-|–|,| et |
	nom = strings.Replace(nom, "'", " ", -1)
	nom = strings.Replace(nom, "’", " ", -1)
	nom = strings.Replace(nom, "–", " ", -1)
	nom = strings.Replace(nom, ",", " ", -1)
	nom = strings.Replace(nom, "|", " ", -1)
	// nom = strings.Replace(nom, "-", " ", -1)
	//#remplace un caractère par blanc _
	nom = strings.Replace(nom, "_", " ", -1)
	//#remplace un caractère par blanc °|/|
	nom = strings.Replace(nom, "°", " ", -1)
	nom = strings.Replace(nom, "/", " ", -1)
	//#remplace ² par .2
	nom = strings.Replace(nom, "²", ".2", -1)
	//#remplace plusieurs '.' par un seul
	re := regexp.MustCompile(`(  +)`)
	nom = re.ReplaceAllLiteralString(nom, " ")
	//#TRIM
	nom = strings.TrimSpace(nom)

	return nom
}

// Substr
func Substr(s string, start, end int) string {
	if start == end {
		return ""
	}
	counter, startIdx := 0, 0
	for i := range s {
		if counter == start {
			startIdx = i
		}
		if counter == end {
			return s[startIdx:i]
		}
		counter++
	}
	return s[startIdx:]
}
