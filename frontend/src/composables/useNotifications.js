// 通知系统 Composable
import { ref } from 'vue'

// 全局通知状态
const notifications = ref([])
let notificationId = 0

/**
 * 通知系统 Hook
 */
export function useNotifications() {
  /**
   * 显示通知
   * @param {string} type - 通知类型: 'success', 'error', 'warning', 'info'
   * @param {string} title - 通知标题
   * @param {string} message - 通知消息
   * @param {number} duration - 显示时长 (毫秒)
   */
  const showNotification = (type, title, message, duration = 4000) => {
    const id = ++notificationId
    const notification = {
      id,
      type,
      title,
      message
    }
    
    notifications.value.push(notification)
    
    // 自动移除通知
    setTimeout(() => {
      removeNotification(id)
    }, duration)
  }

  /**
   * 移除通知
   * @param {number} id - 通知ID
   */
  const removeNotification = (id) => {
    const index = notifications.value.findIndex(n => n.id === id)
    if (index > -1) {
      notifications.value.splice(index, 1)
    }
  }

  /**
   * 清空所有通知
   */
  const clearNotifications = () => {
    notifications.value = []
  }

  return {
    notifications,
    showNotification,
    removeNotification,
    clearNotifications
  }
}