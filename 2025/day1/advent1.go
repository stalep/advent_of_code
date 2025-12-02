package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var MOD = 100

func SolvePasswordA(input []string) int {
	currentValue := 50
	password := 0

	for _, i := range input {
		currentValue, password = SolveOneWordA(i, currentValue, password)
	}

	return password
}

func SolveOneWordA(input string, currentValue, password int) (int, int) {
		movement, _:= strconv.Atoi( input[1:])

		if input[0] == 'L' {
			currentValue -= movement
		} else {
			currentValue += movement
		}

		currentValue = (currentValue + MOD ) % MOD

		if currentValue == 0 {
			password ++
		}

	return currentValue, password
}

func Abs(x int) int {
	if x < 0 {
	return -x
	}
	return x
}

func SolvePasswordB(input []string) int {
	currentValue := 50
	password := 0

	for _, i := range input {
		currentValue, password = SolveOneWordB(i, currentValue, password)
	}

	return password
}

func SolveOneWordB(input string, currentValue, password int) (int, int) {
	m, _:= strconv.Atoi( input[1:])

	if input[0] == 'L' {
		m = m*-1
	}

	password += Abs(m) / 100
	m = m % MOD
	pos := currentValue + m

	switch {
	case pos == 0:
		password++

	case pos > 99:
		password++

	case pos < 0 && currentValue > 0:
		password++
	}

	return (pos + MOD ) % MOD, password
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	currentValueA := 50
	passwordA := 0
	currentValueB := 50
	passwordB := 0


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		currentValueA, passwordA = SolveOneWordA(line, currentValueA, passwordA)
		currentValueB, passwordB = SolveOneWordB(line, currentValueB, passwordB)
	}

	fmt.Printf("Password for the first assignment is: %d\n", passwordA)
	fmt.Printf("Password for the second assignment is: %d\n", passwordB)

}
