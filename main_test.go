package main

import "testing"

// Standardfälle
func TestGetMinuteUp(t *testing.T) {
	res := getMinute(1, 30)
	if res != 2 {
		t.Errorf("getMinute(1, 30) = %d; wanted 2", res)
	}
}

func TestGetMinuteDown(t *testing.T) {
	res := getMinute(1, 15)
	if res != 1 {
		t.Errorf("getMinute(1, 15) = %d; wanted 1", res)
	}
}

// Grenzwerte testen
func TestGetMinuteEdgeCases(t *testing.T) {
	res1 := getMinute(1, 29)
	if res1 != 1 {
		t.Errorf("getMinute(1, 29) = %d; wanted 1", res1)
	}

	res2 := getMinute(1, 30)
	if res2 != 2 {
		t.Errorf("getMinute(1, 30) = %d; wanted 2", res2)
	}
}

// Negative Werte testen
func TestGetMinuteNegative(t *testing.T) {
	res1 := getMinute(-1, 30)
	if res1 != 0 {
		t.Errorf("getMinute(-1, 30) = %d; wanted 0", res1)
	}

	res2 := getMinute(-5, 20)
	if res2 != -5 {
		t.Errorf("getMinute(-5, 20) = %d; wanted -5", res2)
	}
}

// Große Zahlen testen
func TestGetMinuteLargeValues(t *testing.T) {
	res1 := getMinute(1000, 29)
	if res1 != 1000 {
		t.Errorf("getMinute(1000, 29) = %d; wanted 1000", res1)
	}

	res2 := getMinute(1000, 30)
	if res2 != 1001 {
		t.Errorf("getMinute(1000, 30) = %d; wanted 1001", res2)
	}
}
