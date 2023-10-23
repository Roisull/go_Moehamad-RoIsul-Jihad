package calculate

import "testing"

func TestAddition(t *testing.T) {
	result := Addition(3, 4)
	if result != 7 {
		t.Errorf("Expected 7, got %d", result)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(7, 3)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestMult(t *testing.T) {
	result := Mult(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestDiv(t *testing.T) {
	result := Div(8, 2)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}