package config
import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type BotConfig struct {
	ConsumerKey string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	Token string `json:"token"`
	Secret string `json:"secret"`
}

func Parse(filename string) (BotConfig, error) {
	var config BotConfig
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read config file:", err)
		return config, err
	}

	err = json.Unmarshal(jsonString, &config)
	if err != nil {
		log.Println("Failed to json unmarshal:", err)
		return config, nil
	}
	return config, nil
}
