package main

import (
	"fmt"
	"os"
	"os/exec"
)

var installPath = "/usr/local/share/shadowsocksr"
var ssrRepo = "https://github.com/shadowsocksr-backup/shadowsocksr.git"

func main() {
	if len(os.Args) == 1 {
		CommandHelp()
	} else if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			InstallSSR()
		case "uninstall":
			UnInstallSSR()
		case "config":
			ConfigCommand()
		case "start":
			Start()
		case "stop":
			Stop()
		case "hp":
			Hp()
		default:
			CommandHelp()
		}
	}
}

func CommandHelp() {
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
	fmt.Println(Red("    hp             终端 http 代理相关"))
	fmt.Println()
}

func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(Red(string(tmp)))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
