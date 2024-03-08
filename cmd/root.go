package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

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
