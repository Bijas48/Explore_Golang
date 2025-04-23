package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("ezo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("edo e1o ezo eio eko eKo", 10))
}
