package main

import (
	"fmt"

	"github.com/richardbertozzo/track-orders-go/sites"
)

func main() {
	var input string

	fmt.Print("Enter the track code: ")
	fmt.Scanln(&input)

	bytes := sites.GetInfoOrder(&input)

	fmt.Println(string(bytes[:]))
}
