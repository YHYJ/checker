/*
File: root.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:54:23

Description: 程序未带子命令或参数时执行
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "checker",
	Short: "用于进行系统检查",
	Long:  `Checker是适用于Arch Linux的系统检查工具`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// 由main.main调用
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 定义全局Flag
	rootCmd.Flags().BoolP("help", "h", false, "Help for Checker")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scleaner.yaml)")
}
