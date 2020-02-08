package main

import "github.com/nobishino/golang-book/ch03/web"

func main() {
	// svg.Write(os.Stdout)
	web.Listen(8080)
}
