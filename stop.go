package main

import (
	"fmt"
	"os"
)

func Stop() {
	err := os.Chdir(installPath + "/shadowsocks")
	if err != nil {
		fmt.Println(err)
	}
	RunCommand("sudo", "python", "local.py", "-d", "stop")
}
