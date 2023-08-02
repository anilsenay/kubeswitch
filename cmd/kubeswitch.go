package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/anilsenay/kubeswitch/internal/cli"
	"github.com/anilsenay/kubeswitch/internal/kubeconfig"
	"github.com/anilsenay/kubeswitch/internal/models"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

var currentFlag = flag.Bool("current", false, "Print current context")
var configPath = flag.String("config", "", "Path to kubeconfig file")

func main() {
	flag.Parse()
	args := os.Args[1:]

	if *configPath != "" {
		os.Setenv("KUBECONFIG", *configPath)
	}

	config := kubeconfig.NewKubeConfig(
		&kubeconfig.DefaultConfigLoader{},
		&kubeconfig.DefaultConfigParser{},
	)

	err := config.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case *currentFlag, len(args) > 0 && args[0] == "current":
		color.Green(config.CurrentContext())
		return
	default:
		err := switchContext(config.Contexts(), config.CurrentContext(), config.SwitchContext)
		if err != nil {
			fmt.Println(err)
			createBackup(config.Config)
			return
		}
	}
}

func switchContext(contexts []string, currentContext string, switchFn func(string) error) error {
	cli := cli.Cli{
		Contexts:       contexts,
		CurrentContext: currentContext,
	}

	result, err := cli.Run()
	if err != nil {
		return err
	}

	err = switchFn(result)
	if err != nil {
		return err
	}

	fmt.Println("Switched to context:", color.GreenString(result))
	return nil
}

func createBackup(cfg models.Config) {
	b, _ := yaml.Marshal(cfg)
	timestamp := strconv.Itoa(int(time.Now().UnixMilli()))
	err := os.WriteFile("/tmp/kube-config-backup-"+timestamp, b, 0644)
	if err == nil {
		fmt.Println("Backup kubeconfig file to /tmp/kube-config-backup-" + timestamp)
	}
}
