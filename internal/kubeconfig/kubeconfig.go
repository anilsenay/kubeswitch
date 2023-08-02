package kubeconfig

import (
	"fmt"
	"io"

	"github.com/anilsenay/kubeswitch/internal/models"
	"gopkg.in/yaml.v3"
)

type Parser interface {
	ParseConfig(f io.ReadWriteCloser) (models.Config, error)
}

type Loader interface {
	Load() (io.ReadWriteCloser, error)
	Reset() error
}

type kubeConfig struct {
	f      io.ReadWriteCloser
	loader Loader
	parser Parser
	Config models.Config
}

func NewKubeConfig(l Loader, p Parser) *kubeConfig {
	return &kubeConfig{
		loader: l,
		parser: p,
	}
}

func (c *kubeConfig) CurrentContext() string {
	return c.Config.CurrentContext
}

func (c *kubeConfig) Contexts() []string {
	var contexts []string
	for _, ctx := range c.Config.Contexts {
		contexts = append(contexts, ctx.Name)
	}
	return contexts
}

func (c *kubeConfig) Parse() error {
	f, err := c.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load: %v", err)
	}
	c.f = f

	cfg, err := c.parser.ParseConfig(f)
	if err != nil {
		return fmt.Errorf("failed to parse: %v", err)
	}

	c.Config = cfg
	return nil
}

func (c *kubeConfig) SwitchContext(ctx string) error {
	c.Config.CurrentContext = ctx

	err := c.loader.Reset()
	if err != nil {
		return err
	}

	err = yaml.NewEncoder(c.f).Encode(c.Config)
	if err != nil {
		return err
	}

	return nil
}
