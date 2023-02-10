package controldigit

import (
	"fmt"
	// "log"
	"regexp"
	"strconv"
	"strings"
)

type parsedNIF struct {
	number  int
	kind    string
	control string
}

func IsValidNIF(str string) bool {
	// DNI = 7(should add leading zero) or 8 digits + control digit
	// NIF = DNI without control digit
	// NIF K = K + 7 digits + control digit, spaniards younger than 14
	// NIF L = L + 7 digits + control digit, spaniards residents in foreign country
	// NIF M = M + 7 digits + control digit, foreign without NIE
	// Matches optionally starting with k, l or m ignoring case + 7 or 8 numbers + optional control digit
	DNIRegex := regexp.MustCompile(`[(?i)klm(?-i)\d]{1}\d{6,7}\w?`)
	return DNIRegex.MatchString(str)
}

func IsValidNIE(str string) bool {
	DNIRegex := regexp.MustCompile(`^(x|y|z)\d{7}((\w)?)$`)
	return DNIRegex.MatchString(str)
}

func getParsedNIF(str string) parsedNIF {
	nif := parsedNIF{}
	var parsedString string
	if IsValidNIF(str) {
		lastChar := string(str[len(str)-1])
		firstChar := string(str[0:1])
		_, isControl := strconv.Atoi(lastChar)

		if isControl != nil {
			nif.control = lastChar
			parsedString = strings.Replace(str, lastChar, "", 2)
		}
		_, isKind := strconv.Atoi(firstChar)

		if isKind != nil {
			nif.kind = firstChar
			parsedString = strings.Replace(parsedString, firstChar, "", 2)
		}

		// number, err := strconv.Atoi(parsedString)

		// if err != nil {
		// 	log.Fatalf("There was an error parsing %s", nil)
		// }

		// nif.number = number

		fmt.Printf("THE CLEAN STR IS %s \n", lastChar)
		nif.kind = "NIF"

	} else if IsValidNIE(str) {
		nif.kind = "NIE"
	}

	fmt.Printf("THE PARSED NIF IS %+v\n", nif)

	return nif
}

func GetControlDigit(str string) string {

	// remainder := map[int]string{0: "T", 1: "R", 2: "W", 3: "A", 4: "G", 5: "M", 6: "Y", 7: "F", 8: "P", 9: "D", 10: "X", 11: "B", 12: "N", 13: "J", 14: "Z", 15: "S", 16: "Q", 17: "V", 18: "H", 19: "L", 20: "C", 21: "K", 22: "E"}

	getParsedNIF(str)
	// if IsValidNIE(str) {
	// 	firstLetter := str[0:1]
	// 	replacement := map[string]string{"x": "", "y": "1", "z": "2"}

	// 	candidate := strings.Replace(str, firstLetter, replacement[firstLetter], 2)

	// 	result, err := strconv.Atoi(candidate)
	// 	if err != nil {
	// 		log.Fatalf("Error parsing NIE %s", err)
	// 	}
	// 	fmt.Printf("IS VALID NIE: %s \n", remainder[result%23])
	// 	return remainder[result%23]
	// }

	// if IsValidNIF(str) {
	// 	result, err := strconv.Atoi(str)
	// 	if err != nil {
	// 		log.Fatalf("Error parsing NIE %s", err)
	// 	}
	// 	// fmt.Printf("The result is %s \n", remainder[result%23])
	// 	return remainder[result%23]
	// }

	return "match"

}
