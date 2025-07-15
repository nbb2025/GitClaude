<template>
  <div class="side-panel">
    <div class="panel-header">
      <h3>📁 项目文件</h3>
      <div class="repo-indicator">{{ getRepoName(selectedRepo) }}</div>
    </div>
    
    <div class="file-browser">
      <div class="file-tree">
        <div class="tree-item" v-for="item in fileTree" :key="item.path">
          <div 
            v-if="item.type === 'folder'"
            @click="toggleFolder(item.path)"
            :class="['folder-item', { expanded: item.expanded }]"
            :style="{ paddingLeft: (item.level * 20 + 10) + 'px' }"
          >
            <span class="folder-icon">{{ item.expanded ? '📂' : '📁' }}</span>
            <span class="folder-name">{{ item.name }}</span>
          </div>
          <div 
            v-else
            @click="selectFile(item.path)"
            :class="['file-item', { active: selectedFile === item.path }]"
            :style="{ paddingLeft: (item.level * 20 + 10) + 'px' }"
          >
            <span class="file-icon">{{ getFileIcon(item.name) }}</span>
            <span class="file-name">{{ item.name }}</span>
          </div>
        </div>
      </div>
      
      <div class="file-content" v-if="selectedFile">
        <div class="file-header">
          <h4>{{ getFileName(selectedFile) }}</h4>
          <div class="file-type-indicator">
            <span v-if="isTextFile(selectedFile)" class="text-file">文本文件</span>
            <span v-else class="binary-file">二进制文件</span>
          </div>
        </div>
        <div v-if="isTextFile(selectedFile)" class="code-content">{{ fileContent }}</div>
        <div v-else class="non-text-content">
          <div class="non-text-icon">📄</div>
          <p>{{ fileContent }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, ref, watch } from 'vue'
import { repositoryAPI } from '../services/api'
import { useNotifications } from '../composables/useNotifications'
import { getRepoName, getFileName, getFileIcon, isTextFile, buildFileTree } from '../utils'

const props = defineProps({
  selectedRepo: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['file-selected'])

const { showNotification } = useNotifications()

const selectedFile = ref('')
const fileContent = ref('')
const repoFiles = ref([])
const fileTree = ref([])
const expandedFolders = ref(new Set())

const loadRepoFiles = async () => {
  try {
    const files = await repositoryAPI.getFiles()
    repoFiles.value = files
    fileTree.value = buildFileTree(files, expandedFolders.value)
  } catch (error) {
    console.error('加载文件列表失败:', error)
  }
}

const toggleFolder = (folderPath) => {
  if (expandedFolders.value.has(folderPath)) {
    expandedFolders.value.delete(folderPath)
  } else {
    expandedFolders.value.add(folderPath)
  }
  // 重新构建文件树
  fileTree.value = buildFileTree(repoFiles.value, expandedFolders.value)
}

const selectFile = async (filename) => {
  try {
    selectedFile.value = filename
    
    // 只为文本文件加载内容
    if (isTextFile(filename)) {
      const content = await repositoryAPI.readFile(filename)
      fileContent.value = content
    } else {
      fileContent.value = '此文件类型不支持预览'
    }
    
    emit('file-selected', filename)
  } catch (error) {
    console.error('读取文件失败:', error)
    showNotification('error', '读取文件失败', error.toString())
  }
}

// 监听仓库变化，重新加载文件
watch(() => props.selectedRepo, () => {
  if (props.selectedRepo) {
    loadRepoFiles()
  }
}, { immediate: true })

// 暴露方法给父组件
defineExpose({
  loadRepoFiles,
  selectFile,
  selectedFile,
  fileContent
})
</script>

<style scoped>
.side-panel {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.panel-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
}

.panel-header h3 {
  margin: 0;
  font-size: 1.25rem;
  color: white;
}

.repo-indicator {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.25rem 0.75rem;
  border-radius: 15px;
  font-size: 0.8rem;
  color: white;
}

.file-browser {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: white;
}

.file-tree {
  flex: 1;
  overflow-y: auto;
  border-bottom: 1px solid #e5e7eb;
  background: white;
  min-height: 0;
}

.tree-item {
  user-select: none;
}

.folder-item {
  padding: 0.5rem 0.75rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: background-color 0.2s ease;
  font-size: 0.9rem;
  color: #374151;
  border-bottom: 1px solid #f9fafb;
}

.folder-item:hover {
  background: #f3f4f6;
}

.folder-item.expanded {
  background: #f0f9ff;
  color: #0369a1;
}

.folder-icon {
  font-size: 0.8rem;
  width: 16px;
  display: flex;
  justify-content: center;
  flex-shrink: 0;
}

.folder-name {
  font-weight: 500;
  text-align: left;
  flex: 1;
}

.file-item {
  padding: 0.5rem 0.75rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: background-color 0.2s ease;
  font-size: 0.9rem;
  color: #374151;
  border-bottom: 1px solid #f9fafb;
}

.file-item:hover {
  background: #f3f4f6;
}

.file-item.active {
  background: #667eea;
  color: white;
}

.file-icon {
  font-size: 0.8rem;
  width: 16px;
  display: flex;
  justify-content: center;
  flex-shrink: 0;
}

.file-name {
  text-align: left;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.file-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background: white;
  min-height: 0;
}

.file-header {
  padding: 1rem;
  border-bottom: 1px solid #e5e7eb;
  background: #f8fafc;
  flex-shrink: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.file-header h4 {
  margin: 0;
  font-size: 0.95rem;
  color: #374151;
  text-align: left;
}

.file-type-indicator {
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-weight: 500;
}

.text-file {
  background: #dcfce7;
  color: #16a34a;
}

.binary-file {
  background: #fef3c7;
  color: #d97706;
}

.code-content {
  flex: 1;
  padding: 1rem;
  margin: 0;
  overflow: auto;
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.8rem;
  line-height: 1.4;
  background: #f8fafc;
  color: #374151;
  border: none;
  min-height: 0;
  text-align: left;
  white-space: pre-wrap;
}

.non-text-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  background: #f8fafc;
  color: #6b7280;
  text-align: center;
}

.non-text-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  opacity: 0.5;
}

.non-text-content p {
  margin: 0;
  font-size: 0.9rem;
}
</style>