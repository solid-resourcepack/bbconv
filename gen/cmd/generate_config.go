package cmd

import (
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	generateCommand.AddCommand(generateConfigCommand)
}

var generateConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "Generate a config used by the bbconv paper plugin",
	Run: func(cmd *cobra.Command, args []string) {
		InitBBCtx(cmd)
		InitBaseCtx(cmd)
		if err := baseformat.WriteModel(OutDir+"config.json", GetBaseCtx(cmd)); err != nil {
			log.Fatal(err)
		}
	},
}
