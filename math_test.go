package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(10, 10)

	if total != 20 {
		t.Errorf("Expected %v but got %v", 20, total)
	}
}
