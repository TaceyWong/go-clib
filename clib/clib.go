package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	gc "goclib"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.AppHelpTemplate = gc.AppHelpTemplate
	cli.CommandHelpTemplate = gc.CommandHelpTemplate
	cli.SubcommandHelpTemplate = gc.SubcommandHelpTemplate
	cli.FishCompletionTemplate = gc.FishCompletionTemplate
	cli.MarkdownDocTemplate = gc.MarkdownDocTemplate
	app := &cli.App{
		Name:      "clib",
		Usage:     "C语言包管理器",
		Authors:   []*cli.Author{{Name: "Tacey Wong", Email: "xinyong.wang@qq.com"}},
		Copyright: "© 2022 TaceyWong under GNU GENERAL PUBLIC LICENSE Version 3",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "禁止详情输出",
				Value:   true,
			},
		},
		Action: func(ctx *cli.Context) error {
			// TODO: 现实master所在位置
			s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
			s.Color("bgBlack", "bold", "fgRed")
			s.Suffix = "正在查询master"
			s.Start()                   // Start the spinner
			time.Sleep(4 * time.Second) // Run for some time to simulate work
			s.Stop()
			fmt.Println("结果")
			cli.ShowAppHelp(ctx)
			return nil
		},
	}
	app.EnableBashCompletion = true
	app.HideHelpCommand = false
	app.Version = "dev-0.0.1"
	app.HideVersion = false
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		gc.InitCMD, gc.InstallCMD, gc.UninstallCMD, gc.UpdateCMD, gc.UpgradeCMD,
		gc.SearchCMD, gc.BuildCMD, gc.ConfigureCMD}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
