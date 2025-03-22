package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bbconv",
	Short: "bbconv is a BlockBench model converter.",
	Long: `A Fast BlockBench to Minecraft Java ResourcePack converter written in go. 
Complete documentation is available at https://github.com/solid-resourcepack/bbconv`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
