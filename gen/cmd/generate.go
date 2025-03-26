package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCommand)
	generateCommand.PersistentFlags().StringVarP(&BBFile, "source", "s", "./model.bbmodel", "BlockBench model file")
	generateCommand.PersistentFlags().StringVarP(&OutDir, "output", "o", "", "Output directory")
	generateCommand.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "bbconv", "Minecraft ResourcePack namespace")
}

var generateCommand = &cobra.Command{
	Use:   "gen",
	Short: "Generate Minecraft Java ResourcePack definitions by a bbmodel",
	Long:  "Generate a config and Minecraft Java ResourcePack definitions by a bbmodel",
	Run: func(cmd *cobra.Command, args []string) {
		InitBBCtx(cmd)
		InitBaseCtx(cmd)
		for _, subCmd := range cmd.Commands() {
			if !subCmd.Hidden {
				subCmd.SetContext(cmd.Context())
				subCmd.Run(subCmd, args)
			}
		}
	},
}
