package handlers

import (
	"github.com/tyagnii/wallets/internal/db"
)

type Handler struct {
	dbConnector db.DBConnector
}

type PostAmountRequest struct {
	WalletID      string `json:"walletid"`
	OperationType string `json:"operationType"`
	Amount        string `json:"amount"`
}

type WalelteBalance struct {
	WalletID string `json:"walletid"`
	Balance  string `json:"balance"`
}

func NewHandler(dbc db.DBConnector) (Handler, error) {

	return Handler{dbc}, nil
}
