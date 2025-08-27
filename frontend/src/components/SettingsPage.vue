<template>
  <div class="settings-layout">
    <!-- Left Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-icons">
        <!-- Home Icon -->
        <div class="sidebar-icon" title="Home" @click="goToHome">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M3 12L12 3l9 9" />
            <path d="M9 21V9h6v12" />
          </svg>
        </div>
        <!-- Profile Icon -->
        <div class="sidebar-icon" title="Profile" @click="$router.push('/profile/' + user.name)">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <circle cx="12" cy="8" r="4" />
            <path d="M4 21v-2a4 4 0 014-4h8a4 4 0 014 4v2" />
          </svg>
        </div>
        <!-- Groups Icon -->
        <div class="sidebar-icon" title="Groups" @click="$router.push('/mygroups')">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M17 21v-2a4 4 0 00-4-4H7a4 4 0 00-4 4v2" />
            <circle cx="9" cy="7" r="4" />
            <path d="M23 21v-2a4 4 0 00-3-3.87" />
            <path d="M16 3.13a4 4 0 010 7.75" />
          </svg>
        </div>
        <!-- Notification Icon -->
        <div class="sidebar-icon" ref="notifBtn" title="Notifications" @click="handleNotifClick">
          <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M18 8a6 6 0 10-12 0c0 7-3 9-3 9h18s-3-2-3-9" />
            <path d="M13.73 21a2 2 0 01-3.46 0" />
          </svg>
          <span v-if="unreadNotificationCount > 0" class="notif-badge">{{ unreadNotificationCount }}</span>
        </div>
        <!-- Notification Popup -->
        <div v-if="showNotifications" ref="notifPopup" class="notif-popup">
          <div class="notif-popup-header">Notifications</div>
          <ul class="notif-list">
            <li v-if="!notifications || notifications.length === 0" class="notif-item empty">
              No notifications yet.
            </li>
            <li v-for="notif in notifications || []" :key="notif.id" class="notif-item"
              :class="{ unread: !notif.is_read }" @click="markNotificationAsRead(notif.id)">
              <span class="notif-text">
                <b>{{ notif.sender_username }}</b>
                <span>{{ notif.content }}</span>
              </span>
              <span class="notif-time">{{ formatTimestamp(notif.created_at) }}</span>
            </li>
          </ul>
        </div>
        <!-- Logout Icon Button -->
        <div class="sidebar-icon" title="Logout" @click="logout">
          <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M17 16l4-4m0 0l-4-4m4 4H7" />
            <path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4" />
          </svg>
        </div>
      </div>
    </aside>
    <!-- Main Settings Content -->
    <div class="settings-main-content">
      <!-- Settings Card -->
      <div class="settings-card-wrapper">
        <div class="settings-card">
          <h2>Privacy Settings</h2>
          <div class="privacy-section">
            <p class="privacy-status">
              Current Profile Status: {{ isProfilePrivate ? 'Private' : 'Public' }}
            </p>
            <select v-model="selectedPrivacy" @change="updatePrivacy" class="privacy-select">
              <option class="option" value="public">Public</option>
              <option class="option" value="private">Private</option>
            </select>
          </div>

          <div v-if="isProfilePrivate">
            <h2>Follow Requests</h2>
            <div class="invitations-section">
              <button class="action-btn" @click="showInvitations = !showInvitations">
                {{ showInvitations ? 'Hide Invitations' : 'View Invitations' }}
              </button>

              <div v-if="showInvitations">
                <h3>Pending Follow Requests</h3>
                <div v-if="pendingInvitations.length === 0" class="no-invitations">
                  No pending follow requests
                </div>
                <div v-for="invitation in pendingInvitations" :key="invitation.id" class="invitation-card">
                  <span>{{ invitation.username }}</span>
                  <div class="invitation-actions">
                    <button class="accept-btn" @click="acceptInvitation(invitation.username)">Accept</button>
                    <button class="reject-btn" @click="rejectInvitation(invitation.username)">Reject</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ProfileInfo from './ProfileInfo.vue';
