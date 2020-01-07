package main

import "fmt"

func StartHelp() {
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
