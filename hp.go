package main

import "fmt"

func Hp() {

}

func InstallPrivoxy() {
	// 安装 privoxy
	err := RunCommand("sudo", "apt-get", "update")
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
	fmt.Println(Yellow("append \"http_proxy=http://127.0.0.1:8118 https_proxy=http://127.0.0.1:8118\" to your bash profile and run \"source ~/.zshrc\" ro \"source ~/.bashrc\""))
}
