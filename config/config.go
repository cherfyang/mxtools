package config

import (
	"encoding/json"
	"os"
)

var null = map[string]interface{}{}

func GetConfig(key string) map[string]interface{} {
	data, err := os.ReadFile("config.yaml")
	if err != nil {

		return null
	}
	kv := make(map[string]interface{})
	err = json.Unmarshal(data, &kv)
	if err != nil {

		return null
	}
	if kv[key] == nil {

		return null
	}
	return kv[key].(map[string]interface{})
}
