package goclib

import "github.com/urfave/cli/v2"

var UpgradeCMD = &cli.Command{

	Name:      "upgrade",
	Usage:     "升级clib到指定或最新的版本",
	UsageText: "[选项] [名称...]",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "prefix",
			Aliases: []string{"P"},
			Usage:   "改变前缀`目录`(通常/usr/local)",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "强制动作，例如覆写文件",
		},
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "读取私有内容的访问`Token`",
		},
		&cli.StringFlag{
			Name:    "slug",
			Aliases: []string{"S"},
			Usage:   "clib项目的`slug`",
			Value:   "clibs/clib",
		},
		&cli.StringFlag{
			Name:    "tag",
			Aliases: []string{"T"},
			Usage:   "升级的目标`tag`",
			Value:   "latest",
		},
		&cli.Int64Flag{
			Name:    "concurrency",
			Aliases: []string{"CC"},
			Usage:   "设置`并发数目`",
			Value:   12,
		},
	},
	Action: func(ctx *cli.Context) error { return nil },
}
