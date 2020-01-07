package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Config() {
	if len(os.Args) == 2 {
		showConfigHelp()
	} else if len(os.Args) > 2 {
		switch os.Args[2] {
		case "ls":
			RunCommand("sudo", "ls", installPath+"/conf")
		case "sub":
			if len(os.Args) > 3 {
				SubFromUrl(os.Args[3])
			} else {
				fmt.Println(Red("url is required"))
			}
		default:
			showConfigHelp()
		}
	} else {
		showConfigHelp()
	}
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

func ParseSSR(encodeData string) {
	// ssr 协议解析
	configEncodeArray := strings.Split(string(encodeData), "ssr://")
	if len(configEncodeArray) == 2 {
		configString := strings.Replace(string(configEncodeArray[1]), "_", "/", -1)
		configString = strings.Replace(configString, "-", "+", -1)
		decodeConfigString, err := base64.RawStdEncoding.DecodeString(configString)
		if err != nil {
			fmt.Println(err)
		} else {
			WriteSSRConfigFile(string(decodeConfigString))
		}
	}
}
/**
	configString ==> ssr://cmM0LW1kNTowMTIzNDU2NzhAY25oay0xNC5meXZ2di5jb206MzEyNzk=
 */
func WriteSSRConfigFile(configString string) {
	configArray := strings.Split(configString, "/?")
	// 必选参数
	requireParams := strings.Split(configArray[0], ":")
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
