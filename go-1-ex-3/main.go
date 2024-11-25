package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error create eyes.txt:", err)
		return
	}
	defer eyesFile.Close()
	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	LogFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error create dice.log:", err)
		return
	}
	defer LogFile.Close()
	fmt.Fprintln(LogFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(LogFile, "the dice was rolled at", when)
}
