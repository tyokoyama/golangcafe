package main

import (
	"fmt"
	"strconv"

	"github.com/tyokoyama/golangcafe/expertsgo/checkdigit"
)

func main() {
	p := checkdigit.NewLuhn()

	sum, err := p.Generate("411111111111111")
	if err != nil {
		fmt.Printf("Generate Error [%v]\n", err)
		return
	}

	fmt.Printf("Generate Result: [%d]\n", sum)

	if p.Verify("411111111111111" + strconv.Itoa(sum)) {
		fmt.Printf("Verify OK\n")
	} else {
		fmt.Printf("Verify Error\n")
	}
}