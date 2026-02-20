package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	DouBao *DouBaoCfg `yaml:"douBao"`
	BaiDu  *BaiDuCfg  `yaml:"baiDu"`
}
type BaiDuCfg struct {
	BaseUrl string `yaml:"baseUrl"`
	ApiKey  string `yaml:"apiKey"`
}
type DouBaoCfg struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	ApiKey          string `yaml:"apiKey"`
	ModelID         string `yaml:"modelID"`
	BaseUrl         string `yaml:"baseUrl"`
}

var Cfg *Config

func init() {
	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println("加载配置文件报错:", err)
		panic(err)
	}
	// 3. 将 YAML 内容解析到结构体中
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		fmt.Printf("解析 config yaml 失败: %v\n", err)
		panic(err)
	}
}
