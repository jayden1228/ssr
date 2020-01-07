package main

import "fmt"

func showCommandHelp() {
	fmt.Println("SSR 助手工具")
	fmt.Println()
	fmt.Println(Yellow("Usage:"))
	fmt.Println(Blue("    ssr <command> [--parameter1=value1 --parameter2=value2 ...]"))
	fmt.Println()
	fmt.Println(Yellow("Commands:"))
	fmt.Println(Red("    install        安装ssr软件"))
	fmt.Println(Red("    uninstall      卸载ssr软件"))
	fmt.Println(Red("    config         配置ssr服务器"))
	fmt.Println(Red("    start          开始ssr服务器"))
	fmt.Println(Red("    stop           停止ssr服务器"))
	fmt.Println()
}

func showConfigHelp() {
	fmt.Println("SSR 配置")
	fmt.Println()
	fmt.Println(Yellow("Usage:"))
	fmt.Println(Blue("    ssr config command [--parameter1=value1 --parameter2=value2 ...]"))
	fmt.Println()
	fmt.Println(Yellow("Commands:"))
	fmt.Println(Red("    ls                      列出现有的ssr配置文件"))
	fmt.Println(Red("    sub http://demo.com     通过url进行订阅"))
	fmt.Println()
}
