package common

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 全局配置结构
type Config struct {
	ServerURL   string `mapstructure:"server_url"`   // 上报地址
	IntervalSec int    `mapstructure:"interval_sec"` // 心跳间隔（秒）
	LogLevel    string `mapstructure:"log_level"`    // 日志等级
}

var Cfg Config

// LoadConfig 从配置文件加载
func LoadConfig(path string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}
	return nil
}
