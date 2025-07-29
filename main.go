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
	jsonFile := os.Getenv("JSON_FILE")
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		err := lib.CreateJsonFile()
		if err != nil {
			fmt.Printf("Error creating JSON file: %v\n", err)
			os.Exit(1)
		}

	} else if err != nil {
		fmt.Printf("Error checking JSON file: %v\n", err)
		os.Exit(1)
	}

	cmd.Execute()
}
