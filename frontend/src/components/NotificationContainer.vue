<template>
  <div class="notification-container">
    <div 
      v-for="notification in notifications" 
      :key="notification.id"
      :class="['notification', notification.type]"
      @click="removeNotification(notification.id)"
    >
      <div class="notification-icon">
        <span v-if="notification.type === 'success'">✅</span>
        <span v-else-if="notification.type === 'error'">❌</span>
        <span v-else-if="notification.type === 'warning'">⚠️</span>
        <span v-else>ℹ️</span>
      </div>
      <div class="notification-content">
        <div class="notification-title">{{ notification.title }}</div>
        <div class="notification-message">{{ notification.message }}</div>
      </div>
      <button class="notification-close" @click.stop="removeNotification(notification.id)">
        ×
      </button>
    </div>
  </div>
</template>

<script setup>
import { useNotifications } from '../composables/useNotifications'

const { notifications, removeNotification } = useNotifications()
</script>

<style scoped>
.notification-container {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 1000;
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 400px;
}

.notification {
  background: white;
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12);
  border-left: 4px solid;
  display: flex;
  align-items: flex-start;
  gap: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  animation: slideIn 0.3s ease-out;
}

.notification:hover {
  transform: translateX(-4px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
}

.notification.success {
  border-left-color: #10b981;
}

.notification.error {
  border-left-color: #ef4444;
}

.notification.warning {
  border-left-color: #f59e0b;
}

.notification.info {
  border-left-color: #3b82f6;
}

.notification-icon {
  font-size: 20px;
  flex-shrink: 0;
  margin-top: 2px;
}

.notification-content {
  flex: 1;
  min-width: 0;
}

.notification-title {
  font-weight: 600;
  font-size: 14px;
  color: #1f2937;
  margin-bottom: 4px;
  line-height: 1.2;
}

.notification-message {
  font-size: 13px;
  color: #6b7280;
  line-height: 1.4;
  word-break: break-word;
}

.notification-close {
  background: none;
  border: none;
  font-size: 18px;
  color: #9ca3af;
  cursor: pointer;
  padding: 0;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.notification-close:hover {
  background: #f3f4f6;
  color: #374151;
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .notification-container {
    left: 10px;
    right: 10px;
    top: 10px;
    max-width: none;
  }
}
</style>