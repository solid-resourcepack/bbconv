package cmd

import (
	"context"
	"github.com/solid-resourcepack/bbconv/reader"
	"github.com/solid-resourcepack/bbconv/types"
	"github.com/spf13/cobra"
)

func InitBBCtx(cmd *cobra.Command) {
	if cmd.Context() != nil && cmd.Context().Value("bb-data") != nil {
		return
	}
	ctx := context.WithValue(context.Background(), "bb-data", reader.ReadBBModel(BBFile))
	cmd.SetContext(ctx)
}

func GetBBCtx(cmd *cobra.Command) *types.BlockBenchModel {
	if cmd.Context() == nil && cmd.Context().Value("bb-data") == nil {
		return nil
	}
	return cmd.Context().Value("bb-data").(*types.BlockBenchModel)
}
