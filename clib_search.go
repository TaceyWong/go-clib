package goclib

import "github.com/urfave/cli/v2"

var SearchCMD = &cli.Command{

	Name:      "search",
	Usage:     "搜索包",
	UsageText: "[选项] [查询...]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "no-color",
			Aliases: []string{"n"},
			Usage:   "关闭彩色输出",
		},
		&cli.BoolFlag{
			Name:    "skip-cache",
			Aliases: []string{"c"},
			Usage:   "跳过搜索缓存",
		},
		&cli.BoolFlag{
			Name:    "json",
			Aliases: []string{"j"},
			Usage:   "生成JSON序列化输出",
		},
	},
	Action: func(ctx *cli.Context) error { return nil },
}
