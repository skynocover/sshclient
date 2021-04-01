package globals

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var Setting Config

type Config struct {
	Hosts []Host `json:"hosts"`
}

type Host struct {
	Domain   string `json:"domain"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfiguration(filePath string) error {

	newconfig := func() error {
		Setting = Config{Hosts: []Host{
			{
				Domain:   "",
				User:     "",
				Password: "",
			},
		}}
		jsonByte, _ := json.Marshal(Setting)
		if err := ioutil.WriteFile(filePath, jsonByte, 0755); err != nil {
			return err
		}

		return errors.New("initial the new config")
	}

	if _, err := os.Stat(filePath); err != nil {
		return newconfig()
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return newconfig()
	}
	err = json.Unmarshal(content, &Setting)
	if err != nil {
		return newconfig()
	}
	return nil
}

// Save 儲存設定檔
func (c *Config) Save() error {
	content, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("setting.json", content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
