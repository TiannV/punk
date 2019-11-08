package common

// @CreateTime: Nov 8, 2019 4:18 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 8, 2019 4:41 PM
// @Description: 工具方法

import "os"

// PathExists 判断路径是否已经存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}
