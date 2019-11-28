cli
====

雑に cli をつくるための go ライブラリ

## Example

```
package main

import (
	"fmt"
	"os"

	"github.com/thamaji/cli"
)

func main() {
	command := &cli.Command{
		Name:        "hoge",
		Description: "Hoge CLI",
		Version:     "v1.0.0",
		ArgsUsage:   "ID [ID...]",
		Options: []cli.Option{
			&cli.StringOption{
				Name:        "username",
				Short:       "u",
				Description: "set username",
			},
			&cli.StringOption{
				Name:        "password",
				Short:       "p",
				Description: "set password",
			},
		},
		Action: func(ctx *cli.Context) error {
			if len(ctx.Args()) < 1 {
				return ctx.ShowHelp(os.Stdout)
			}

			username, err := ctx.StringOrInput("username")
			if err != nil {
				return err
			}

			password, err := ctx.StringOrPassword("password")
			if err != nil {
				return err
			}

			fmt.Println("username", username)
			fmt.Println("password", password)

			for _, id := range ctx.Args() {
				fmt.Println(id)
			}

			return nil
		},
	}

	if err := command.Run(os.Args, cli.ShowHelp(os.Stdout)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
```

```
$ go run main.go --help
NAME:
  hoge - Hoge CLI

USAGE:
  hoge [OPTIONS] ID [ID...]

OPTIONS:
  -u,--username=string set username
  -p,--password=string set password
  -v,--version         show version
  -h,--help            show help

VERSION:
  v1.0.0
```
