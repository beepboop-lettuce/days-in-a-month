package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter the name of a month")
		return
	}

	year := strings.ToLower(os.Args[1])

	if year != "january" || year != "february" || year != "march" || year != "april" || year != "may" || year != "june" || year != "july" || year != "august" || year != "september" || year != "october" || year != "november" || year != "december" {
		fmt.Printf("%q is not a month.\n", year)
		return
	}

}
