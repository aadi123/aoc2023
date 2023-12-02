package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const MAX_R = 12
const MAX_G = 13
const MAX_B = 14

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getValidGameSum(filePath string) int {
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameRegexp := regexp.MustCompile(`Game (?P<game>\d+):`)
		match := gameRegexp.FindStringSubmatch(line)
		gameNo := match[1]
		fmt.Println(gameNo)
	}
	return sum
}
