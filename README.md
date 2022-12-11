# maven
[![编译状态](https://github.ruijc.com:20443/api/badges/dronestock/maven/status.svg)](https://github.ruijc.com:20443/dronestock/maven)
[![Golang质量](https://goreportcard.com/badge/github.com/dronestock/maven)](https://goreportcard.com/report/github.com/dronestock/maven)
![版本](https://img.shields.io/github/go-mod/go-version/dronestock/maven)
![仓库大小](https://img.shields.io/github/repo-size/dronestock/maven)
![最后提交](https://img.shields.io/github/last-commit/dronestock/maven)
![授权协议](https://img.shields.io/github/license/dronestock/maven)
![语言个数](https://img.shields.io/github/languages/count/dronestock/maven)
![最佳语言](https://img.shields.io/github/languages/top/dronestock/maven)
![星星个数](https://img.shields.io/github/stars/dronestock/maven?style=social)

Drone持续集成Maven插件，功能

- 测试
- 打包
- 发布

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
steps:
  - name: 发布到Maven仓库
  image: dronestock/maven
  settings:
    username: xxx
    password: xxx
    token: xxx
```


更多使用教程，请参考[使用文档](https://www.dronestock.tech/plugin/stock/maven)

## 交流

![微信群](https://www.dronestock.tech/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢

[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://www.jetbrains.com/?from=dronestock/maven)
