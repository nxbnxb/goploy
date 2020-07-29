<p align=center>
    <img src="./banner.png" alt="logo" title="logo" />
</p>

<p align="center">
  <a href="#">
      <img src="https://img.shields.io/badge/readme%20style-standard-brightgreen.svg">
  </a>
  <a href="#">
      <img src="https://img.shields.io/badge/give%20me-a%20star-green.svg">
    </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg">
  </a>
</p>

一个web部署系统工具!

支持各种代码发布、回滚，可以通过web来一键完成！

完整的安装指引，上手无难度！

[DEMO](http://49.234.66.193:3000) admin:admin!@# (不一定能打开，看心情)

[帮助文档](https://zhenorzz.gitee.io/goploy)

## 内容列表

- [背景](#背景)
- [安装](#安装)
- [使用](#使用)
- [原理图](#原理图)
- [反向代理](#反向代理)
- [后端开发说明](#后端开发说明)
- [前端开发说明](#前端开发说明)
- [相关仓库](#相关仓库)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)

## 背景
大部分流行的发布工具功能虽然强大，但是配置较为麻烦，上手难度大，中小型项目难以使用。

## 安装
1. 安装mysql
2. 下载最新的release

## 使用
1. 运行./goploy or goploy.exe or goploy.mac (运行不了可能需要 chmod a+x, 再不行就发issue或加群)
2. 根据安装指引填入相关的配置
3. web http://ip:port  (账号:密码 admin:admin!@#)

## 预览
![预览](./preview.png)

## 原理图
![原理图](./goploy.png)

## 反向代理
nginx反向代理记得升级ws协议
```nginx
proxy_set_header Upgrade         $http_upgrade;
proxy_set_header Connection      "upgrade";
```

## 后端开发说明
1. 安装go，必须是1.13以上
2. 项目使用go mod管理
3. 运行go run main.go
4. 或者使用gin（可以特更新代码，改变就生效）

## 前端开发说明
1. cd web
2. 修改.env.development
3. npm run dev

## 相关仓库

- [gin](https://github.com/codegangsta/gin) - GO hot reload。
- [element-ui](https://github.com/ElemeFE/element) - 前端UI样式。

## 如何贡献

非常欢迎你的加入! （QQ群903750786）

[提一个Issue](https://github.com/zhenorzz/goploy/issues/new) 

提交一个 Pull Request.

## 使用许可

[MIT](LICENSE) © zhenorzz
