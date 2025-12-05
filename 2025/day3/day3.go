package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func TotalOutputJoltage(input []string, task2 bool) int {
	total := 0
	for _,line := range input {
		if task2 {
		total += findJoltagePart2(line)
		} else {
		total += findJoltage(line)
		}
	}
	return total
}

func findJoltage(bank string) int {
	tmp := 0
	max10 := 1
	max1 := 1
	for i := 0; i < len(bank)-1; i++ {
		tmp = int(bank[i]-'0')
		if tmp > max10 {
			max1 = 1
			max10 = tmp
			continue
		}
		if tmp > max1 {
			max1 = tmp
		}
	}
		tmp = int(bank[len(bank)-1]-'0')
	if(tmp > max1) {
		max1 = tmp
	}
	
	return max10 * 10 + max1
}

type JoltBank struct {
	Joltage [12]int
}
// 0123456789ABCDE
// 234234234234278
func (j *JoltBank) readNewValue(value, numLeftInInput int) {
	for i := 0; i < len(j.Joltage); i++ {
		if numLeftInInput >= (11-i) && j.Joltage[i] < value {
			j.Joltage[i] = value
			j.clearRemainingJolts(i)
			break
		}
	}
}

func (j *JoltBank) clearRemainingJolts(index int) {
	for index++; index < len(j.Joltage); index++ {
		j.Joltage[index] = 1
	}
}

func (j *JoltBank) summarize() int {
	sum := 0
	for i := 0; i < len(j.Joltage); i++ {
		sum += int( math.Pow(float64(10), float64(11-i))) * j.Joltage[i]
	}
	return sum
}

func findJoltagePart2(bank string) int {
	// the bank can now store 12 digits
	// can only update joltage[i] when i > len(bank)-12o
	joltBank := JoltBank{Joltage: [12]int{}}
	bankLenght := len(bank)
	tmp := 0
	for i := range bankLenght {
		tmp = int(bank[i]-'0')
		joltBank.readNewValue(tmp, bankLenght - i - 1)
	}

	return joltBank.summarize()
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltageTask1 := 0
	totalJoltageTask2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			totalJoltageTask1 += findJoltage(line)
			totalJoltageTask2 += findJoltagePart2(line)
		}
	}

	fmt.Printf("Total Joltage task1: %d\n", totalJoltageTask1)
	fmt.Printf("Total Joltage task2: %d\n", totalJoltageTask2)

}
