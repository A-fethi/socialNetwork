<!-- src/components/UserLogin.vue -->
<template>
  <div class="login-container">
    <div class="login-card">
      <div class="logo-container">
        <img src="https://logosandtypes.com/wp-content/uploads/2023/03/astro-framework.svg" alt="Astro Logo" class="logo" />
      </div>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Email/Username</label>
          <input type="text" v-model="username" id="username" placeholder="Enter your Email/Username" required />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input type="password" v-model="password" id="password" placeholder="Enter your Password" required />
        </div>
        <button type="submit" class="btn">Login</button>
      </form>
      <p class="signup-link"> Don't have an account? <a href="/register">Sign up</a>
      </p>
    </div>
  </div>
</template>

<script>




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
      try {
        const response = await fetch('http://localhost:8080/api/auth/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            username: this.username,
            password: this.password
          }),
          credentials: 'include'
        });

        if (!response.ok) {
          throw new Error('Login failed');
        }

        const data = await response.json();
        console.log("dataaaaaaaA1", data);
        
        localStorage.setItem('xyz', data.xyz);
        this.$router.push('/home');
      } catch (error) {
      this.showNotification('Login failed. Please check your credentials.', 'error');
        console.error('Error during login:', error);
      }
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
  background-color: rgba(0, 0, 0, 0.7);
  background-image: url('https://learn.zone01oujda.ma/assets/img/01-trame.svg');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  font-family: 'Arial', sans-serif;
  position: relative;
}

.login-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.login-card {
  background-color: rgba(17, 17, 17, 0.95);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  padding: 40px;
  border-radius: 12px;
  width: 100%;
  max-width: 400px;
  text-align: center;
  position: relative;
  z-index: 2;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
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
  color: #e0e0e0;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background-color: rgba(255, 255, 255, 0.05);
  color: #ffffff;
  box-sizing: border-box;
  transition: all 0.3s ease;
}

input:focus {
  outline: none;
  border-color: #6c63ff;
  background-color: rgba(255, 255, 255, 0.1);
  box-shadow: 0 0 0 2px rgba(108, 99, 255, 0.2);
}

input::placeholder {
  color: rgba(255, 255, 255, 0.5);
}

button {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  background-color: #6c63ff;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 500;
}

button:hover {
  background-color: #5a52e1;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(108, 99, 255, 0.3);
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
  margin-top: 20px;
  font-size: 14px;
  color: #e0e0e0;
}

.signup-link a {
  color: #6c63ff;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
}

.signup-link a:hover {
  color: #5a52e1;
  text-decoration: none;
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

.logo-container {
  margin-bottom: 30px;
  display: flex;
  justify-content: center;
}

.logo {
  width: 120px;
  height: auto;
  transition: transform 0.3s ease;
  filter: brightness(0) invert(1);
}

.logo:hover {
  transform: scale(1.05);
}
</style>
