package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
func moveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "move",
		Short: "moves a file from source dir to a dest directory",
		Run: func(cmd *cobra.Command, args []string) {
			sourceFile := args[0]
			destDir := args[1]
			err := os.Rename(sourceFile, destDir+"/"+sourceFile)
			if err != nil {
				log.Fatalf("failed moving file: %s", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "File %s moved successfully\n", sourceFile)
		},
	}
}

func init() {
	RootCmd().AddCommand(moveCmd())

}

// doesnt meet hte requirements ; the move implementation does
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("Couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("Couldn't copy to dest from source: %v", err)
	}

	inputFile.Close() // for Windows, close before trying to remove: https://stackoverflow.com/a/64943554/246801

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't remove source file: %v", err)
	}
	return nil
}
