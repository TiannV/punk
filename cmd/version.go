package cmd

// @CreateTime: Nov 8, 2019 10:38 AM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 8, 2019 10:47 AM
// @Description: Modify Here, Please

import (
	"punk/common"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version number of punk",
	Run: func(cmd *cobra.Command, args []string) {
		version, repo := common.Version, common.Repo
		InfoMessage.Printf("Punk version: v%s \n", version)
		InfoMessage.Printf("Source code is here: %s \n", repo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
