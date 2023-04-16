package cmd

import (
	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/connectors"
	"github.com/skowrons/todo-meister/pkg/connectors/github"
	"github.com/skowrons/todo-meister/pkg/connectors/terminal"
	"github.com/skowrons/todo-meister/pkg/exec"
	"github.com/spf13/cobra"
)

type GenerateConfig struct {
	Path      string
	Connector connectors.ConnectorType
}

func NewGenerateCmd() *cobra.Command {
	var cfg GenerateConfig

	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "Get all todos in your project and create issues with the provided connector.",
		Aliases: []string{"gen", "g"},
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.LoadConfig()
			switch cfg.Connector {
			case connectors.GitHubConnector:
				exec.BaseExec(conf, cfg.Path, github.NewGitHubConnector())
			case connectors.CountConnector:
				exec.BaseExec(conf, cfg.Path, terminal.NewCountTerminalConnector())
			default:
				exec.BaseExec(conf, cfg.Path, terminal.NewTableTerminalConnector())
			}
		},
	}

	var con string
	cmd.Flags().StringVarP(&cfg.Path, "path", "p", ".", "The path to the folder that should be scanned.")
	cmd.Flags().StringVarP(&con, "connector", "c", "github", "Define the used connector.")
	cfg.Connector = connectors.ConnectorType(con)

	return cmd
}
