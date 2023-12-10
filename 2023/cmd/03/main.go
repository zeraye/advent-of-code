package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSymbol(char byte) bool {
	return (char < '0' || char > '9') && char >= '!' && char <= '~' && char != '.'
}

// check if symbol is around number
func symbolAroundDigit(x, y int, lines []string) bool {
	if x-1 >= 0 && y-1 >= 0 && isSymbol(lines[x-1][y-1]) {
		return true
	}
	if x-1 >= 0 && isSymbol(lines[x-1][y]) {
		return true
	}
	if x-1 >= 0 && y+1 < len(lines[x-1]) && isSymbol(lines[x-1][y+1]) {
		return true
	}
	if y-1 >= 0 && isSymbol(lines[x][y-1]) {
		return true
	}
	if y+1 < len(lines[x]) && isSymbol(lines[x][y+1]) {
		return true
	}
	if x+1 < len(lines) && y-1 >= 0 && isSymbol(lines[x+1][y-1]) {
		return true
	}
	if x+1 < len(lines) && isSymbol(lines[x+1][y]) {
		return true
	}
	if x+1 < len(lines) && y+1 < len(lines[x+1]) && isSymbol(lines[x+1][y+1]) {
		return true
	}
	return false
}

// check if symbol is around number
// (x, y) are coords of first element of number
func symbolAroundNumber(x, y, length int, lines []string) (bool, int) {
	isSymbolAround := false
	stringNumber := ""
	i := 0
	for i < length {
		if symbolAroundDigit(x, y+i, lines) {
			isSymbolAround = true
		}
		stringNumber += string(lines[x][y+i])
		i++
	}
	number, _ := strconv.Atoi(stringNumber)
	return isSymbolAround, number
}

func main() {
	file, err := os.ReadFile("cmd/03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	lines := strings.Split(input, "\n")

	// Part One
	sum := 0
	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}
		j := 0
		length := 0
		for j < len(lines[i]) {
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				length++
			} else {
				if length > 0 {
					isSymbolAround, number := symbolAroundNumber(i, j-length, length, lines)
					if isSymbolAround {
						sum += number
					}
				}
				length = 0
			}
			j++
		}
		// chack again for the last row
		if length > 0 {
			isSymbolAround, number := symbolAroundNumber(i, j-length, length, lines)
			if isSymbolAround {
				sum += number
			}
		}
	}
	fmt.Println(sum)

	// Part Two
	// TODO
}
