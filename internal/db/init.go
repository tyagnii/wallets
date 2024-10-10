package db

import (
	"context"
	"fmt"

	"github.com/tyagnii/wallets/config"
	"github.com/tyagnii/wallets/gen/ent"

	_ "github.com/lib/pq"
)

var ConnectionString string

func init() {
	// client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	ConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.Env.DBHostName,
		config.Env.DBPort,
		config.Env.DBUser,
		config.Env.DBHostName,
		config.Env.DBPassword)

	client, err := ent.Open("postgres", ConnectionString)
	if err != nil {
		fmt.Println("ERROR due init DB. Couldn't create db schema", err)
	} else {
		client.Schema.Create(context.Background())
	}
}
