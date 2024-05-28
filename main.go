package main

import (
	"fmt"
	"tohexa/files"
)

// "fmt"

func main() {
	// str := os.Args[1]
	str := "356"
	hx := tohexa.ToHexadecimal(str)
	fmt.Sprintln(hx)
}
