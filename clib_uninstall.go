package goclib

import "github.com/urfave/cli/v2"

var UninstallCMD = &cli.Command{

	Name:      "uninstall",
	Usage:     "卸载一个或多个包",
	UsageText: "[选项] [名称...]",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "prefix",
			Aliases: []string{"P"},
			Usage:   "改变前缀目录",
			Value:   "/usr/local",
		},
	},
	Action: func(ctx *cli.Context) error { return nil },
}
