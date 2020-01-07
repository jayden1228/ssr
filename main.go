package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"os/exec"
)

var installPath = "/usr/local/share/shadowsocksr"
var ssrRepo = "https://github.com/shadowsocksr-backup/shadowsocksr.git"

func main() {
	if len(os.Args) == 1 {
		showCommandHelp()
	} else if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			InstallSSR()
		case "uninstall":
			UnInstallSSR()
		case "config":
			Config()
		case "start":
			Start()
		case "stop":
			Stop()
		default:
			showCommandHelp()
		}
	}
}

func Start() {
	if len(os.Args) > 2 {
		RunCommand("sudo", "python", installPath+"/shadowsocks/local.py", "-d", "start", "-c", installPath+"/conf/"+os.Args[2]+".json")
	} else {
		fmt.Println(Red("config name is required"))
	}
}

func Stop() {
	RunCommand("sudo", "python", installPath+"/shadowsocks/local.py", "-d", "stop")
}

func UnInstallSSR() {
	prompt := promptui.Prompt{
		Label:     "Delete Resource",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		return
	}

	if result == "y" {
		RunCommand("sudo", "rm", "-rf", installPath)
		RunCommand("sudo", "apt", "remove", "-y", "privoxy")
	}
}
func InstallSSR() {
	// 克隆ssr库
	err := RunCommand("sudo", "git", "clone", "-b", "manyuser", ssrRepo, installPath)
	if err != nil {
		return
	}

	// 安装 privoxy
	err = RunCommand("sudo", "apt-get", "update")
	if err != nil {
		return
	}
	err = RunCommand("sudo", "apt-get", "-y", "install", "privoxy")
	if err != nil {
		return
	}
	// 配置 privoxy
	err = RunCommand("sudo", "sh", "-c", "\"sudo echo -e 'forward-socks5 / 127.0.0.1:1080 .' >> /etc/privoxy/config\"")
	if err != nil {
		return
	}
	err = RunCommand("sudo", "service", "privoxy", "restart")
	if err != nil {
		return
	}
	err = RunCommand("sudo", "sh", "-c", "\"sudo echo -e 'alias hp=\"http_proxy=http://127.0.0.1:8118 https_proxy=http://127.0.0.1:8118\"' >> ~/.zshrc\"")
	if err != nil {
		return
	}
	err = RunCommand("chsh", "-s", "/bin/zsh", "source", "~/.zshrc")
	if err != nil {
		return
	}

	err = RunCommand("sudo", "mkdir", "/usr/local/share/shadowsocksr/conf/")
	if err != nil {
		return
	}
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
