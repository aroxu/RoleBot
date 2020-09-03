package utils

import (
	"github.com/B1ackAnge1/RoleBot/structure"
	"github.com/BurntSushi/toml"
	"log"
)

//LoadConfig get rawConfig string and returns error, prefix, token
func LoadConfig(rawConfig string) (errLoadFailed error, prefix string, token string) {
	var config structure.Config
	_, err := toml.Decode(rawConfig, &config)
	if err != nil {
		log.Fatal(err)
		return err, "", ""
	}
	return nil, config.Prefix, config.Token
}