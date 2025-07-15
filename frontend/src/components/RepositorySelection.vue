<template>
  <div class="step-content">
    <div class="step-header">
      <h2>🎯 选择代码仓库</h2>
      <p>选择您要进行AI代码管理的本地项目目录</p>
    </div>
    
    <div class="selection-card">
      <div v-if="!selectedRepo" class="empty-state">
        <div class="empty-icon">📁</div>
        <h3>选择项目目录</h3>
        <p>点击下方按钮选择您的代码项目目录</p>
        <button @click="selectRepository" class="btn btn-primary btn-large">
          <span class="btn-icon">📂</span>
          选择仓库目录
        </button>
      </div>
      
      <div v-else class="selected-state">
        <div class="success-icon">✅</div>
        <h3>已选择项目</h3>
        <div class="repo-path">{{ selectedRepo }}</div>
        <div class="action-buttons">
          <button @click="selectRepository" class="btn btn-success">
            <span class="btn-icon">🔄</span>
            重新选择
          </button>
          <button @click="handleNext" class="btn btn-primary">
            <span class="btn-icon">➡️</span>
            下一步
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'
import { repositoryAPI } from '../services/api'
import { useNotifications } from '../composables/useNotifications'

const props = defineProps({
  selectedRepo: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:selectedRepo', 'next'])

const { showNotification } = useNotifications()

const selectRepository = async () => {
  console.log('selectRepository 函数被调用')
  try {
    console.log('开始调用 SelectRepository...')
    const repo = await repositoryAPI.select()
    console.log('SelectRepository 返回结果:', repo)
    if (repo) {
      emit('update:selectedRepo', repo)
      console.log('设置 selectedRepo 为:', repo)
      showNotification('success', '仓库选择成功', `已选择: ${repo}`)
    } else {
      console.log('用户取消了选择或返回空值')
    }
  } catch (error) {
    console.error('选择仓库失败:', error)
    showNotification('error', '选择仓库失败', error.toString())
  }
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

.selection-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  padding: 3rem;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 600px;
  margin: 0 auto;
}

.empty-state, .selected-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}

.empty-icon, .success-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
}

.empty-state h3, .selected-state h3 {
  font-size: 1.5rem;
  margin: 0;
  color: #1f2937;
}

.empty-state p {
  color: #6b7280;
  margin: 0;
}

.repo-path {
  background: #f3f4f6;
  padding: 1rem;
  border-radius: 8px;
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.9rem;
  color: #374151;
  word-break: break-all;
  max-width: 100%;
}

.action-buttons {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
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

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none !important;
}

.btn-icon {
  font-size: 0.9rem;
}
</style>