package gitlab

import "github.com/skowrons/todo-meister/pkg/meister"

func NewGitLabConnector() *gitlab {
	return &gitlab{
		APIURL:    "",
		AccessKey: "",
		Project:   "",
	}
}

type gitlab struct {
	APIURL    string
	AccessKey string
	Project   string
}

func (g *gitlab) Exec(entries []meister.Entry) error {
	return nil
}
