package easy_page

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	RunMode string   `json:"run_mode" yaml:"run_mode"`
	Port    string   `json:"port" yaml:"port"`
	Plugins []string `json:"plugins" yaml:"plugins"`
	Db      struct {
		ServerName string `json:"server_name" yaml:"server_name"`
		Url        string `json:"url" yaml:"url"`
	} `json:"db" yaml:"db"`
}

func LoadConfig(path string) (*Config, error) {
	dataBs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(dataBs, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
