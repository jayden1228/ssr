package main

import "github.com/manifoldco/promptui"

func UnInstallSSR() {
	prompt := promptui.Prompt{
		Label:     "是否确认卸载",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return
	}
	if result == "y" {
		RunCommand("sudo", "rm", "-rf", installPath)
		RunCommand("sudo", "apt", "remove", "-y", "privoxy")
		//TODO 删除privoxy 配置
	}
}
