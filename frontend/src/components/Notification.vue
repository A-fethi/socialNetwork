<template>
    <div v-if="visible" :class="['notification', type]" :style="notificationStyle">
      <span>{{ message }}</span>
    </div>
  </template>
  
  <script>
  export default {
    name: 'AppNotification',
    props: ['message', 'type', 'visible'],
    computed: {
      notificationStyle() {
        return {
          right: '20px',
          top: '20px',
          animation: this.visible ? 'fadeIn 0.5s ease' : 'fadeOut 0.5s ease',
        };
      }
    }
  };
  </script>
  
  <style scoped>
  .notification {
    position: fixed;
    right: 20px;
    top: 20px;
    max-width: 300px;
    padding: 15px 25px;
    border-radius: 8px;
    font-size: 16px;
    color: white;
    z-index: 9999;
    display: flex;
    justify-content: center;
    align-items: center;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  
  .notification.success {
    background-color: #4caf50;
  }
  
  .notification.error {
    background-color: #f44336;
  }
  
  .notification span {
    display: block;
    font-weight: bold;
  }
  
  /* Animations */
  @keyframes fadeIn {
    0% {
      opacity: 0;
      transform: translateY(-20px);
    }
    100% {
      opacity: 1;
      transform: translateY(0);
    }
  }
  
  @keyframes fadeOut {
    0% {
      opacity: 1;
      transform: translateY(0);
    }
    100% {
      opacity: 0;
      transform: translateY(-20px);
    }
  }
  </style>
  