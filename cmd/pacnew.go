/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 11:24:15

Description: 程序子命令'pacnew'时执行
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/checker/function"
)

// pacnewCmd represents the pacnew command
var pacnewCmd = &cobra.Command{
	Use:   "pacnew",
	Short: "检查对比软件包新旧配置文件",
	Long:  `检查对比已安装软件包的新旧配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		function.PacnewChecker()
	},
}

func init() {
	pacnewCmd.Flags().BoolP("help", "h", false, "Help for pacnew")
	rootCmd.AddCommand(pacnewCmd)
}
