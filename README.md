## ssr 使用
1. 把执行文件拷入系统路径 `mv ssr /usr/local/bin/`
2. 赋予执行权限 `chmod +x /usr/local/bin/ssr`
3. 初始化安装 `ssr install`
4. 添加配置 `ssr config add` 或者 `ssr config sub`
5. 查看现有的配置 `ssr config ls`
6. 启动服务器 `ssr start config_name`
7. 如果需要终端使用 http 代理，可以运行 `ssr hp install` 来安装 privoxy
