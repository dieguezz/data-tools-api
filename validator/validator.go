package nifvalidator

import "regexp"

func IsValidNIF(str string) bool {
	// DNI = 7(should add leading zero) or 8 digits + control digit
	// NIF = DNI without control digit
	// NIF K = K + 7 digits + control digit, spaniards younger than 14
	// NIF L = L + 7 digits + control digit, spaniards residents in foreign country
	// NIF M = M + 7 digits + control digit, foreign without NIE
	// Matches optionally starting with k, l or m ignoring case + 7 or 8 numbers + optional control digit
	DNIRegex := regexp.MustCompile(`^[(?i)klm(?-i)\d]{1}\d{6,7}\w?$`)
	return DNIRegex.MatchString(str)
}

func IsValidNIE(str string) bool {
	DNIRegex := regexp.MustCompile(`^(x|y|z)\d{7}((\w)?)$`)
	return DNIRegex.MatchString(str)
}
