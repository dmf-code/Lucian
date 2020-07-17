package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {

	var cmd = &cobra.Command{
		Short: "welcome use cobra command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	cmd.Execute()
}
