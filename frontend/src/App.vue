<template>
  <div id="app">
    <!-- 顶部导航 -->
    <nav class="navbar">
      <div class="nav-brand">
        <h1>🔥 GitClaude</h1>
        <span class="nav-subtitle">版本管理驱动的Claude Code</span>
      </div>
      <div class="nav-steps">
        <div 
          v-for="(step, index) in steps" 
          :key="index"
          :class="['step-indicator', { active: currentStep === index, completed: index < currentStep }]"
          @click="goToStep(index)"
        >
          <div class="step-number">{{ index + 1 }}</div>
          <div class="step-label">{{ step.label }}</div>
        </div>
      </div>
    </nav>

    <!-- 主内容区域 -->
    <main class="main-content">
      <!-- 步骤1: 仓库选择 -->
      <RepositorySelection 
        v-if="currentStep === 0"
        v-model:selected-repo="selectedRepo"
        @next="nextStep"
      />

      <!-- 步骤2: 初始化Git仓库 -->
      <GitInitialization 
        v-if="currentStep === 1"
        :selected-repo="selectedRepo"
        v-model:git-config-exists="gitConfigExists"
        v-model:git-initialized="gitInitialized"
        @previous="prevStep"
        @next="nextStep"
      />

      <!-- 步骤3: AI对话界面 -->
      <ChatInterface 
        v-if="currentStep === 2"
        :selected-repo="selectedRepo"
        ref="chatInterface"
      />
    </main>
    
    <!-- 通知组件 -->
    <NotificationContainer />
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import RepositorySelection from './components/RepositorySelection.vue'
import GitInitialization from './components/GitInitialization.vue'
import ChatInterface from './components/ChatInterface.vue'
import NotificationContainer from './components/NotificationContainer.vue'
import { gitAPI } from './services/api'

// 步骤定义
const steps = [
  { label: '选择仓库' },
  { label: '初始化Git' },
  { label: 'AI对话' }
]

// 状态管理
const currentStep = ref(0)
const selectedRepo = ref('')
const gitInitialized = ref(false)
const gitConfigExists = ref(false)
const chatInterface = ref(null)

// 检查Git配置是否存在
const checkGitConfig = async () => {
  try {
    if (selectedRepo.value) {
      const exists = await gitAPI.checkExists()
      gitConfigExists.value = exists
      if (exists) {
        gitInitialized.value = true
      }
    }
  } catch (error) {
    console.error('检查Git配置失败:', error)
    gitConfigExists.value = false
  }
}

// 步骤导航
const nextStep = () => {
  if (currentStep.value < steps.length - 1) {
    currentStep.value++
    if (currentStep.value === 2 && chatInterface.value) {
      chatInterface.value.refreshHistory()
  }
}
}

const prevStep = () => {
  if (currentStep.value > 0) {
    currentStep.value--
  }
}

const goToStep = (stepIndex) => {
  if (stepIndex === 0 || 
      (stepIndex === 1 && selectedRepo.value) || 
      (stepIndex === 2 && (gitInitialized.value || gitConfigExists.value))) {
    currentStep.value = stepIndex
    if (stepIndex === 2 && chatInterface.value) {
      chatInterface.value.refreshHistory()
    }
  }
}

// 监听仓库选择变化
watch(selectedRepo, async (newRepo) => {
  if (newRepo) {
    await checkGitConfig()
  }
})

onMounted(() => {
  // 初始化逻辑
})
</script>

<style scoped>
* {
  box-sizing: border-box;
}

#app {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  margin: 0;
  padding: 0;
}

/* 顶部导航 */
.navbar {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
}

.nav-brand h1 {
  margin: 0;
  font-size: 1.8rem;
  font-weight: 700;
  background: linear-gradient(45deg, #667eea, #764ba2);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.nav-subtitle {
  font-size: 0.9rem;
  color: #666;
  margin-left: 0.5rem;
}

.nav-steps {
  display: flex;
  gap: 1rem;
}

.step-indicator {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  border-radius: 25px;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #e5e7eb; /* 默认灰色背景 */
}

.step-indicator.active {
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
  transform: scale(1.05);
}

.step-indicator.completed {
  background: #10b981;
  color: white;
}

.step-number {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.8rem;
  font-weight: bold;
}

.step-label {
  font-size: 0.9rem;
  font-weight: 500;
}

/* 主内容 */
.main-content {
  flex: 1;
  padding: 2rem;
  overflow-y: auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .navbar {
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
  }
  
  .nav-steps {
    flex-wrap: wrap;
    justify-content: center;
  }
  
  .main-content {
    padding: 1rem;
  }
}
</style>