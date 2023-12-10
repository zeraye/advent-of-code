package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == ',' || r == ';'
}

func main() {
	file, err := os.ReadFile("cmd/02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	lines := strings.Split(input, "\n")

	// Part One
	maxColorMap := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}

	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		colonSplit := strings.Split(line, ":")

		id, _ := strconv.Atoi(strings.Split(colonSplit[0], " ")[1])
		sum += id

		for _, cube := range strings.FieldsFunc(colonSplit[1], split) {
			countColor := strings.Split(strings.TrimSpace(cube), " ")
			count, _ := strconv.Atoi(countColor[0])
			color := countColor[1]
			if count > maxColorMap[color] {
				sum -= id
				break
			}
		}
	}
	fmt.Println(sum)

	// Part Two
	sum = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		minColorMap := map[string]int{
			"red": 0, "green": 0, "blue": 0,
		}

		colonSplit := strings.Split(line, ":")

		for _, cube := range strings.FieldsFunc(colonSplit[1], split) {
			countColor := strings.Split(strings.TrimSpace(cube), " ")
			count, _ := strconv.Atoi(countColor[0])
			color := countColor[1]
			if count > minColorMap[color] {
				minColorMap[color] = count
			}
		}
		sum += minColorMap["red"] * minColorMap["green"] * minColorMap["blue"]
	}
	fmt.Println(sum)
}
