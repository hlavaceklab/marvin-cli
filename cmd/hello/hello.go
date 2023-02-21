package hello

import (
	"fmt"
	"github.com/martinhlavacek/marvin/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "hello",
	Short:   "Say: Hello",
	Aliases: []string{"h"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
