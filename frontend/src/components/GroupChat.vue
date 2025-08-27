<template>
  <div class="group-chat">
    <div class="chat-header">
      <h2>{{ groupName }}</h2>
    </div>

    <div class="messages-container" ref="messagesContainer">
      <div v-for="message in messages" :key="message.id" class="message" :class="{
        'own-message': message.sender_id === currentUserId,
        'unread-message': !message.is_read && message.sender_id !== currentUserId
      }" @click="markMessageAsRead(message)">
        <div class="message-header">
          <span class="username">{{ message.username }}</span>
          <span class="timestamp">{{ formatTime(message.created_at) }}</span>
        </div>
        <div class="message-content">{{ message.content }}</div>
      </div>
    </div>

    <div class="message-input">
      <input type="text" v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type a message...">
      <button @click="sendMessage" :disabled="!newMessage.trim()">Send</button>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted, nextTick, watch } from 'vue';
import { groupWebSocket } from '../services/groupWebSocket';
import { fetchInfo } from '@/api/getInfo';

export default {

  name: 'GroupChat',
  props: {
    groupId: {
      type: String,
      required: true
    },
    groupName: {
      type: String,
      required: true
    }
  },
  beforeMount() {
    this.fetchInfo();
  },
  methods: {
    async fetchInfo() {
      const userData = await fetchInfo();
      this.currentUserId = userData.id;
    }
  },
  setup(props) {
    const messages = ref([]);
    const newMessage = ref('');
    const messagesContainer = ref(null);
    const currentUserId = ref(null);
    const isTabActive = ref(true);

    const formatTime = (timestamp) => {
      const date = new Date(timestamp);
      const now = new Date();
      const diffInSeconds = Math.floor((now - date) / 1000);

      if (diffInSeconds < 60) {
        return 'just now';
      }

      const diffInMinutes = Math.floor(diffInSeconds / 60);
      if (diffInMinutes < 60) {
        return `${diffInMinutes}m ago`;
      }

      const diffInHours = Math.floor(diffInMinutes / 60);
      if (diffInHours < 24) {
        return `${diffInHours}h ago`;
      }

      const diffInDays = Math.floor(diffInHours / 24);
      if (diffInDays < 7) {
        return `${diffInDays}d ago`;
      }

      return date.toLocaleDateString();
    };

    const scrollToBottom = async () => {
      await nextTick();
      if (messagesContainer.value) {
        messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
      }
    };

    const handleInitialMessages = (initialMessages) => {
      if (!isTabActive.value) return;
      messages.value = initialMessages.reverse();
      scrollToBottom();
    };

    const handleNewMessage = (message) => {
      if (!isTabActive.value) return;
      messages.value.push(message);
      scrollToBottom();
    };

    const sendMessage = () => {
      if (newMessage.value.trim()) {
        groupWebSocket.sendMessage(newMessage.value.trim());
        newMessage.value = '';
      }
    };

    const fetchUserInfo = async () => {
      try {
        const userData = await fetchInfo();
        currentUserId.value = userData.id;
      } catch (error) {
        console.error('Error fetching user info:', error);
      }
    };

    const handleVisibilityChange = () => {
      isTabActive.value = document.visibilityState === 'visible';
      console.log('Tab visibility changed:', isTabActive.value);

      if (isTabActive.value) {

        connectWebSocket();
      } else {

        console.log('Tab hidden, disconnecting from group:', props.groupId);
        groupWebSocket.disconnect();
      }
    };

    const connectWebSocket = () => {

      messages.value = [];


      groupWebSocket.disconnect();


      console.log('Connecting to group:', props.groupId);
      groupWebSocket.connect(props.groupId);
      groupWebSocket.onInitialMessages(handleInitialMessages);
      groupWebSocket.onMessage(handleNewMessage);

    };


    watch(() => props.groupId, (newGroupId, oldGroupId) => {
      if (newGroupId && newGroupId !== oldGroupId) {
        console.log('Group ID changed from', oldGroupId, 'to', newGroupId);
        connectWebSocket();
      }
    });

    onMounted(async () => {
      await fetchUserInfo();
      connectWebSocket();


      document.addEventListener('visibilitychange', handleVisibilityChange);
    });

    onUnmounted(() => {
      console.log('Disconnecting from group:', props.groupId);
      groupWebSocket.disconnect();
      groupWebSocket.removeInitialHandler(handleInitialMessages);
      groupWebSocket.removeMessageHandler(handleNewMessage);


      document.removeEventListener('visibilitychange', handleVisibilityChange);
    });

    return {
      messages,
      newMessage,
      messagesContainer,
      currentUserId,
      formatTime,
      sendMessage
    };
  }
};
</script>

<style scoped>
.group-chat {
  display: flex;
  flex-direction: column;
  height: 600px;
  background-color: #f5f5f5;
  border-radius: 8px;
  overflow: hidden;
}

.chat-header {
  padding: 1rem;
  background-color: #4a90e2;
  color: white;
  flex-shrink: 0;
}

.chat-header h2 {
  margin: 0;
  font-size: 1.2rem;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  min-height: 0;
}

.message {
  max-width: 70%;
  padding: 0.8rem;
  border-radius: 8px;
  background-color: white;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  word-wrap: break-word;
}

.own-message {
  align-self: flex-end;
  background-color: #4a90e2;
  color: white;
}

.message-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.9rem;
}

.username {
  font-weight: bold;
}

.timestamp {
  color: #666;
  font-size: 0.8rem;
}

.own-message .timestamp {
  color: #e0e0e0;
}

.message-content {
  word-break: break-word;
  white-space: pre-wrap;
}

.message-input {
  display: flex;
  gap: 0.5rem;
  padding: 1rem;
  background-color: white;
  border-top: 1px solid #e0e0e0;
  flex-shrink: 0;
}

.message-input input {
  flex: 1;
  padding: 0.8rem;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
  font-size: 1rem;
}

.message-input button {
  padding: 0.8rem 1.5rem;
  background-color: #4a90e2;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.2s;
}

.message-input button:hover {
  background-color: #357abd;
}

.message-input button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

/* Custom scrollbar styles */
.messages-container::-webkit-scrollbar {
  width: 6px;
}

.messages-container::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.messages-container::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 3px;
}

.messages-container::-webkit-scrollbar-thumb:hover {
  background: #555;
}

.unread-message {
  background-color: #e3f2fd;
  border-left: 4px solid #2196f3;
  cursor: pointer;
}

.unread-message:hover {
  background-color: #bbdefb;
}
</style>