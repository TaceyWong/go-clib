package goclib

import "github.com/urfave/cli/v2"

var InstallCMD = &cli.Command{

	Name:      "install",
	Aliases:   []string{"i"},
	Usage:     "安装一个或多个包",
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
			Usage:   "安装开发依赖",
		},
		&cli.BoolFlag{
			Name:    "save",
			Aliases: []string{"S"},
			Usage:   "保存clib.json或package.json种的依赖",
		},
		&cli.BoolFlag{
			Name:    "save-dev",
			Aliases: []string{"D"},
			Usage:   "保存clib.json或package.json种的开发依赖",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "强制动作，例如覆写文件",
		},
		&cli.BoolFlag{
			Name:    "skip-cache",
			Aliases: []string{"c"},
			Usage:   "安装的时候跳过缓存",
		},
		&cli.BoolFlag{
			Name:    "global",
			Aliases: []string{"g"},
			Usage:   "全局安装，不写入输出目录(缺省deps/)",
		},
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "读取私有内容的访问`Token`",
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
