package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ProjectLocation string
}

func (c *Config) LoadConfig() *Config {
	// 从YAML文件读取数据
	data, err := ioutil.ReadFile("cfg.yaml")
	if err != nil {
		log.Printf("Error reading config file: %v. Creating new config file...", err)
		c.SaveConfig()
		return c
	}
	c2 := &Config{}
	err = yaml.Unmarshal(data, c2)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("ProjectConfig: %+v", c2)
	//c.ProjectLocation = c2.ProjectLocation
	return c2
}

func (c *Config) SaveConfig() {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile("cfg.yaml", data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
