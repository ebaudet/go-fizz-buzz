package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ebaudet/go-fizz-buzz/db/mock"
	db "github.com/ebaudet/go-fizz-buzz/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestGetFizzBuzz(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchString(t, recorder.Body, "\"1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz,16.\"")
			},
		},
		{
			name: "BadRequest",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
			},
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"int1":  3,
				"int2":  5,
				"limit": 16,
				"str1":  "Fizz",
				"str2":  "Buzz",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().IncrementRequest(gomock.Any(), gomock.Any()).Times(1).Return(db.FizzbuzzStatistic{}, sql.ErrTxDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := newTestServer(t, store)
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
