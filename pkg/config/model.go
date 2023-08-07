package config

type RootConfig struct {
	InputPath string     `yaml:"input-path"`
	PipeLine  []PipeLine `yaml:"pipeline"`
}

type PipeLine struct {
	Name       string   `yaml:"name"`
	Type       string   `yaml:"type"`
	FileFormat string   `yaml:"file-format"`
	NoHeader   bool     `yaml:"no-header"`
	Headers    []Header `yaml:"headers"`
	Template   string   `yaml:"template"`
}

type Header struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	DataFormat string `yaml:"data-format"`
}
