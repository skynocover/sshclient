package order

import (
	"ssh/src/globals"
	"ssh/src/ssh"
	"strconv"
	"fmt"
)

func Do(input string){
	switch input {
	case "fix":
	default:
		i,err := strconv.Atoi(input)
		if err!=nil {
			fmt.Println("input convert to int failed, err: ", err)
			return
		}
		if i>=len(globals.Setting.Hosts) {
			fmt.Printf("input above the config length, config only %d now ", len(globals.Setting.Hosts))
			return
		}

		host := globals.Setting.Hosts[i]
		ssh.Dial(host.Domain,host.User,host.Password)
	}
}