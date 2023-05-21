package main

import (
	"fmt"
	"generate-id-use-autoincrement/cmd/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
