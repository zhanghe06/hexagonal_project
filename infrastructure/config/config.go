package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"sync"

	"gopkg.in/yaml.v2"
)

type ServerConf struct {
	Lang string `yaml:"lang"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type MysqlConf struct {
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbHost          string `yaml:"db_host"`
	DbPort          string `yaml:"db_port"`
	DbName          string `yaml:"db_name"`
	Charset         string `yaml:"charset"`
	Timeout         string `yaml:"timeout"`
	TimeoutRead     string `yaml:"timeout_read"`
	TimeoutWrite    string `yaml:"timeout_write"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type Config struct {
	Server ServerConf
	Mysql  MysqlConf
}

var (
	configOnce sync.Once
	config     *Config
)

// NewConfig .
func NewConfig() *Config {
	configOnce.Do(func() {
		projectPath := ""
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("get config file path failure")
		}
		filePath := path.Dir(filename)
		projectPath = path.Dir(path.Dir(filePath))
		configFilePath := path.Join(projectPath, "server", "config", "config.yaml")
		//configFilePath := "/sysvol/conf/sap-cert-mgt.yaml"
		file, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			panic(fmt.Sprintf("load %v failed: %v", configFilePath, err))
		}

		err = yaml.Unmarshal(file, &config)
		if err != nil {
			panic(fmt.Sprintf("unmarshal yaml file failed: %v", err))
		}
		//fmt.Printf("%v", config)
	})

	return config
}
