package cmd

import (
	"fmt"
	"os"

	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/connectors"
	"github.com/skowrons/todo-meister/pkg/exec"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List all todos in your project.",
		Long:    "List is a alias for the command generate --connector cli",
		Aliases: []string{"ls"},
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
