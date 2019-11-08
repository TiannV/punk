package common

// @CreateTime: Nov 7, 2019 11:45 AM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 8, 2019 3:37 PM
// @Description: 错误信息输出

var (
	// ErrFolderIsExsit 目录已经存在
	ErrFolderIsExsit = "folder is already exsit, don't need create repeatly\n"
	// ErrNoConfigKey punk不存在该配置项
	ErrNoConfigKey = "Punk don't have the key which you specified\n"
	// ErrConfigBroken 配置文件损坏
	ErrConfigBroken = "The config file was broken\n"
	// ErrConfigGetByOneKey 获取配置时仅能一次指定一个关键词
	ErrConfigGetByOneKey = "You should specify one key which in config list\n"
)
