package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Author struct {
	Name  string
	Email string
}

type Command struct {
	Name        string
	Aliases     []string
	ArgsUsage   string
	Description string
	Options     []Option
	Commands    []*Command
	Action      func(*Context) error

	Copyright string
	Version   string
	NoHelp    bool
}

func (command *Command) Run(args []string, defaultAction func(*Context) error) error {
	if defaultAction == nil {
		defaultAction = ShowHelp(os.Stdout)
	}
	return command.run(nil, args, defaultAction)
}

func (command *Command) run(parent *Context, args []string, defaultAction func(*Context) error) error {
	args = args[1:]

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

	context := &Context{
		parent:  parent,
		command: command,
		options: map[string]interface{}{},
		args:    []string{},
	}

	options := map[string]Option{}
	for _, option := range command.Options {
		option.SetDefaultValue(context.options)

		for _, keyword := range option.Keywords() {
			options[keyword] = option
		}
	}

	i := 0

PARSE_OPTIONS:
	for i < len(args) {
		switch {
		default:
			break PARSE_OPTIONS

		case args[i] == "--":
			// end of option list
			i++
			break PARSE_OPTIONS

		case len(args[i]) >= 3 && args[i][0:2] == "--":
			// long option

			if j := strings.Index(args[i], "="); j >= 0 {
				// --key=value
				key := args[i][:j]
				value := args[i][j+1:]
				i++

				option, ok := options[key]
				if !ok {
					return errors.New("unknown option: " + key)
				}

				if _, err := option.Apply(context.options, value); err != nil {
					return err
				}

			} else {
				// --key value
				key := args[i]
				i++

				option, ok := options[key]
				if !ok {
					return errors.New("unknown option: " + key)
				}

				n, err := option.Apply(context.options, args[i:]...)
				i += n
				if err != nil {
					return err
				}
			}

		case len(args[i]) >= 2 && args[i][0] == '-':
			// short option
			arg := args[i][1:]
			i++

			for j := 0; j < len(arg); j++ {
				key := "-" + string(arg[j])

				option, ok := options[key]
				if !ok {
					return errors.New("unknown option: " + key)
				}

				if j == len(arg)-1 {
					n, err := option.Apply(context.options, args[i:]...)
					if err != nil {
						return err
					}

					i += n

					break
				}

				n, err := option.Apply(context.options, arg[j+1:])
				if err != nil {
					return err
				}

				if n > 0 {
					break
				}
			}
		}
	}

	if command.Version != "" && context.IsSet("version") {
		fmt.Fprintln(os.Stdout, command.Version)
		return nil
	}

	if !command.NoHelp && context.IsSet("help") {
		return context.ShowHelp(os.Stdout)
	}

	context.args = args[i:]

	// sub command
	if len(context.args) > 0 {
		for _, subcommand := range command.Commands {
			if subcommand.Name == context.args[0] {
				return subcommand.run(context, context.args, defaultAction)
			}

			for _, alias := range subcommand.Aliases {
				if alias == context.args[0] {
					return subcommand.run(context, context.args, defaultAction)
				}
			}
		}
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
		if err := context.ShowHelp(out); err != nil {
			return err
		}
		return errors.New("invalid arguments")
	}
}
