package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"stock_win/cmd"
)

var rootCmd = &cobra.Command{Use: "server"}

func main() {
	rootCmd.AddCommand(cmd.Scrape)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
