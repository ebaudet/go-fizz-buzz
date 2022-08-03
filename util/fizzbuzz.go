package util

import (
	"errors"
	"fmt"
)

type FizzParams struct {
	Int1 int
	Int2 int
	Str1 string
	Str2 string
}

func FizzBuzz(i int, f FizzParams) (string, error) {
	var output string

	if f.Int1 == 0 {
		return "", errors.New("int1 cannot be 0")
	}
	if f.Int2 == 0 {
		return "", errors.New("int2 cannot be 0")
	}

	if i%f.Int1 == 0 {
		output += f.Str1
	}
	if i%f.Int2 == 0 {
		output += f.Str2
	}
	if output == "" {
		output = fmt.Sprintf("%d", i)
	}
	return output, nil
}
