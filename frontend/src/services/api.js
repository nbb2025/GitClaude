// API 服务层 - 封装所有 Wails API 调用
import {
  SelectRepository,
  InitializeLocalGitRepo,
  CallClaudeCode,
  GetRepoFiles,
  ReadFile,
  WriteFile,
  GetCommitHistory,
  CheckGitRepoExists,
  CommitChanges,
  GetSelectedRepo
} from '../../wailsjs/go/main/App'

/**
 * 仓库相关 API
 */
export const repositoryAPI = {
  /**
   * 选择仓库
   */
  async select() {
    return await SelectRepository()
  },

  /**
   * 获取当前选中的仓库
   */
  async getSelected() {
    return await GetSelectedRepo()
  },

  /**
   * 获取仓库文件列表
   */
  async getFiles() {
    return await GetRepoFiles()
  },

  /**
   * 读取文件内容
   */
  async readFile(filename) {
    return await ReadFile(filename)
  },

  /**
   * 写入文件内容
   */
  async writeFile(filename, content) {
    return await WriteFile(filename, content)
  }
}

/**
 * Git 相关 API
 */
export const gitAPI = {
  /**
   * 初始化本地 Git 仓库
   */
  async initialize() {
    return await InitializeLocalGitRepo()
  },

  /**
   * 检查 Git 仓库是否存在
   */
  async checkExists() {
    return await CheckGitRepoExists()
  },

  /**
   * 提交更改
   */
  async commit(message) {
    return await CommitChanges(message)
  },

  /**
   * 获取提交历史
   */
  async getHistory() {
    return await GetCommitHistory()
  }
}

/**
 * Claude 相关 API
 */
export const claudeAPI = {
  /**
   * 调用 Claude Code
   */
  async callClaudeCode(prompt) {
    return await CallClaudeCode(prompt)
  }
}