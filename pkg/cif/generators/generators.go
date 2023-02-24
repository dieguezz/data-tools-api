package cifgenerators

import (
	"fmt"
	"math/rand"

	cifcontroldigit "github.com/dieguezz/nif-tools/pkg/cif/control-digit"
)

func numberFromRange(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func GenerateCIF() string {
	availableLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "P", "Q", "S", "K", "L", "M", "X"}
	num := numberFromRange(1000000, 9999999)
	letter := availableLetters[rand.Intn(len(availableLetters))]
	control := cifcontroldigit.GetCIFControlDigit(fmt.Sprintf("%s%v", letter, num))
	return fmt.Sprintf("%v%v%v", letter, num, control)
}
