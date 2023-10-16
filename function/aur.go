/*
File: aur.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-27 12:59:16

Description: 子命令`aur`功能函数
*/

package function

func AURChecker() error {
	aurArgs := []string{"-c", `pacman -Qmq | parallel 'result=$(package-query -AQ -f "%v" "{}" | uniq -d | wc -l); [ $result -eq 0  ] && echo "{}"'`}
	if err := RunCommand("bash", aurArgs); err != nil {
		return err
	}
	return nil
}
