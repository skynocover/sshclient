package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"ssh/src/globals"
	"ssh/src/order"
)

var configFile = flag.String("config", "setting.json", "customize the config file name")

func main() {

	flag.Parse()

	if err := globals.LoadConfiguration(*configFile); err != nil {
		log.Fatal(err)
	}

	for {
		CallClear() //清空

		globals.Setting.Show()
		fmt.Println(`What is your command? input Order to connect || "m" to modify || "h" to help || "exit" to leave`)
		var input string
		fmt.Scanln(&input)
		switch input {
		case "exit":
			return
		case "h":
			fmt.Println("modify   : modify the configuration")
			fmt.Println("{any int}: ssh in to machine")
			fmt.Println("Press any key to continue")
			fmt.Scanln(&input)
			continue

		default:
			if err := order.Do(input); err != nil {
				fmt.Println(err)
				fmt.Println("Press any key to continue")
				fmt.Scanln(&input)
			}
		}
	}
}

func CallClear() { //清空用
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	Value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		Value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
