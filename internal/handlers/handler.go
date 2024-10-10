package handlers

import (
	"fmt"

	"github.com/tyagnii/wallets/config"
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
	conn, err := ent.Open("postgres", config.ConnectionString)
	if err != nil {
		fmt.Println("ERROR due create handler", err)
	}
	return Handler{dbConnector: db.Connector{Client: conn}}, nil
}
