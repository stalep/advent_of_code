package main

import (
	"testing"
)

func TestDay4Task1(t *testing.T) {

	paperGrid := ReadFileGeneratePaperGrid("test_example")

	// check if the grid is properly populated
	cases := []Paper {
		{0,0,false, false},
		{0,1, true, false},
		{8,9, true, false},
		{9,7, true, false},
		{9,8, false, false},
		{9,9, false, false},
	}
	
	for _, p := range cases {
		if paperGrid[p.Y][p.X].HasPaper != p.HasPaper {
			t.Errorf("Expected paper[%d][%d] to be %t, but was %t", p.Y, p.X, p.HasPaper, paperGrid[p.Y][p.X].HasPaper)
		}
	}

	//find adjacent papers
	want := 13

	got := FindNumberOfAccessablePapers(paperGrid)

	if want != got {
		t.Errorf("Wanted to find %d accessible papers, but found %d\n", want, got)
	}
}


func TestDay4Task2(t *testing.T) {

	paperGrid := ReadFileGeneratePaperGrid("test_example")
	// Task 2
	want := 43
	got := 1
	totalRemovedRolls := 0

	for i := 0; i < got; i = 0 {
		got = FindNumberOfAccessablePapers(paperGrid)
		totalRemovedRolls += got
		ClearPaperRollsMarkedForRemoval(paperGrid)
	}

	if want != totalRemovedRolls {
		t.Errorf("Wanted to find %d accessible papers, but found %d\n", want, totalRemovedRolls)
	}

}
