package problem1

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
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
	return sum;
}

func ExtractFirstAndLastNumerics(input string) (int) {
	
	var first int = -1 
	var last int = -1
	for _, char := range input {
		charAsString := string(byte(char))
		numeral, err := strconv.Atoi(charAsString)
		if err != nil {
			continue
		}

		last = numeral
		if first == -1 {
			first = numeral
		}
		
	}

	outputStr := fmt.Sprintf("%d%d", first, last)
	output, err := strconv.Atoi(outputStr)	
	if err != nil {
		panic(err)
	}
	return output 
}
