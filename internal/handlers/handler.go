package handlers

import (
	"github.com/tyagnii/wallets/internal/db"
)

type Handler struct {
	dbConnector db.Connector
}

func NewHandler() (Handler, error) {
	return Handler{}, nil
}
