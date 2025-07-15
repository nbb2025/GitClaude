# Wails 文档参考

## 项目介绍
Wails 是一个用 Go 和 Web 技术构建跨平台桌面应用的框架。

## 重要提醒
- 在GitBash环境中运行wails命令可能会有路径问题
- 建议使用Windows命令行或PowerShell来运行wails相关命令
- 开发流程：先运行 `wails dev` 生成绑定文件，然后前端才能正常构建

## 安装要求
- Go 1.21+
- NPM (Node 15+)

## 支持平台
- Windows 10/11 (AMD64/ARM64)
- macOS 10.15+ (AMD64)
- Linux (AMD64/ARM64)

## 安装步骤
1. 安装 Go 和 NPM
2. 运行 `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
3. 验证安装 `wails doctor`

## 主要 CLI 命令

### 开发模式
```bash
wails dev
```
- 构建并运行应用
- 绑定 Go 代码到前端
- 使用 Vite 监视和自动重建
- 在 http://localhost:34115 提供本地 web 服务器
- **重要**: 首次运行会生成 wailsjs 绑定文件

### 构建应用
```bash
wails build
```
- 编译生产就绪的二进制文件
- 支持跨平台构建

### 系统诊断
```bash
wails doctor
```
- 验证开发环境准备情况

### 生成项目
```bash
wails init
```
- 生成新项目
- 支持远程 GitHub 模板

### 其他命令
- `wails generate`: 创建项目模板
- `wails update`: 更新 Wails CLI 版本
- `wails version`: 显示当前 CLI 版本

## 构建选项
- `-debug`: 保留调试信息
- `-platform`: 指定构建目标平台
- `-obfuscated`: 混淆应用
- `-upx`: 压缩最终二进制文件

## 平台特定依赖
- macOS: 需要 Xcode 命令行工具
- Windows: 需要 WebView2 运行时
- Linux: 需要 gcc, GTK3, WebKit 库

## 可选工具
- UPX: 用于压缩
- NSIS: 用于 Windows 安装程序

## 故障排除
如果 `wails` 命令缺失，确保 `go/bin` 目录在系统 PATH 中。

## 项目结构
- `app.go`: 后端应用逻辑
- `main.go`: 应用入口点
- `frontend/`: 前端代码目录
- `wails.json`: 项目配置文件
- `wailsjs/`: 自动生成的 JS 绑定

## 开发流程
1. 修改后端代码 (app.go)
2. 运行 `wails dev` 生成绑定并启动开发服务器
3. 前端代码会自动重新加载
4. 使用 `wails build` 构建最终应用