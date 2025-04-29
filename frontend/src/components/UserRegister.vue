<template>
    <div class="register-container">
      <div class="register-card">
        <h1>Create Account</h1>
        <form @submit.prevent="handleRegister">
          <div class="form-group" v-for="(field, key) in fields" :key="key">
            <label :for="key">{{ field.label }}</label>
            <input
              :type="field.type"
              :id="key"
              v-model="form[key]"
              :placeholder="field.placeholder"
              required
            />
          </div>
          <button type="submit" class="btn">Register</button>
          <p class="login-link">
            Already have an account?
            <a href="/login">Login</a>
          </p>
        </form>
      </div>
    </div>
  </template>
  
  <script>
import router from '@/router';

  export default {
    props: ['showNotification'],
    name: "UserRegister",
    data() {
      return {
        form: {
          username: '',
          email: '',
          firstname: '',
          lastname: '',
          birthday: '',
        },
        fields: {
          username: { label: "Username", type: "text", placeholder: "johndoe" },
          email: { label: "Email", type: "email", placeholder: "you@example.com" },
            password: { label: "Password", type: "password", placeholder: "********" },
          firstname: { label: "First Name", type: "text", placeholder: "John" },
          lastname: { label: "Last Name", type: "text", placeholder: "Doe" },
          bio : { label: "Bio", type: "text", placeholder: "Tell us about yourself" },
          Date: { label: "Birthday", type: "date", placeholder: "" }
        }
      };
    },
    methods: {
      handleRegister() {
        fetch('http://localhost:8080/api/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
          body: JSON.stringify(this.form),
        })
          .then(res => {
            if (!res.ok) throw new Error("Registration failed");
            return res.json();
          })
          .then(data => {
            this.showNotification("Registration successful!", "success");
            console.log("Success:", data);
            router.push('/');
          })
          .catch(err => {
            this.showNotification("Registration failed. Please try again.", "error");
            console.error("Error:", err);
          });
      }
    }
  };
  </script>
  
  <style scoped>
  .register-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background: #eef1f7;
    font-family: 'Arial', sans-serif;
  }
  
  .register-card {
    background: white;
    padding: 40px;
    border-radius: 10px;
    width: 100%;
    max-width: 450px;
    box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
  }
  
  h1 {
    font-size: 26px;
    margin-bottom: 25px;
    color: #333;
    text-align: center;
  }
  
  .form-group {
    margin-bottom: 20px;
  }
  
  label {
    font-size: 14px;
    display: block;
    margin-bottom: 6px;
    color: #555;
  }
  
  input {
    width: 100%;
    padding: 10px;
    font-size: 15px;
    border: 1px solid #ccc;
    border-radius: 6px;
    background-color: #f9f9f9;
  }
  
  input:focus {
    outline: none;
    border-color: #6c63ff;
    background: white;
  }
  
  .btn {
    width: 100%;
    padding: 12px;
    background-color: #6c63ff;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 16px;
    cursor: pointer;
    transition: background 0.3s;
  }
  
  .btn:hover {
    background-color: #5b52d8;
  }
  
  .login-link {
    text-align: center;
    margin-top: 20px;
    font-size: 14px;
  }
  
  .login-link a {
    color: #6c63ff;
    text-decoration: none;
  }
  
  .login-link a:hover {
    text-decoration: underline;
  }
  </style>
  