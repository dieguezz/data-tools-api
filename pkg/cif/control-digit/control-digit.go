package cifcontroldigit

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CalculateCIFLetter(num int32) string {
	remainder := map[int32]string{0: "T", 1: "R", 2: "W", 3: "A", 4: "G", 5: "M", 6: "Y", 7: "F", 8: "P", 9: "D", 10: "X", 11: "B", 12: "N", 13: "J", 14: "Z", 15: "S", 16: "Q", 17: "V", 18: "H", 19: "L", 20: "C", 21: "K", 22: "E"}
	return remainder[num%23]
}

func GetCIFControlDigit(str string) string {
	sumOfEven := 0
	sumOfOdds := 0
	firstChar := str[0:1]
	parsedString := str[1:]

	for i, char := range parsedString {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatalf("Error converting to int: %v", err)
		}

		if i%2 != 0 {
			sumOfEven += num
		} else {
			foo := fmt.Sprintf("%d", num*2)
			subSum := 0
			for _, subChar := range foo {
				num, err := strconv.Atoi(string(subChar))
				if err != nil {
					log.Fatalf("Error converting to int: %v", err)
				}

				subSum += num
			}

			sumOfOdds += subSum
		}
	}

	res := sumOfEven + sumOfOdds

	result := 10 - res%10

	control := ""
	if strings.Contains("PQRSNW", firstChar) {
		control = strings.ReplaceAll(fmt.Sprintf("%q", 64+result), "'", "")
	} else {
		if strings.Contains("ABCDEFGHJUV", firstChar) {
			if result == 10 {
				control = "0"
			} else {
				control = fmt.Sprintf("%v", result)
			}
		} else {
			// I am not sure this case is correct
			control = CalculateCIFLetter(int32(result))
		}
	}
	return control

}
