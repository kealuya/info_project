package configs

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	config map[string]map[string]string
	once   sync.Once
)

func init() {
	if config == nil {
		_, err := loadConfig()
		if err != nil {
			panic(fmt.Sprintf("加载配置失败: %v", err))
		}
		fmt.Println("加载配置成功")
	}
}

// LoadConfig 加载配置
func loadConfig() (map[string]map[string]string, error) {
	var err error

	once.Do(func() {
		// 初始化配置
		err = initConfig()
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}

// initConfig 初始化配置
func initConfig() error {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 添加配置文件路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")  // 优先级2
	viper.AddConfigPath("./configs") // 优先级3

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 将配置解析到结构体中
	// 创建新的配置实例
	config = make(map[string]map[string]string)
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	return nil
}

// GetConfig 获取配置
func GetConfig() map[string]map[string]string {
	return config
}
