package cmd

import (
	"context"
	"github.com/solid-resourcepack/bbconv/baseformat"
	"github.com/solid-resourcepack/bbconv/bbformat"
	"github.com/spf13/cobra"
)

func InitBaseCtx(cmd *cobra.Command) {
	if GetBBCtx(cmd) == nil {
		return
	}
	if cmd.Context() != nil && cmd.Context().Value("base-data") != nil {
		return
	}
	bbModel := cmd.Context().Value("bb-data").(*bbformat.Model)
	baseModel := baseformat.BBToBase(bbModel)
	ctx := context.WithValue(cmd.Context(), "base-data", baseModel)
	cmd.SetContext(ctx)
}

func GetBaseCtx(cmd *cobra.Command) *baseformat.Model {
	if cmd.Context() == nil || cmd.Context().Value("base-data") == nil {
		return nil
	}
	return cmd.Context().Value("base-data").(*baseformat.Model)
}
