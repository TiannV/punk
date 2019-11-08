package cmd

// @CreateTime: Nov 6, 2019 7:23 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 8, 2019 3:40 PM
// @Description: 读写配置

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"punk/common"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config specified args for punk",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all config item",
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range viper.AllKeys() {
			InfoMessage.Printf("%s : ", item)
			OKMessage.Printf(" %s \n", viper.Get(item))
		}
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set specified config for punk",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if viper.IsSet(args[0]) { // 判断配置文件中是否有此项目
			viper.Set(args[0], args[1])
			OKMessage.Printf("%s has been changed to %s", args[0], args[1])
			viper.WriteConfig()
		} else {
			ErrorMessage.Println("")
		}
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get specified config of punk",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			ErrorMessage.Printf("%s", common.ErrConfigGetByOneKey)
		} else {
			configValue := viper.Get(args[0])
			if configValue != nil {
				OKMessage.Println(viper.Get(args[0]))
			} else {
				WarningMessage.Printf("%s", common.ErrNoConfigKey)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
