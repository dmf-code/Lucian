package main

import (
	"app/cmd"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Short: "welcome use cobra command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	rootCmd.AddCommand(cmd.MigrationCmd)
	rootCmd.Execute()

}
