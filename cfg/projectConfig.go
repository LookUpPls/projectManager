package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type openMethod struct {
	name   string `yaml:"name"`
	method string `yaml:"method"`
}
type ProjectConfig struct {
	A    string       `yaml:"a"`
	B    int          `yaml:"b"`
	open []openMethod `yaml:"open"`
}

func (c *ProjectConfig) GetOpenMethod(name string) string {
	for _, open := range c.open {
		if open.name == name {
			return open.method
		}
	}
	return ""
}
func (c *ProjectConfig) SetOpenMethod(name string, method string) {
	for _, open := range c.open {
		if open.name == name {
			open.method = method
			return
		}
	}
	c.open = append(c.open, openMethod{name, method})
}
func (c *ProjectConfig) LoadConfig(path string) *ProjectConfig {
	// 从YAML文件读取数据
	data, err := ioutil.ReadFile(path + ".projectManager.yaml")
	if err != nil {
		log.Printf("Error reading config file: %v. Creating new config file...", err)
		// 创建新的配置
		c.SaveConfig(path)
		return c
	}
	c2 := &ProjectConfig{}
	err = yaml.Unmarshal(data, c2)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("ProjectConfig: %+v", c2)
	return c2
}

func (c *ProjectConfig) SaveConfig(path string) {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(path+".projectManager.yaml", data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
