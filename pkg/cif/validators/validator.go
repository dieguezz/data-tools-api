package cifvalidator

import (
	cifcontroldigit "github.com/dieguezz/nif-tools/pkg/cif/control-digit"
)

func IsValidCIF(str string) bool {
	control := cifcontroldigit.GetCIFControlDigit(str[:len(str)-1])
	if control == string(str[len(str)-1]) {
		return true
	}
	return false
}
