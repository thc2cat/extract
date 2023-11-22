package main

// Basic ip extrator
// read stdin, output filteredinput if error to output
// 2023/07/20 : V0.1
// 2023/11/07 : V0.2 -ipv6 + net.Parse , -match pattern

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

var (
	// Please verify regexp in _test.go before blindly using it !
	// emails := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// DomainUrl := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	// words := regexp.MustCompile(`[\p{L}]+`) // Without numbers
	// words := regexp.MustCompile("\\P{M}+") // With numbers ?

	MAC   = regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)
	URL   = regexp.MustCompile(`https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)`)
	EMAIL = regexp.MustCompile("[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")
	IP    = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	IP6 = regexp.MustCompile(`((([0-9a-fA-F]{1,4}|)([:]{1,2})([0-9a-fA-F]{1,4}|[:]{1,2})){1,8})`)
	// Catch a lot more than ipv6 address (have to be verified with net.Parse )
)

func main() {

	var (
		re  *regexp.Regexp
		arg string
	)

	if os.Args == nil {
		os.Exit(-1)
	}
	if len(os.Args) == 1 { // Just the Name
		arg = "-ip4"
	} else {
		arg = os.Args[1]
	}

	switch arg {
	case "-help":
		Usage()
		os.Exit(0)
	case "-url":
		re = URL
	case "-email":
		re = EMAIL
	case "-ip4":
		re = IP
	case "-ip6":
		re = IP6
	case "-mac":
		re = MAC
	case "-match":
		if len(os.Args) == 2 {
			re = regexp.MustCompile(os.Args[2])
		}
	default:
		Usage()
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		for _, element := range match(text, re) {
			switch arg {
			case "-ip4", "-ip6":
				if r := net.ParseIP(element); r != nil {
					fmt.Println(element)
				}
			default:
				fmt.Println(element)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

}

// Extracting re.FindAllString func for regex testing
func match(text string, re *regexp.Regexp) []string {
	return re.FindAllString(text, -1)
}

func Usage() {
	if os.Args != nil {
		fmt.Printf("Usage: %s [-url|-email|(-ip4 default)|-ip6|-mac\n", os.Args[0])
	}
}