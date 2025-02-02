package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var inputOne int
	var err error
	for {
		inputOne, err = askUser()
		if err == nil && inputOne > 0 {
			break
		}
		log.Println(err)
	}
	results := missing(inputOne)
	fmt.Println(results)
}

// ask user for a positive number
//
// returns -1 if error is present else nil
func askUser() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a positive number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return -1, fmt.Errorf("error reading input: %v", err)
	}
	cleanedInput, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1, fmt.Errorf("error while converting integer: %v", err)
	}
	if cleanedInput <= 1 {
		return -1, fmt.Errorf("negative input or less than 2: %d is not a positive number", cleanedInput)
	}
	return cleanedInput, nil
}

func missing(n int, arr []int) int {
	slices.Sort(arr)
	
}
