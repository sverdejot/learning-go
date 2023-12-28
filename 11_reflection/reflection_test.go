package reflex

import (
	"slices"
	"testing"
)

type Person struct {
	Name  string
	Birth Date
}

type Date struct {
	Day   int
	Month string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with single string field",
			struct {
				Name string
			}{"Samuel"},
			[]string{"Samuel"},
		},
		{
			"struct with two string fields",
			struct {
				Name    string
				Surname string
			}{"Samuel", "Verdejo"},
			[]string{"Samuel", "Verdejo"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Samuel", 25},
			[]string{"Samuel"},
		},
		{
			"struct with nested structs",
			Person{"Samuel", Date{4, "March"}},
			[]string{"Samuel", "March"},
		},
		{
			"pointer instead of copy",
			&Person{"Samuel", Date{4, "March"}},
			[]string{"Samuel", "March"},
		},
		{
			"slices",
			[]Date{
				{4, "March"},
				{28, "April"},
			},
			[]string{"March", "April"},
		},
		{
			"arrays",
			[2]Date{
				{4, "March"},
				{28, "April"},
			},
			[]string{"March", "April"},
		},
	}

	for _, cse := range cases {
		t.Run(cse.Name, func(t *testing.T) {
			var got []string

			walk(cse.Input, func(input string) {
				got = append(got, input)

			})

			if !slices.Equal(cse.ExpectedCalls, got) {
				t.Errorf("got %v want %v", got, cse.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Samuel": "Verdejo",
			"Martin": "Fowler",
		}
		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Verdejo")
		assertContains(t, got, "Fowler")
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Date)

		go func() {
			ch <- Date{4, "March"}
			ch <- Date{28, "April"}
			close(ch)
		}()

		var got []string
		want := []string{"March", "April"}

		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with funcs", func(t *testing.T) {
		fn := func() (Date, Date) {
			return Date{4, "March"}, Date{28, "April"}
		}

		var got []string
		want := []string{"March", "April"}

		walk(fn, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, stack []string, got string) {
	t.Helper()
	contains := false

	for _, x := range stack {
		if x == got {
			contains = true
		}
	}

	if !contains {
		t.Errorf("%s is not in %q", got, stack)
	}
}
