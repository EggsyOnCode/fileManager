package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func chDirCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "chdir",
		Short: "changes current dir to the specified dir",
		Run: func(cmd *cobra.Command, args []string) {
			dirName := strings.Join(args, " ")
			err := os.Chdir(dirName )
			if err != nil {
				log.Fatalf("failed switching dir: %s", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Directory %s changed successfully\n", dirName)
		},
	}
}

func init() {
	RootCmd().AddCommand(chDirCmd())
}
