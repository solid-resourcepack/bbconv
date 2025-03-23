package cmd

import (
	"context"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/spf13/cobra"
)

func InitBBCtx(cmd *cobra.Command) {
	if cmd.Context() != nil && cmd.Context().Value("bb-data") != nil {
		return
	}
	ctx := context.WithValue(context.Background(), "bb-data", bbformat.ReadBBModel(BBFile))
	cmd.SetContext(ctx)
}

func GetBBCtx(cmd *cobra.Command) *bbformat.Model {
	if cmd.Context() == nil && cmd.Context().Value("bb-data") == nil {
		return nil
	}
	return cmd.Context().Value("bb-data").(*bbformat.Model)
}
