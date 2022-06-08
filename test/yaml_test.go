package test

import (
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const configFile = "app.yaml"
const defaultFilePath = "/etc/app/"
const etcdPath = "/config/{id}/dblog/app.yaml"

type Setting struct{}

//从etcd读取配置
func applyEtcdSetting(addr string, id string, setting *Setting) error {
	cfg := strings.ReplaceAll(etcdPath, "{id}", id)
	viper.AddRemoteProvider("etcd", addr, cfg)
	viper.SetConfigFile("yaml")
	err := viper.ReadRemoteConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(setting)
}

//从本地文件读取配置
func applyFileSetting(configFilePath string, setting *Setting) error {
	fullPath := path.Join(configFilePath, configFile)
	viper.SetConfigFile(fullPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(setting)
}

func TestYaml(t *testing.T) {
	pflag.String("configPath", defaultFilePath, "config file path")
	pflag.String("etcd", "http://127.0.0.1:4001", "etcd server address, eg. http://127.0.0.1:4001")
	pflag.String("id", "id:12345", "server identity")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	fmt.Println(viper.GetString("configPath"))
	fmt.Println(viper.GetString("etcd"))
	fmt.Println(viper.GetString("id"))
	// fmt.Println(viper.GetString(""))
}
