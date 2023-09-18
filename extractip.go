package main

// Basic ip extrator
// read stdin, output filteredinput if error to output
// 2023/07/20 : V0.1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	// Please verify regexp in _test.go before blindly using it !
	// emails := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// DomainUrl := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	// words := regexp.MustCompile(`[\p{L}]+`) // Without numbers
	// words := regexp.MustCompile("\\P{M}+") // With numbers ?

	URL = regexp.MustCompile(`https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)`)

	EMAIL = regexp.MustCompile("[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")
	IP    = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
)

func main() {

	var re *regexp.Regexp
	switch os.Args[1] {
	case "-url":
		re = URL
	case "-email":
		re = EMAIL
	case "-ip":
		re = IP

	default:
		re = IP
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		for _, element := range match(text, re) {
			fmt.Println(element)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

// Extracting re.FindAllString func for regex testing
func match(text string, re *regexp.Regexp) []string {
	return re.FindAllString(text, -1)
}
