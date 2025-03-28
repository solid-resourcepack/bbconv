package cmd

import (
	"github.com/solid-resourcepack/bbconv/mcformat"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
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
		path, err := filepath.Abs(filepath.Join(OutDir, "assets", Namespace))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Writing minecraft assets to directory %s\n", path)
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
