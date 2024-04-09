package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func editCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "edits a file",
		Long:  "first arg should be the fileName, second should be the mode in which file needs to be edited; if mode is append then type 0 and then 3rd arg should be the text u want to add; if the mode is write_at then use 1 and then the point in the file to write at and then 4th arg should be the text ",
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			mode := args[1]
			modeBit, error := strconv.Atoi(mode)
			if error != nil {
				fmt.Println("parse error during str to int conversion")
			}
			if modeBit == 0 {
				filePtr, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				text := args[2]
				if _, err := filePtr.Write([]byte(text)); err != nil {
					log.Fatal(err)
				}
				defer filePtr.Close()
				fmt.Fprintf(cmd.OutOrStdout(), "File %s opened and text appended successfully\n", filename)
			} else {
				filePtr, err := os.OpenFile(filename, os.O_RDWR, 0644)
				if err != nil {
					log.Fatal(err)
				}
				point, error := strconv.Atoi(args[2])
				if error != nil {
					fmt.Println("parse error during str to int conversion")
				}
				text := args[3]
				textByes := []byte(text)
				filePtr.WriteAt(textByes, int64(point))
				defer filePtr.Close()
				fmt.Fprintf(cmd.OutOrStdout(), "File %s opened and text written at point successfully\n", filename)
			}

		},
	}
}

func init() {
	RootCmd().AddCommand(editCmd())

}
