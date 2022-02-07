package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter the name of a month")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0

	days := 28

	month := os.Args[1]

	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a mont.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
