package main

import (
	"fmt"
	"os"

	"github.com/ronthesoul/tago/cmd"
	"github.com/ronthesoul/tago/pkg"
)

func main() {
	if _, err := os.Stat(pkg.CSVFile); os.IsNotExist(err) {
		err := pkg.CreateCSVFile()
		if err != nil {
			fmt.Printf("Error creating CSV file: %v\n", err)
			os.Exit(1)
		}

	} else if err != nil {
		fmt.Printf("Error checking CSV file: %v\n", err)
		os.Exit(1)
	}

	cmd.Execute()
}
