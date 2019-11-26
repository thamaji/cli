package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Command struct {
	Name        string
	Aliases     []string
	ArgsUsage   string
	Description string
	Options     []Option
	Commands    []*Command
	Action      func(*Context) error

	Version string
	NoHelp  bool
}

func (command *Command) Run(args []string, defaultAction func(*Context) error) error {
	return command.run(nil, args, defaultAction)
}

func (command *Command) run(parent *Context, args []string, defaultAction func(*Context) error) error {
	args = args[1:]

	keywords := map[string]*Node{}

	if command.Version != "" {
		command.Options = append(command.Options, &BoolOption{
			Name:        "version",
			Short:       "v",
			Description: "show version",
		})
	}

	if !command.NoHelp {
		command.Options = append(command.Options, &BoolOption{
			Name:        "help",
			Short:       "h",
			Description: "show help",
		})
	}

	for _, option := range command.Options {
		for _, keyword := range option.Keywords() {
			keywords[keyword] = &Node{Type: T_Option, Value: option}
		}
	}

	for _, command := range command.Commands {
		keywords[command.Name] = &Node{Type: T_Command, Value: command}
		for _, alias := range command.Aliases {
			keywords[command.Name] = &Node{Type: T_Command, Value: alias}
		}
	}

	context := &Context{
		parent:  parent,
		command: command,
		options: map[string]interface{}{},
		args:    []string{},
	}

	for _, option := range command.Options {
		option.SetDefaultValue(context.options)
	}

	nodes, next := parse(keywords, args)
	var subcommand *Command

	i := 0
	for i < len(nodes) {
		switch nodes[i].Type {
		case T_Command:
			subcommand = nodes[i].Value.(*Command)
			i++

			if len(context.args) > 0 {
				// sub command の前に arg があったらとりあえずエラーにしておく
				// そして、sub command は常に最後の node であるはず
				return errors.New("invalid arguments: " + strings.Join(context.args, " "))
			}

		case T_Value:
			arg := nodes[i].Value.(string)
			i++

			context.args = append(context.args, arg)

		case T_Option:
			Option := nodes[i].Value.(Option)
			i++

			n, err := Option.Parse(nodes[i:], context.options)
			if err != nil {
				return err
			}

			i += n
		}
	}

	if command.Version != "" && context.IsSet("version") {
		fmt.Fprintln(os.Stdout, command.Version)
		return nil
	}

	if !command.NoHelp && context.IsSet("help") {
		context.ShowHelp(os.Stdout)
		return nil
	}

	if subcommand != nil {
		return subcommand.run(context, next, defaultAction)
	}

	if command.Action == nil {
		if defaultAction != nil {
			return defaultAction(context)
		}

		return nil
	}

	if err := command.Action(context); err != nil {
		return err
	}

	return nil
}

func ShowHelp(out io.Writer) func(*Context) error {
	return func(context *Context) error {
		context.ShowHelp(out)
		return nil
	}
}
