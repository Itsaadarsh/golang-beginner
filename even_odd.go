package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your AGE : ")
	scanner.Scan()
	inputNum, err := strconv.Atoi(scanner.Text())
	if err == nil {
		if inputNum%2 == 0 {
			fmt.Printf("\n%v is EVEN", inputNum)
		} else {
			fmt.Printf("\n%v is ODD", inputNum)
		}
	} else {
		fmt.Printf("\nEnter only integer type")
	}
}
