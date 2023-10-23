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

var rootCmd = &cobra.Command{
	Use:   "checker",
	Short: "For system checks",
	Long:  `checker is a system checking tool for Arch Linux.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for checker")
}
