package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	var age uint8
	MINIMUM_AGE := 10

	fmt.Println("Welcome to my guessing game")

	fmt.Print("What is your name > ")
	fmt.Scan(&name)

	fmt.Print("How old are you > ")
	fmt.Scan(&age)

	if age < uint8(MINIMUM_AGE) {
		fmt.Println("You aren't old enough to play")
		return
	}

	msg := "Welcome to my game " + name
	fmt.Println(msg)

	fmt.Printf("Which is better, the RTX 3080 or RTX 3090 > ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()
	answer = strings.TrimSpace(answer)

	fmt.Println("You picked ", answer)

	var xr = []int{1, 2, 3}
	xr = append(xr, 10)
	fmt.Println(xr)

}
