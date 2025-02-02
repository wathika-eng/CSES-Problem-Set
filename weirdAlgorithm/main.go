package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	results := divide(inputOne)
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
	if cleanedInput <= 0 {
		return -1, fmt.Errorf("negative input: %d is not a positive number", cleanedInput)
	}
	return cleanedInput, nil
}

// a recursive function
func divide(n int) []int {
	// make a slice allocated space of 10^6
	nums := make([]int, 0, 1000000)
	// append the first number (userinput)
	nums = append(nums, n)
	// run recursivey till n is less than 1
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (n * 3) + 1
		}
		// append the results of each function call
		nums = append(nums, n)
	}
	// return final results as a slice
	return nums
}
