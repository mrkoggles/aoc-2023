package problem1_test

import (
	"aoc-2023/problem1"
	"testing"
)


type testCase struct {
	filePath string; 
	expectedOutput int;
}

var testCases = []testCase{
	{ "sample-pt1.txt", 142 },
	{ "dataset-pt1.txt", 54601 },
}

func TestProblem1(t *testing.T) {
	for _, test := range testCases {
		actual := problem1.Solve(test.filePath)
		if actual != test.expectedOutput {
			t.Fatalf("Expected %d but was actually: %d", test.expectedOutput, actual)
		}
	}
}




type extractionTestCase struct {
	input string;
	expected int;
}


var extractionTestCases = []extractionTestCase {
	{ "12", 12 }, 
	{ "abc5de1", 51 },
}
func TestExtract(t *testing.T) {
	for _, test := range extractionTestCases {
		
		actual := problem1.ExtractFirstAndLastNumerics(test.input)	
		if actual != test.expected {
			t.Fatalf("Expected: %d but was actually: %d", test.expected, actual);
		}	
	}
}
