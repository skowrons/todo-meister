package github

import (
	"os"

	"github.com/skowrons/todo-meister/pkg/config"
	"github.com/skowrons/todo-meister/pkg/meister"
)

const (
	EnvAccessKey = "ACCESS_KEY"
	EnvOwner     = "OWNER"
	EnvRepo      = "REPO"
)

func NewGitHubConnector() *github {
	return &github{
		AccessKey: os.Getenv(EnvAccessKey),
		Owner:     os.Getenv(EnvOwner),
		Repo:      os.Getenv(EnvRepo),
	}
}

type github struct {
	AccessKey string
	Owner     string
	Repo      string
}

func (g *github) Exec(config *config.Config, entries []meister.Entry) error {
	getIssuesByLabel(g, config.ManageLabel)
	return nil
}
