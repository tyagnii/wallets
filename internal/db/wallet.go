package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tyagnii/wallets/gen/ent"
	"github.com/tyagnii/wallets/gen/ent/wallet"

	_ "github.com/lib/pq"
)

type PGConnector struct {
	*ent.Client
}

// NewDBConnector returns PGconnector
func NewDBConnector(dsn string) (DBConnector, error) {

	conn, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return PGConnector{Client: conn}, nil
}

// CreateWallet returns wallet UUID
func (c PGConnector) CreateWallet(ctx context.Context) (id string, err error) {
	uid := uuid.New()
	id = uid.String()
	w, err := c.Wallet.Create().
		SetUUID(id).Save(ctx)
	if err != nil {
		return "", err
	}

	return w.UUID, nil
}

// GetWalletBalance returns current wallet balance
func (c PGConnector) GetWalletBalance(ctx context.Context, uuid string) (balance int, err error) {
	w, err := c.Wallet.Query().
		Where(wallet.UUIDEQ(uuid)).
		First(ctx)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}

// ChangeWalletBalance adds or subtract some amount from wallet
func (c PGConnector) ChangeWalletBalance(ctx context.Context, uuid string, op string, amount int) (balance int, err error) {
	w, err := c.Wallet.Query().
		Where(wallet.UUIDEQ(uuid)).First(ctx)
	if err != nil {
		return 0, err
	}

	switch op {
	case "DEPOSIT":
		w, err = w.Update().SetBalance(w.Balance + amount).Save(ctx)
		if err != nil {
			return 0, err
		}
	case "WITHDRAW":
		if w.Balance < amount {
			return 0, fmt.Errorf("insufficient funds")
		}
		w, err = w.Update().SetBalance(w.Balance - amount).Save(ctx)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("unsupported wallet operation")
	}

	return w.Balance, nil
}
