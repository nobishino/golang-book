package typeconversion

import "fmt"

func Example01() {
	var f float64 = 10.1
	i := int(f)
	fmt.Printf("%T, %v\n", f, f)
	f = 5.3
	fmt.Printf("%T, %v\n", i, i)
	//Output:
	//float64, 10.1
	//int, 10
}
