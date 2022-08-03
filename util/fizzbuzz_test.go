package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T) {
	f := FizzParams{
		Int1: 3,
		Int2: 5,
		Str1: "fizz",
		Str2: "buzz",
	}
	require.Equal(t, fizzBuzzNoError(t, 0, f), "fizzbuzz")
	require.Equal(t, fizzBuzzNoError(t, 1, f), "1")
	require.Equal(t, fizzBuzzNoError(t, 3, f), "fizz")
	require.Equal(t, fizzBuzzNoError(t, 5, f), "buzz")
	require.Equal(t, fizzBuzzNoError(t, 15, f), "fizzbuzz")

	f = FizzParams{
		Int1: 6,
		Int2: 7,
		Str1: "alice",
		Str2: "bob",
	}
	require.Equal(t, fizzBuzzNoError(t, 0, f), "alicebob")
	require.Equal(t, fizzBuzzNoError(t, 1, f), "1")
	require.Equal(t, fizzBuzzNoError(t, 6, f), "alice")
	require.Equal(t, fizzBuzzNoError(t, 7, f), "bob")
	require.Equal(t, fizzBuzzNoError(t, 42, f), "alicebob")

	f = FizzParams{
		Int1: 0,
		Int2: 5,
		Str1: "fizz",
		Str2: "buzz",
	}
	result, err := FizzBuzz(1, f)
	require.Error(t, err)
	require.Empty(t, result)

	f = FizzParams{
		Int1: 3,
		Int2: 0,
		Str1: "fizz",
		Str2: "buzz",
	}
	result, err = FizzBuzz(1, f)
	require.Error(t, err)
	require.Empty(t, result)
}

func fizzBuzzNoError(t *testing.T, i int, f FizzParams) string {
	result, err := FizzBuzz(i, f)
	require.NoError(t, err)
	return result
}
