package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Ch01_03() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
