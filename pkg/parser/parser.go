package nifparser

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	nifcontroldigit "github.com/dieguezz/nif-tools/pkg/control-digit"
	nifvalidator "github.com/dieguezz/nif-tools/pkg/validator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ParsedNIF struct {
	Number  int
	Kind    string
	Control string
}

func GetParsedNIF(str string) (ParsedNIF, error) {
	nif := ParsedNIF{}
	var parsedString string = str
	if nifvalidator.LooksLikeNIF(str) {
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
		nif.Control = nifcontroldigit.GetNIFControlDigit(int32(number))

		return nif, nil

	} else if nifvalidator.LooksLikeNIE(str) {
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

		prefixMapping := map[string]string{"X": "", "Y": "1", "Z": "2"}

		parsedString = strings.Replace(str, firstChar, prefixMapping[firstChar], 2)
		number, err := strconv.Atoi(parsedString)

		if err != nil {
			log.Fatalf("There was an error parsing %s with error %v", parsedString, nil)
		}

		nif.Number = number
		nif.Control = nifcontroldigit.GetNIFControlDigit(int32(nif.Number))
		return nif, nil

	}
	return nif, status.Error(codes.InvalidArgument, "Invalid NIF")
}
