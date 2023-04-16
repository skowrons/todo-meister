package config

type Config struct {
	Ignores map[string]any
	Fields  []Field
}

type Field struct {
	// Text is the text that will be matched bei todo-meister
	Text string

	// Labels represent the labels concept used by many issue tools.
	// In the end the connectors will decide what to do with the labels.
	// e.x. Jira -> priority:minor could be transalted to the native priority field of a jira issue
	Labels []string
}

// TODO: load from config file
func LoadConfig() *Config {
	return StdConfig()
}

func StdConfig() *Config {
	return &Config{
		Ignores: map[string]any{
			"node_modules": "",
			"exclude":      "",
			"venv":         "",
			".git":         "",
		},
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
