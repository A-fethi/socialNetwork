<template>
  <div class="register-container">
    <div class="register-card">
      <div class="logo-container">
        <img
          src="https://logosandtypes.com/wp-content/uploads/2023/03/astro-framework.svg"
          alt="Astro Logo"
          class="logo"
        />
      </div>
      <form @submit.prevent="handleRegister">
        <div class="form-group" v-for="(field, key) in fields" :key="key">
          <label :for="key">{{ field.label }}</label>
          <input
            :type="field.type"
            :id="key"
            v-model="form[key]"
            :placeholder="field.placeholder"
            :required="field.required !== false"
          />
        </div>
        <div class="form-group">
          <label for="avatar">Avatar</label>
          <input
            type="file"
            id="avatar"
            @change="onAvatarChange"
            accept="image/*"
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
import router from "@/router";

export default {
  props: ["showNotification"],
  name: "UserRegister",
  data() {
    return {
      form: {
        email: "",
        password: "",
        firstname: "",
        lastname: "",
        bio: "",
        birthday: "",
        nickname: "",
      },
      avatarFile: null,
      fields: {
        
        email: {
          label: "Email",
          type: "email",
          placeholder: "test@example.com",
        },
        password: {
          label: "Password",
          type: "password",
          placeholder: "********",
        },
        firstname: { label: "First Name", type: "text", placeholder: "First Name" },
        lastname: { label: "Last Name", type: "text", placeholder: "Last Name" },
        nickname: {
          label: "Nickname",
          type: "text",
          placeholder: "Nickname (optional)",
          required: false,
        },
        bio: {
          label: "Bio",
          type: "text",
          placeholder: "Tell us about yourself",
          required: false,
        },
        birthday: { label: "Birthday", type: "date", placeholder: "" },
      },
    };
  },
  
  methods: {
    onAvatarChange(event) {
      this.avatarFile = event.target.files[0];
    },
    validateForm() {
      const errors = [];

      if (!/^[a-zA-Z0-9_]{3,20}$/.test(this.form.username)) {
        errors.push("Username must be 3-20 alphanumeric characters.");
      }

      if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(this.form.email)) {
        errors.push("Invalid email format.");
      }

      
      
      

      if (!/^[a-zA-Z]{1,30}$/.test(this.form.firstname)) {
        errors.push("First name must contain only letters.");
      }

      if (!/^[a-zA-Z]{1,30}$/.test(this.form.lastname)) {
        errors.push("Last name must contain only letters.");
      }

      
      if (!this.form.birthday) {
        errors.push("Date of birth is required.");
      } else {
        const birthDate = new Date(this.form.birthday);
        const today = new Date();
        const minAge = 13;
        const maxAge = 100;
        
        
        if (isNaN(birthDate.getTime())) {
          errors.push("Invalid date of birth format.");
        } else {
          
          if (birthDate > today) {
            errors.push("Date of birth cannot be in the future.");
          }
          
          
          const age = today.getFullYear() - birthDate.getFullYear();
          const monthDiff = today.getMonth() - birthDate.getMonth();
          const dayDiff = today.getDate() - birthDate.getDate();
          
          
          const actualAge = monthDiff < 0 || (monthDiff === 0 && dayDiff < 0) ? age - 1 : age;
          
          if (actualAge < minAge) {
            errors.push(`You must be at least ${minAge} years old to register.`);
          }
          if (actualAge > maxAge) {
            errors.push(`Date of birth indicates an age over ${maxAge} years.`);
          }
        }
      }

      if (this.avatarFile) {
        const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
        if (!allowedTypes.includes(this.avatarFile.type)) {
          errors.push("Invalid avatar type (jpg/jpeg/png only).");
        }

        if (this.avatarFile.size > 2 * 1024 * 1024) {
          errors.push("Avatar must be under 2MB.");
        }
      }

      return errors;
    },
    handleRegister() {
      this.errors = this.validateForm();
      if (this.errors.length > 0) {
        this.showNotification(
          "Please fix the following errors: " + this.errors.join(", "),
          "error"
        );
        return;
      }
      const formData = new FormData();
      
      if (this.avatarFile) {
        formData.append("avatar", this.avatarFile);
      }
      
      const userJson = JSON.stringify(this.form);
      formData.append("user", userJson);
      fetch("http://localhost:8080/api/auth/register", {
        method: "POST",
        credentials: "include",
        body: formData,
      })
        .then((res) => {
          if (!res.ok) throw new Error("Registration failed");
          return res.json();
        })
        .then((data) => {
          this.showNotification("Registration successful!", "success");
          console.log("Success:", data);
          router.push("/");
        })
        .catch((err) => {
          this.showNotification(
            "Registration failed. Please try again.",
            "error"
          );
          console.error("Error:", err);
        });
    },
  },
};
</script>

<style scoped>
/* General styles for the register page */
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: rgba(0, 0, 0, 0.7);
  background-image: url("https://learn.zone01oujda.ma/assets/img/01-trame.svg");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  font-family: "Arial", sans-serif;
  position: relative;
}

.register-container::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1;
}

.register-card {
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

.login-link {
  margin-top: 20px;
  font-size: 14px;
  color: #e0e0e0;
}

.login-link a {
  color: #6c63ff;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
}

.login-link a:hover {
  color: #5a52e1;
  text-decoration: none;
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

/* Responsive Design */
@media (max-width: 600px) {
  .register-card {
    padding: 30px;
    max-width: 90%;
  }

  button {
    font-size: 14px;
    padding: 10px;
  }
}
</style>
