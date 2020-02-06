package tempconv

import "fmt"

func ExampleFeetAndMeter() {
	fmt.Println(FeetToMeter(Feet(1.0)))
	//Output:
	//0.3048 m
	fmt.Println(MeterToFeet(Meter(1.0)))
	//3.28084 feet
}
