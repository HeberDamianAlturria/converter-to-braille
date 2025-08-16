package braille

import (
	"fmt"
	"testing"
)

type TestCase struct {
	testName string
	matrix   BrailleMatrix
	expected rune
}

func TestBrailleMatrixToRune(t *testing.T) {
	testCases := []TestCase{
		{
			testName: "All dots down",
			matrix:   BrailleMatrix{{0, 0, 0, 0}, {0, 0, 0, 0}},
			expected: '\u2800',
		},
		{
			testName: "Dot 1 up",
			matrix:   BrailleMatrix{{1, 0, 0, 0}, {0, 0, 0, 0}},
			expected: '\u2801',
		},
		{
			testName: "Dot 2 up",
			matrix:   BrailleMatrix{{0, 1, 0, 0}, {0, 0, 0, 0}},
			expected: '\u2802',
		},
		{
			testName: "Dot 3 up",
			matrix:   BrailleMatrix{{0, 0, 1, 0}, {0, 0, 0, 0}},
			expected: '\u2804',
		},
		{
			testName: "Dot 4 up",
			matrix:   BrailleMatrix{{0, 0, 0, 1}, {0, 0, 0, 0}},
			expected: '\u2840',
		},
		{
			testName: "Dot 5 up",
			matrix:   BrailleMatrix{{0, 0, 0, 0}, {1, 0, 0, 0}},
			expected: '\u2808',
		},
		{
			testName: "Dot 6 up",
			matrix:   BrailleMatrix{{0, 0, 0, 0}, {0, 1, 0, 0}},
			expected: '\u2810',
		},
		{
			testName: "Dot 7 up",
			matrix:   BrailleMatrix{{0, 0, 0, 0}, {0, 0, 1, 0}},
			expected: '\u2820',
		},
		{
			testName: "Dot 8 up",
			matrix:   BrailleMatrix{{0, 0, 0, 0}, {0, 0, 0, 1}},
			expected: '\u2880',
		},
		{
			testName: "All dots up",
			matrix:   BrailleMatrix{{1, 1, 1, 1}, {1, 1, 1, 1}},
			expected: '\u28FF',
		},
		{
			testName: "Random pattern",
			matrix:   BrailleMatrix{{1, 0, 1, 0}, {1, 1, 0, 0}},
			expected: '\u281D',
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.matrix.ToRune()
			fmt.Printf("Result: %s\n", string(result))
			if result != tc.expected {
				t.Errorf("Expected %U, but got %U", tc.expected, result)
			}
		})
	}
}
