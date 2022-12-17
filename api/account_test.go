package api

import (
	"database/sql"
	"bytes"
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"net/http/httptest"
	"github.com/techschool/simplebank/db/mock"
	"github.com/stretchr/testify/require"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/techschool/simplebank/util"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/db/sqlc"
	"time"
)

func TestGetAccountAPI(t *testing.T){
	user, _ := randomUser(t)
	account := randomAccount(user.Username)
	testCases :=[]struct{
		name string
		accountID int64
		setupAuth func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs func(store *mockdb.MockStore)
		checkResonse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().
				GetAccount(gomock.Any(),gomock.Eq(account.ID)).
				Times(1).
				Return(account,nil)
			},
			checkResonse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t,recorder.Body,account)
			},
		},
		{
			name: "UnauthorizedUser",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().
				GetAccount(gomock.Any(),gomock.Eq(account.ID)).
				Times(1).
				Return(account,nil)
			},
			checkResonse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Any()).
				Times(0)
			},
			checkResonse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NotFound",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().
				GetAccount(gomock.Any(),gomock.Eq(account.ID)).
				Times(1).
				Return(db.Account{},sql.ErrNoRows)
			},
			checkResonse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().
				GetAccount(gomock.Any(),gomock.Eq(account.ID)).
				Times(1).
				Return(db.Account{},sql.ErrConnDone)
			},
			checkResonse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			accountID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().
				GetAccount(gomock.Any(),gomock.Any).
				Times(0)
			},
			checkResonse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		//TODO: add more cases
	}

	

	for i:= range testCases{
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T){
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			
			//start test server and send request 
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d" , tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t,err)

			tc.setupAuth(t,request, server.tokenMaker)
			server.router.ServeHTTP(recorder,request)
			tc.checkResonse(t,recorder)
		})
	}
}

func randomAccount(owner string) db.Account{
	return db.Account{
		ID:util.RandomeInt(1,1000),
		Owner:owner,
		Balance:util.RandomMoney(),
		Currency:util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account){
	data, err := ioutil.ReadAll(body)
	require.NoError(t,err)

	var gotAccount db.Account
	err = json.Unmarshal(data,&gotAccount)
	require.NoError(t,err)
	require.Equal(t, account, gotAccount)
}