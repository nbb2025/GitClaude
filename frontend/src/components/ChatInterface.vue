<template>
  <div class="step-content chat-step">
    <div class="chat-layout">
      <!-- 左侧: 文件浏览 -->
      <FileBrowser 
        :selected-repo="selectedRepo"
        @file-selected="onFileSelected"
        ref="fileBrowser"
      />

      <!-- 中间: 对话区域 -->
      <div class="chat-panel">
        <div class="panel-header">
          <h3>💬 Claude对话</h3>
          <div class="chat-status">{{ isProcessing ? '正在处理...' : '就绪' }}</div>
        </div>
        
        <div class="chat-messages" ref="chatMessages">
          <div v-if="chatHistory.length === 0" class="welcome-message">
            <div class="welcome-icon">🤖</div>
            <h4>欢迎使用GitClaude！</h4>
            <p>在下方输入您的需求，我将帮助您修改代码并自动保存版本。</p>
          </div>
          
          <div v-for="(message, index) in chatHistory" :key="index" class="message">
            <div :class="['message-bubble', message.type === '用户' ? 'user' : 'assistant']">
              <div class="message-header">
                <span class="message-author">{{ message.type }}</span>
                <span class="message-time">{{ message.timestamp }}</span>
              </div>
              <div class="message-content">{{ message.content }}</div>
            </div>
          </div>
        </div>
        
        <div class="chat-input">
          <div class="input-container">
            <textarea 
              v-model="currentPrompt" 
              @keydown="handleKeydown"
              placeholder="告诉我您想要修改的内容... (回车发送，Shift+回车换行)"
              rows="3"
              :disabled="isProcessing"
              ref="promptInput"
            ></textarea>
            <button 
              @click="sendToClaudeCode" 
              :disabled="isProcessing || !currentPrompt.trim()" 
              class="btn btn-primary"
            >
              {{ isProcessing ? '处理中...' : '发送' }}
            </button>
          </div>
        </div>
      </div>

      <!-- 右侧: 历史记录 -->
      <div class="side-panel">
        <div class="panel-header">
          <h3>📚 提交历史</h3>
          <button @click="refreshHistory" class="btn btn-small btn-outline-white">刷新</button>
        </div>

        <div class="history-section">
          <div class="history-list">
            <div v-if="commitHistory.length === 0" class="empty-history">
              <div class="empty-icon">📝</div>
              <p>暂无提交历史</p>
            </div>
            <div v-for="(commit, index) in commitHistory" :key="index" class="commit-item">
              <div class="commit-icon">📝</div>
              <div class="commit-content">{{ commit }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, ref, onMounted, onUnmounted, nextTick } from 'vue'
import FileBrowser from './FileBrowser.vue'
import { claudeAPI, gitAPI } from '../services/api'
import { useNotifications } from '../composables/useNotifications'

const props = defineProps({
  selectedRepo: {
    type: String,
    required: true
  }
})

const { showNotification } = useNotifications()

const fileBrowser = ref(null)
const chatMessages = ref(null)
const promptInput = ref(null)

const currentPrompt = ref('')
const isProcessing = ref(false)
const chatHistory = ref([])
const commitHistory = ref([])

let historyRefreshInterval = null

// 定期刷新历史记录
const startHistoryRefresh = () => {
  if (historyRefreshInterval) {
    clearInterval(historyRefreshInterval)
  }
  historyRefreshInterval = setInterval(() => {
    refreshHistory()
  }, 5000) // 每5秒刷新一次
}

const stopHistoryRefresh = () => {
  if (historyRefreshInterval) {
    clearInterval(historyRefreshInterval)
    historyRefreshInterval = null
  }
}

const refreshHistory = async () => {
  try {
    const history = await gitAPI.getHistory()
    commitHistory.value = history
  } catch (error) {
    console.error('获取提交历史失败:', error)
  }
}

const onFileSelected = (filename) => {
  // 文件选择回调，可以在这里处理文件选择逻辑
  console.log('文件已选择:', filename)
}

const handleKeydown = (event) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendToClaudeCode()
  }
}

const sendToClaudeCode = async () => {
  if (!currentPrompt.value.trim()) return
  
  const prompt = currentPrompt.value.trim()
  
  // 添加用户消息到历史
  chatHistory.value.push({
    type: '用户',
    content: prompt,
    timestamp: new Date().toLocaleTimeString()
  })
  
  isProcessing.value = true
  currentPrompt.value = ''
  
  try {
    const result = await claudeAPI.callClaudeCode(prompt)
    
    // 添加Claude响应到历史
    chatHistory.value.push({
      type: 'Claude',
      content: result,
      timestamp: new Date().toLocaleTimeString()
    })
    
    // 刷新文件列表和提交历史
    if (fileBrowser.value) {
      await fileBrowser.value.loadRepoFiles()
    }
    await refreshHistory()
    
    // 如果有选中的文件，重新加载内容
    if (fileBrowser.value && fileBrowser.value.selectedFile) {
      await fileBrowser.value.selectFile(fileBrowser.value.selectedFile)
    }
    
    showNotification('success', 'Claude响应完成', '代码修改已自动提交')
    
  } catch (error) {
    console.error('Claude Code执行失败:', error)
    chatHistory.value.push({
      type: '错误',
      content: '执行失败: ' + error,
      timestamp: new Date().toLocaleTimeString()
    })
    showNotification('error', 'Claude执行失败', error.toString())
  } finally {
    isProcessing.value = false
    
    // 滚动到底部
    nextTick(() => {
      scrollToBottom()
    })
  }
}

