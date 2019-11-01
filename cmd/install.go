package cmd

// @CreateTime: Jul 6, 2019 7:19 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: tianwei
// @Last Modified Time: Nov 1, 2019 5:27 PM
// @Description: 为系统安装指定的软件

import (
	"fmt"
	"os/exec"
	"punk/utils"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version string

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.AddCommand(listPackageCmd)
	installCmd.AddCommand(python3Cmd)
	python3Cmd.Flags().StringVarP(&version, "version", "v", "3.8.0", "python3 version (required)")
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

// python3安装，仅支持macOS和Linux操作系统
var python3Cmd = &cobra.Command{
	Use:   "python3",
	Short: "install the python3 with the version of flag",
	Long:  `install the python3 with the version of flag`,
	Run: func(cmd *cobra.Command, args []string) {
		isInstalled := exec.Command("python3", "-V")
		output, err := isInstalled.Output()
		if err == nil {
			color.Blue("您已安装python3，版本号为： " + string(output))
			color.Yellow("确认继续安装?(Y/N)")
			for {
				var confirm string
				fmt.Scanln(&confirm)
				if confirm == "Y" {
					break
				}
				if confirm == "N" {
					return
				}
			}
		}
		color.Red(version)

		if runtime.GOOS == "darwin" {
			suffix := "-macosx10.6.pkg"
			if version >= "3.8.0" {
				suffix = "-macosx10.9.pkg"
			}
			pkgName := "python-" + version + suffix
			pkgPath := "https://www.python.org/ftp/python/" + version + "/" + pkgName
			fmt.Println(pkgPath)
			err = utils.CmdExec("wget -P pkgs/ " + pkgPath)
			if err != nil {
				return
			}
			err = utils.CmdExec("sudo installer -pkg pkgs/" + pkgName + " -target /")
			if err != nil {
				return
			}

			export := "export PATH=/Library/Frameworks/Python.framework/Versions/" + version[0:3] + "/bin:${PATH}"
			err = utils.CmdExec("echo " + export + " >> ~/.bash_profile && source ~/.bash_profile")
			if err != nil {
				return
			}
			color.Green("python " + version + " has installed successfilly, restart terminal to take effect")
		} else if runtime.GOOS == "linux" {
			pkgName := "Python-" + version + ".tgz"
			pkgPath := "https://www.python.org/ftp/python/" + version + "/" + pkgName
			fmt.Println(pkgPath)
			err = utils.CmdExec("wget -P pkgs/ " + pkgPath)
			if err != nil {
				return
			}

			if version >= "3.7.0" {
				err = utils.CmdExec("yum install libffi-devel -y")
				if err != nil {
					return
				}
			}

			err = utils.CmdExec("yum -y install zlib* && mkdir -p /usr/local/python3")
			if err != nil {
				return
			}

			err = utils.CmdExec("tar -zxvf pkgs/" + pkgName + " -C /usr/local/python3")
			if err != nil {
				return
			}

			err = utils.CmdExec("cd /usr/local/python3/Python-" + version + " && ./configure --prefix=/usr/local/python3 && make && make install")
			if err != nil {
				return
			}

			err = utils.CmdExec("rm -f /usr/bin/python3 && ln -s /usr/local/python3/bin/python3 /usr/bin/python3")
			if err != nil {
				return
			}

			err = utils.CmdExec("echo export PATH=$PATH:$HOME/bin:/usr/local/python3/bin >> ~/.bash_profile")
			if err != nil {
				return
			}

			err = utils.CmdExec("source ~/.bash_profile")
			if err != nil {
				return
			}

			color.Green("python " + version + "has installed successfilly")
		} else {
			color.Red("暂不支持maxOS和Linux以外的操作系统")
		}
	},
}
