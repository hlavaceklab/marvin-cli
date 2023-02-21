package cmd

import (
	_ "github.com/martinhlavacek/marvin/cmd/hello"
	"github.com/martinhlavacek/marvin/cmd/root"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
