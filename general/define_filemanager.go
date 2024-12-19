/*
File: define_filemanager.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-24 16:41:17

Description: 文件管理
*/

package general

import (
	"io"
	"os"
	"path/filepath"
)

// FileExist 判断文件是否存在
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 文件存在返回 true，否则返回 false
func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// FolderEmpty 判断文件夹是否为空
//
//   - 包括隐藏文件
//
// 参数：
//   - dir: 文件夹路径
//
// 返回：
//   - 文件夹为空返回 true，否则返回 false
func FolderEmpty(dir string) bool {
	text, err := os.Open(dir)
	if err != nil {
		return true
	}
	defer text.Close()

	if _, err = text.Readdir(1); err == io.EOF {
		return true
	}
	return false
}

// CreateFile 创建文件，包括其父目录
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 错误信息
func CreateFile(file string) error {
	if FileExist(file) {
		return nil
	}
	// 创建父目录
	parentPath := filepath.Dir(file)
	if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
		return err
	}
	// 创建文件
	if _, err := os.Create(file); err != nil {
		return err
	}

	return nil
}

// WriteFile 写入内容到文件，文件不存在则创建，不自动换行
//
// 参数：
//   - filePath: 文件路径
//   - content: 内容
//   - mode: 写入模式，追加('a', O_APPEND, 默认)或覆盖('t', O_TRUNC)
//
// 返回：
//   - 错误信息
func WriteFile(filePath, content, mode string) error {
	// 确定写入模式
	writeMode := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	if mode == "t" {
		writeMode = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	// 将内容写入文件
	file, err := os.OpenFile(filePath, writeMode, 0666)
	if err != nil {
		return err
	}
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
}

// EmptyFile 清空文件内容，文件不存在则创建
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 错误信息
func EmptyFile(file string) error {
	// 打开文件，如果不存在则创建，文件权限为读写
	text, err := os.OpenFile(file, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer text.Close()

	// 清空文件内容
	if err := text.Truncate(0); err != nil {
		return err
	}
	return nil
}

// DeleteFile 删除文件，如果目标是文件夹则包括其下所有文件
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 错误信息
func DeleteFile(filePath string) error {
	if !FileExist(filePath) {
		return nil
	}
	return os.RemoveAll(filePath)
}
