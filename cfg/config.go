package cfg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	SpaceLocation string
}

func loadConfigPath() string {
	home, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("找不到用户home路径:", err)
		home = ".\\"
	} else {
		home += "\\projectManager"
	}
	return home
}
func loadHomePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("找不到用户home路径:", err)
		home = ".\\"
	} else {
		home += "\\"
	}
	return home
}
func loadExecutablePath() string {
	home, err := os.Executable()
	if err != nil {
		fmt.Println("找不到用户home路径:", err)
		home = ".\\"
	}
	fmt.Println(home)
	return home
}
func (c *Config) LoadConfig() *Config {
	// 从YAML文件读取数据
	home := loadHomePath()
	data, err := ioutil.ReadFile(home + ".projectManager.yaml")
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
	//log.Printf("ProjectConfig: %+v", c2)
	//c.SpaceLocation = c2.SpaceLocation
	return c2
}

func (c *Config) SaveConfig() {
	home := loadHomePath()
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(home+".projectManager.yaml", data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
