package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tyagnii/wallets/config"
	"github.com/tyagnii/wallets/gen/ent"
)

var dbc Connector

// docker
// docker run --name postgres -e POSTGRES_PASSWORD=password -d postgres
func TestMain(m *testing.M) {
	config.ConnectionString = "host=127.0.0.1 port=5432 user=postgres password=password sslmode=disable"

	err = InitDBSchema()
	if err != nil {
		fmt.Printf("error during init db shema: %s", err.Error())
	}

	entc, err := ent.Open("postgres", config.ConnectionString)
	if err != nil {
		fmt.Printf("error during connecting to db: %s", err.Error())
	}

	dbc = Connector{Client: entc}

	code := m.Run()

	os.Exit(code)

}

// Benchmarks
var id string
var err error

func BenchmarkCreateWallet(b *testing.B) {
	c := context.Background()
	b.Run("Create wallets", func(b *testing.B) {
		id, err = dbc.CreateWallet(c)
		require.NoError(b, err)
	})

}

func BenchmarkGetWalletBalance(b *testing.B) {
	c := context.Background()
	b.Run("Get balance", func(b *testing.B) {
		_, err := dbc.GetWalletBalance(c, id)
		require.NoError(b, err)

	})
}
