package goclib

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

/*
 init
*/

var InitCMD = &cli.Command{
	Name:  "init",
	Usage: "初始化一个新项目",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "manifest",
			Aliases: []string{"M"},
			Usage:   "给出清单文件的清单",
			Value:   "clib.json",
		},
	},
	Action: initAction,
}

func initAction(ctx *cli.Context) error {
	manifest := &survey.Input{
		Message: "配置清单文件",
		Default: "clib.json",
	}
	manifestA := ""
	survey.AskOne(manifest, &manifestA)
	var qs = []*survey.Question{
		{
			Name:     "pkg_name",
			Prompt:   &survey.Input{Message: "包名?", Default: "base"},
			Validate: survey.Required,
		},
		{
			Name:     "pkg_version",
			Prompt:   &survey.Input{Message: "版本号?", Default: "0.1.0"},
			Validate: survey.Required,
		},
		{
			Name:     "pkg_repo",
			Prompt:   &survey.Input{Message: "仓库?", Default: "0.1.0"},
			Validate: survey.Required,
		},
		{
			Name:     "pkg_keywords",
			Prompt:   &survey.Input{Message: "包关键词（以逗号分割）:", Default: "clib-package"},
			Validate: survey.Required,
		},
		{
			Name: "pkg_license",
			Prompt: &survey.Select{Message: "选择协议:",
				Options: []string{"MIT", "GNU3", "other"}},
		},
		{
			Name:     "pkg_desc",
			Prompt:   &survey.Input{Message: "包描述?", Default: "clib-package"},
			Validate: survey.Required,
		},

		{
			Name:     "pkg_makefile",
			Prompt:   &survey.Input{Message: "Makefile?", Default: "Makefile"},
			Validate: survey.Required,
		},
	}
	answers := struct {
		Name        string `survey:"pkg_name"`
		Version     string `survey:"pkg_version"`
		Description string `survey:"pkg_desc"`
		License     string `survey:"pkg_license"`
		Keywords    string `survey:"pkg_keywords"`
		Makefile    string `survey:"pkg_makefile"`
		Repo        string `survey:"pkg_repo"`
	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(answers)
	return nil
}
