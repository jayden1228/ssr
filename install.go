package main

import "fmt"

func InstallSSR() {
	// 克隆ssr库
	err := RunCommand("sudo", "git", "clone", "-b", "manyuser", ssrRepo, installPath)
	if err != nil {
		fmt.Println(err)
	}
	err = RunCommand("sudo", "mkdir", "/usr/local/share/shadowsocksr/conf/")
	if err != nil {
		fmt.Println(err)
	}
}
