package controldigit

import (
	"fmt"
	"log"
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
	DNIRegex := regexp.MustCompile(`^[(?i)klm(?-i)\d]{1}\d{6,7}\w?$`)
	return DNIRegex.MatchString(str)
}

func IsValidNIE(str string) bool {
	DNIRegex := regexp.MustCompile(`^(x|y|z)\d{7}((\w)?)$`)
	return DNIRegex.MatchString(str)
}

func GetParsedNIF(str string) parsedNIF {
	nif := parsedNIF{}
	var parsedString string
	if IsValidNIF(str) {
		lastChar := string(str[len(str)-1])
		firstChar := string(str[0:1])
		_, isControl := strconv.Atoi(lastChar)

		if isControl != nil {
			nif.control = strings.ToUpper(lastChar)
			parsedString = strings.Replace(str, lastChar, "", 2)
		}
		_, isKind := strconv.Atoi(firstChar)

		if isKind != nil {
			nif.kind = fmt.Sprintf("NIF%s", strings.ToUpper(firstChar))
			parsedString = strings.Replace(parsedString, firstChar, "", 2)
		} else {
			nif.kind = "NIF"
		}

		number, err := strconv.Atoi(parsedString)

		if err != nil {
			log.Fatalf("There was an error parsing %v", nil)
		}

		nif.number = number
		nif.control = GetControlDigit(number)
	
		return nif

	} else if IsValidNIE(str) {
		lastChar := string(str[len(str)-1])
		firstChar := string(str[0:1])
		
		nif.kind = fmt.Sprintf("NIE%s", strings.ToUpper(firstChar))
		
		_, isControl := strconv.Atoi(lastChar)

		if isControl != nil {
			nif.control = lastChar
			parsedString = strings.Replace(parsedString, firstChar, "", 2)
			parsedString = strings.Replace(parsedString, lastChar, "", 2)

			number, err := strconv.Atoi(parsedString)
			
			if err != nil {
				log.Fatalf("There was an error parsing string with value %v", nil)
			}

			nif.number = number

			return nif
		}

		parsedString = strings.Replace(str, firstChar, "", 2)
		number, err := strconv.Atoi(parsedString)
		
		if err != nil {
			log.Fatalf("There was an error parsing %s with error %v", parsedString, nil)
		}
	
		nif.number = number
		nif.control = GetControlDigit(nif.number)
		return nif

	}

	return nif
}

func GetControlDigit(num int) string {

	remainder := map[int]string{0: "T", 1: "R", 2: "W", 3: "A", 4: "G", 5: "M", 6: "Y", 7: "F", 8: "P", 9: "D", 10: "X", 11: "B", 12: "N", 13: "J", 14: "Z", 15: "S", 16: "Q", 17: "V", 18: "H", 19: "L", 20: "C", 21: "K", 22: "E"}
	return remainder[num % 23]
	

}
