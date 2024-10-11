/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tyagnii/wallets/config"
	"github.com/tyagnii/wallets/internal/db"
	"github.com/tyagnii/wallets/internal/handlers"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		// read params
		err := config.ReadConfig("config.env")
		if err != nil {
			os.Exit(1)
		}

		// init db schema
		err = db.InitDBSchema()
		if err != nil {
			err = fmt.Errorf("serve error: %w", err)
			fmt.Println(err)
			os.Exit(1)
		}

		// create pg connector
		// keep conn string for testing purposes
		conn, err := db.NewDBConnector(config.ConnectionString)
		if err != nil {
			err = fmt.Errorf("serve error: %w", err)
			fmt.Println(err)
			os.Exit(1)
		}

		// create handler instance
		h, err := handlers.NewHandler(conn)
		if err != nil {
			err = fmt.Errorf("serve error: %w", err)
			fmt.Println(err)
			os.Exit(1)
		}

		// run router with handler created above
		r := handlers.NewRouter(h)
		r.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
