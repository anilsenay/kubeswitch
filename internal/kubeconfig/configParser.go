package kubeconfig

import (
	"fmt"
	"io"

	"github.com/anilsenay/kubeswitch/internal/models"
	"gopkg.in/yaml.v3"
)

type DefaultConfigParser struct{}

func (c *DefaultConfigParser) ParseConfig(f io.ReadWriteCloser) (models.Config, error) {
	var v models.Config
	if err := yaml.NewDecoder(f).Decode(&v); err != nil {
		return models.Config{}, fmt.Errorf("failed to decode: %v", err)
	}

	return v, nil
}
