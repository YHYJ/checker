/*
File: aur.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 13:33:21

Description: 程序子命令'aur'时执行
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/checker/function"
)

// aurCmd represents the aur command
var aurCmd = &cobra.Command{
	Use:   "aur",
	Short: "检查AUR包同步情况",
	Long:  `检查本地安装的AUR包是否和源保持同步`,
	Run: func(cmd *cobra.Command, args []string) {
		function.AURChecker()
	},
}

func init() {
	aurCmd.Flags().BoolP("help", "h", false, "Help for aur")
	rootCmd.AddCommand(aurCmd)
}
