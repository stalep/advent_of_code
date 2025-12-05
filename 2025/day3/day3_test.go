package main

import "testing"

func TestDay3Part1(t *testing.T) {
	input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}

	want := 357

	got := TotalOutputJoltage(input, false)

	if want != got {
		t.Errorf("Task1, wanted %d, got %d", want, got)
	}

	// Task 2
	want = 3121910778619

	got = TotalOutputJoltage(input, true)

	if want != got {
		t.Errorf("Task2, wanted %d, got %d", want, got)
	}

}
