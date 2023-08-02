package kubeconfig

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type DefaultConfigLoader struct {
	file *os.File
}

func (c *DefaultConfigLoader) Load() (io.ReadWriteCloser, error) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	cfgPath := os.Getenv("KUBECONFIG")
	if cfgPath != "" {
		if cfgPath == "~" {
			cfgPath = dir
		} else if strings.HasPrefix(cfgPath, "~/") {
			cfgPath = filepath.Join(dir, cfgPath[2:])
		}
	} else {
		defaultPath, err := c.defaultConfigPath()
		if err != nil {
			return nil, err
		}
		cfgPath = defaultPath
	}

	f, err := os.OpenFile(cfgPath, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("kubeconfig file not found: %v", err)
		}
		return nil, fmt.Errorf("failed to open file")
	}

	c.file = f
	return f, nil
}

func (c *DefaultConfigLoader) Reset() error {
	if err := c.file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %v", err)
	}

	_, err := c.file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed to seek in file: %v", err)
	}

	return nil
}

func (c *DefaultConfigLoader) defaultConfigPath() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("HOME environment variable not set")
	}
	return filepath.Join(home, ".kube", "config"), nil
}
