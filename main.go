package main

import (
	"fmt"
	"os"

	"github.com/EggsyOnCode/fileManager/cmd"
)

func main() {
	must(cmd.RootCmd().Execute())
}
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
