package connectors

import (
	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/meister"
)

type Connector interface {
	Exec(*config.Config, []meister.Entry) error
}

type ConnectorType string

func (ct *ConnectorType) String() string {
	return string(*ct)
}

const (
	TerminalConnector ConnectorType = "terminal"
	CountConnector    ConnectorType = "count"
	GitHubConnector   ConnectorType = "github"
)
