<template>
  <div class="groups-layout">
    <!-- Left Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-icons">
        <!-- Home Icon -->
        <div class="sidebar-icon" title="Home" @click="goToHome">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24"><path d="M3 12L12 3l9 9"/><path d="M9 21V9h6v12"/></svg>
        </div>
        <!-- Profile Icon -->
        <div class="sidebar-icon" title="Profile" @click="goToProfile">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24"><circle cx="12" cy="8" r="4"/><path d="M4 21v-2a4 4 0 014-4h8a4 4 0 014 4v2"/></svg>
        </div>
        <!-- Groups Icon -->
        <div class="sidebar-icon active" title="Groups" @click="$router.push('/mygroups')">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24"><path d="M17 21v-2a4 4 0 00-4-4H7a4 4 0 00-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 00-3-3.87"/><path d="M16 3.13a4 4 0 010 7.75"/></svg>
        </div>
        <!-- Notification Icon -->
        <div class="sidebar-icon" ref="notifBtn" title="Notifications" @click="handleNotifClick">
          <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
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
            <li v-for="notif in notifications || []" :key="notif.id" class="notif-item" :class="{ unread: !notif.is_read }" @click="markNotificationAsRead(notif.id)">
              <span class="notif-text">
                <b>{{ notif.sender_username }}</b>
                <span>{{ notif.content }}</span>
              </span>
              <span class="notif-time">{{ formatTimestamp(notif.created_at) }}</span>
            </li>
          </ul>
        </div>
        <!-- Settings Icon -->
        <div class="sidebar-icon" title="Settings" @click="$router.push('/settings')">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 01-2.83 2.83l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-4 0v-.09a1.65 1.65 0 00-1-1.51 1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83-2.83l.06-.06a1.65 1.65 0 00.33-1.82 1.65 1.65 0 00-1.51-1H3a2 2 0 010-4h.09a1.65 1.65 0 001.51-1 1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 012.83-2.83l.06.06a1.65 1.65 0 001.82.33h.09A1.65 1.65 0 008.91 3H9a2 2 0 014 0v.09a1.65 1.65 0 001 1.51h.09a1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 2.83l-.06.06a1.65 1.65 0 00-.33 1.82v.09a1.65 1.65 0 001.51 1H21a2 2 0 010 4h-.09a1.65 1.65 0 00-1.51 1z"/></svg>
        </div>
        <!-- Logout Icon -->
        <div class="sidebar-icon" title="Logout" @click="logout">
          <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24"><path d="M17 16l4-4m0 0l-4-4m4 4H7"/><path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4"/></svg>
        </div>
      </div>
    </aside>
    <!-- Main Content -->
    <div class="groups-main-content">
      <div class="groups-card-wrapper">
        <h1 class="groups-title">My Groups</h1>
        <div v-if="myGroups.length === 0" class="no-groups">
          <p>You haven't joined any groups yet.</p>
          <button class="explore-btn" @click="goToForum">Explore Groups</button>
        </div>
        <div v-else class="groups-grid">
          <div v-for="group in myGroups" :key="group.id" class="group-card">
            <div class="group-info">
              <h3>{{ group.title }}</h3>
              <p>{{ group.description }}</p>
              <button class="show-invitations-btn" @click="toggleInvitations(group.id)">
                {{ showInvitations[group.id] ? 'Hide Invitations' : 'Show Invitations' }}
              </button>
            </div>
            <div class="group-actions">
              <button class="view-btn" @click="viewGroup(group.id)">View Group</button>
              <button v-if="!group.is_owner" class="leave-btn" @click="leaveGroup(group.id)">Leave Group</button>
            </div>
            <div class="pending-invitations" v-if="showInvitations[group.id]">
              <h4>Pending Invitations</h4>
              <div class="invitations-list">
                <div v-for="invitation in pendingInvitations" :key="invitation.id" class="invitation-card">
                  <div class="invitation-info">
                    <h3>{{ invitation.title }}</h3>
                    <p>{{ invitation.description }}</p>
                    <p class="inviter">Invited by: {{ invitation.username }}</p>
                  </div>
                  <div class="invitation-actions">
                    <button class="accept-btn" @click="handleInvitation(group.id, invitation.id, 'accept')">Accept</button>
                    <button class="refuse-btn" @click="handleInvitation(group.id, invitation.id, 'refuse')">Refuse</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <button class="back-btn" @click="goBack">Back to Forum</button>
      </div>
    </div>
  </div>
</template>

<script>
import { fetchInfo } from '@/api/getInfo';
import { notificationWebSocket } from '../services/notificationWebSocket.js';

