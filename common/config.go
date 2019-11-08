package common

// @CreateTime: Nov 6, 2019 7:26 PM
// @Author: ant1wv2
// @Contact: ant1wv2@gmail.com
// @Last Modified By: ant1wv2
// @Last Modified Time: Nov 7, 2019 4:13 PM
// @Description: 利用Viper进行配置文件的读写

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	configFile     = ".punk.json"
	configFileName = ".punk" // 配置文件尽量还是隐藏为好
	defaultIPURL   = "ifconfig.me"
)

var (
	// HomeDir punk配置文件存放根路径
	HomeDir               string
	defaultAdvanceProject = []string{"0-idea", "1-requirement", "2-architecture", "3-prototype", "4-design", "5-code", "6-review", "7-summary", "8-bp", "9-operate"}
	defaultNormalProject  = []string{"resource", "code", "docs"}
)

// Config 成员标签必须大写，否则marshal的时候无法转换
type Config struct {
	// 默认指令查询或下载地址
	IPURL string `json:"ip_url"`
	// 目录列表
	Advance []string `json:"advance_project"`
	Normal  []string `json:"normal_project"`
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	HomeDir = home
}

// IsConfigFileExsit 检查配置文件是否存在
func IsConfigFileExsit() bool {
	if fileRef, err := os.Stat(path.Join(HomeDir, configFileName)); err != nil || fileRef.IsDir() {
		return false
	}
	return true
}

// ReadConfigFile 读取配置
func ReadConfigFile() (*Config, error) {
	viper.SetConfigType("json")
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(HomeDir)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		IPURL:   viper.GetString("ip_url"),
		Advance: viper.GetStringSlice("advance_project"),
		Normal:  viper.GetStringSlice("normal_project"),
	}, nil
}

// WriteConfigFile 写入配置
func WriteConfigFile(cfg *Config) error {
	_, err := os.Create(path.Join(HomeDir, configFile))
	if err != nil {
		fmt.Println("thisi is why")
		return err
	}
	viper.AddConfigPath(HomeDir)
	viper.SetConfigName(".punk")
	viper.SetConfigType("json")
	viper.Set("ip_url", cfg.IPURL)
	viper.Set("advance_project", cfg.Advance)
	viper.Set("normal_project", cfg.Normal)
	fmt.Println(viper.GetString(cfg.IPURL))
	return viper.WriteConfig()
}

// WriteDefaultConfigFile 当配置文件不存在时调用写入默认配置
func WriteDefaultConfigFile() error {
	return WriteConfigFile(&Config{
		IPURL:   defaultIPURL,
		Advance: defaultAdvanceProject,
		Normal:  defaultNormalProject,
	})
}
