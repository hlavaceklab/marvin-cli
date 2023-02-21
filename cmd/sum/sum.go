package sum

import (
	"fmt"
	"github.com/martinhlavacek/marvin/cmd/root"
	"github.com/spf13/cobra"
)

var CmdFlagA int
var CmdFlagB int

var Cmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum of 2 numbers",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(CmdFlagA + CmdFlagB)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().IntVarP(
		&CmdFlagA,
		"a",
		"a",
		0,
		"1st number",
	)
	Cmd.MarkPersistentFlagRequired("a")
	Cmd.PersistentFlags().IntVarP(
		&CmdFlagB,
		"b",
		"b",
		0,
		"2nd number")
	Cmd.MarkPersistentFlagRequired("b")
}
