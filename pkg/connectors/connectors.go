package connectors

import "github.com/skowrons/todo-meister/pkg/meister"

type Connector interface {
	Exec([]meister.Entry) error
}
