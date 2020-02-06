package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/nobishino/golang-book/ch02/tempconv"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		process(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Problem 2-2: %v\n", err)
				os.Exit(1)
			}
			process(f)
			f.Close()
		}
	}
}

func process(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "q" {
			fmt.Println("OK, Quitting the program...")
			os.Exit(0)
		}
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Println("(To Quit The Program, type 'q' and press Enter.)")
			continue
		}
		printTemperatureConversion(value)
		printWeightConversion(value)
		printLengthConversion(value)
	}
}

func printTemperatureConversion(value float64) {
	f := tempconv.Fahrenheit(value)
	c := tempconv.Celsius(value)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func printWeightConversion(value float64) {
	killgram := tempconv.KillGram(value)
	pound := tempconv.Pound(value)
	fmt.Printf("%s = %s, %s = %s\n", killgram, tempconv.KillGramToPound(killgram), pound, tempconv.PoundToKillGram(pound))
}

func printLengthConversion(value float64) {
	meter := tempconv.Meter(value)
	feet := tempconv.Feet(value)
	fmt.Printf("%s = %s, %s = %s\n", meter, tempconv.MeterToFeet(meter), feet, tempconv.FeetToMeter(feet))
}
