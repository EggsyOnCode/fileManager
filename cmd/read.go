package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func readCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "read",
		Short: "reads a file and returns the content",
		Long:  "if sequential read then only specify the filename as the 1st arg; otherwise if reading is to be done between 2 points then provide them as 2nd and 3rd args",
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			if len(args) == 1 {
				content, err := os.ReadFile(filename)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "File %s content read successfully: \n %+v\n", filename, string(content))
				return
			}
			start, error := strconv.Atoi(args[1])
			if error != nil {
				fmt.Println("parse error during str to int conversion")
			}
			end, error := strconv.Atoi(args[2])
			if error != nil {
				fmt.Println("parse error during str to int conversion")
			}
			filePtr, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			numBytes := end - start

			// Create a buffer to hold the read data
			buffer := make([]byte, numBytes)

			bytesRead, err := filePtr.ReadAt(buffer, int64(start))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Read %d bytes from position %d: %s\n", bytesRead, start, buffer)
		},
	}
}

func init() {
	RootCmd().AddCommand(readCmd())

}
