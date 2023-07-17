package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// // 要大写才能被第三方包序列化
type openMethod struct {
	Name   string `yaml:"Name"`
	Method string `yaml:"Method"`
}
type ProjectConfig struct {
	A    string       `yaml:"a"`
	B    int          `yaml:"b"`
	Open []openMethod `yaml:"Open"`
}

func (c *ProjectConfig) GetOpenMethod(name string) string {
	for _, open := range c.Open {
		if open.Name == name {
			return open.Method
		}
	}
	return ""
}
func (c *ProjectConfig) SetOpenMethod(name string, method string) {
	for i, open := range c.Open {
		if open.Name == name {
			// 要用数组下表改 ，否则修改不上
			c.Open[i].Method = method
			return
		}
	}
	c.Open = append(c.Open, openMethod{name, method})
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
	//log.Printf("ProjectConfig: %+v", c2)
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
