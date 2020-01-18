package solutions

import (
	"bufio"
	"fmt"
	"os"
)

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
				fmt.Printf("%d times:\t%s in ", inc.Count, line)
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
		fmt.Println(input.Text(), f.Name())
		if s := counts[input.Text()]; s == nil {
			counts[input.Text()] = &inclusion{Count: 1, Files: make(map[string]int)}
			counts[input.Text()].Files[f.Name()]++
		} else {
			counts[input.Text()].Count++
			counts[input.Text()].Files[f.Name()]++
		}
	}
}
