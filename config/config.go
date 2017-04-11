package config

import (
	"github.com/mgreau/picon/types"

	"gopkg.in/yaml.v2"
)

// Parse the picon.yml file
func Parse(data []byte) (*types.Picon, error) {
	c := new(types.Picon)
	err := yaml.Unmarshal(data, c)
	return c, err
}
