<!-- src/components/UserLogin.vue -->
<template>
  <div class="login-container">
    <div class="login-card">
      <h1>Login</h1>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Username</label>
          <input type="text" v-model="username" id="username" required />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input type="password" v-model="password" id="password" required />
        </div>
        <button type="submit" class="btn">Login</button>
      </form>
      <p class="signup-link"> Don't have an account? <a href="/register">Sign up</a>
      </p>
    </div>
  </div>
</template>

<script>
import router from '@/router';



export default {
  props: ['showNotification'],
  name: 'UserLogin',
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
   async handleLogin() {
  
      fetch('http://localhost:8080/api/login', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json' 
  },
  body: JSON.stringify({
    username: this.username,
    password: this.password
  }),
  credentials: 'include'
})
  .then(response => {
    if (response.ok) {
      return response.json();
    }
    throw new Error('Login failed');
  })
  .then(data => {
   router.push('/home'); 
    console.log('Success:', data);
  })
  .catch(error => {
    this.showNotification('Login failed. Please try again.', 'error');
    console.error('Error:', error);
  });

}


  }
};
</script>

<style scoped>
/* General styles for the login page */
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f4f7fa;
  font-family: 'Arial', sans-serif;
}

.login-card {
  background-color: #ffffff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 40px;
  border-radius: 8px;
  width: 100%;
  max-width: 400px;
  text-align: center;
}

h1 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}

.form-group {
  margin-bottom: 20px;
  text-align: left;
}

label {
  display: block;
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border-radius: 4px;
  border: 1px solid #ddd;
  background-color: #f9f9f9;
  box-sizing: border-box;
}

input:focus {
  outline: none;
  border-color: #6c63ff;
  background-color: #fff;
}

button {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  background-color: #6c63ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #5a52e1;
}

.forgot-password {
  margin-top: 20px;
  font-size: 14px;
}

.forgot-password a {
  color: #6c63ff;
  text-decoration: none;
}

.forgot-password a:hover {
  text-decoration: underline;
}

.signup-link {
  margin-top: 15px;
  font-size: 14px;
}

.signup-link a {
  color: #6c63ff;
  text-decoration: none;
}

.signup-link a:hover {
  text-decoration: underline;
}

/* Responsive Design */
@media (max-width: 600px) {
  .login-card {
    padding: 30px;
    max-width: 90%;
  }

  h1 {
    font-size: 22px;
  }

  button {
    font-size: 14px;
    padding: 10px;
  }
}
</style>
