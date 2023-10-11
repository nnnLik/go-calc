package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var isProcessing bool = true
	scanner := bufio.NewScanner(os.Stdin)

	for isProcessing {
		var userInput string

		fmt.Print("Enter an expression (1 + 2): ")

		scanner.Scan()
		userInput = scanner.Text()

		if userInput == "exit" {
			isProcessing = false
			fmt.Println("Exiting the program.")
			continue
		}

		result, err := calculateExpression(userInput)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}
