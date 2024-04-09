package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "fileManager",
		Long: "A cli tool that is a wrapper over the native os file management services written in GO",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Welcome to FileManger Cli")
			fmt.Fprintln(cmd.OutOrStdout(), "A cli tool that is a wrapper over the native os file management services written in GO")
		},
	}
	cmd.AddCommand(createCmd())
	cmd.AddCommand(deleteCmd())
	cmd.AddCommand(mkdirCmd())
	cmd.AddCommand(chDirCmd())
	cmd.AddCommand(moveCmd())
	cmd.AddCommand(openCmd())
	cmd.AddCommand(editCmd())
	cmd.AddCommand(readCmd())
	
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func TestingRoot() *cobra.Command {
	return RootCmd()
}
