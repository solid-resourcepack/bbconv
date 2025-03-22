package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Generate ",
	Long:  "Generate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
	},
}