export default {
  name: 'MyGroups',
  data() {
    return {
      myGroups: [],
      pendingInvitations: [],
      showInvitations: {},
      user: {
        name: ''
      },
      showNotifications: false,
      notifications: [],
      unreadNotificationCount: 0
    }
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
  async mounted() {
    
    this.setupNotificationWebSocket();
    
    await this.fetchMyGroups();
    await this.fetchPendingInvitations();
    document.addEventListener('click', this.handleNotifClose);
    this.fetchNotifications();
  },
  async created(){
    const userRes = await fetch("http://localhost:8080/api/info", {
      method: "GET",
      credentials: "include",
    });
    const userData = await userRes.json();
    console.log("userData", userData);

    this.user.name = userData.Username;
  },
  methods: {
    setupNotificationWebSocket() {
      
      notificationWebSocket.onNotification('my-groups', (notification) => {
        console.log('Received real-time notification:', notification);
        
        
        this.fetchNotifications();
      });
    },
    async fetchMyGroups() {
      try {
        const response = await fetch('http://localhost:8080/api/mygroups', {
          credentials: 'include'
        });
        if (response.ok) {
          this.myGroups = await response.json();
        }
      } catch (error) {
        console.error('Error fetching groups:', error);
      }
    },
    async fetchPendingInvitations(groupId) {
      try {
        const response = await fetch(`http://localhost:8080/api/GetInvitations?group_id=${groupId}`, {
          credentials: 'include'
        });
        if (response.ok) {
          this.pendingInvitations = await response.json();
        }
      } catch (error) {
        console.error('Error fetching pending invitations:', error);
      }
    },
    async toggleInvitations(groupId) {
      Object.keys(this.showInvitations).forEach(id => {
        if (id !== groupId.toString()) {
          this.showInvitations[id] = false;
        }
      });
      this.showInvitations[groupId] = !this.showInvitations[groupId];
      if (this.showInvitations[groupId]) {
        await this.fetchPendingInvitations(groupId);
      }
    },
    goBack() {
      this.$router.push('/home');
    },
    goToForum() {
      this.$router.push('/home');
    },
    goToHome() {
      this.$router.push('/home');
    },
    goToProfile() {
      this.$router.push(`/profile/${this.user.name}`);
    },
    logout() {
      fetch('http://localhost:8080/api/logout', {
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
    viewGroup(groupId) {
      this.$router.push(`/group/${groupId}`);
    },
    async handleInvitation(groupId, userId, action) {
      try {
        const endpoint = action === 'accept' 
          ? 'http://localhost:8080/api/acceptgroupmember'
          : 'http://localhost:8080/api/removememberfromgroup';
        const response = await fetch(endpoint, {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            group_id: groupId,
            user_id: userId
          })
        });
        if (response.ok) {
          await this.fetchPendingInvitations(groupId);
        } else {
          console.error('Failed to handle invitation:', await response.text());
        }
      } catch (error) {
        console.error('Error handling invitation:', error);
      }
    },
    async leaveGroup(groupId) {
      try {
        
        const userRes = await fetch("http://localhost:8080/api/info", {
          method: "GET",
          credentials: "include",
        });
        
        if (!userRes.ok) {
          console.error("Failed to get user information");
          return;
        }
        
        const userData = await userRes.json();
        
        const response = await fetch("http://localhost:8080/api/removememberfromgroup", {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            group_id: groupId,
            user_id: userData.id, 
          }),
        });
        
        if (response.ok) {
          console.log("Successfully left the group");
          
          await this.fetchMyGroups();
        } else {
          const errorText = await response.text();
          console.error("Failed to leave group:", errorText);
        }
      } catch (error) {
        console.error("Error leaving group:", error);
      }
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
        }, 2000); 
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
    }
  },
  beforeUnmount() {
    
    notificationWebSocket.removeNotificationHandler('my-groups');
    document.removeEventListener('click', this.handleNotifClose);
  }
}
</script>

