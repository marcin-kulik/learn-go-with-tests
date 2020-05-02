package iteration

import "testing"

func TestRepeat(t *testing.T) {

	checkIfValid := func(t *testing.T, repeated string, expected string) {

		t.Helper()

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("iterate 2 times", func(t *testing.T) {
		repeated := Repeat(2, "a")
		expected := ("aa")

		checkIfValid(t, repeated, expected)
	})

	t.Run("iterate 5 times", func(t *testing.T) {
		repeated := Repeat(5, "a")
		expected := ("aaaaa")

		checkIfValid(t, repeated, expected)
	})

}
