package learning

import "fmt"

func Example1() {
	m := make(map[string]Person)
	m["President"] = Person{0, "Trump"}
	fmt.Println(m["President"].Name)
	m["President"] = Person{1, "Obama"}
	fmt.Println(m["President"].Name)
	// m["President"].Name = "Shino"
	// cannot assign to struct field m["President"].Name in map

	//Output:
	//Trump
	//Obama
}

func Example2() {
	m := make(map[string]*Person) //Value = pointer of struct Person
	m["President"] = &Person{0, "Trump"}
	fmt.Println(m["President"].Name)
	m["President"] = &Person{1, "Obama"}
	fmt.Println(m["President"].Name)
	m["President"].Name = "Shino"
	fmt.Println(m["President"].Name)
	//Output:
	//Trump
	//Obama
	//Shino
}

func Example0() {
	fmt.Println("Hi")
	//Output:
	//Hi
}
