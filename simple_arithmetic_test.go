package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Add was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestAddTable(t *testing.T) {
	testTable := []struct{ a, b, result int }{
		{1, 1, 2},
		{5, 5, 10},
		{2, 3, 5},
	}

	for _, test := range testTable {
		total := Add(test.a, test.b)
		t.Logf("Testing add %d, %d", test.a, test.b)
		if total != test.result {
			t.Errorf("Addition failed for %d, %d. Got: %d Expected: %d", test.a, test.b, total, test.result)
		}
	}
}

func TestSubtract(t *testing.T) {
	total := Subtract(5, 5)
	if total != 0 {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestMultiply(t *testing.T) {
	total := Multiply(5, 5)
	if total != 25 {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 25)
	}
}

func TestDivide(t *testing.T) {
	total := Divide(5, 5)
	if total != 1 {
		t.Errorf("Divide was incorrect, got: %d, want: %d.", total, 1)
	}
}
