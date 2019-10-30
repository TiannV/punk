package cmd

// @CreateTime: Jul 6, 2019 7:19 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Jul 7, 2019 11:38 PM
// @Description: 为系统安装指定的软件

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.AddCommand(listPackageCmd)
}

// 父级命令
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install softwares on indicated os",
	Long:  `Install softwares on indicated operate system`,
}

// 子命令
var listPackageCmd = &cobra.Command{
	Use:   "list",
	Short: "List softwares that could be installed",
	Long:  `List softwares which punk support to install`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("编写punk install list对应的代码")
	},
}
