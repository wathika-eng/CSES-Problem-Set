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
	inputOne := askUser()
	if inputOne <= 0 {
		log.Fatal("err")
	}
	results := divide(inputOne)
	fmt.Println(results)
}

func askUser() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	cleanedInput, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}
	return cleanedInput
}

func divide(n int) []int {
	nums := make([]int, 0, 1000000)
	nums = append(nums, n)
	for n > 1 {

		if n%2 == 0 {
			n /= 2
		} else {
			n = (n * 3) + 1
		}
		nums = append(nums, n)
	}

	return nums
}
