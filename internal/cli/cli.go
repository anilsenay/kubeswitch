package cli

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type Cli struct {
	Contexts       []string
	CurrentContext string
	Result         string
}

type item struct {
	Label     string
	IsCurrent bool
}

func (i item) String() string {
	return i.Label
}

func (c *Cli) Run() (string, error) {
	contexts := make([]item, len(c.Contexts))
	pos := 0
	for i, context := range c.Contexts {
		if context == c.CurrentContext {
			pos = i
		}
		contexts[i] = item{
			Label:     context,
			IsCurrent: context == c.CurrentContext,
		}
	}

	searcher := func(input string, index int) bool {
		ctx := contexts[index]
		name := strings.Replace(strings.ToLower(ctx.Label), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label: "Select Context",
		Items: contexts,
		Templates: &promptui.SelectTemplates{
			Active:   `{{if .IsCurrent}} {{ "✔" | green | bold }} {{ . | green | bold }} {{ "(current)" | green | bold }} {{else}}  {{ "✔" | green | bold }} {{ . | cyan | bold }} {{end}}`,
			Inactive: `{{if .IsCurrent}} {{ . | green }} {{ "(current)" | green }} {{else}} {{ . }} {{end}}`,
			Selected: `{{ "✔" | green | bold }} {{ "Selected" | bold }}: {{ . | cyan }}`,
		},
		CursorPos: pos,
		Searcher:  searcher,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %v", err)
	}

	return result, nil
}
