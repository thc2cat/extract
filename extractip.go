package main

// Basic UTF8 mime decoder using Go libs
// read stdin, output decoded or raw input if error to output
// 2023/03/30 : V0.2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		submatchall := re.FindAllString(text, -1)
		for _, element := range submatchall {
			fmt.Println(element)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
