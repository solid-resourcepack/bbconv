package cmd

import (
	"github.com/solid-resourcepack/bbconv/mcformat"
	"github.com/spf13/cobra"
	"log"
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
		pack, err := mcformat.BaseToMc(GetBaseCtx(cmd), Namespace)
		if err != nil {
			log.Fatal(err)
		}
		err = mcformat.WritePackData(pack, OutDir, Namespace)
		if err != nil {
			log.Fatal(err)
		}
	},
}
