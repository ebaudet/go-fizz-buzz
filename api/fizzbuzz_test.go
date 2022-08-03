package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetFizzBuzz(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"int1":  3,
				"int2":  5,
				"limit": 16,
				"str1":  "Fizz",
				"str2":  "Buzz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchString(t, recorder.Body, "\"1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz,16.\"")
			},
		},
		{
			name: "BadRequest",
			body: gin.H{},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Int1TooSmall",
			body: gin.H{
				"int1":  0,
				"int2":  5,
				"limit": 16,
				"str1":  "Fizz",
				"str2":  "Buzz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Int2TooSmall",
			body: gin.H{
				"int1":  3,
				"int2":  0,
				"limit": 16,
				"str1":  "Fizz",
				"str2":  "Buzz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "LimitTooSmall",
			body: gin.H{
				"int1":  3,
				"int2":  5,
				"limit": 0,
				"str1":  "Fizz",
				"str2":  "Buzz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "LimitTooHigh",
			body: gin.H{
				"int1":  3,
				"int2":  5,
				"limit": 100000000000,
				"str1":  "Fizz",
				"str2":  "Buzz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Str1IsRequired",
			body: gin.H{
				"int1":  3,
				"int2":  5,
				"limit": 16,
				"str2":  "Buzz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Str2IsRequired",
			body: gin.H{
				"int1":  3,
				"int2":  5,
				"limit": 16,
				"str1":  "Fizz",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/fizzbuzz"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchString(t *testing.T, body *bytes.Buffer, expected string) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	require.Equal(t, string(data), expected)
}
