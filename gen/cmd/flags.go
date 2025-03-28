package cmd

import "github.com/spf13/cobra"

var BBFile string
var OutDir string
var Namespace string

func InitFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&BBFile, "source", "s", "./model.bbmodel", "BlockBench model file")
	cmd.PersistentFlags().StringVarP(&OutDir, "output", "o", "./", "Output directory")
	cmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "bbconv", "Minecraft ResourcePack namespace")

}
