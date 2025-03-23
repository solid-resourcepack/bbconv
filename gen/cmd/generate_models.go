package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	generateCommand.AddCommand(generateModelsCommand)
}

var generateModelsCommand = &cobra.Command{
	Use:   "models",
	Short: "Generate Minecraft Java ResourcePack definitions by a bbmodel",
	Run: func(cmd *cobra.Command, args []string) {
		InitBBCtx(cmd)
		InitBaseCtx(cmd)
	},
}
