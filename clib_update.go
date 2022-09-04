package goclib

import "github.com/urfave/cli/v2"

var UpdateCMD = &cli.Command{

	Name:      "update",
	UsageText: "[选项] [包名...]",
	Aliases:   []string{"up"},
	Usage:     "更新一个或多个包",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "out",
			Aliases: []string{"o"},
			Usage:   "改变输出`目录`",
		},
		&cli.PathFlag{
			Name:    "prefix",
			Aliases: []string{"P"},
			Usage:   "改变前缀`目录`(通常/usr/local)",
		},
		&cli.BoolFlag{
			Name:    "dev",
			Aliases: []string{"d"},
			Usage:   "构建测试依赖",
		},
		&cli.StringFlag{
			Name: "token",
			Aliases: []string{"t"},
			Usage: "读取私有内容的访问`Token`",
		},
		&cli.Int64Flag{
			Name:    "concurrency",
			Aliases: []string{"CC"},
			Usage:   "设置`并发数目`",
			Value:   12,
		},
	},
	Action:    func(ctx *cli.Context) error { return nil },
}
