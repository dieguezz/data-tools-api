package main

import (
	"fmt"

	nifgenerators "github.com/dieguezz/nif-tools/pkg/generators"
)

func main() {
	// res := nifgenerators.GenerateNIF(nifgenerators.NIFOptions{UnderAge: false, LivingOutside: false, ForeignWithoutNIE: false})
	fmt.Println(nifgenerators.GenerateCIF())
	// fmt.Println(res)
}
