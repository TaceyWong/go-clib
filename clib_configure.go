package goclib

import "github.com/urfave/cli/v2"

var ConfigureCMD = &cli.Command{

	Name:      "configure",
	Usage:     "配置一个或多个包",
	UsageText: "[选项] [名称...]",
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
			Usage:   "配置开发依赖",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "强制动作，例如覆写文件",
		},
		&cli.BoolFlag{
			Name:    "cflags",
			Aliases: []string{"flags"},
			Usage:   "不配合，仅输出编译器flags",
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