<style scoped>
.groups-layout {
  display: flex;
  min-height: 100vh;
  background: linear-gradient(145deg, #e9ecef 0%, #dee2e6 100%);
}
.sidebar {
  width: 70px;
  background: #23263a;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem 0 1rem 0;
  border-radius: 1.5rem 0 0 1.5rem;
  box-shadow: 2px 0 16px rgba(35,38,58,0.08);
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
.sidebar-icon.active, .sidebar-icon:hover {
  background: #35395a;
  opacity: 1;
}
.groups-main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem 2rem 2rem 1.5rem;
  min-width: 0;
  margin-left: 70px;
}
.groups-card-wrapper {
  width: 100%;
  max-width: 1100px;
  margin: 0 auto;
  padding: 2rem 0;
}
.groups-title {
  font-size: 2.2rem;
  font-weight: 700;
  color: #23263a;
  margin-bottom: 2rem;
  letter-spacing: -0.5px;
}
.no-groups {
  text-align: center;
  padding: 3rem;
  background: rgba(255,255,255,0.85);
  border-radius: 1rem;
  border: 1px solid #e5e7eb;
  box-shadow: 0 4px 24px rgba(35,38,58,0.07);
}
.no-groups p {
  color: #6b7280;
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
}
.explore-btn {
  padding: 0.9rem 2rem;
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: #fff;
  border: none;
  border-radius: 50px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 4px 16px rgba(99,102,241,0.13);
}
.explore-btn:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(99,102,241,0.23);
}
.groups-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 2rem;
}
.group-card {
  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(10px);
  border-radius: 1.25rem;
  padding: 2rem 1.5rem 1.5rem 1.5rem;
  box-shadow: 0 8px 32px rgba(35,38,58,0.10);
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
  transition: transform 0.2s, box-shadow 0.2s;
  position: relative;
}
.group-card:hover {
  transform: translateY(-4px) scale(1.01);
  box-shadow: 0 16px 40px rgba(35,38,58,0.13);
}
.group-info h3 {
  color: #23263a;
  font-size: 1.3rem;
  font-weight: 700;
  margin: 0 0 0.5rem 0;
}
.group-info p {
  color: #6b7280;
  font-size: 1rem;
  margin: 0 0 0.75rem 0;
}
.group-actions {
  display: flex;
  gap: 0.75rem;
  margin-top: 0.5rem;
}
.view-btn, .leave-btn {
  padding: 0.7rem 1.5rem;
  border-radius: 50px;
  border: none;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.2s;
  box-shadow: 0 2px 8px rgba(99,102,241,0.08);
}
.view-btn {
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: #fff;
}
.view-btn:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: translateY(-2px);
}
.leave-btn {
  background: linear-gradient(135deg, #ef4444, #dc2626);
  color: #fff;
}
.leave-btn:hover {
  background: linear-gradient(135deg, #dc2626, #b91c1c);
  transform: translateY(-2px);
}
.show-invitations-btn {
  margin-top: 0.5rem;
  padding: 0.6rem 1.2rem;
  background: linear-gradient(135deg, #10b981, #22d3ee);
  color: #fff;
  border: none;
  border-radius: 50px;
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, transform 0.2s;
  box-shadow: 0 2px 8px rgba(16,185,129,0.08);
}
.show-invitations-btn:hover {
  background: linear-gradient(135deg, #22d3ee, #10b981);
  transform: translateY(-2px);
}
.pending-invitations {
  margin-top: 1.25rem;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}
.pending-invitations h4 {
  color: #23263a;
  font-size: 1.1rem;
  margin-bottom: 0.75rem;
  font-weight: 600;
}
.invitations-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
.invitation-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.1rem 1.5rem;
  background: rgba(248,250,252,0.95);
  border-radius: 1rem;
  border: 1px solid #e5e7eb;
  box-shadow: 0 2px 8px rgba(99,102,241,0.07);
  transition: transform 0.2s, box-shadow 0.2s;
}
.invitation-card:hover {
  transform: translateX(5px);
  box-shadow: 0 6px 16px rgba(99,102,241,0.13);
}
.invitation-info h3 {
  color: #23263a;
  font-size: 1.05rem;
  margin: 0 0 0.25rem 0;
  font-weight: 600;
}
.invitation-info p {
  color: #6b7280;
  font-size: 0.97rem;
  margin: 0;
}
.inviter {
  font-size: 0.93rem;
  color: #6b7280;
  margin-top: 0.5rem;
  font-style: italic;
}
.invitation-actions {
  display: flex;
  gap: 0.5rem;
}
.accept-btn, .refuse-btn {
  padding: 0.6rem 1.3rem;
  border-radius: 50px;
  border: none;
  cursor: pointer;
  font-size: 0.97rem;
  font-weight: 600;
  transition: all 0.2s;
}
.accept-btn {
  background: linear-gradient(135deg, #10b981, #22d3ee);
  color: #fff;
}
.accept-btn:hover {
  background: linear-gradient(135deg, #22d3ee, #10b981);
  transform: translateY(-2px);
}
.refuse-btn {
  background: linear-gradient(135deg, #ef4444, #dc2626);
  color: #fff;
}
.refuse-btn:hover {
  background: linear-gradient(135deg, #dc2626, #b91c1c);
  transform: translateY(-2px);
}
.back-btn {
  margin: 2.5rem auto 0 auto;
  display: block;
  padding: 0.8rem 2rem;
  background: #e5e7eb;
  border: none;
  border-radius: 50px;
  font-size: 1rem;
  font-weight: 600;
  color: #23263a;
  cursor: pointer;
  transition: background 0.2s, transform 0.2s;
  box-shadow: 0 2px 8px rgba(35,38,58,0.07);
}
.back-btn:hover {
  background: #d1d5db;
  transform: translateY(-2px);
}
@media (max-width: 900px) {
  .groups-card-wrapper {
    padding: 1rem 0;
  }
  .groups-grid {
    grid-template-columns: 1fr;
    gap: 1.2rem;
  }
}
@media (max-width: 768px) {
  .sidebar {
    width: 100%;
    min-height: auto;
    height: 70px;
    flex-direction: row;
    padding: 0;
    border-radius: 0;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
  }

  .sidebar-icons {
    flex-direction: row;
    justify-content: space-around;
    margin: 0;
    padding: 0 1rem;
    position: relative;
  }

  .groups-main-content {
    margin-left: 0;
    margin-top: 70px;
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
</style> 