package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
func mkdirCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mkdir",
		Short: "creates a new dir in the current directory",
		Run: func(cmd *cobra.Command, args []string) {
			dirName := strings.Join(args, " ")
			err := os.Mkdir(dirName, 0755)
			if err != nil {
				log.Fatalf("failed creating dir: %s", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Directory %s created successfully\n", dirName)
		},
	}
}

func init() {
	RootCmd().AddCommand(mkdirCmd())
}
