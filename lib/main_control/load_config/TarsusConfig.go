package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type TarsusServerConfig struct {
	Server struct {
		Project  string `yaml:"Project"`
		Database struct {
			URL      string `yaml:"URL"`
			Username string `yaml:"Username"`
			Password string `yaml:"Password"`
		} `yaml:"Database"`
	} `yaml:"Server"`
}

func LoadConfig() TarsusServerConfig {
	// 读取配置文件
	data, err := os.ReadFile("tarsus.config.yaml")
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
	fmt.Printf("读取的内容:\n%s\n", string(data))

	// 解析配置文件
	var config TarsusServerConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}

	fmt.Printf("总配置: %+v\n", config)

	// 打印配置信息
	fmt.Printf("项目配置: %s\n", config.Server.Project)
	fmt.Printf("数据库URL: %s\n", config.Server.Database.URL)
	fmt.Printf("数据库用户名: %s\n", config.Server.Database.Username)
	fmt.Printf("数据库密码: %s\n", config.Server.Database.Password)
	return config
}
