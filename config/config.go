package config

import (
	_ "embed"
	"github.com/HarryWang29/echo_mind/pkg/util"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

//go:embed config.yaml
var configYaml string

type Config struct {
	DataSource *DataSourceConfig `yaml:"data_source"`
	Wechat     *WechatConfig     `yaml:"wechat"`
}

type DataSourceConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type WechatConfig struct {
	WatchDir []struct {
		Dir string `yaml:"dir"`
		Key string `yaml:"key"`
	} `yaml:"watch_dir"`
	WatchID []struct {
		Id  string `yaml:"id"`
		Key string `yaml:"key"`
	} `yaml:"watch_id"`

	Path      string            `yaml:"-"`
	WatchInfo []WechatWatchInfo `yaml:"-"`
}
type WechatWatchInfo struct {
	Hash string `yaml:"-"`
	Id   string `yaml:"-"`
	Path string `yaml:"-"`
	Key  string `yaml:"-"`
}

func NewConfig() *Config {
	cfg := &Config{}
	err := yaml.Unmarshal([]byte(configYaml), &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func GetWechatConfig(cfg *Config) *WechatConfig {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cfg.Wechat.Path = path.Join(dir, "Library", "Containers", "com.tencent.xinWeChat", "Data", "Library", "Application Support", "com.tencent.xinWeChat", "2.0b4.0.9")
	dirs := make(map[string]string)
	for _, v := range cfg.Wechat.WatchDir {
		dirs[v.Dir] = v.Key
	}
	for _, v := range cfg.Wechat.WatchID {
		hex := util.HashHex(util.MD5, v.Id)
		dirs[hex] = v.Key
	}
	cfg.Wechat.WatchInfo = make([]WechatWatchInfo, 0, len(dirs))
	for k, v := range dirs {
		cfg.Wechat.WatchInfo = append(cfg.Wechat.WatchInfo, WechatWatchInfo{
			Id:   "",
			Hash: k,
			Path: path.Join(cfg.Wechat.Path, k),
			Key:  v,
		})
	}
	if len(cfg.Wechat.WatchInfo) == 0 {
		panic("watch dir or id is empty")
	}
	return cfg.Wechat
}

func GetDataSourceConfig(cfg *Config) *DataSourceConfig {
	return cfg.DataSource
}
