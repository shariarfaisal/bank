package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/shariarfaisal/bank/db/mock"
)

func TestCreateTransfer(t *testing.T) {

	type request struct{}

	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:          "OK",
			buildStubs:    func(store *mockdb.MockStore) {},
			checkResponse: func(recorder *httptest.ResponseRecorder) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			store := mockdb.NewMockStore(ctrl)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			request := httptest.NewRequest(http.MethodPost, "/transfers", nil)
			server.router.ServeHTTP(recorder, request)
		})
	}
}
