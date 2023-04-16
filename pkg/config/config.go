package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Ignores map[string]any `yaml:"ignores"`
	Fields  []Field        `yaml:"fields"`

	// ManageLabel will the label used by todo meister to track if this issue
	// is managed by todo meister. Defaults to managed-by::tm
	ManageLabel string `yaml:"manageLabel"`
}

type Field struct {
	// Text is the text that will be matched bei todo-meister
	Text string `yaml:"text"`

	// Labels represent the labels concept used by many issue tools.
	// In the end the connectors will decide what to do with the labels.
	// e.x. Jira -> priority:minor could be transalted to the native priority field of a jira issue
	Labels []string `yaml:"labels"`
}

func LoadConfig() *Config {
	cfg, err := readConfigFile()
	if err != nil {
		return StdConfig()
	}

	return cfg
}

// readConfigFile attempts to read the config file from disk.
func readConfigFile() (*Config, error) {
	for _, name := range []string{"tm.yaml", "tm.config.yaml"} {
		if data, err := os.ReadFile(name); err == nil {
			var cfg *Config
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				// TODO: maybe provide a better message to the cli user
				return nil, os.ErrNotExist
			}
			return cfg, nil
		}
	}

	// Config file not found
	return nil, os.ErrNotExist
}

func StdConfig() *Config {
	return &Config{
		Ignores: map[string]any{
			"node_modules": "",
			"exclude":      "",
			"venv":         "",
			".git":         "",
		},
		ManageLabel: "managed-by::tm",
		Fields: []Field{
			{
				Text: "TODO:",
				Labels: []string{
					"type::refactor",
				},
			},
		},
	}
}
