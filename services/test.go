package services

import "errors"

type CasesString struct {
	name     string
	data     [][]string
	expected string
}

type CasesNumber struct {
	name     string
	data     [][]string
	expected int64
	err      error
}

func getDefaultMatrix() [][]string {
	return [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
}

// the cases where math is involved are a bit more complex
func getCasesNumber(values []int64) []CasesNumber {

	return []CasesNumber{
		{
			name:     "Success",
			data:     getDefaultMatrix(),
			expected: values[0],
			err:      nil,
		},
		{
			name:     "Negative number",
			data:     [][]string{{"1", "2"}, {"3", "-4"}},
			expected: values[1],
			err:      nil,
		},
		{
			name:     "Conversion error",
			data:     [][]string{{"1", "W"}, {"3", "4"}},
			expected: 0,
			err:      errors.New("Error converting W to integer"),
		},
		{
			name: "Big matrix",
			data: [][]string{
				{"1", "2", "3", "4", "5", "6"},
				{"7", "8", "9", "10", "11", "12"},
				{"13", "14", "15", "16", "17", "18"},
				{"19", "20", "21", "22", "23", "24"},
				{"25", "26", "27", "28", "29", "30"},
				{"31", "32", "33", "34", "35", "36"},
			},
			expected: values[2],
			err:      nil,
		},
		{
			name:     "Matrix with zero",
			data:     [][]string{{"1", "0"}, {"3", "4"}},
			expected: values[3],
			err:      nil,
		},
	}
}

// the string cases are simple and have no errors
func getCasesString(name []string, expected []string) []CasesString {
	return []CasesString{
		{
			name:     name[0],
			data:     getDefaultMatrix(),
			expected: expected[0],
		},
		{
			name:     name[1],
			data:     [][]string{{"9", "8", "7", "0"}, {"6", "5", "4", "0"}, {"3", "2", "1", "0"}, {"0", "0", "0", "0"}},
			expected: expected[1],
		},
	}
}