import { fetchInfo } from '../api/getInfo.js';
import { notificationWebSocket } from '../services/notificationWebSocket.js';

export default {
  name: "MySettings",
  data() {
    return {
      selectedPrivacy: 'public',
      isProfilePrivate: false,
      showInvitations: false,
      pendingInvitations: [],
      user: {
        name: ''
      },
      showNotifications: false,
      notifications: [],
      unreadNotificationCount: 0
    };
  },

  beforeRouteEnter(to, from, next) {
    fetch("http://localhost:8080/api/info", {
      method: "GET",
      credentials: "include",
    })
      .then(res => {
        if (res.ok) {
          next();
        } else {
          next('/login');
        }
      })
      .catch(() => {
        next('/login');
      });
  },

  methods: {
    async updatePrivacy() {
      try {
        const res = await fetch(`http://localhost:8080/api/updateprivacy`, {
          method: 'POST',
          credentials: 'include',
          body: JSON.stringify({
            privacy: this.selectedPrivacy
          }),
        });

        if (res.ok) {
          this.isProfilePrivate = this.selectedPrivacy === 'private';
        } else {
          console.error('Failed to update privacy');
        }
      } catch (err) {
        console.error('Failed to update privacy:', err);
      }
    },
    async fetchInvitations() {
      try {
        const res = await fetch(`http://localhost:8080/api/getinvitationsfollow`, {
          method: 'GET',
          credentials: 'include'
        });
        const data = await res.json();
        this.pendingInvitations = data || [];
      } catch (err) {
        console.error('Failed to fetch invitations:', err);
        this.pendingInvitations = [];
      }
    },
    async acceptInvitation(invitationId) {
      try {
        const res = await fetch(`http://localhost:8080/api/accepteinvi`, {
          method: 'POST',
          credentials: 'include',
          body: JSON.stringify({
            follower_id: invitationId
          })
        });
        if (res.ok) {
          this.pendingInvitations = this.pendingInvitations.filter(inv => inv.username !== invitationId);

        }
      } catch (err) {
        console.error('Failed to accept invitation:', err);
      }
    },

    async rejectInvitation(invitationId) {
      try {
        const res = await fetch(`http://localhost:8080/api/followers?action=rejectInvitation&profileUser=${invitationId}`, {
          method: 'POST',
          credentials: 'include'
        });
        if (res.ok) {
          this.pendingInvitations = this.pendingInvitations.filter(inv => inv.id !== invitationId);
        }
      } catch (err) {
        console.error('Failed to reject invitation:', err);
      }
    },
    async checkprivacy() {
      const res = await fetch("http://localhost:8080/api/checkmyprivacy", {
        credentials: 'include',
      })
      const data = await res.json()
      if (data.privacy === "private") {
        this.isProfilePrivate = true
        this.selectedPrivacy = 'private'
      } else {
        this.isProfilePrivate = false
        this.selectedPrivacy = 'public'
      }
    },
    logout() {
      fetch('http://localhost:8080/api/auth/logout', {
        method: 'POST',
        credentials: 'include'
      })
        .then(res => {
          if (res.ok) {
            this.$router.push('/login');
          } else {
            console.error('Logout failed');
            this.$router.push('/login');
          }
        })
        .catch(err => {
          console.error('Logout error:', err);
          this.$router.push('/login');
        });
    },

    goToHome() {
      this.$router.push('/home');
    },
    async handleNotifClick() {
      this.showNotifications = !this.showNotifications;
      if (this.showNotifications) {
        await this.fetchNotifications();
      }
    },
    async fetchNotifications() {
      try {
        const res = await fetch('http://localhost:8080/api/notifications', {
          method: 'GET',
          credentials: 'include'
        });
        if (res.ok) {
          const data = await res.json();
          this.notifications = data;
          this.unreadNotificationCount = data.filter(n => !n.is_read).length;
        } else {
          this.notifications = [];
          this.unreadNotificationCount = 0;
        }
      } catch (err) {
        console.error('Failed to fetch notifications:', err);
        this.notifications = [];
        this.unreadNotificationCount = 0;
      }
    },
    async markNotificationAsRead(notificationId) {
      console.log(notificationId);

      try {
        
        const notification = this.notifications.find(n => n.id === notificationId);
        if (!notification) return;

        
        setTimeout(async () => {
          try {
            const res = await fetch(`http://localhost:8080/api/markasread`, {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json'
              },
              credentials: 'include',
              body: JSON.stringify({
                notificationId,
                is_read: 1
              })
            });
            console.log("res", res);
            if (res.ok) {
              notification.is_read = true;
              
              this.unreadNotificationCount = this.notifications.filter(n => !n.is_read).length;
            }
          } catch (error) {
            console.error('Error marking notification as read:', error);
          }
        }, 300); 
      } catch (error) {
        console.error('Error marking notification as read:', error);
      }
    },
    formatTimestamp(timestamp) {
      if (!timestamp) return '';
      const date = new Date(timestamp);
      const now = new Date();
      const diff = now - date;

      if (diff < 60000) return 'just now';
      if (diff < 3600000) return `${Math.floor(diff / 60000)}m ago`;
      if (diff < 86400000) return `${Math.floor(diff / 3600000)}h ago`;
      if (diff < 604800000) return `${Math.floor(diff / 86400000)}d ago`;

      return date.toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric'
      });
    },
    handleNotifClose(event) {
      if (this.showNotifications &&
        !this.$refs.notifBtn.contains(event.target) &&
        !this.$refs.notifPopup.contains(event.target)) {
        this.showNotifications = false;
      }
    },
    setupNotificationWebSocket() {
      
      notificationWebSocket.onNotification('settings-page', (notification) => {
        console.log('Received real-time notification:', notification);
        
        
        this.fetchNotifications();
      });
    },
  },
  async created() {
    const userRes = await fetch("http://localhost:8080/api/info", {
      method: "GET",
      credentials: "include",
    });
    const userData = await userRes.json();
    this.user.name = userData.Username;
    console.log("userData", this.user.name);
    await this.checkprivacy();
    await this.fetchInvitations();
  },
  mounted() {
    
    this.setupNotificationWebSocket();
    
    this.fetchNotifications();
    document.addEventListener('click', this.handleNotifClose);
  },
  beforeUnmount() {
    
    notificationWebSocket.removeNotificationHandler('settings-page');
    document.removeEventListener('click', this.handleNotifClose);
  }
};
</script>
<style scoped>
.settings-layout {
  display: flex;
  min-height: 100vh;
  background: #eaf1fb;
  font-family: 'Inter', sans-serif;
}

