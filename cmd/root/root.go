package root

import (
	"github.com/martinhlavacek/marvin/version"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var RootCmd = &cobra.Command{
	Use:   "marvin",
	Short: "Marvin CLI tools, " + version.Version,
}

func init() {
	RootCmd.PersistentFlags().BoolVar(
		&RootCmdFlagJson,
		"json",
		false,
		"Output format to json",
	)
}
