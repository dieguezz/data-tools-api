package nifvalidator

import (
	"regexp"
	"strconv"
	"strings"

	nifcontroldigit "github.com/dieguezz/nif-tools/pkg/control-digit"
)

func LooksLikeNIF(str string) bool {
	// DNI = 7(should add leading zero) or 8 digits + control digit
	// NIF = DNI without control digit
	// NIF K = K + 7 digits + control digit, spaniards younger than 14
	// NIF L = L + 7 digits + control digit, spaniards residents in foreign country
	// NIF M = M + 7 digits + control digit, foreign without NIE
	// Matches optionally starting with k, l or m ignoring case + 7 or 8 numbers + optional control digit

	DNIRegex := regexp.MustCompile(`^[(?i)klm(?-i)\d]{1}\d{6,7}\w?$`)
	return DNIRegex.MatchString(str)
}

func LooksLikeNIE(str string) bool {
	NIERegex := regexp.MustCompile(`^(x|X|y|Y|z|Z)\d{7}((\w)?)$`)
	return NIERegex.MatchString(str)
}

func IsValidNIF(str string) bool {

	match := LooksLikeNIF(str)
	firstChar := string(str[0:1])
	lastChar := string(str[len(str)-1])
	numStr := str
	_, isControl := strconv.Atoi(lastChar)

	if isControl != nil {
		numStr = numStr[:len(numStr)-1]
	} else {
		return false
	}

	_, isLetter := strconv.Atoi(firstChar)

	if isLetter != nil {
		numStr = strings.Replace(str, firstChar, "", 2)
	}

	num, err := strconv.Atoi(numStr)

	if err != nil {
		return false
	}

	control := nifcontroldigit.GetNIFControlDigit(int32(num))

	return match && strings.EqualFold(control, lastChar)
}

func IsValidNIE(str string) bool {
	match := LooksLikeNIE(str)
	firstChar := string(str[0:1])
	lastChar := string(str[len(str)-1])
	_, isLetter := strconv.Atoi(firstChar)
	numStr := str
	if isLetter != nil {
		prefixMapping := map[string]string{"X": "", "Y": "1", "Z": "2"}
		numStr = strings.Replace(str, firstChar, prefixMapping[firstChar], 2)
	}
	_, isControl := strconv.Atoi(lastChar)

	if isControl != nil {
		numStr = numStr[:len(numStr)-1]
	} else {
		return false
	}

	num, err := strconv.Atoi(numStr)

	if err != nil {
		return false
	}

	control := nifcontroldigit.GetNIFControlDigit(int32(num))

	return match && strings.EqualFold(control, lastChar)

}
