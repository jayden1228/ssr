package main

import (
	"fmt"
	"os"
	"strings"
)

func Start() {
	if len(os.Args) > 2 {
		err := os.Chdir(installPath + "/shadowsocks")
		if err != nil {
			fmt.Println(err)
		}
		name := os.Args[2]
		if !strings.Contains(name, ".json") {
			name += ".json"
		}
		RunCommand("sudo", "python", "local.py", "-d", "start", "-c", installPath+"/conf/"+name)
	} else {
		fmt.Println(Red("config name is required"))
	}
}


