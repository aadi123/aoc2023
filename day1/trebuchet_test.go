package main

import "testing"

func TestGetCaliBrationSum(t *testing.T) {
	want := 142
	if got := getCalibrationSum("./day1_test.txt"); got != want {
		t.Errorf("getCalibrationSum() = %d, want %d", got, want)
	}
}
