// @CreateTime: Jun 28, 2019 6:04 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 8, 2019 3:09 PM
// @Description: punk根命令

package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"punk/common"
)

const configFileName = ".punk"

var (
	// HomeDir 操作系统用户家目录
	HomeDir string
	// ErrorMessage 出现错误时，显示的颜色
	ErrorMessage = color.New(color.FgRed)
	// OKMessage 成功正确操作时，显示绿色
	OKMessage = color.New(color.FgGreen)
	// WarningMessage 出现警告时，显黄色
	WarningMessage = color.New(color.FgYellow)
	// InfoMessage 一般信息时，显示的蓝色
	InfoMessage = color.New(color.FgHiBlue)
)

// punk命令行工具简介
var rootCmd = &cobra.Command{
	Use:   "punk",
	Short: "A set of useful tools for saving your life.", // Short信息出现在命令介绍中
	Long: `
	 ________  ___  ___  ________   ___  __       
	|\   __  \|\  \|\  \|\   ___  \|\  \|\  \     
	\ \  \|\  \ \  \\\  \ \  \\ \  \ \  \/  /__   
	 \ \   ____\ \  \\\  \ \  \\ \  \ \   ___  \  
	  \ \  \___|\ \  \\\  \ \  \\ \  \ \  \\ \  \ 
	   \ \__\    \ \_______\ \__\\ \__\ \__\\ \__\
	    \|__|     \|_______|\|__| \|__|\|__| \|__|
	`,
	// Run: func(cmd *cobra.Command, args []string) { }, // 根命令仅需输出帮助信息即可
}

// Execute 函数会将所有子命令加入到punk体系下，并会自动设置对应的flags，此函数仅执行一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}
}

// init 初始化punk
func init() {
	// 每次执行命令时都会调用
	cobra.OnInitialize(initConfig)
	// 自定义颜色输出
	cobra.AddTemplateFunc("StyleHeading", color.New(color.BgHiMagenta, color.FgHiWhite).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
	).Replace(usageTemplate)
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	rootCmd.SetUsageTemplate(usageTemplate)
}

// initConfig 初始化设置
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	HomeDir = home

	if ok := common.IsConfigFileExsit(); !ok {
		// 不存在配置文件时，重新写入一个默认配置
		err := common.WriteDefaultConfigFile()
		if err != nil {
			ErrorMessage.Printf("%s", err)
		}
	} else {
		_, err := common.ReadConfigFile()
		if err != nil {
			ErrorMessage.Printf("%s", common.ErrConfigBroken)
		}
	}
}
