package controldigit

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ParsedNIF struct {
	Number  int
	Kind    string
	Control string
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

func GetParsedNIF(str string) (ParsedNIF, error) {
	nif := ParsedNIF{}
	var parsedString string = str
	if IsValidNIF(str) {
		lastChar := string(str[len(str)-1])
		firstChar := string(str[0:1])
		_, isControl := strconv.Atoi(lastChar)

		if isControl != nil {
			nif.Control = strings.ToUpper(lastChar)
			parsedString = strings.Replace(str, lastChar, "", 2)
		}
		_, isKind := strconv.Atoi(firstChar)

		if isKind != nil {
			nif.Kind = fmt.Sprintf("NIF%s", strings.ToUpper(firstChar))
			parsedString = strings.Replace(parsedString, firstChar, "", 2)
		} else {
			nif.Kind = "NIF"
		}

		number, err := strconv.Atoi(parsedString)

		if err != nil {
			log.Fatalf("There was an error parsing %v", err)
		}

		nif.Number = number
		nif.Control = GetControlDigit(int32(number))

		return nif, nil

	} else if IsValidNIE(str) {
		lastChar := string(str[len(str)-1])
		firstChar := string(str[0:1])

		nif.Kind = fmt.Sprintf("NIE%s", strings.ToUpper(firstChar))

		_, isControl := strconv.Atoi(lastChar)

		if isControl != nil {
			nif.Control = lastChar
			parsedString = strings.Replace(parsedString, firstChar, "", 2)
			parsedString = strings.Replace(parsedString, lastChar, "", 2)

			number, err := strconv.Atoi(parsedString)

			if err != nil {
				log.Fatalf("There was an error parsing string with value %v", nil)
			}

			nif.Number = number

			return nif, nil
		}

		parsedString = strings.Replace(str, firstChar, "", 2)
		number, err := strconv.Atoi(parsedString)

		if err != nil {
			log.Fatalf("There was an error parsing %s with error %v", parsedString, nil)
		}

		nif.Number = number
		nif.Control = GetControlDigit(int32(nif.Number))
		return nif, nil

	}
	return nif, status.Error(codes.InvalidArgument, "Invalid NIF")
}

func GetControlDigit(num int32) string {

	remainder := map[int32]string{0: "T", 1: "R", 2: "W", 3: "A", 4: "G", 5: "M", 6: "Y", 7: "F", 8: "P", 9: "D", 10: "X", 11: "B", 12: "N", 13: "J", 14: "Z", 15: "S", 16: "Q", 17: "V", 18: "H", 19: "L", 20: "C", 21: "K", 22: "E"}
	return remainder[num%23]

}
