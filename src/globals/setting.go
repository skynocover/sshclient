package globals

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Setting Config

type Config struct {
	Hosts []Host `json:"hosts"`
}

type Host struct {
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfiguration(filePath string) error {
	newconfig := func() error {
		Setting = Config{Hosts: []Host{
			{
				Name:     "",
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

const lable = "Current Store Host"

// Show ...
func (c *Config) Show() {

	orderlen := strings.Count("Order", "")
	namelen := strings.Count("Name", "")
	domainlen := strings.Count("Domain", "")
	userlen := strings.Count("User", "")
	passwordlen := strings.Count("Password", "")
	for _, host := range Setting.Hosts {
		temp := strings.Count(host.Name, "")
		if temp > namelen {
			namelen = temp
		}
		temp = strings.Count(host.Domain, "")
		if temp > domainlen {
			domainlen = temp
		}
		temp = strings.Count(host.User, "")
		if temp > userlen {
			userlen = temp
		}
		temp = strings.Count(host.Password, "")
		if temp > passwordlen {
			passwordlen = temp
		}
	}

	var totallen = orderlen + namelen + domainlen + userlen + passwordlen + 4

	for i := 0; i < (totallen-strings.Count(lable, ""))/2; i++ {
		fmt.Printf("=")
	}
	fmt.Printf("%s", lable)
	for i := 0; i < (totallen-strings.Count(lable, ""))/2; i++ {
		fmt.Printf("=")
	}
	fmt.Println("")

	// fmt.Println("==============Current Store Host============")

	fmt.Printf("Order  ")
	fmt.Printf("Name")
	for i := 0; i < namelen-strings.Count("Name", "")+2; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("Domain")
	for i := 0; i < domainlen-strings.Count("Domain", "")+2; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("User")
	for i := 0; i < userlen-strings.Count("User", "")+2; i++ {
		fmt.Printf(" ")
	}

	fmt.Println("Password")

	for i, host := range Setting.Hosts {
		ii := fmt.Sprintf("%d", i)
		fmt.Printf("%s", ii)
		for i := 0; i < orderlen-strings.Count(ii, "")+2; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", host.Name)
		for i := 0; i < namelen-strings.Count(host.Name, "")+2; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", host.Domain)
		for i := 0; i < domainlen-strings.Count(host.Domain, "")+2; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", host.User)
		for i := 0; i < userlen-strings.Count(host.User, "")+2; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", host.Password)
		fmt.Println("")
	}
	fmt.Println("")
}
