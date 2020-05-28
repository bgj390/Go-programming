// This is an exersice for http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2
// The purpose is to use at least two libraries other than rand, fmt, time or string

// This program takes preferably a long'ish' txt-file with the scanner,
// and orders the words alphabetically. Example: cat (insert txt-file here) | go run scan.go
 
// sources:
// https://stackoverflow.com/questions/49429269/how-can-i-trim-whitespaces-in-go-from-a-slice-after-split/49429437
// https://gobyexample.com/sorting
// https://gobyexample.com/line-filters

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)


func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()		
		var words = strings.Split(txt, " ")
		
		sort.Strings(words)
		fmt.Println(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
