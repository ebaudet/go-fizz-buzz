package util

import "fmt"

type FizzParams struct {
	Int1 int
	Int2 int
	Str1 string
	Str2 string
}

func FizzBuzz(i int, f FizzParams) string {
	var output string

	if i%f.Int1 == 0 {
		output += f.Str1
	}
	if i%f.Int2 == 0 {
		output += f.Str2
	}
	if output == "" {
		output = fmt.Sprintf("%d", i)
	}
	return output
}
