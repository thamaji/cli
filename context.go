package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/thamaji/tablewriter"
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

func (context Context) UserConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, context.Name()), nil
}

func (context Context) UserCacheDir() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, context.Name()), nil
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
		return ReadInputBool(name)
	}

	return context.parent.BoolOrInput(name)
}

func (context Context) BoolOrPassword(name string) (bool, error) {
	v, ok := context.options[name]
	if ok {
		return v.(bool), nil
	}

	if context.parent == nil {
		return ReadPasswordBool(name)
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
		return ReadInputString(name)
	}

	return context.parent.StringOrInput(name)
}

func (context Context) StringOrPassword(name string) (string, error) {
	v, ok := context.options[name]
	if ok {
		return v.(string), nil
	}

	if context.parent == nil {
		return ReadPasswordString(name)
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
		return ReadInputInt(name)
	}

	return context.parent.IntOrInput(name)
}

func (context Context) IntOrPassword(name string) (int, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int), nil
	}

	if context.parent == nil {
		return ReadPasswordInt(name)
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
		return ReadInputInt32(name)
	}

	return context.parent.Int32OrInput(name)
}

func (context Context) Int32OrPassword(name string) (int32, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int32), nil
	}

	if context.parent == nil {
		return ReadPasswordInt32(name)
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
		return ReadInputInt64(name)
	}

	return context.parent.Int64OrInput(name)
}

func (context Context) Int64OrPassword(name string) (int64, error) {
	v, ok := context.options[name]
	if ok {
		return v.(int64), nil
	}

	if context.parent == nil {
		return ReadPasswordInt64(name)
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
		return ReadInputFloat32(name)
	}

	return context.parent.Float32OrInput(name)
}

func (context Context) Float32OrPassword(name string) (float32, error) {
	v, ok := context.options[name]
	if ok {
		return v.(float32), nil
	}

	if context.parent == nil {
		return ReadPasswordFloat32(name)
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
		return ReadInputFloat64(name)
	}

	return context.parent.Float64OrInput(name)
}

func (context Context) Float64OrPassword(name string) (float64, error) {
	v, ok := context.options[name]
	if ok {
		return v.(float64), nil
	}

	if context.parent == nil {
		return ReadPasswordFloat64(name)
	}

	return context.parent.Float64OrPassword(name)
}

func (context *Context) Name() string {
	if context.parent == nil {
		return context.command.Name
	}
	return context.parent.Name() + " " + context.command.Name
}

func (context *Context) ShowHelp(out io.Writer) error {
	fmt.Fprintln(out, "NAME:")
	name := context.Name()

	if context.command.Description != "" {
		name += " - " + context.command.Description
	}

	fmt.Fprintln(out, " ", name)

	fmt.Fprintln(out)
	fmt.Fprintln(out, "USAGE:")
	usage := context.Name()

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
			name := strings.Join(append([]string{command.Name}, command.Aliases...), ",")
			tw.Add(" ", name, command.Description)
		}
		tw.Flush()
	}

	if len(context.command.Options) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "OPTIONS:")
		tw := tablewriter.New(out)
		for _, option := range context.command.Options {
			help := option.Help()
			tw.Add(" ", help[0], help[1])
		}
		tw.Flush()
	}

	if context.command.Copyright != "" {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "COPYRIGHT:")
		fmt.Fprintln(out, " ", context.command.Copyright)
	}

	if context.command.Version != "" {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "VERSION:")
		fmt.Fprintln(out, " ", context.command.Version)
	}

	return nil
}
