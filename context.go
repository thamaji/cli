package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/thamaji/tablewriter"
	"golang.org/x/crypto/ssh/terminal"
)

type Context struct {
	parent *Context

	command *Command
	options map[string]interface{}
	args    []string
}

func (context Context) Args() []string {
	return context.args
}

func (context Context) IsSet(name string) bool {
	if _, ok := context.options[name]; ok {
		return true
	}

	if context.parent == nil {
		return false
	}

	return context.parent.IsSet(name)
}

func (context Context) Bool(name string) bool {
	v, ok := context.options[name]
	if ok {
		return v.(bool)
	}

	if context.parent == nil {
		return false
	}

	return context.parent.Bool(name)
}

func (context Context) BoolOr(name string, value bool) bool {
	v, ok := context.options[name]
	if ok {
		return v.(bool)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.BoolOr(name, value)
}

func (context Context) BoolOrInput(name string) (bool, error) {
	v, ok := context.options[name]
	if ok {
		return v.(bool), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return false, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return false, err
		}

		v, err := strconv.ParseBool(ans)
		if err != nil {
			return false, err
		}

		return v, nil
	}

	return context.parent.BoolOrInput(name)
}

func (context Context) BoolOrPassword(name string) (bool, error) {
	v, ok := context.options[name]
	if ok {
		return v.(bool), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return false, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return false, err
		}

		v, err := strconv.ParseBool(string(ans))
		if err != nil {
			return false, err
		}

		return v, nil
	}

	return context.parent.BoolOrInput(name)
}

func (context Context) String(name string) string {
	v, ok := context.options[name]
	if ok {
		return v.(string)
	}

	if context.parent == nil {
		return ""
	}

	return context.parent.String(name)
}

func (context Context) StringOr(name string, value string) string {
	v, ok := context.options[name]
	if ok {
		return v.(string)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.StringOr(name, value)
}

func (context Context) StringOrInput(name string) (string, error) {
	v, ok := context.options[name]
	if ok {
		return v.(string), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return "", errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return "", err
		}

		return ans, nil
	}

	return context.parent.StringOrInput(name)
}

func (context Context) StringOrPassword(name string) (string, error) {
	v, ok := context.options[name]
	if ok {
		return v.(string), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return "", errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return "", err
		}

		return string(ans), nil
	}

	return context.parent.StringOrPassword(name)
}

func (context Context) Int(name string) int {
	v, ok := context.options[name]
	if ok {
		return v.(int)
	}

	if context.parent == nil {
		return 0
	}

	return context.parent.Int(name)
}

func (context Context) IntOr(name string, value int) int {
	v, ok := context.options[name]
	if ok {
		return v.(int)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.IntOr(name, value)
}

func (context Context) IntOrInput(name string) (int, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseInt(ans, 10, 64)
		if err != nil {
			return 0, err
		}

		return int(v), nil
	}

	return context.parent.IntOrInput(name)
}

func (context Context) IntOrPassword(name string) (int, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseInt(string(ans), 10, 64)
		if err != nil {
			return 0, err
		}

		return int(v), nil
	}

	return context.parent.IntOrPassword(name)
}

func (context Context) Int32(name string) int32 {
	v, ok := context.options[name]
	if ok {
		return v.(int32)
	}

	if context.parent == nil {
		return 0
	}

	return context.parent.Int32(name)
}

func (context Context) Int32Or(name string, value int32) int32 {
	v, ok := context.options[name]
	if ok {
		return v.(int32)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.Int32Or(name, value)
}

func (context Context) Int32OrInput(name string) (int32, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int32), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseInt(ans, 10, 32)
		if err != nil {
			return 0, err
		}

		return int32(v), nil
	}

	return context.parent.Int32OrInput(name)
}

func (context Context) Int32OrPassword(name string) (int32, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int32), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseInt(string(ans), 10, 32)
		if err != nil {
			return 0, err
		}

		return int32(v), nil
	}

	return context.parent.Int32OrPassword(name)
}

func (context Context) Int64(name string) int64 {
	v, ok := context.options[name]
	if ok {
		return v.(int64)
	}

	if context.parent == nil {
		return 0
	}

	return context.parent.Int64(name)
}

func (context Context) Int64Or(name string, value int64) int64 {
	v, ok := context.options[name]
	if ok {
		return v.(int64)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.Int64Or(name, value)
}

func (context Context) Int64OrInput(name string) (int64, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int64), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseInt(ans, 10, 64)
		if err != nil {
			return 0, err
		}

		return int64(v), nil
	}

	return context.parent.Int64OrInput(name)
}

