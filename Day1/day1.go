package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var digitmap = map[string]int{
	"one"	:	1,
	"two"	:	2,
	"three"	:	3,
	"four"	:	4,
	"five"	:	5,
	"six"	:	6,
	"seven"	:	7,
	"eight"	:	8,
	"nine"	:	9,
}

func Solve() int {
	lines := readFile("day1/input.txt")

	sum := 0
	for _, line := range lines {
		sum = sum + buildInt(line)
	}

	return sum
}

func buildInt(inputStr string) int {
	runes := []rune(inputStr)
	first := 0
	last := 0

	for i := range runes {
		parsedVal, err := strconv.Atoi(string(runes[i]))
		if err == nil {
			if first == 0 {
				first = parsedVal
			}
			last = parsedVal
		} else {
			var buildDigitStr string
			for j := i; j < len(runes); j++ {
				buildDigitStr += string(runes[j])
				mapValue, found := digitmap[buildDigitStr]
				if found {
					if first == 0 {
						first = mapValue
					}
					last = mapValue
				}
			}
		}
	}

	buildInt, _ := strconv.Atoi(strconv.Itoa(first)+strconv.Itoa(last))

	return buildInt
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}