<template>
  <div class="step-content">
    <div class="step-header">
      <h2>⚙️ 初始化本地Git仓库</h2>
      <p>为您的项目创建专用的版本控制系统</p>
    </div>
    
    <div class="init-card">
      <div class="info-section">
        <div class="info-item">
          <strong>项目路径:</strong> {{ selectedRepo }}
        </div>
        <div class="info-item">
          <strong>Git仓库位置:</strong> {{ selectedRepo }}\.gitclaude
        </div>
        <div class="info-item">
          <strong>说明:</strong> 将在项目目录下创建 .gitclaude 文件夹管理版本历史
        </div>
        <div v-if="gitConfigExists" class="info-item status-item">
          <strong>状态:</strong> 
          <span class="status-existing">检测到现有Git仓库配置</span>
        </div>
      </div>
      
      <!-- 主要操作区域：居中布局 -->
      <div class="main-actions">
        <button @click="handlePrevious" class="btn btn-primary main-action-btn">
          <span class="btn-icon">⬅️</span>
          上一步
        </button>
        <div v-if="!gitConfigExists">
          <button @click="initializeGitRepo" class="btn btn-primary main-action-btn">
            <span class="btn-icon">🚀</span>
            初始化Git仓库
          </button>
        </div>
        <div v-else>
          <button @click="handleNext" class="btn btn-success main-action-btn">
            <span class="btn-icon">➡️</span>
            使用现有仓库
          </button>
        </div>
      </div>
      
      <!-- 重新初始化按钮（居中，较小，警告色） -->
      <div v-if="gitConfigExists" class="reinit-section">
        <button @click="reinitializeGitRepo" class="btn btn-warning">
          <span class="btn-icon">🔄</span>
          重新初始化仓库
        </button>
      </div>
      
      <!-- 初始化成功后的继续按钮（居中，较小，主色） -->
      <div v-if="gitInitialized && !gitConfigExists" class="continue-section">
        <button @click="handleNext" class="btn btn-primary">
          <span class="btn-icon">➡️</span>
          开始AI对话
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'
import { gitAPI } from '../services/api'
import { useNotifications } from '../composables/useNotifications'

const props = defineProps({
  selectedRepo: {
    type: String,
    required: true
  },
  gitConfigExists: {
    type: Boolean,
    default: false
  },
  gitInitialized: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:gitConfigExists', 'update:gitInitialized', 'previous', 'next'])

const { showNotification } = useNotifications()

const initializeGitRepo = async () => {
  try {
    await gitAPI.initialize()
    emit('update:gitInitialized', true)
    emit('update:gitConfigExists', true)
    showNotification('success', 'Git仓库初始化成功', '本地版本控制已启用')
  } catch (error) {
    console.error('初始化Git仓库失败:', error)
    showNotification('error', '初始化Git仓库失败', error.toString())
  }
}

const reinitializeGitRepo = async () => {
  try {
    await gitAPI.initialize()
    emit('update:gitInitialized', true)
    emit('update:gitConfigExists', true)
    showNotification('success', 'Git仓库重新初始化成功', '本地版本控制已重置')
  } catch (error) {
    console.error('重新初始化Git仓库失败:', error)
    showNotification('error', '重新初始化Git仓库失败', error.toString())
  }
}

const handlePrevious = () => {
  emit('previous')
}

const handleNext = () => {
  emit('next')
}
</script>

<style scoped>
.step-content {
  max-width: 1200px;
  margin: 0 auto;
}

.step-header {
  text-align: center;
  margin-bottom: 2rem;
  color: white;
}

.step-header h2 {
  font-size: 2.5rem;
  margin: 0 0 0.5rem 0;
  font-weight: 700;
}

.step-header p {
  font-size: 1.1rem;
  opacity: 0.9;
  margin: 0;
}

.init-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  padding: 3rem;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 600px;
  margin: 0 auto;
  align-items: center;
}

.info-section {
  background: #f8fafc;
  border-radius: 12px;
  padding: 1.5rem;
  margin-bottom: 2rem;
  text-align: left;
}

.info-item {
  margin-bottom: 1rem;
  font-size: 0.95rem;
  line-height: 1.6;
  color: #374151;
}

.info-item:last-child {
  margin-bottom: 0;
}

.info-item strong {
  color: #1f2937;
}

.status-existing {
  color: #10b981;
  font-weight: 500;
}

.main-actions {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  width: 100%;
  margin-bottom: 2rem;
}

.main-action-btn {
  min-width: 180px;
  justify-content: center;
}

.reinit-section {
  text-align: center;
  margin-top: 1rem;
}

.continue-section {
  text-align: center;
  margin-top: 2rem;
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

.btn-large {
  padding: 1rem 2rem;
  font-size: 1.1rem;
}

.btn-primary {
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.btn-success {
  background: linear-gradient(45deg, #10b981, #059669);
  color: white;
}

.btn-success:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(16, 185, 129, 0.4);
}

.btn-warning {
  background: linear-gradient(45deg, #f59e0b, #ef4444);
  color: white;
  border: none;
}

.btn-warning:hover:not(:disabled) {
  background: linear-gradient(45deg, #ef4444, #f59e0b);
  color: white;
  box-shadow: 0 8px 25px rgba(245, 158, 11, 0.2);
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none !important;
}

.btn-icon {
  font-size: 0.9rem;
}
</style>