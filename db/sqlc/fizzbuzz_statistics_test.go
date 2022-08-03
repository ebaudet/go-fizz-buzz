package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrementRequest(t *testing.T) {
	request := json.RawMessage(`{
	"int1": 3,
	"int2": 5,
	"limit": 15,
	"str1": "fizz",
	"str2": "buzz"
	}`)

	stat, err := testQueries.IncrementRequest(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, stat)

	require.JSONEq(t, string(request), string(stat.Request))
	count := stat.Count

	request = json.RawMessage(`{
		"int2": 5,
		"str1": "fizz",
		"limit": 15,
		"str2": "buzz",
		"int1": 3
		}`)
	stat, err = testQueries.IncrementRequest(context.Background(), request)
	require.NoError(t, err)
	require.JSONEq(t, string(request), string(stat.Request))
	require.Equal(t, count+1, stat.Count)
}

func TestGetMostUsedRequest(t *testing.T) {
	result, err := testQueries.GetMostUsedRequest(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
