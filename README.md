## ssr 使用
1. 构建代码 `go build .` 或者直接 [下载](https://dev.tencent.com/u/dongkaipo/p/ssr/git/releases) 编译好的 Ubuntu 版本
2. 把执行文件拷入系统路径 `mv ssr /usr/local/bin/`
3. 赋予执行权限 `chmod +x /usr/local/bin/ssr`
4. 初始化安装 `ssr install`
5. 添加配置 `ssr config add` 或者 `ssr config sub`
6. 查看现有的配置 `ssr config ls`
7. 启动服务器 `ssr start config_name`
8. 如果需要终端使用 http 代理，可以运行 `ssr hp install` 来安装 privoxy
