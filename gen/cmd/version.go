package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the current bbconv version",
	Long:  "Print the current bbconv version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("bbconv v1.0.0")
	},
}
