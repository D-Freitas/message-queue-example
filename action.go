package main

import (
	"fmt"
)

var currentBalance = 1000

func PrintPayload(payload int) {
	currentBalance = currentBalance - 10
	fmt.Println("currentBalance: ", currentBalance)
}
