package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func InstallSSR() {
	prompt := promptui.Prompt{
		Label:     "是否安装 privoxy 来支持命令行 http 代理",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(Red(err.Error()))
		return
	}
	// 克隆ssr库
	err = RunCommand("sudo", "git", "clone", "-b", "manyuser", ssrRepo, installPath)
	if err != nil {
		fmt.Println(err)
	}
	err = RunCommand("sudo", "mkdir", "/usr/local/share/shadowsocksr/conf/")
	if err != nil {
		fmt.Println(err)
	}
	if result == "y" {
		InstallPrivoxy()
	}
}