.sidebar {
  width: 70px;
  background: #23263a;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem 0 1rem 0;
  border-radius: 1.5rem 0 0 1.5rem;
  box-shadow: 2px 0 16px rgba(35, 38, 58, 0.08);
  min-height: 100vh;
  z-index: 100;
  justify-content: flex-start;
  position: fixed;
  left: 0;
  top: 0;
}

.sidebar-icons {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  width: 100%;
  align-items: center;
  position: relative;
  margin-bottom: 2.5rem;
}

.sidebar-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 1rem;
  color: #fff;
  font-size: 1.5rem;
  opacity: 0.7;
  cursor: pointer;
  transition: background 0.2s, opacity 0.2s;
}

.sidebar-icon:hover {
  background: #35395a;
  opacity: 1;
}

.settings-main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem 2rem 2rem 1.5rem;
  min-width: 0;
  margin-left: 70px;
}

.settings-card-wrapper {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
}

.settings-card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(16px);
  border-radius: 1.5rem;
  padding: 2.5rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  position: relative;
  animation: fadeIn 0.5s ease;
}

.settings-card h2 {
  color: #1f2937;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  letter-spacing: -0.5px;
}

.privacy-section {
  margin-bottom: 2.5rem;
  padding: 1.75rem;
  background: #f8fafc;
  border-radius: 1rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s ease;
}

