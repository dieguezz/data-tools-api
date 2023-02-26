package main

import (
	"fmt"

	cifvalidator "github.com/dieguezz/nif-tools/pkg/cif/validators"
)

func main() {
	fmt.Println(cifvalidator.IsValidCIF("A71255764"))
}
