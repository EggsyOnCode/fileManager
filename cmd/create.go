package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "creates a new file in the current directory",
		Run: func(cmd *cobra.Command, args []string) {
			filename := strings.Join(args, " ")
			filePtr, err := os.Create(filename)
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "File %s created successfully\n", filename)
			defer filePtr.Close()
		},
	}
}

func init() {
	RootCmd().AddCommand(createCmd())
}
