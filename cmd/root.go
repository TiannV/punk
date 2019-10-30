// @CreateTime: Jun 28, 2019 6:04 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Jul 7, 2019 11:38 PM
// @Description: punk根命令

package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var cfgFile string

// punk命令行工具简介
var rootCmd = &cobra.Command{
	Use:   "punk",
	Short: "A set of useful tools for saving your life.", // Short信息出现在命令介绍中
	// Run: func(cmd *cobra.Command, args []string) { }, // 根命令仅需输出帮助信息即可
}

// Execute 函数会将所有子命令加入到punk体系下，并会自动设置对应的flags，此函数仅执行一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("s%", err)
		os.Exit(1)
	}
}

// init 初始化punk
func init() {

}

// initConfig 自定义配置文件读写
func initConfig() {
	if cfgFile != "" {
		// viper.SetConfigFile(cfgFile)
	} else {
		// 
	}
}
