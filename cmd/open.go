package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func openCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "open",
		Short: "opens a file in the application and then closes it",
		Run: func(cmd *cobra.Command, args []string) {
			filename := strings.Join(args, " ")
			filePtr, err := os.OpenFile(filename, os.O_RDWR, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer filePtr.Close()
			fmt.Fprintf(cmd.OutOrStdout(), "File %s opened in application successfully\n", filename)
		},
	}
}

func init() {
	RootCmd().AddCommand(openCmd())

}
