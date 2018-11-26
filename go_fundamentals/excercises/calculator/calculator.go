package main

import (
	"fmt"
	"go_bootcamp/go_fundamentals/excercises/operations"
	"os"
	"strconv"
)

func main() {
	x, xErr := strconv.Atoi(os.Args[1])
	y, yErr := strconv.Atoi(os.Args[3])

	if xErr != nil {
		fmt.Println(xErr)
		return
	}

	if yErr != nil {
		fmt.Println(yErr)
		return
	}

	switch operand := os.Args[2]; operand {
	case "+":
		operations.Sum(x, y)
	case "-":
		operations.Sub(x, y)
	case "*":
		operations.Mul(x, y)
	case "/":
		operations.Div(x, y)
	default:
		fmt.Printf("Operand %s is not valid", operand)

	}
}
