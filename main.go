package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the first number? ")
	firstNum, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Couldn't get first number")
		return
	}
	firstNum = strings.TrimSpace(firstNum)

	fmt.Print("What is the second number? ")
	secondNum, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Couldn't get second number")
		return
	}
	secondNum = strings.TrimSpace(secondNum)

	firstInt, err := strconv.Atoi(firstNum)
	if err != nil {
		fmt.Print("Couldn't convert first num to Int")
		return
	}

	secondInt, err := strconv.Atoi(secondNum)
	if err != nil {
		fmt.Print("Couldn't convert second num to Int")
		return
	}

	sum := firstInt + secondInt
	difference := firstInt - secondInt
	product := firstInt * secondInt
	quotient := firstInt / secondInt

	fmt.Printf("%v + %v = %v\n", firstInt, secondInt, sum)
	fmt.Printf("%v - %v = %v\n", firstInt, secondInt, difference)
	fmt.Printf("%v * %v = %v\n", firstInt, secondInt, product)
	fmt.Printf("%v / %v = %v\n", firstInt, secondInt, quotient)
}
