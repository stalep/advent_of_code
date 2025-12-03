package main

import "testing"


func TestInvalidIDs(t *testing.T) {
	productIDs := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	want := 1227775554

	addedInvalidIDs := FindAllInvalidIDs(productIDs, false)

	if addedInvalidIDs != want {
		t.Errorf("wanted %d, got %d", want, addedInvalidIDs)
	}

	// now test task 2
	want = 4174379265

	addedInvalidIDs = FindAllInvalidIDs(productIDs, true)

	if addedInvalidIDs != want {
		t.Errorf("wanted %d, got %d", want, addedInvalidIDs)
	}

}
