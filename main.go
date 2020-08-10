package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {

	numberIndex, _ := strconv.Atoi(os.Args[1])
	fibonacciNumber := calcFibonacci(numberIndex)
	fmt.Println("Calculated number at index " + strconv.Itoa(numberIndex) + " is " + fibonacciNumber.String())
}

func calcFibonacci(index int) big.Int {

	returnValue := big.NewInt(1)
	if index == 0 {
		returnValue = big.NewInt(0)
	} else if index == 1 {
		returnValue = big.NewInt(1)
	} else {
		if index == 0 {
			firstCheck := index / 2
			secondCheck := firstCheck - 1

			firstFibonacci := calcFibonacci(firstCheck)
			secondFibonacci := calcFibonacci(secondCheck)

			var firstPart = big.NewInt(0)
			firstPart.Mul(&secondFibonacci, big.NewInt(2))

			var secondPart = big.NewInt(0)
			secondPart.Add(firstPart, &firstFibonacci)

			returnValue = returnValue.Mul(secondPart, &firstFibonacci)

		} else {
			firstCheck := (index + 1) / 2
			secondCheck := firstCheck - 1

			firstFibonacci := calcFibonacci(firstCheck)
			secondFibonacci := calcFibonacci(secondCheck)

			var firstSquared = big.NewInt(0)
			var secondSquared = big.NewInt(0)

			firstSquared.Exp(&firstFibonacci, big.NewInt(2), nil)
			secondSquared.Exp(&secondFibonacci, big.NewInt(2), nil)

			var total big.Int
			total = *total.Add(firstSquared, secondSquared)
			returnValue = &total
		}

	}

	return *returnValue
}
