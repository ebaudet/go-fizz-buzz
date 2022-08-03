package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/ebaudet/go-fizz-buzz/db/mock"
	db "github.com/ebaudet/go-fizz-buzz/db/sqlc"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestGetStatistics(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				result := db.FizzbuzzStatistic{
					Request:   json.RawMessage(`{"int1":3,"int2":5,"limit":15,"str1":"fizz","str2":"buzz"}`),
					Count:     int64(42),
					UpdatedAt: time.Time{},
					CreatedAt: time.Time{},
				}
				store.EXPECT().GetMostUsedRequest(gomock.Any()).Times(1).Return(result, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchJson(t, recorder.Body, json.RawMessage(`{"request":{"int1":3,"int2":5,"limit":15,"str1":"fizz","str2":"buzz"},"hints":42}`))
			},
		},
		{
			name: "OKEmpty",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetMostUsedRequest(gomock.Any()).Times(1).Return(db.FizzbuzzStatistic{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchJson(t, recorder.Body, json.RawMessage(`{"hints":0, "request":null}`))
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetMostUsedRequest(gomock.Any()).Times(1).Return(db.FizzbuzzStatistic{}, sql.ErrTxDone)
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

			url := "/statistics"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchJson(t *testing.T, body *bytes.Buffer, expected json.RawMessage) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	require.JSONEq(t, string(data), string(expected))
}
