package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"generate-id-use-autoincrement/internal/application/usecase"
	"generate-id-use-autoincrement/internal/infrastructure"
	"github.com/spf13/cobra"
	"strconv"
)

var findUserCmd = &cobra.Command{
	Use: "find-user",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return errors.New(fmt.Sprintf("fail to find-user command: %v", "id is empty"))
		}

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("fail to find-user command: %v", err))
		}

		ctx := cmd.Context()
		db := cmd.Context().Value("db").(*sql.DB)
		r := infrastructure.NewUserRepositoryInSqlite(db)
		u := usecase.NewFindUserUsecase(r)

		foundUser, err := u.Execute(ctx, id)
		if err != nil {
			return errors.New(fmt.Sprintf("fail to find-user command: %v", err))
		}

		fmt.Printf("found user: id=%v, name=%v \n", foundUser.Id().Value(), foundUser.Name())

		return nil
	},
}

func FindUserCmd() *cobra.Command {
	return findUserCmd
}
