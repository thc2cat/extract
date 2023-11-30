package main

// Basic ip extrator
// read stdin, output filteredinput if error to output
// 2023/07/20 : V0.1
// 2023/11/07 : V0.2 -ipv6 + net.Parse , -match pattern
// 2023/11/28 : v1.3 -ipv4p ( print private address, default ignore )
// 2023/11/28 : v1.4 : + uniqs map ( avoiding "| sort -u" )

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
	// Privates = regexp.MustCompile(`^(10|172\.(1[6789]|2[0-9]|3[01])|192\.168\.`)
	Privates = regexp.MustCompile(`^(0\.|240\.|255\.|224\.|169\.254|127\.|10\.|(172\.1[6-9]\.)|(172\.2[0-9]\.)(172\.3[0-1]\.)|(192\.168\.))`)

	// map of uniq elements founds ( avoid |sort -u|)
	uniqs = make(map[string]bool)
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
	case "-ip4p":
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
			case "-ip4", "-ip6", "-ip4p":
				if r := net.ParseIP(element); r != nil { // Parseable
					if arg == "-ip4p" { // print also Privates
						uniqs[element] = true
					} else {
						if t := match(element, Privates); len(t) == 0 {
							uniqs[element] = true
						}
					}
				}
			default:
				uniqs[element] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	for key := range uniqs {
		fmt.Println(key)
	}

}

// Extracting re.FindAllString func for regex testing
func match(text string, re *regexp.Regexp) []string {
	return re.FindAllString(text, -1)
}

func Usage() {
	if os.Args != nil {
		fmt.Printf("Usage: %s [-url|-email|-mac|-ip6|-ip4[p(ublic)] ]\n", os.Args[0])
	}
}
