package main

import (
	"fmt"
	"os"
)

func Hp() {
	if len(os.Args) == 2 {
		HpHelp()
	} else if len(os.Args) > 2 {
		switch os.Args[2] {
		case "install":
			InstallPrivoxy()
		default:
			HpHelp()
		}
	} else {
		HpHelp()
	}
}

func HpHelp() {
	fmt.Println("命令行 hp http 代理配置工具")
	fmt.Println()
	fmt.Println(Yellow("Usage:"))
	fmt.Println(Blue("    ssr hp <command> [--parameter1=value1 --parameter2=value2 ...]"))
	fmt.Println()
	fmt.Println(Yellow("Commands:"))
	fmt.Println(Red("    install        安装 http 代理软件"))
	fmt.Println()
}

func InstallPrivoxy() {
	// 安装 privoxy
	err := RunCommand("sudo", "apt-get", "update")
	if err != nil {
	}
	err = RunCommand("sudo", "apt-get", "-y", "install", "privoxy")
	if err != nil {
		return
	}
	// 配置 privoxy
	err = RunCommand("sudo", "sh", "-c", "\"\"sudo echo -e 'forward-socks5 / 127.0.0.1:1080 .' >> /etc/privoxy/config\"\"")
	if err != nil {
		return
	}
	err = RunCommand("sudo", "service", "privoxy", "restart")
	if err != nil {
		return
	}
	fmt.Println(Yellow("append \"http_proxy=http://127.0.0.1:8118 https_proxy=http://127.0.0.1:8118\" to your bash profile and run \"source ~/.zshrc\" ro \"source ~/.bashrc\""))
}
