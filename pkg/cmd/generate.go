package cmd

import (
	"fmt"
	"os"

	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/connectors"
	"github.com/skowrons/todo-meister/pkg/exec"
	"github.com/spf13/cobra"
)

func NewGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "generate",
		Short:   "Get all todos in your project and create issues with the provided connector.",
		Aliases: []string{"gen", "g"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("You need to provide the path as the first and only arg.")
				os.Exit(1)
			}
			root := args[0]
			cfg := config.LoadConfig()
			exec.BaseExec(cfg, root, connectors.NewTerminalConnector())
		},
	}
}
