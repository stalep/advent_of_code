package main

import "testing"

func TestRead(t *testing.T) {
	t.Run("test assignment part one", func(t *testing.T) {

		input := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

		password := SolvePasswordA(input)
		want := 3

		if password != want {
			t.Errorf("got %d, wanted %d", password, want)
		}
	})

	cases := []struct{
		Input []string
		Want int
		Description string
	}{
		{
		[]string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82", "R268"},
		9,
		"test assignment part two",
		},
		{
		[]string{"R50", "R50", "L50", "L50", "R75", "L50", "L25", "L75", "R50"},
			6,
			"test assignment part two B",
		},
		{
		[]string{"L499", "L51"},
			6,
			"test assignment part two C",
		},
	{
		[]string{"R50", "R50", "L50", "L50"},
			2,
			"test assignment part two D",
		},
	{
		[]string{"R50", "R50", "L50", "L50", "R75"},
			3,
			"test assignment part two E",
		},
	{
		[]string{"R50", "R50", "L50", "L50", "R75", "L50"},
			4,
			"test assignment part two F",
		},
	{
		[]string{"R50", "R50", "L50", "L50", "R75", "L50", "L25"},
			4,
			"test assignment part two G",
		},
	{
		[]string{"R50", "R50", "L50", "L50", "R75", "L50", "L25", "L499"},
			9,
			"test assignment part two H",
		},
	{
		[]string{"R499", "R51"},
			6,
			"test assignment part two J",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			password := SolvePasswordB(test.Input)

			if password != test.Want {
				t.Errorf("got %d, wanted %d", password, test.Want)
			}
		})
	}

}
