package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Config struct {
	Server string
	Local  string
	Mysql  struct {
		Port int
		Host string
	}
}

func yamlTest() {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("read yaml file faild!")
		return
	}
	conf := Config{}
	if err = yaml.Unmarshal(data, &conf); err != nil {
		fmt.Println("Unmarshal faild")
		//return
	}

	fmt.Println(conf)
	fmt.Println(strings.Repeat("-", 100))
	t := make(map[interface{}]interface{})
	if err = yaml.Unmarshal(data, &t); err != nil {
		fmt.Println("Unmarshal map faild")
		return
	}
	fmt.Println(t)
}

func main() {
	yamlTest()
}
