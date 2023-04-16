package cmd

import (
	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/connectors/terminal"
	"github.com/skowrons/todo-meister/pkg/exec"
	"github.com/spf13/cobra"
)

type ListConfig struct {
	Path  string
	Count bool
}

func NewListCmd() *cobra.Command {
	var cfg ListConfig

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all todos in your project.",
		Long:    "List is a alias for the command generate --connector cli",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.LoadConfig()

			if cfg.Count {
				exec.BaseExec(conf, cfg.Path, terminal.NewCountTerminalConnector())
				return
			}

			exec.BaseExec(conf, cfg.Path, terminal.NewTableTerminalConnector())
		},
	}

	cmd.Flags().StringVar(&cfg.Path, "path", ".", "The path to the folder that should be scanned.")
	cmd.Flags().BoolVarP(&cfg.Count, "count", "c", false, "Define if you only what the count of all matched keywords.")

	return cmd
}
