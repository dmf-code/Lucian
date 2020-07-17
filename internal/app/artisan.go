package main

import (
	"app/cmd"
	"app/utils/mysqlTools"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init()  {
	// 加载.env配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化Mysql连接池
	if !mysqlTools.GetInstance().InitDataPool() {
		log.Println("init database mysqlTools failure...")
		os.Exit(1)
	}
}

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
