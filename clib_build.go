package goclib

import "github.com/urfave/cli/v2"

var BuildCMD = &cli.Command{

	Name:      "build",
	Usage:     "构建一个或多个包",
	UsageText: "[选项] [包名...]",
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
			Name:    "global",
			Aliases: []string{"g"},
			Usage:   "使用全局目标",
		},
		&cli.PathFlag{
			Name:    "clean",
			Aliases: []string{"C"},
			Usage:   "构建之前`清理目标`",
			Value:   "clean",
		},
		&cli.PathFlag{
			Name:    "test",
			Aliases: []string{"T"},
			Usage:   "`测试目标`而非构建",
			Value:   "test",
		},
		&cli.BoolFlag{
			Name:    "dev",
			Aliases: []string{"d"},
			Usage:   "构建测试依赖",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "强制动作，例如覆写文件",
		},
		&cli.BoolFlag{
			Name:    "skip-cache",
			Aliases: []string{"c"},
			Usage:   "配置时跳过缓存",
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
