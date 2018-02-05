package main

import "fmt"

var TheInt int

func init() {
	TheInt = 42
}

func main() {
	fmt.Printf("The int is: %d\n", TheInt)
}
