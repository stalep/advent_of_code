package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindAllInvalidIDs(input string, task2 bool) int {
	allInvalidIDs := 0

	for pair := range strings.SplitSeq(input, ",") {
		allInvalidIDs += parseIDpair(pair, task2)
	}

	return allInvalidIDs
}

func parseIDpair(pair string, task2 bool) int {
	invalidIDs := 0

	ids := strings.Split(pair, "-")
	startId, _:= strconv.Atoi(ids[0]);
	endId, endErr:= strconv.Atoi(ids[1]);
	if endErr != nil {
		// fmt.Printf("failed to get endId %d from %s, original pair: %s, err %s\n", endId, ids[1], pair, endErr)
		//the last number contains a \n, lets remove it and try again
		endId, _= strconv.Atoi(strings.Trim(ids[1], "\n"));
	}

	for ; startId <= endId; startId++ {
		if(task2) {
			invalidIDs += verifyID2(startId)
		} else {
			invalidIDs += verifyID(startId)
		}
	}

	return invalidIDs
}

func verifyID(id int) int {

	strId := strconv.Itoa(id)
	if len(strId) % 2 != 0 {
		return 0
	}

	middle := len(strId) / 2

	if strId[:middle] == strId[middle:] {
		return id
	}

	return 0
}

func verifyID2(id int) int {
	strId := strconv.Itoa(id)
	//first check if all chars are identical this should cover all numbers up to 4 digits
	if len(strId) > 1 && allSameChar(strId) {
		return id
	}
	for i := 2; i <= len(strId) / 2; i++ {
		if len(strId) % i == 0 && allSameStrings( divideString(strId, i)) {
			return id
		}
	}

	return 0
}

func divideString(mystr string, size int) []string {
	var parts []string
	partSize := len(mystr) / size
	for i := range size {
		start := i * partSize
		end := start + partSize
		if i == size-1 {
			end = len(mystr)
		}
		parts = append(parts, mystr[start:end])
	}
	return parts
}

func allSameStrings(input []string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] != input[0] {
			return false
		}
	}
	return true
}

func allSameChar(input string) bool {
	for i:= 1; i < len(input); i++ {
		if input[0] != input[i] {
			return false
		}
	}
	return true
}

func main() {

	b, err := os.ReadFile("input")
	if err != nil {
		fmt.Print(err)
	}

	allInvalidIds := FindAllInvalidIDs(string(b), false)

	fmt.Printf("Task1 all invalid IDs: %d\n", allInvalidIds)

	allInvalidIds = FindAllInvalidIDs(string(b), true)

	fmt.Printf("Task2 all invalid IDs: %d\n", allInvalidIds)
}
