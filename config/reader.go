package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Webroot    string
	ListenPort int
	ListenAddr string
}

func (c *Config) Load(fname string) error {
	yamlf, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlf, c)
	return err
}
