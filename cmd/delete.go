package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "deletes an existing file in the current directory",
		Run: func(cmd *cobra.Command, args []string) {
			filename := strings.Join(args, " ")
			err := os.Remove(filename)
			if err != nil {
				log.Fatal("file doesn't exist", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "File %s deleted successfully\n", filename)
		},
	}
}

func init() {
	RootCmd().AddCommand(deleteCmd())
}
