package cmd

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "generate-id-use-autoincrement",
	Short: "This command is a sample for generating IDs using auto-increment.",

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		db, err = sql.Open("sqlite3", "./storage/db")
		if err != nil {
			return err
		}
		if err := db.Ping(); err != nil {
			return err
		}

		ctx := context.WithValue(cmd.Context(), "db", db)
		cmd.SetContext(ctx)

		return nil
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		db.Close()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(CreateUserCmd())
	rootCmd.AddCommand(FindUserCmd())
}
