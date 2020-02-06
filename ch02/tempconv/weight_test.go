package tempconv

import "fmt"

func ExamplePToKg() {
	fmt.Println(PoundToKillGram(Pound(1.0)))
	//Output:
	//0.453592 kg
	fmt.Println(KillGramToPound(KillGram(1.0)))
	//2.20462 pound
}
