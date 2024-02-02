<div align="center">
<h1>PalWorld Server Status</h1>
<h3>幻兽帕鲁服务器状态监控</h3>
</div>

[![Go-Zero](https://img.shields.io/badge/Go--Zero-v1.6.0-brightgreen.svg)](https://go-zero.dev/)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/go-admin-team/go-admin)

[前端](https://github.com/KiClover/palworld-status-ui)

[资源监控(后端)](https://github.com/KiClover/palworld-status-rpc-monitor)

[服务器监控(后端)](https://github.com/KiClover/palworld-status-rpc-server)
## 简介
优雅的PalWorld服务器监控系统

## 特性
- 前后端分离，分布式部署
### 前端
- Vue3 + Vite
- TypeScript
- Arco Design
- Nuxt3
- PWA/移动端适配
- 风滚草源码
### 后端
- 基于Go-Zero微服务框架
- 低占用，高性能


## 开始使用
### 前端页面 StatusUI
> 不建议把前端和游戏服务器放在同一台部署

更新依赖
```bash
pnpm i
```
运行
```bash
pnpm dev
```
编译
```bash
pnpm build
```

### 服务器本体资源监控 MonitorStatus

>该服务请置于运行PalWorldServerServer(幻兽帕鲁服务器)的服务器上运行

更新依赖
```bash
go mod tidy
```
运行
```bash
go run palworldmonitor.go -f etc/palworldmonitor.yaml
```


