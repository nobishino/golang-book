package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Ch01_03() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func NaiveJoin(words []string) string {
	var result, separator string
	for _, v := range words {
		result += separator + v
		separator = " "
	}
	return result
}
