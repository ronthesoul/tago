/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"tago/cmd"
	lib "tago/pkg"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	csvFile := os.Getenv("CSV_FILE")
	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
		err := lib.CreateJsonFile()
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
