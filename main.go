/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/ronthesoul/tago/cmd"
	"github.com/ronthesoul/tago/pkg"

	"github.com/joho/godotenv"
)

func ensureEnvFile() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		content := "CSV_FILE=/home/ron/tasks.csv\n"
		err := os.WriteFile(".env", []byte(content), 0644)
		if err != nil {
			fmt.Println("Failed to create .env:", err)
		} else {
			fmt.Println(".env file created with default CSV_FILE")
		}
	}
}

func init() {
	ensureEnvFile()
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		os.Exit(1)
	}
}

func main() {
	csvFile := os.Getenv("CSV_FILE")
	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
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
