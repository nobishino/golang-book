package solutions

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//Ignoring error caused by input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup2() {
	counts := make(map[string]*inclusion)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, inc := range counts {
			if inc.Count > 1 {
				fmt.Printf("%d\t%s", inc.Count, line)
				for key := range inc.Files {
					fmt.Printf(" %v", key)
				}
				fmt.Println()
			}
		}
	}
}

type inclusion struct {
	Count int
	Files map[string]int
}

func countLines(f *os.File, counts map[string]*inclusion) {
	input := bufio.NewScanner(f)
	for input.Scan() {

		if s := counts[input.Text()]; s == nil {
			counts[input.Text()] = &inclusion{Count: 0, Files: make(map[string]int)}
			counts[input.Text()].Files[f.Name()]++
		} else {
			counts[input.Text()].Count++
			counts[input.Text()].Files[f.Name()]++
		}
	}
}
