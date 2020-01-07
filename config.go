package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sparrc/go-ping"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	Server     string `json:"server"`
	ServerPort int64  `json:"server_port"`
}

func ConfigCommand() {
	if len(os.Args) == 2 {
		ConfigHelp()
	} else if len(os.Args) > 2 {
		switch os.Args[2] {
		case "ls":
			RunCommand("sudo", "ls", installPath+"/conf")
		case "sub":
			if len(os.Args) > 3 {
				SubFromUrl(os.Args[3])
			} else {
				fmt.Println(Red("ssr config sub http://demo.com 通过url进行订阅"))
			}
		case "add":
			if len(os.Args) > 3 {
				ParseSSR(string(os.Args[3]))
			} else {
				fmt.Println(Red("ssr config add ssr://cmM0LW... 通过ssr协议进行添加"))
			}
		case "edit":
			if len(os.Args) > 3 {
				EditConfigFileCustom(string(os.Args[3]))
			} else {
				fmt.Println(Red("ssr config new name 通过手动进行添加"))
			}
		case "ping":
			PingConfigs()
		default:
			ConfigHelp()
		}
	} else {
		ConfigHelp()
	}
}

func ConfigHelp() {
	fmt.Println("SSR 配置")
	fmt.Println()
	fmt.Println(Yellow("Usage:"))
	fmt.Println(Blue("    ssr config command [--parameter1=value1 --parameter2=value2 ...]"))
	fmt.Println()
	fmt.Println(Yellow("Commands:"))
	fmt.Println(Red("    ls                      列出现有的ssr配置文件"))
	fmt.Println(Red("    ping                    获取所有配置服务器延迟"))
	fmt.Println(Red("    sub http://demo.com     通过url进行订阅"))
	fmt.Println(Red("    add ssr://cmM0LW...     通过ssr协议添加"))
	fmt.Println(Red("    edit name               手动编辑配置文件，如果没有该配置则创建"))
	fmt.Println()
}

func SubFromUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	encodeString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(string(encodeString))
	if err != nil {
		panic(err)
	}
	configs := strings.Split(string(decodeBytes), "\n")
	for _, value := range configs {
		ParseSSR(value)
	}
}

/**
configString ==> ssr://cmM0LW1kNTowMTIzNDU2NzhAY25oay0xNC5meXZ2di5jb206MzEyNzk=
*/
func ParseSSR(encodeData string) {
	// ssr 协议解析
	configEncodeArray := strings.Split(string(encodeData), "ssr://")
	if len(configEncodeArray) == 2 {
		configString := strings.Replace(string(configEncodeArray[1]), "_", "/", -1)
		configString = strings.Replace(configString, "-", "+", -1)
		decodeConfigString, err := base64.RawStdEncoding.DecodeString(configString)
		if err != nil {
			fmt.Println(Red("解析ssr协议失败"))
		} else {
			WriteSSRConfigFile(string(decodeConfigString))
		}
	}
}

func WriteSSRConfigFile(configString string) {
	configArray := strings.Split(configString, "/?")
	// 必选参数
	requireParams := strings.Split(configArray[0], ":")
	if len(requireParams) < 6 {
		fmt.Println(Red("ssr协议不正确"))
		return
	}
	password, err := base64.RawStdEncoding.DecodeString(requireParams[5])
	if err != nil {
		fmt.Println("decode password failed")
		return
	}
	// 可选参数
	//_ := strings.Split(configArray[1], "&")
	// 拼凑配置
	config :=
		`{
    		"server": "` + requireParams[0] + `",
    		"server_ipv6": "::",
    		"server_port": ` + requireParams[1] + `,
    		"local_address": "0.0.0.0",
    		"local_port": 1080,

    		"password": "` + string(password) + `",
    		"method": "` + requireParams[3] + `",
    		"protocol": "auth_sha1_v4",
    		"protocol_param": "",
    		"obfs": "tls1.2_ticket_auth",
    		"obfs_param": "",
    		"speed_limit_per_con": 0,
    		"speed_limit_per_user": 0,

			"additional_ports" : {}, 
    		"additional_ports_only" : false, 
    		"timeout": 120,
    		"udp_timeout": 60,
    		"dns_ipv6": false,
    		"connect_verbose_info": 0,
    		"redirect": "",
    		"fast_open": false
		}`

	domainArray := strings.Split(requireParams[0], ".")
	err = ioutil.WriteFile(installPath+"/conf/"+domainArray[0]+".json", []byte(config), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func EditConfigFileCustom(name string) {
	if !strings.Contains(name, ".json") {
		name += ".json"
	}
	cmd := exec.Command("sudo", "vim", installPath+"/conf/"+name)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(Red(err.Error()))
	}
}

func GetConfigs() map[string]string {
	files, err := ioutil.ReadDir(installPath + "/conf/")
	if err != nil {
		log.Fatal(err)
	}
	servers := make(map[string]string)
	for _, file := range files {
		jsonFile, err := os.Open(installPath + "/conf/" + file.Name())
		defer jsonFile.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		var config Config
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &config)
		if config.Server != "" {
			key := strings.Replace(file.Name(), ".json", "", -1)
			servers[key] = config.Server
		}
	}
	return servers

}
func PingConfigs() {
	services := GetConfigs()
	for key, server := range services {
		pinger, err := ping.NewPinger(server)
		if err != nil {
			panic(err)
		}
		pinger.SetPrivileged(true)
		pinger.Count = 1
		pinger.Run()                 // blocks until finished
		stats := pinger.Statistics() // get send/receive/rtt stats
		fmt.Println(key + " => " + stats.AvgRtt.String())
	}
}
