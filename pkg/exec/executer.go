package exec

import (
	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/connectors"
	"github.com/skowrons/todo-meister/pkg/meister"
)

// BaseExec is the basic use case of todo-meister.
func BaseExec(cfg *config.Config, root string, con connectors.Connector) {
	index := meister.Walk(cfg, root)

	entries := make([]meister.Entry, 0)
	for _, path := range index {
		entries = append(entries, meister.ScanFile(path, cfg)...)
	}

	con.Exec(entries)
}
