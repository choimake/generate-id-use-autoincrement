package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"generate-id-use-autoincrement/internal/application/usecase"
	"generate-id-use-autoincrement/internal/infrastructure"
	"github.com/spf13/cobra"
)

var createUserCmd = &cobra.Command{
	Use: "create-user",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return errors.New(fmt.Sprintf("fail to create-user command :%v", "name is empty"))
		}

		name := args[0]

		ctx := cmd.Context()
		db := cmd.Context().Value("db").(*sql.DB)
		r := infrastructure.NewUserRepositoryInSqlite(db)
		u := usecase.NewCreateUserUsecase(r)

		if err := u.Execute(ctx, name); err != nil {
			return errors.New(fmt.Sprintf("fail to create-user command :%v", err))
		}

		return nil
	},
}

func CreateUserCmd() *cobra.Command {
	return createUserCmd
}
