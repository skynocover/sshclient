package order

import (
	"fmt"
	"ssh/src/globals"
	"ssh/src/ssh"
	"strconv"
	"strings"
)

var order string

func Do(input string) error {
	switch input {
	case "m":

		fmt.Println("Choose the Order to modify")
		fmt.Scanln(&order)
		i, err := strconv.Atoi(order)
		if err != nil {
			return fmt.Errorf("input convert to int failed, err: %v", err)
		}
		if i >= len(globals.Setting.Hosts) {
			return fmt.Errorf("input above the config length, config only %d now ", len(globals.Setting.Hosts))
		}
		fmt.Println("Use {parameter}={config} to modify, for example: name=server")
		var modify string
		fmt.Scanln(&modify)
		inputarr := strings.Split(modify, "=")
		if len(inputarr) != 2 {
			return fmt.Errorf("input formate failed")
		}
		switch inputarr[0] {
		case "name":
			globals.Setting.Hosts[i].Name = inputarr[1]
		case "domain":
			globals.Setting.Hosts[i].Domain = inputarr[1]
		case "user":
			globals.Setting.Hosts[i].User = inputarr[1]
		case "password":
			globals.Setting.Hosts[i].Password = inputarr[1]
		default:
			return fmt.Errorf("input %s not found", inputarr[0])
		}
		return globals.Setting.Save()
	default:
		i, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("input convert to int failed, err: %v", err)
		}
		if i >= len(globals.Setting.Hosts) {
			return fmt.Errorf("input above the config length, config only %d now ", len(globals.Setting.Hosts))
		}

		host := globals.Setting.Hosts[i]
		ssh.Dial(host.Domain, host.User, host.Password)
	}
	return nil
}
