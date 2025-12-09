package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IngredientID struct {
	Start, Stop int
}

func (id *IngredientID) NumberOfIDs() int {
	return id.Stop - id.Start + 1
}

func (id *IngredientID) CheckIfNewIDsOverlap(start, stop int) int {
	if (start >= id.Start && start <= id.Stop) && (stop >= id.Start && stop <= id.Stop) {
		return 1
	}
	// start is outside, but stop is inside, then we change current
	if start < id.Start && (id.Start <= stop && stop <= id.Stop) {
		id.Start = start
		return 2
		// check if start and stop is outside
	} else if start < id.Start && stop > id.Stop {
		id.Start = start
		id.Stop = stop
		return 2
		// check if start is inside and stop is greater
	} else if start >= id.Start && start <= id.Stop && stop > id.Start {
		id.Stop = stop
		return 2
	}

	return 0
}

var idDivider = "-"

func CheckFresh(ingredientID IngredientID, id int) bool {
	return ingredientID.Start <= id && id <= ingredientID.Stop
}

func CheckFreshAll(ids []IngredientID, id int) bool {
	for _,i := range ids {
		if CheckFresh(i, id) {
			return true
		}
	}
	return false
}

func ReadDBTask1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer file.Close()

	var ingredientIDs []IngredientID

	scanner := bufio.NewScanner(file)

	numberOfAvailableIngredients := 0

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) > 0 && strings.Contains(input, idDivider) {
			ingredientIDs = append(ingredientIDs, CreateIngredientID(strings.Split(input, idDivider)))
		} else if(len(input) > 0) {
		//we're assuming that we're now reading IDs that we need to check
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Could not convert %s to int", input)
			}

			if CheckFreshAll(ingredientIDs, id) {
				numberOfAvailableIngredients++
			}
		}

	}
		fmt.Printf("Found %d fresh Ingredients\n", numberOfAvailableIngredients)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %s", err)
	}
}


func ReadDBTask2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer file.Close()

	var ingredientIDs = make([]IngredientID, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) > 0 && strings.Contains(input, idDivider) {
			ingredientIDs = CreateIngredientIDWithChecks(ingredientIDs, strings.Split(input, idDivider))
		} 
	}

	fmt.Printf("Total number of available Ingredients: %d\n", TotalnumberOfIngredientIDs(ingredientIDs))

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %s", err)
	}
}


func CreateIngredientID(input []string) IngredientID {
	start, _:= strconv.Atoi( input[0])
	stop, _:= strconv.Atoi( input[1])
	return IngredientID{Start: start, Stop: stop}
}

func CreateIngredientIDWithChecks(existingIDs []IngredientID, input []string) []IngredientID {
	start, _:= strconv.Atoi( input[0])
	stop, _:= strconv.Atoi( input[1])

	var status int

	for i := len(existingIDs)-1; i >= 0; i-- {
		// if check is true, we've updated the start/stop with new values
		// and need to check if the new values invalidate the existing ones
		status = existingIDs[i].CheckIfNewIDsOverlap(start, stop); 
		if status == 2 {
			return CheckForOverlaps(existingIDs, i)
		} else if (status == 1) {
			// start stop is inside another element, return quickly
			return existingIDs
		}
	}
	// we've not added anything, no overlaps, lets just append it
	if status == 0 {
		existingIDs = append(existingIDs, IngredientID{Start: start, Stop: stop})
	}
	return existingIDs
}

func CheckForOverlaps(existingIDs []IngredientID, currentIndex int) []IngredientID {
	id := existingIDs[currentIndex]
	for i := currentIndex-1; i >= 0; i-- {
		if existingIDs[i].CheckIfNewIDsOverlap(id.Start, id.Stop) == 2 {
			existingIDs = append(existingIDs[:currentIndex], existingIDs[currentIndex+1:]...)
			existingIDs = CheckForOverlaps(existingIDs, i)
		}
	}

	return existingIDs
}

func TotalnumberOfIngredientIDs(existingIDs []IngredientID) int {
	total := 0
	for _, id := range existingIDs {
		total += id.NumberOfIDs()
	}

	return total
}

func main() {
	ReadDBTask1("input")
	ReadDBTask2("input")
}
