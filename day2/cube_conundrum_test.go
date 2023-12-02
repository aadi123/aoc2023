package main

import "testing"

func TestGetValidGameSum(t *testing.T) {
	want := 8
	if got := getValidGameSum("./day2_test.txt"); got != want {
		t.Errorf("getValidGameSum() = %d, want %d", got, want)
	}
}
