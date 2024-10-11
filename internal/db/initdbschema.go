package db

import (
	"context"
	"fmt"

	"github.com/tyagnii/wallets/config"
	"github.com/tyagnii/wallets/gen/ent"
)

func InitDBSchema() error {
	ctx := context.Background()

	client, err := ent.Open("postgres", config.ConnectionString)
	if err != nil {
		return err
	}

	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	return nil
}
