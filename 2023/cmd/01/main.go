package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.ReadFile("cmd/01/input.txt")
	if err != nil {

		log.Fatal(err)
	}

	input := string(file)
	lines := strings.Split(input, "\n")

	stringToInt := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	}
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitReg := regexp.MustCompile("[1-9]")

	sum := 0
	for _, line := range lines {
		indices := digitReg.FindAllStringIndex(line, -1)
		for _, digit := range digits {
			spelledDigitReg := regexp.MustCompile(digit)
			indices = append(indices, spelledDigitReg.FindAllStringIndex(line, -1)...)
		}
		sort.Slice(indices, func(i, j int) bool {
			return indices[i][0] < indices[j][0]
		})
		if len(indices) > 0 {
			sum += 10*stringToInt[line[indices[0][0]:indices[0][1]]] +
				stringToInt[line[indices[len(indices)-1][0]:indices[len(indices)-1][1]]]
		}

	}

	fmt.Println(sum)
}
