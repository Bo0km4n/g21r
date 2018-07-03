package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		return
	}
	str := format(os.Args[1])
	if err := clipboard.WriteAll(str); err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func format(input string) string {
	return addCRLFAfterPeriod(replaceCRLF(input))
}

func replaceCRLF(input string) string {
	return strings.Replace(input, "\n", " ", -1)
}

func addCRLFAfterPeriod(input string) string {
	rep := regexp.MustCompile(`\. `)
	str := rep.ReplaceAllString(input, ".\n")
	return str
}
