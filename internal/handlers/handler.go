package handlers

import (
	"fmt"

	"github.com/tyagnii/wallets/gen/ent"
	"github.com/tyagnii/wallets/internal/db"
)

type Handler struct {
	dbConnector db.Connector
}

type PostAmountRequest struct {
	WalletID      string `json:"walletid"`
	OperationType string `json:"operationType"`
	Amount        string `json:"amount"`
}

func NewHandler() (Handler, error) {
	conn, err := ent.Open("postgres", db.ConnectionString)
	if err != nil {
		fmt.Println("ERROR due create handler", err)
	}
	return Handler{dbConnector: db.Connector{Client: conn}}, nil
}
