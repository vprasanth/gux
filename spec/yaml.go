package spec

import (
	"gopkg.in/yaml.v2"
)

// GuxYaml Lint this!
type GuxConfig struct {
	Version string
	Session []struct {
		Name   string
		Window []struct {
			Layout     string
			Name       string
			WorkingDir string `yaml:"workingDir"`
			Panes      []struct {
				Command string
			}
		}
	}
}

type Pane struct {
	Command string
}

type Session struct {
	Name   string
	Window []Window
}

type Window struct {
	Layout     string
	Name       string
	WorkingDir string `yaml:"workingDir"`
	Panes      []struct {
		Command string
	}
}

func (g *GuxConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, g)
}
