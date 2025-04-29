<template>
    <div>Loading...</div>
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
  