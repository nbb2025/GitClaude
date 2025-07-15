# 🔥 GitClaude

> **AI驱动的本地Git仓库管理工具**

GitClaude是一个基于Wails框架构建的跨平台桌面应用程序，它将Claude AI与本地Git版本控制系统完美结合，为开发者提供智能化的代码管理和版本控制体验。

![GitClaude Banner](https://img.shields.io/badge/Wails-v2-blue?style=for-the-badge&logo=go)
![Vue3](https://img.shields.io/badge/Vue-3-green?style=for-the-badge&logo=vue.js)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)

## ✨ 核心特性

### 🎯 智能工作流程
- **三步式操作**: 仓库选择 → Git初始化 → AI对话
- **一键式体验**: 简化的用户界面，专注于核心功能
- **自动化管理**: 每次AI对话后自动提交代码更改

### 🤖 Claude AI集成
- **代码修改**: 通过自然语言描述修改代码
- **智能理解**: Claude AI理解项目结构和代码逻辑
- **版本追踪**: 每次修改都有完整的提交历史

### 🗂️ 文件管理
- **树形结构**: 直观的文件夹树形展示
- **智能预览**: 自动识别并预览文本文件
- **实时更新**: 文件更改后自动刷新显示

### 🔔 现代化UI
- **响应式设计**: 适配不同屏幕尺寸
- **动画过渡**: 流畅的用户体验
- **通知系统**: 实时状态反馈

## 🚀 快速开始

### 环境要求

- **Go**: 1.21+
- **Node.js**: 15+
- **NPM**: 包管理器
- **Git**: 版本控制系统

### 安装步骤

1. **克隆仓库**
```bash
git clone https://github.com/nbb2025/GitClaude.git
cd GitClaude
```

2. **安装Wails CLI**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

3. **验证环境**
```bash
wails doctor
```

4. **安装依赖**
```bash
cd frontend
npm install
cd ..
```

5. **开发模式运行**
```bash
wails dev
```

6. **构建生产版本**
```bash
wails build
```

## 📁 项目结构

```
GitClaude/
├── frontend/                   # 前端Vue3应用
│   ├── src/
│   │   ├── components/        # Vue组件
│   │   │   ├── RepositorySelection.vue
│   │   │   ├── GitInitialization.vue
│   │   │   ├── ChatInterface.vue
│   │   │   ├── FileBrowser.vue
│   │   │   └── NotificationContainer.vue
│   │   ├── composables/       # 组合式API
│   │   │   └── useNotifications.js
│   │   ├── services/          # API服务层
│   │   │   └── api.js
│   │   ├── utils/             # 工具函数
│   │   │   └── index.js
│   │   └── App.vue            # 主应用组件
│   ├── package.json
│   └── vite.config.js
├── app.go                      # Go后端应用逻辑
├── main.go                     # 应用入口点
├── wails.json                  # Wails配置文件
├── go.mod                      # Go模块文件
└── README.md                   # 项目文档
```

## 🎮 使用指南

### 1. 选择仓库
- 点击"选择仓库目录"按钮
- 选择你要管理的代码项目目录
- 系统会自动检测项目结构

### 2. 初始化Git
- 应用会在项目目录下创建`.gitclaude`文件夹
- 这是独立的本地Git仓库，不会影响原项目的Git历史
- 支持使用现有配置或重新初始化

### 3. AI对话
- 在对话框中用自然语言描述你的需求
- Claude AI会理解并修改代码
- 每次修改都会自动创建Git提交
- 查看右侧的提交历史记录

## 🔧 配置说明

### Wails配置 (wails.json)
```json
{
  "name": "GitClaude",
  "outputfilename": "GitClaude",
  "frontend:install": "npm install",
  "frontend:build": "npm run build",
  "frontend:dev:watcher": "npm run dev",
  "frontend:dev:serverUrl": "auto"
}
```

### 开发环境
- 在GitBash环境中运行wails命令可能会有路径问题
- 建议使用Windows命令行或PowerShell
- 开发流程：先运行 `wails dev` 生成绑定文件，然后前端才能正常构建

## 🌟 技术架构

### 前端技术栈
- **Vue 3**: 渐进式JavaScript框架
- **Composition API**: 现代化的组件逻辑组织
- **Vite**: 快速的前端构建工具
- **CSS3**: 现代化样式和动画

### 后端技术栈
- **Go**: 高性能的系统编程语言
- **Wails v2**: 现代化的跨平台桌面应用框架
- **Git**: 版本控制系统集成

### 核心功能实现
```go
// 主要API接口
func (a *App) SelectRepository() (string, error)
func (a *App) InitializeLocalGitRepo() error
func (a *App) CallClaudeCode(prompt string) (string, error)
func (a *App) GetRepoFiles() ([]string, error)
func (a *App) GetCommitHistory() ([]string, error)
```

## 🔒 安全性

- **本地存储**: 所有数据都存储在本地，不会上传到云端
- **独立仓库**: 使用单独的`.gitclaude`目录，不影响原项目
- **访问控制**: 只能访问用户明确选择的目录

## 🚧 开发计划

### 即将推出
- [ ] 多语言支持
- [ ] 主题定制
- [ ] 插件系统
- [ ] 命令行工具
- [ ] 云同步功能

### 长期规划
- [ ] 团队协作功能
- [ ] 代码审查集成
- [ ] 持续集成支持
- [ ] 性能优化

## 🤝 贡献指南

我们欢迎各种形式的贡献！

1. **Fork** 本仓库
2. **创建** 功能分支 (`git checkout -b feature/AmazingFeature`)
3. **提交** 更改 (`git commit -m 'Add some AmazingFeature'`)
4. **推送** 到分支 (`git push origin feature/AmazingFeature`)
5. **开启** Pull Request

## 📄 许可证

本项目采用 MIT 许可证。查看 [LICENSE](LICENSE) 文件了解更多信息。

## 🙏 致谢

- [Wails](https://wails.io/) - 优秀的跨平台桌面应用框架
- [Vue.js](https://vuejs.org/) - 渐进式JavaScript框架
- [Anthropic Claude](https://claude.ai/) - 强大的AI助手

---

⭐ 如果这个项目对你有帮助，请给我们一个星标！

📚 更多文档和教程即将发布，请关注项目更新。

🔥 **GitClaude - 让AI为你的代码管理保驾护航！**