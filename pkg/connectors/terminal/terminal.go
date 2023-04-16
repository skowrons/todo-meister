// package terminal will provide ways to print every extracted entry to the console.
package terminal

import (
	"fmt"

	"github.com/rodaine/table"
	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/meister"
)

// This cli connector will print everything as a table to the terminal.
func NewTableTerminalConnector() *terminal {
	return &terminal{}
}

type terminal struct{}

func (t *terminal) Exec(config *config.Config, entries []meister.Entry) error {
	tbl := table.New("Type", "Path", "Index", "Comment")

	for _, entry := range entries {
		tbl.AddRow(entry.KeyWord, entry.FilePath, entry.LineIndex, entry.Comment)
	}

	tbl.Print()

	return nil
}

// This cli connector will print the count of matched entries to the terminal.
func NewCountTerminalConnector() *terminalCount {
	return &terminalCount{}
}

type terminalCount struct{}

func (t *terminalCount) Exec(cfg *config.Config, entries []meister.Entry) error {
	fmt.Printf("%d\n", len(entries))
	return nil
}
