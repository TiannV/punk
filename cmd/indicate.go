package cmd

// @CreateTime: Jul 6, 2019 9:20 AM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Oct 31, 2019 11:30 AM
// @Description: 指出命令，譬如指出本地ip，当前文件夹大小等等

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// 父级命令
var indicateCmd = &cobra.Command{
	Use:   "indicate",
	Short: "Inidicate the anwser that you want to know",
	Long: `You can indicate ip address of current host. 
Get the time/date of region which you indicate. 
etc...`,
}

// 子命令
var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get ipv4 address of current host",
	Long: `Get ipv4 address of current host from internet,
so you should make sure your host online`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://ifconfig.me/ip")
		if err != nil {
			color.Red("Error Messgage: %s", err)
			os.Exit(1)
		} else {
			defer resp.Body.Close()
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				color.Red("Error Message: %s", err)
				os.Exit(1)
			}
			color.Yellow("Your Pulbic IP Address is:\n %s", string(contents))
		}
	},
}

var getWeatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get weather info from internet",
	Long:  "Get weather info from internet by indicating region",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("这里需要添加获取天气信息的代码")
	},
}

func init() {
	rootCmd.AddCommand(indicateCmd)
	indicateCmd.AddCommand(getIPCmd)
	indicateCmd.AddCommand(getWeatherCmd)
}
