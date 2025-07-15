// 工具函数集合

/**
 * 从路径中提取仓库名
 * @param {string} path - 路径
 * @returns {string} 仓库名
 */
export const getRepoName = (path) => {
  if (!path) return ''
  return path.split('\\').pop() || path.split('/').pop()
}

/**
 * 从路径中提取文件名
 * @param {string} path - 路径
 * @returns {string} 文件名
 */
export const getFileName = (path) => {
  if (!path) return ''
  return path.split('\\').pop() || path.split('/').pop()
}

/**
 * 获取文件图标
 * @param {string} filename - 文件名
 * @returns {string} 图标字符
 */
export const getFileIcon = (filename) => {
  const ext = filename.split('.').pop()?.toLowerCase()
  const iconMap = {
    js: '🟨',
    ts: '🔷',
    vue: '💚',
    go: '🐹',
    py: '🐍',
    java: '☕',
    html: '🌐',
    css: '🎨',
    json: '📄',
    md: '📝',
    txt: '📄',
    yml: '⚙️',
    yaml: '⚙️',
    xml: '📄',
    sql: '🗃️'
  }
  return iconMap[ext] || '📄'
}

/**
 * 检查文件是否为文本文件
 * @param {string} filename - 文件名
 * @returns {boolean} 是否为文本文件
 */
export const isTextFile = (filename) => {
  const textExtensions = [
    'txt', 'md', 'markdown', 'js', 'ts', 'jsx', 'tsx', 'vue', 'html', 'css', 'scss', 'sass',
    'json', 'xml', 'yaml', 'yml', 'toml', 'ini', 'cfg', 'conf', 'config',
    'py', 'java', 'go', 'rs', 'c', 'cpp', 'h', 'hpp', 'cs', 'php', 'rb', 'pl', 'sh', 'bash',
    'sql', 'log', 'dockerfile', 'gitignore', 'gitattributes', 'editorconfig',
    'svg', 'csv', 'tsv', 'properties', 'env', 'example'
  ]
  
  const ext = filename.split('.').pop()?.toLowerCase()
  return ext && textExtensions.includes(ext)
}

/**
 * 截断字符串
 * @param {string} s - 字符串
 * @param {number} length - 最大长度
 * @returns {string} 截断后的字符串
 */
export const truncateString = (s, length) => {
  if (s.length <= length) {
    return s
  }
  return s.substring(0, length) + '...'
}

/**
 * 构建文件树
 * @param {Array} files - 文件列表
 * @param {Set} expandedFolders - 已展开的文件夹
 * @returns {Array} 文件树数组
 */
export const buildFileTree = (files, expandedFolders) => {
  const tree = []
  const folderMap = new Map()
  
  // 构建文件夹结构
  files.forEach(file => {
    const parts = file.split(/[\/\\]/)
    let currentPath = ''
    
    parts.forEach((part, index) => {
      const parentPath = currentPath
      currentPath = currentPath ? `${currentPath}/${part}` : part
      
      if (index === parts.length - 1) {
        // 这是文件
        tree.push({
          type: 'file',
          name: part,
          path: file,
          level: index,
          parentPath: parentPath
        })
      } else {
        // 这是文件夹
        if (!folderMap.has(currentPath)) {
          folderMap.set(currentPath, {
            type: 'folder',
            name: part,
            path: currentPath,
            level: index,
            expanded: expandedFolders.has(currentPath),
            parentPath: parentPath
          })
        }
      }
    })
  })
  
  // 合并文件夹和文件
  const allItems = [...Array.from(folderMap.values()), ...tree]
  
  // 按路径深度和名称排序，确保文件夹内容正确排列
  allItems.sort((a, b) => {
    // 首先按父路径排序
    if (a.parentPath !== b.parentPath) {
      return a.parentPath.localeCompare(b.parentPath)
    }
    // 同一父级下，文件夹优先
    if (a.type !== b.type) {
      return a.type === 'folder' ? -1 : 1
    }
    // 最后按名称排序
    return a.name.localeCompare(b.name)
  })
  
  // 构建层级结构的显示列表
  const visibleItems = []
  
  const addItemsRecursively = (parentPath, level) => {
    // 查找当前级别的所有项目
    const currentLevelItems = allItems.filter(item => 
      item.parentPath === parentPath && item.level === level
    )
    
    currentLevelItems.forEach(item => {
      visibleItems.push(item)
      
      // 如果是展开的文件夹，递归添加其子项
      if (item.type === 'folder' && item.expanded) {
        addItemsRecursively(item.path, level + 1)
      }
    })
  }
  
  // 从根级别开始构建
  addItemsRecursively('', 0)
  
  return visibleItems
}