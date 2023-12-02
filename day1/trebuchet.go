package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCalibrationSum(filePath string) int {
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstNum rune
		var secondNum rune

		// find first number
		for _, c := range line {
			if unicode.IsNumber(c) {
				firstNum = c
			}
		}

		// find second number
		runes := []rune(line)
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsNumber(runes[i]) {
				secondNum = runes[i]
			}
		}
		calibrationStr := string(secondNum) + string(firstNum)
		calibrationNum, e := strconv.ParseInt(calibrationStr, 0, 32)
		check(e)
		sum += int(calibrationNum)
	}
	return sum
}
