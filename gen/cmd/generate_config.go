package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	generateCommand.AddCommand(generateConfigCommand)
}

var generateConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "Generate a config used by the bbconv paper plugin",
	Run: func(cmd *cobra.Command, args []string) {
		InitBBCtx(cmd)
		fmt.Println("config called")
	},
}
