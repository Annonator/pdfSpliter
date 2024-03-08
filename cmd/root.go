package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringP("author", "a", "", "author of the application")

	err := viper.BindPFlag("author", rootCommand.PersistentFlags().Lookup("author"))
	if err != nil {
		return
	}

	viper.SetDefault("author", "Andreas Pohl")
}

func initConfig() {
	viper.SetConfigFile(".pdf.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

var rootCommand = &cobra.Command{
	Use:   "pdf",
	Short: "pdf is a tool for working with PDF files",
	Long:  `pdf is a tool for working with PDF files. It provides several subcommands for working with PDF files, such as splitting, merging, and encrypting.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