.privacy-section:hover {
  transform: translateY(-3px);
}

.privacy-status {
  font-size: 1.1rem;
  color: #4f46e5;
  font-weight: 600;
  margin-bottom: 1.25rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.privacy-status::before {
  content: '🔒';
  font-size: 1.2rem;
}

.privacy-select {
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: #ffffff;
  padding: 0.9rem 2.5rem;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
  appearance: none;
  width: 100%;
  max-width: 300px;
  position: relative;
  text-align: center;
}

.privacy-select:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(79, 70, 229, 0.3);
}

.invitations-section {
  padding: 1.75rem;
  background: #f8fafc;
  border-radius: 1rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.action-btn {
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: #ffffff;
  padding: 0.9rem 2rem;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
  margin-bottom: 1.75rem;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.action-btn:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(79, 70, 229, 0.3);
}

.invitations-section h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 1.25rem;
}

.invitation-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  margin-bottom: 1rem;
  background: #ffffff;
  border-radius: 0.75rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.invitation-card:hover {
  transform: translateX(5px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.1);
}

.invitation-card span {
  font-size: 1rem;
  font-weight: 500;
  color: #1f2937;
}

.invitation-actions {
  display: flex;
  gap: 0.75rem;
}

.accept-btn,
.reject-btn {
  padding: 0.6rem 1.5rem;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  font-size: 0.9rem;
  font-weight: 600;
  transition: all 0.3s ease;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.accept-btn {
  background: linear-gradient(90deg, #22c55e, #16a34a);
  color: #ffffff;
}

.accept-btn:hover {
  background: linear-gradient(90deg, #16a34a, #15803d);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(22, 163, 74, 0.3);
}

.reject-btn {
  background: linear-gradient(90deg, #ef4444, #dc2626);
  color: #ffffff;
}

.reject-btn:hover {
  background: linear-gradient(90deg, #dc2626, #b91c1c);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(220, 38, 38, 0.3);
}

.no-invitations {
  color: #6b7280;
  font-size: 1rem;
  font-weight: 500;
  text-align: center;
  padding: 1.5rem;
  background: #f1f5f9;
  border-radius: 0.75rem;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(15px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 900px) {
  .settings-layout {
    flex-direction: column;
  }

  .sidebar {
    flex-direction: row;
    width: 100vw;
    height: 60px;
    border-radius: 0 0 1.5rem 1.5rem;
    padding: 0 1rem;
    box-shadow: 0 2px 16px rgba(35, 38, 58, 0.08);
    position: fixed;
    top: 0;
    left: 0;
    min-height: auto;
  }

  .sidebar-icons {
    flex-direction: row;
    gap: 2rem;
    width: 100%;
    justify-content: center;
    margin-bottom: 0;
  }

  .settings-main-content {
    padding: 1rem 0.5rem;
    margin-left: 0;
    margin-top: 60px;
  }

  .settings-card-wrapper,
  .settings-card {
    max-width: 100%;
    margin: 1rem auto;
  }

  .settings-card {
    padding: 1.5rem;
  }

  .settings-card h2 {
    font-size: 1.5rem;
  }

  .privacy-section,
  .invitations-section {
    padding: 1.5rem;
  }

  .notif-popup {
    top: 60px;
    left: 50%;
    transform: translateX(-50%);
    width: 90%;
    max-width: 350px;
  }
}

@media (max-width: 600px) {
  .settings-main-content {
    padding: 0.5rem 0.2rem;
  }

  .settings-card-wrapper,
  .settings-card {
    padding: 1rem;
  }

  .settings-card h2 {
    font-size: 1.5rem;
  }

  .privacy-select,
  .action-btn {
    padding: 0.8rem 1.5rem;
  }

  .invitation-card {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
    padding: 1rem;
  }

  .invitation-actions {
    width: 100%;
    justify-content: space-between;
  }

  .accept-btn,
  .reject-btn {
    width: 48%;
    text-align: center;
  }

  .sidebar-icons {
    gap: 1rem;
  }

  .sidebar-icon {
    width: 40px;
    height: 40px;
  }
}

@media (max-width: 480px) {
  .settings-card,
  .privacy-section,
  .invitations-section {
    padding: 1rem;
    border-radius: 1rem;
  }

  .settings-card h2 {
    font-size: 1.25rem;
  }

  .privacy-status {
    font-size: 1rem;
  }

  .privacy-select,
  .action-btn {
    padding: 0.7rem 1.25rem;
    font-size: 0.9rem;
  }

  .invitation-card {
    padding: 0.875rem;
  }

  .invitation-card span {
    font-size: 0.9rem;
  }

  .accept-btn,
  .reject-btn {
    padding: 0.5rem 1rem;
    font-size: 0.8rem;
  }

  .no-invitations {
    padding: 1rem;
    font-size: 0.9rem;
  }

  .sidebar-icons {
    gap: 0.5rem;
  }

  .sidebar-icon {
    width: 36px;
    height: 36px;
  }
}

@media (max-width: 360px) {
  .sidebar-icons {
    gap: 0.25rem;
  }

  .sidebar-icon {
    width: 32px;
    height: 32px;
  }

  .settings-card {
    padding: 0.875rem;
  }

  .privacy-section,
  .invitations-section {
    padding: 0.875rem;
  }
}

.sidebar-icon[title="Notifications"] {
  position: relative;
}

.notif-badge {
  position: absolute;
  top: 4px;
  right: 4px;
  background: #ef4444;
  color: white;
  border-radius: 50%;
  width: 18px;
  height: 18px;
  font-size: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  border: 2px solid #fff;
  z-index: 1;
}

.notif-popup {
  position: absolute;
  top: 70px;
  left: 80px;
  width: 350px;
  background: #fff;
  border-radius: 1.1rem;
  box-shadow: 0 8px 32px rgba(35, 38, 58, 0.18);
  z-index: 5000;
  padding: 0.5rem 0 0.5rem 0;
  animation: modalPopIn 0.22s cubic-bezier(.4, 2, .6, 1) both;
}

.notif-popup-header {
  font-size: 1.1rem;
  font-weight: 700;
  color: #4f46e5;
  padding: 1rem 1.5rem 0.5rem 1.5rem;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.notif-list {
  list-style: none;
  margin: 0;
  padding: 0.5rem 0;
  max-height: 400px;
  overflow-y: auto;
}

.notif-item {
  display: flex;
  align-items: flex-start;
  gap: 0.7rem;
  padding: 0.8rem 1.5rem;
  font-size: 0.98rem;
  color: #23263a;
  border-bottom: 1px solid #f3f4f8;
  transition: background 0.18s;
  position: relative;
  cursor: pointer;
}

.notif-item:last-child {
  border-bottom: none;
}

.notif-item:hover {
  background: #f8fafc;
}

.notif-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.notif-text b {
  color: #4f46e5;
  font-weight: 600;
}

.notif-text span {
  color: #4b5563;
  line-height: 1.4;
}

.notif-time {
  font-size: 0.85rem;
  color: #6b7280;
  margin-left: 0.5rem;
  white-space: nowrap;
}

.notif-item.empty {
  color: #b0b3c6;
  text-align: center;
  justify-content: center;
  font-size: 1rem;
  padding: 2rem 1.5rem;
}

.notif-item.unread {
  background: #f0f7ff;
  border-left: 3px solid #4f46e5;
}

.notif-item.unread::after {
  content: '';
  position: absolute;
  top: 50%;
  right: 1rem;
  transform: translateY(-50%);
  width: 8px;
  height: 8px;
  background: #4f46e5;
  border-radius: 50%;
}

@keyframes modalPopIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }

  to {
    opacity: 1;
    transform: scale(1);
  }
}

.option {
  background: #f3f4f6;
  color: #1f2937;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  font-size: 0.9rem;
}
.option:hover {
  background: #e5e7eb;
}
</style>