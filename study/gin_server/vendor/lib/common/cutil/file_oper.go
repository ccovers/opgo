/*
 * author:        liujun
 * created:       2018-04-23 14:43
 * last modified: 2018-04-23 14:43
 * filename:      file_oper.go
 * description:   文件操作工具类
 */
package cutil

import (
	"os"
)

func FileExist(fpath string) (bool, error) {
	_, err := os.Stat(fpath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
