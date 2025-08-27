<template>
  <div id="app">
    <router-view :showNotification="showNotification" />
    <AppNotification :message="notification.message" :type="notification.type" :visible="notification.visible" />
  </div>
</template>

<script>
import AppNotification from './components/Notification.vue';

export default {
  name: 'App',
  components: { AppNotification },
  data() {
    return {
      notification: {
        message: '',
        type: 'success',
        visible: false
      }
    };
  },
  methods: {
    showNotification(message, type = 'success') {
      this.notification.message = message;
      this.notification.type = type;
      this.notification.visible = true;
      setTimeout(() => {
        this.notification.visible = false;
      }, 4000);
    }
  }
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>