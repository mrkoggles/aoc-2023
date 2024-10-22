package problem1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(filePath string) int {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	var sum int = 0
	for scanner.Scan() {
		result := ExtractFirstAndLastNumerics(scanner.Text())
		sum = sum + result
	}
	return sum
}

func ExtractFirstAndLastNumerics(input string) int {

	var first int = -1
	var last int = -1
	for i := 0; i < len(input); i++ {

		numeral, err := PrefixIsNumeral(input[i:])
		if err != nil {
			continue
		}

		last = numeral
		if first == -1 {
			first = numeral
		}
	}

	if (first == -1) || (last == -1) {
		panic("Could not extract first and last numerics")
	}

	outputStr := fmt.Sprintf("%d%d", first, last)
	output, err := strconv.Atoi(outputStr)
	if err != nil {
		panic(err)
	}
	return output
}

var numerals = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func PrefixIsNumeral(input string) (int, error) {

	for index, numeral := range numerals {
		indexAsString := strconv.Itoa(index)
		isMatch := strings.HasPrefix(input, indexAsString) || strings.HasPrefix(input, numeral)
		if isMatch {
			return index, nil
		}
	}
	return 0, fmt.Errorf(`Could not find a prefix that matched`)
}
