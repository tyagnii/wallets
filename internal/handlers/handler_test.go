package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mock_db "github.com/tyagnii/wallets/gen/mock"
	"go.uber.org/mock/gomock"
)

func TestMain(m *testing.M) {
}

func TestNewHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mdb := mock_db.NewMockDBConnector(mockCtrl)
	_, err := NewHandler(mdb)
	require.NoError(t, err)
}

func TestWallet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mdb := mock_db.NewMockDBConnector(mockCtrl)
	h, _ := NewHandler(mdb)

	r := NewRouter(h)

	// Test CreateWallet
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/wallets/create", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "*", w.Body)

	// Test ChangeBalance
	pr := PostAmountRequest{
		WalletID:      w.Body.String(),
		OperationType: "DEPOSIT",
		Amount:        "100",
	}

	prJSON, _ := json.Marshal(pr)
	req, _ = http.NewRequest("POST", "/api/v1/wallet", strings.NewReader(string(prJSON)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(prJSON), w.Body.String())

	// Test GetBalance
	req, _ = http.NewRequest("GET", "/api/v1/wallet/"+pr.WalletID, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(prJSON), w.Body.String())
}
