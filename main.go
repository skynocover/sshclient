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

	var input string
	for {
		CallClear() //清空
		fmt.Println("==============Current Store Host============")

		for i, host := range globals.Setting.Hosts {
			fmt.Printf("Order: %d, %+v\n", i, host)
		}
		fmt.Println("What is your command? input Order to connect or exit to leave")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		order.Do(input)
	}
}

func CallClear() { //清空用
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
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
