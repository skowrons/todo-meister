package connectors

import (
	"github.com/rodaine/table"
	"github.com/skowrons/todo-meister/pkg/meister"
)

// This cli connector will print everything as a table to the terminal.
func NewTerminalConnector() *terminal {
	return &terminal{}
}

type terminal struct{}

func (t *terminal) Exec(entries []meister.Entry) error {
	tbl := table.New("Type", "Path", "Index", "Comment")

	for _, entry := range entries {
		tbl.AddRow(entry.KeyWord, entry.FilePath, entry.LineIndex, entry.Comment)
	}
	
	tbl.Print()

	return nil
}