// 滚动到聊天底部
const scrollToBottom = () => {
  if (chatMessages.value) {
    chatMessages.value.scrollTop = chatMessages.value.scrollHeight
  }
}

onMounted(() => {
  startHistoryRefresh()
  refreshHistory()
})

onUnmounted(() => {
  stopHistoryRefresh()
})

// 暴露方法给父组件
defineExpose({
  refreshHistory
})
</script>

<style scoped>
.chat-step {
  height: calc(100vh - 200px);
}

.chat-layout {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  gap: 1.5rem;
  height: 100%;
}

.chat-panel, .side-panel {
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

.chat-status {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.25rem 0.75rem;
  border-radius: 15px;
  font-size: 0.8rem;
  color: white;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 1.5rem;
  background: white;
}

.welcome-message {
  text-align: center;
  padding: 3rem 1rem;
  color: #6b7280;
}

.welcome-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.welcome-message h4 {
  margin: 0 0 0.5rem 0;
  color: #374151;
}

.welcome-message p {
  color: #6b7280;
}

.message {
  margin-bottom: 1.5rem;
  display: flex;
  flex-direction: column;
}

.message-bubble {
  max-width: 85%;
  padding: 1rem;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  word-wrap: break-word;
  width: fit-content;
}

.message-bubble.user {
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
  align-self: flex-end;
  margin-left: auto;
}

.message-bubble.assistant {
  background: #f8fafc;
  color: #374151;
  border: 1px solid #e5e7eb;
  align-self: flex-start;
  margin-right: auto;
}

.message-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.8rem;
  opacity: 0.8;
}

.message-bubble.user .message-header {
  color: rgba(255, 255, 255, 0.8);
}

.message-bubble.assistant .message-header {
  color: #6b7280;
}

.message-content {
  white-space: pre-wrap;
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.9rem;
  line-height: 1.5;
  color: inherit;
}

.chat-input {
  padding: 1.5rem;
  border-top: 1px solid #e5e7eb;
  background: #f8fafc;
}

.input-container {
  display: flex;
  gap: 1rem;
  align-items: flex-end;
}

.input-container textarea {
  flex: 1;
  padding: 1rem;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  resize: vertical;
  font-family: inherit;
  font-size: 0.95rem;
  line-height: 1.5;
  transition: border-color 0.3s ease;
  background: white;
  color: #374151;
  min-height: 60px;
}

.input-container textarea:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-container textarea::placeholder {
  color: #9ca3af;
}

.history-section {
  height: 100%;
  display: flex;
  flex-direction: column;
  background: white;
}

.empty-history {
  text-align: center;
  padding: 3rem 1rem;
  color: #6b7280;
}

.empty-history .empty-icon {
  font-size: 2rem;
  margin-bottom: 1rem;
}

.empty-history p {
  color: #6b7280;
  margin: 0;
}

.history-list {
  flex: 1;
  overflow-y: auto;
  background: white;
}

.commit-item {
  padding: 1rem;
  border-bottom: 1px solid #f3f4f6;
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
  font-size: 0.85rem;
}

.commit-icon {
  font-size: 0.8rem;
  margin-top: 0.1rem;
  color: #667eea;
}

.commit-content {
  flex: 1;
  font-family: 'JetBrains Mono', monospace;
  color: #374151;
  line-height: 1.4;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  font-size: 0.95rem;
  font-weight: 500;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  text-decoration: none;
}

.btn-small {
  padding: 0.5rem 1rem;
  font-size: 0.85rem;
}

.btn-primary {
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.btn-outline-white {
  background: transparent;
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.5);
}

.btn-outline-white:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.2);
  border-color: white;
  transform: translateY(-2px);
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none !important;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .chat-layout {
    grid-template-columns: 1fr 2fr;
    grid-template-rows: auto 1fr;
  }
  
  .chat-layout .side-panel:first-child {
    grid-column: 1 / -1;
    max-height: 200px;
  }
  
  .chat-layout .side-panel:last-child {
    grid-column: 1 / -1;
    max-height: 200px;
  }
}

@media (max-width: 768px) {
  .chat-layout {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
}
</style>