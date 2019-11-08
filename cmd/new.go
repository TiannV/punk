package cmd

// @CreateTime: Oct 31, 2019 11:26 AM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 8, 2019 5:05 PM
// @Description: 新建功能（创建组织树等）

import (
	"fmt"
	"os"
	"path"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"punk/common"
)

var (
	dirFlag string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create something new",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "create a new project with specify dir tree.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Project Type",
			Items: []string{"advance", "normal"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			ErrorMessage.Printf("%s", err)
			return
		}

		result += "_project"                      // 配置文件中是xx_project
		dirStruct := viper.Get(result).([]string) // 通过类型断言来转换为[]string类型
		// currentPath, pathErr := filepath.Abs(filepath.Dir(os.Args[0])) // 获取punk当前运行的路径
		// if pathErr != nil {
		// 	ErrorMessage.Printf("%s", pathErr)
		// 	return
		// }

		for _, subDir := range dirStruct {
			targetPath := path.Join(dirFlag, subDir)
			if ok, _ := common.PathExists(targetPath); ok {
				// 如果目录存在不做任何事情
			} else {
				os.Mkdir(targetPath, 0777)
				InfoMessage.Printf("%s \n", targetPath)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.AddCommand(createProjectCmd)
	createProjectCmd.Flags().StringVarP(&dirFlag, "dir", "d", ".", "specify the path u want to create project")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
