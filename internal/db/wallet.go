package db

import (
	"context"

	"github.com/tyagnii/wallets/gen/ent"
	"github.com/tyagnii/wallets/gen/ent/wallet"
)

type Connector struct {
	*ent.Client
}

// CreateWallet returns wallet UUID
func (c *Connector) CreateWallet(ctx context.Context) (uuid string, err error) {
	w, err := c.Wallet.Create().
		SetUUID(uuid).Save(ctx)
	if err != nil {
		return "", err
	}

	return w.UUID, nil
}

// GetWalletBalance returns current wallet balance
func (c *Connector) GetWalletBalance(ctx context.Context, uuid string) (balance int, err error) {
	w, err := c.Wallet.Query().
		Where(wallet.UUIDEQ(uuid)).
		First(ctx)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}

// ChangeWalletBalance adds or subtract some amount from wallet
func (c *Connector) ChangeWalletBalance(ctx context.Context, uuid string, amount int) (balance int, err error) {
	w, err := c.Wallet.Query().
		Where(wallet.UUIDEQ(uuid)).First(ctx)
	if err != nil {
		return 0, err
	}

	w, err = w.Update().SetBalance(w.Balance + amount).Save(ctx)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}
