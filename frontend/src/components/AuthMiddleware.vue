<template>
    <div class="loading">Loading...</div>
  </template>
  
  <script>
  
  export default {
    name: 'AuthMiddleware',  
    data() {
      return {
        isLoggedIn: false,  
      };
    },
    methods: {
      async checkLogin() {
        try {
          const response = await fetch('http://localhost:8080/middle', {
            method: 'GET',
            credentials: 'include',
          });
  
          const data = await response.json();
  
          if (data.message === 'Login successful') {
            this.isLoggedIn = true;
            this.$router.push('/home'); 
          } else {
            this.isLoggedIn = false;
            this.$router.push('/login');
          }
        } catch (error) {
          console.error('Error during login check:', error);
          this.$router.push('/login'); 
        }
      }
    },
    mounted() {
      this.checkLogin();
    }
  };
  </script>

<style scoped>
.loading {
  display: flex;
  flex-direction: column;
  gap: 10px;
  font-size: 24px;
  color: #333;
  font-weight: bold;
  text-align: center;
  background-color: #f9f9f9;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.loading::after {
  content: "";
  width: 48px;
  height: 48px;
  border: 6px solid #ccc;
  border-top: 6px solid #333;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  display: block;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>
  