func (context Context) Int64OrPassword(name string) (int64, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int64), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseInt(string(ans), 10, 64)
		if err != nil {
			return 0, err
		}

		return int64(v), nil
	}

	return context.parent.Int64OrPassword(name)
}

func (context Context) Float32(name string) float32 {
	v, ok := context.options[name]
	if ok {
		return v.(float32)
	}

	if context.parent == nil {
		return 0
	}

	return context.parent.Float32(name)
}

func (context Context) Float32Or(name string, value float32) float32 {
	v, ok := context.options[name]
	if ok {
		return v.(float32)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.Float32Or(name, value)
}

func (context Context) Float32OrInput(name string) (float32, error) {
	v, ok := context.options[name]
	if ok {
		return v.(float32), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseFloat(ans, 32)
		if err != nil {
			return 0, err
		}

		return float32(v), nil
	}

	return context.parent.Float32OrInput(name)
}

func (context Context) Float32OrPassword(name string) (float32, error) {
	v, ok := context.options[name]
	if ok {
		return v.(float32), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseFloat(string(ans), 32)
		if err != nil {
			return 0, err
		}

		return float32(v), nil
	}

	return context.parent.Float32OrPassword(name)
}

func (context Context) Float64(name string) float64 {
	v, ok := context.options[name]
	if ok {
		return v.(float64)
	}

	if context.parent == nil {
		return 0
	}

	return context.parent.Float64(name)
}

func (context Context) Float64Or(name string, value float64) float64 {
	v, ok := context.options[name]
	if ok {
		return v.(float64)
	}

	if context.parent == nil {
		return value
	}

	return context.parent.Float64Or(name, value)
}

func (context Context) Float64OrInput(name string) (float64, error) {
	v, ok := context.options[name]
	if ok {
		return v.(float64), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := readline(os.Stdin)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseFloat(ans, 64)
		if err != nil {
			return 0, err
		}

		return float64(v), nil
	}

	return context.parent.Float64OrInput(name)
}

func (context Context) Float64OrPassword(name string) (float64, error) {
	v, ok := context.options[name]
	if ok {
		return v.(float64), nil
	}

	if context.parent == nil {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			return 0, errors.New("stdin is not a terminal")
		}

		fmt.Fprint(os.Stdout, name+": ")
		ans, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(os.Stdout)
		if err != nil {
			return 0, err
		}

		v, err := strconv.ParseFloat(string(ans), 64)
		if err != nil {
			return 0, err
		}

		return float64(v), nil
	}

	return context.parent.Float64OrPassword(name)
}

func (context *Context) name() string {
	if context.parent == nil {
		return context.command.Name
	}
	return context.parent.name() + " " + context.command.Name
}

func (context *Context) ShowHelp(out io.Writer) error {
	fmt.Fprintln(out, "NAME:")
	name := context.name()

	if context.command.Description != "" {
		name += " - " + context.command.Description
	}

	fmt.Fprintln(out, " ", name)

	fmt.Fprintln(out)
	fmt.Fprintln(out, "USAGE:")
	usage := context.name()

	if len(context.command.Options) > 0 {
		usage += " [OPTIONS]"
	}

	if len(context.command.Commands) > 0 {
		usage += " COMMAND"
	} else {
		if context.command.ArgsUsage != "" {
			usage += " " + context.command.ArgsUsage
		}
	}

	fmt.Fprintln(out, " ", usage)

	if len(context.command.Commands) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "COMMANDS:")
		tw := tablewriter.New(out)
		for _, command := range context.command.Commands {
			tw.Add(" ", command.Name, command.Description)
		}
		tw.Flush()
	}

	if len(context.command.Options) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "OPTIONS:")
		tw := tablewriter.New(out)
		for _, Option := range context.command.Options {
			help := Option.Help()
			tw.Add(" ", help[0], help[1])
		}
		tw.Flush()
	}

	return nil
}

func readline(r io.Reader) (string, error) {
	var bytes [1]byte
	var buf []byte

	for {
		size, err := r.Read(bytes[:])
		if size > 0 {
			switch bytes[0] {
			case '\n':
				return string(buf), nil
			case '\r':
				// remove \r from passwords on Windows
			default:
				buf = append(buf, bytes[0])
			}
			continue
		}
		if err != nil {
			if err == io.EOF && len(buf) > 0 {
				return string(buf), nil
			}
			return string(buf), err
		}
	}
}
