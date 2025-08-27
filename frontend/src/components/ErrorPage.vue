<template>
  <div class="error-container">
    <div class="error-content">
      <div class="error-icon">
        <svg width="120" height="120" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
      </div>
      <h1 class="error-title">404</h1>
      <h2 class="error-subtitle">Page Not Found</h2>
      <p class="error-message">Oops! The page you're looking for doesn't exist or has been moved.</p>
      <div class="error-actions">
        <button class="primary-btn" @click="$router.push('/')">
          <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M3 12l9-9 9 9"/>
            <path d="M9 21V9h6v12"/>
          </svg>
          Back to Home
        </button>
        <button class="secondary-btn" @click="$router.go(-1)">
          <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M19 12H5M12 19l-7-7 7-7"/>
          </svg>
          Go Back
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ErrorPage'
}
</script>

<style scoped>
.error-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #eaf1fb;
  padding: 2rem;
}

.error-content {
  background: #fff;
  border-radius: 1.5rem;
  padding: 3rem;
  text-align: center;
  box-shadow: 0 4px 24px rgba(35,38,58,0.08);
  max-width: 500px;
  width: 100%;
  animation: fadeIn 0.5s ease-out;
}

.error-icon {
  color: #4f46e5;
  margin-bottom: 1.5rem;
}

.error-title {
  font-size: 4rem;
  font-weight: 700;
  color: #23263a;
  margin: 0;
  line-height: 1;
  background: linear-gradient(90deg, #23263a 0%, #4f46e5 100%);
  -webkit-text-fill-color: transparent;
}

.error-subtitle {
  font-size: 1.5rem;
  color: #4f46e5;
  margin: 0.5rem 0 1rem;
  font-weight: 600;
}

.error-message {
  color: #6b7280;
  font-size: 1.1rem;
  margin-bottom: 2rem;
  line-height: 1.5;
}

.error-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
}

.primary-btn, .secondary-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.8rem 1.5rem;
  border-radius: 0.7rem;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.primary-btn {
  background: linear-gradient(90deg, #23263a 60%, #4f46e5 100%);
  color: #fff;
}

.primary-btn:hover {
  background: linear-gradient(90deg, #4f46e5 60%, #23263a 100%);
  transform: translateY(-1px);
}

.secondary-btn {
  background: #f3f4f8;
  color: #23263a;
}

.secondary-btn:hover {
  background: #e5e7eb;
  transform: translateY(-1px);
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Responsive Design */
@media (max-width: 600px) {
  .error-content {
    padding: 2rem;
  }
  
  .error-title {
    font-size: 3rem;
  }
  
  .error-subtitle {
    font-size: 1.3rem;
  }
  
  .error-message {
    font-size: 1rem;
  }
  
  .error-actions {
    flex-direction: column;
  }
  
  .primary-btn, .secondary-btn {
    width: 100%;
    justify-content: center;
  }
}
</style>