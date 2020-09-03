package main

import (
	"fmt"
	"github.com/B1ackAnge1/RoleBot/utils"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello World!")
	rawConfig, err := ioutil.ReadFile("config.toml") // just pass the file name
	if err != nil {
		fmt.Println("Error while load config file: " + err.Error())
		return
	}
	fmt.Println(utils.LoadConfig(string(rawConfig)))
}
