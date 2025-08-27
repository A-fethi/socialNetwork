<template>
  <div class="profile-layout">
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
        <div class="sidebar-icon active" title="Profile" @click="goToProfile">
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
        <!-- Settings Icon (only if isOwner) -->
        <div v-if="isOwner" class="sidebar-icon" title="Settings" @click="goToSettings">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <circle cx="12" cy="12" r="3" />
            <path
              d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 01-2.83 2.83l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-4 0v-.09a1.65 1.65 0 00-1-1.51 1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83-2.83l.06-.06a1.65 1.65 0 00.33-1.82 1.65 1.65 0 00-1.51-1H3a2 2 0 010-4h.09a1.65 1.65 0 001.51-1 1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 012.83-2.83l.06.06a1.65 1.65 0 001.82.33h.09A1.65 1.65 0 0011 3.09V3a2 2 0 014 0v.09a1.65 1.65 0 001 1.51h.09a1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 2.83l-.06.06a1.65 1.65 0 00-.33 1.82v.09A1.65 1.65 0 0020.91 11H21a2 2 0 010 4h-.09a1.65 1.65 0 00-1.51 1z" />
          </svg>
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

    <!-- Main Content Area -->
    <div class="profile-main-content">
      <!-- Profile Card -->
      <div class="profile-card-wrapper">
        <div class="profile-card">
          <!-- Profile Image -->
          <button v-if="isOwner" class="settings-icon-btn" @click="goToSettings">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path
                d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z">
              </path>
              <circle cx="12" cy="12" r="3"></circle>
            </svg>
          </button>
          <div class="profile-image">
            <img
              :src="user.avatar"
              alt="Profile Picture"
              @click="showFullImage(user.avatar)"
              class="clickable-image"
              style="cursor:pointer"
            />
            <p class="username">@{{ user.username }}</p>
            <!-- Follow/Unfollow/Pending buttons -->
            <button v-if="!isOwner && user.follow_status === 'accepted'" @click="unfollow(user.username)"
              class="follow-btn following">
              Unfollow
            </button>
            <button v-else-if="!isOwner && user.follow_status === 'pending'" class="follow-btn pending" disabled>
              Pending
            </button>
            <button v-else-if="!isOwner && user.follow_status === 'not_following'" @click="follow(user.username)"
              class="follow-btn">
              Follow
            </button>
          </div>
          <div class="profile-details">
            <h2>{{ user.first_name }} {{ user.last_name }}</h2>
            <p v-if="user.date_of_birth" class="dob">
              {{ formatDate(user.date_of_birth) }}
            </p>
            <p v-if="
              user.email &&
              (isOwner || isFollowing || user.privacy === 'public')
            ">
              {{ user.email }}
            </p>
            <p v-if="user.nickname" class="nickname"><strong>Nickname: </strong>{{ user.nickname }}</p>
            <p v-if="user.bio" class="bio">{{ user.bio }}</p>
          </div>
          <div class="profile-stats">
            <div class="stat">
              <strong class="clickable" @click="toggleFollowersList">{{
                user.followers_count
              }}</strong>
              <span>Followers</span>
            </div>
            <div class="stat">
              <strong class="clickable" @click="toggleFollowingList">{{
                user.following_count
              }}</strong>
              <span>Following</span>
            </div>
            <div class="stat">
              <strong>{{ user.posts }}</strong>
              <span>Posts</span>
            </div>
          </div>
          <div class="followers-list" v-if="showFollowersList && canViewFollowData">
            <h3>Followers</h3>
            <ul>
              <li v-for="name in user.follower_usernames" :key="name">
                {{ name }}
              </li>
            </ul>
          </div>
          <div class="following-list" v-if="showFollowingList && canViewFollowData">
            <h3>Following</h3>
            <ul>
              <li v-for="name in user.following_usernames" :key="name">
                {{ name }}
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Posts Section -->
      <div v-if="canViewPosts" class="user-posts">
        <h3>Posts</h3>
        <div v-if="Array.isArray(posts) && posts.length > 0">
          <div v-for="post in posts" :key="post.id" class="post">
            <div class="post-header">
              <div class="post-author">
                <img
                  :src="user.avatar"
                  alt="Author Avatar"
                  class="post-avatar clickable-image"
                  @click="showFullImage(user.avatar)"
                  style="cursor:pointer"
                />
                <div class="post-author-info">
                  <h4>{{ post.title }}</h4>
                  <span class="post-author-name">{{ user.first_name }} {{ user.last_name }}</span>
                </div>
              </div>
              <small class="post-date">{{ formatDate(post.creation_date) }}</small>
            </div>
            <div class="post-content">
              <p>{{ post.content }}</p>
              <img
                v-if="post.image"
                :src="'http://localhost:8080/uploads/' + post.image"
                alt="Post Image"
                class="post-image clickable-image"
                @click="showFullImage('http://localhost:8080/uploads/' + post.image)"
                style="cursor:pointer"
              />
            </div>
            <div v-if="post.comments_count > 0" class="post-actions">
              <button @click="toggleComments(post.id)">
                {{ visibleCommentsPostId === post.id ? 'Hide Comments' : 'View Comments' }}
              </button>
              <span class="comments-count">{{ post.comments_count }} comments</span>
            </div>
            <div v-if="visibleCommentsPostId === post.id" class="comments-section">
              <ul>
                <li v-for="comment in Comments" :key="comment.id" class="comment-item">
                  <div class="comment-header">
                    <img v-if="comment.Avatar" :src="comment.Avatar" alt="Comment Avatar" class="comment-avatar" />
                    <div class="comment-info">
                      <strong class="comment-author">{{ comment.author }}</strong>
                      <small class="comment-date">{{ formatDate(comment.creation_date) }}</small>
                    </div>
                  </div>
                  <p class="comment-text">{{ comment.comment }}</p>
                  <img
                    v-if="comment.image"
                    :src="'http://localhost:8080/uploads/' + comment.image"
                    alt="Comment Image"
                    class="comment-image clickable-image"
                    @click="showFullImage('http://localhost:8080/uploads/' + comment.image)"
                    style="cursor:pointer"
                  />
                </li>
              </ul>
            </div>
          </div>
        </div>
        <div v-else>
          <p>No posts available.</p>
        </div>
      </div>
      <div v-else-if="user.privacy === 'private' && !isOwner" class="private-message">
        <div class="lock-icon">🔒</div>
        <h3>This account is private</h3>
        <p v-if="isPending">Your follow request is pending approval.</p>
        <p v-else>Follow this account to see their posts.</p>
      </div>
    </div>

    <!-- Image Modal -->
    <div v-if="showImageModal" class="image-modal" @click.self="closeImageModal">
      <img :src="selectedImage" class="modal-image" />
      <button class="close-modal-btn" @click="closeImageModal">&times;</button>
    </div>
  </div>
</template>

<script>
import { notificationWebSocket } from '../services/notificationWebSocket.js';

export default {
  name: "ProfileInfo",
  data() {
    return {
      isOwner: false,
      isFollowing: false,
      isPending: false,
      showFollowersList: false,
      showFollowingList: false,
      posts: [],
      Comments: [],
      currentUserId: null,
      visibleCommentsPostId: null,
      showNotifications: false,
      notifications: [],
      unreadNotificationCount: 0,
      user: {
        image: "",
        username: "",
        nickname: "",
        avatar: "",
        first_name: "",
        last_name: "",
        bio: "",
        followers_count: 0,
        following_count: 0,
        follower_usernames: [],
        following_usernames: [],
        posts: 0,
        privacy: "public",
        date_of_birth: "",
        email: "",
        follow_status: ""
      },
      CurrentUsername: "",
      imagePreview: null,
      showImageModal: false,
      selectedImage: null,
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
  computed: {
    canViewPosts() {
      return this.user.privacy === 'public' ||
        (this.user.privacy === 'private' && this.isFollowing) ||
        this.isOwner;
    },
    canViewFollowData() {
      return this.user.privacy === 'public' ||
        (this.user.privacy === 'private' && this.isFollowing) ||
        this.isOwner;
    }
  },
  watch: {
    '$route'(to, from) {
      
      if (to.params.id !== from.params.id) {
        this.resetComponentState();
        this.initializeProfile(to.params.id);
      }
    }
  },
  async created() {
    const userId = this.$route.params.id;
    const currentUserId = await this.getData();
    
    if (!currentUserId) {
      console.error("User is not logged in. Redirecting to login page.");
      this.$router.push("/login");
      return;
    }

    this.CurrentUsername = currentUserId;
    this.initializeProfile(userId, currentUserId);
  },
  methods: {
    showFullImage(imageUrl) {
      this.selectedImage = imageUrl;
      this.showImageModal = true;
    },
    closeImageModal() {
      this.showImageModal = false;
      this.selectedImage = null;
    },
    resetComponentState() {
      this.showFollowersList = false;
      this.showFollowingList = false;
      this.posts = [];
      this.Comments = [];
      this.visibleCommentsPostId = null;
      this.showNotifications = false;
      this.notifications = [];
      this.unreadNotificationCount = 0;
    },
    async initializeProfile(userId, currentUserId = null) {
      if (!currentUserId) {
        currentUserId = await this.getData();
        this.CurrentUsername = currentUserId;
      }
      
      this.fetchuserData(userId);
      this.isOwner = currentUserId === userId;
    },
    goToProfile() {
      this.$router.push(`/profile/${this.CurrentUsername}`);
    },
    goToSettings() {
      this.$router.push('/settings');
    },
    async toggleComments(postId) {
      if (this.visibleCommentsPostId === postId) {
        
        this.visibleCommentsPostId = null;
        this.Comments = [];
      } else {
        
        this.visibleCommentsPostId = postId;
        await this.getComments(postId);
      }
    },
    getComments(postId) {
      return fetch(`http://localhost:8080/api/getcomments?post_id=${postId}`, {
        method: 'GET',
        credentials: 'include'
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Failed to fetch comments');
          }
          return response.json();
        })
        .then(comments => {
          console.log('Comments fetched:', comments);
          for (let i = 0; i < comments.length; i++) {
            if (!comments[i].Avatar) {
              comments[i].Avatar = `https://api.dicebear.com/7.x/avataaars/svg?seed=${comments[i].Author}`;
            } else {
              comments[i].Avatar = 'http://localhost:8080/uploads/' + comments[i].Avatar;
            }
          }
          this.Comments = comments || [];
          console.log('Comments:', this.Comments);
        })
        .catch(error => {
          console.error('Error fetching comments:', error);
          this.showNotification(error.message || 'Error fetching comments', 'error');
        });
    },
    async getData() {
      try {
        const response = await fetch('http://localhost:8080/api/info', {
          method: 'GET',
          credentials: 'include',
        });

        if (!response.ok) {
          if (response.status === 404) {
            this.$router.push('/404');
            return;
          } else if (response.status === 401) {
            this.$router.push('/');
            return;
          }
          throw new Error('Failed to fetch data');
        }

        const data = await response.json();
        return data.Username;
      } catch (error) {
        console.error('Error fetching data:', error);
        this.$router.push('/404');
        throw error;
      }
    },
    checkIfFollowing() {
      if (!this.user.username) return;

      fetch(`http://localhost:8080/api/followers?action=isFollowing&profileUser=${this.user.username}`, {
        method: "GET",
        credentials: "include"
      })
        .then(response => {
          if (!response.ok) {
            if (response.status === 404) {
              this.$router.push('/404');
              return;
            }
            throw new Error("Failed to fetch follow status");
          }
          return response.json();
        })
        .then(data => {
          this.isFollowing = data.isFollowing && data.status === "accepted";
          this.isPending = data.isFollowing && data.status === "pending";
        })
        .catch(error => {
          console.error("Error checking follow status:", error);
          this.$router.push('/404');
          
          this.isFollowing = false;
          this.isPending = false;
        });
    },

    goToHome() {
      this.$router.push('/');
    },

    goToSettings() {
      this.$router.push('/settings');
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

    async fetchuserData(id) {
      try {
        const response = await fetch(`http://localhost:8080/api/userinfo?user_id=${id}`, {
          method: "GET",
          credentials: "include"
        });

        if (!response.ok) {
          const errorText = await response.text();
          if (response.status === 404) {
            this.$router.push('/404');
            return;
          } else if (response.status === 401) {
            this.$router.push('/');
            return;
          }
          throw new Error(errorText);
        }

        const data = await response.json();
        console.log("User data fetched:", data);
        this.user = data;
        if (!this.user.avatar) {
          this.user.avatar = `https://api.dicebear.com/7.x/avataaars/svg?seed=${data.username}`;
        } else {
          this.user.avatar = 'http://localhost:8080/uploads/' + this.user.avatar;
        }

        
        if (this.canViewPosts) {
          await this.fetchUserPosts(data.username);
        }

        
        await this.checkIfFollowing();

      } catch (error) {
        console.error("Error fetching user data:", error);
        this.showNotification(error.message || "Error fetching user data", "error");
        this.$router.push('/404');
      }
    },

    fetchUserPosts(username) {
      fetch(`http://localhost:8080/api/ownposts?username=${username}`, {
        method: "GET",
        credentials: "include"
      })
        .then(response => {
          if (!response.ok) {
            if (response.status === 404) {
              this.$router.push('/404');
              return;
            }
            throw new Error("Failed to fetch user posts");
          }
          return response.json()
        })
        .then(posts => {
          this.posts = posts || [];
          console.log("User posts fetched:", this.posts);
        })
        .catch(error => {
          console.error("Error fetching user posts:", error);
          this.$router.push('/404');
        });
    },

    toggleFollowersList() {
      if (this.canViewFollowData) {
        this.showFollowersList = !this.showFollowersList;
        this.showFollowingList = false;
      }
    },

    toggleFollowingList() {
      if (this.canViewFollowData) {
        this.showFollowingList = !this.showFollowingList;
        this.showFollowersList = false;
      }
    },

    async follow(profileUser) {
      try {
        const response = await fetch(`http://localhost:8080/api/followers?action=follow&profileUser=${profileUser}`, {
          method: 'POST',
          credentials: 'include'
        });

        if (!response.ok) {
          throw new Error('Follow failed');
        }

        
        await response.text();

        
        if (this.user.privacy === 'private') {
          this.isPending = true;
          this.isFollowing = false;
        } else {
          this.isFollowing = true;
          this.isPending = false;
        }

        
        await this.fetchuserData(this.user.username);

        
        setTimeout(() => {
          this.checkIfFollowing();
        }, 500);

      } catch (err) {
        console.error('Follow error:', err);
        this.isFollowing = false;
        this.isPending = false;
      }
    },

    async unfollow(profileUser) {
      try {
        const response = await fetch(`http://localhost:8080/api/followers?action=unfollow&profileUser=${profileUser}`, {
          method: 'POST',
          credentials: 'include'
        });

        if (!response.ok) {
          throw new Error('Unfollow failed');
        }

        
        this.isFollowing = false;
        this.isPending = false;

        
        await this.fetchuserData(this.user.username);

        
        if (!this.canViewPosts) {
          this.posts = [];
        }

        
        setTimeout(() => {
          this.checkIfFollowing();
        }, 300);

      } catch (err) {
        console.error('Unfollow error:', err);
        
      }
    },

    formatDate(dateString) {
      if (!dateString) return '';
      const date = new Date(dateString);
      return date.toLocaleDateString();
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
        }, 3000); 
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
      
      notificationWebSocket.onNotification('profile-info', (notification) => {
        console.log('Received real-time notification:', notification);
        
        
        this.fetchNotifications();
      });
    },
  },
  mounted() {
    
    this.setupNotificationWebSocket();
    
    this.fetchNotifications();
    document.addEventListener('click', this.handleNotifClose);
  },
  beforeUnmount() {
    
    notificationWebSocket.removeNotificationHandler('profile-info');
    document.removeEventListener('click', this.handleNotifClose);
  }
};
</script>

<style scoped>
/* Layout */
.profile-layout {
  display: flex;
  min-height: 100vh;
  background: #eaf1fb;
  font-family: 'Inter', sans-serif;
}

/* Sidebar */
.sidebar {
  width: 70px;
  background: #23263a;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem 0 1rem 0;
  border-radius: 1.5rem 0 0 1.5rem;
  box-shadow: 2px 0 16px rgba(35, 38, 58, 0.08);
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  z-index: 100;
}

.sidebar-icons {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  width: 100%;
  align-items: center;
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

.sidebar-icon.active,
.sidebar-icon:hover {
  background: #35395a;
  opacity: 1;
}

/* Main Area */
.profile-main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem 2rem 2rem 1.5rem;
  min-width: 0;
}

.profile-container {
  font-family: "Inter", sans-serif;
  min-height: 100vh;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(145deg, #e9ecef 0%, #dee2e6 100%);
  position: relative;
  overflow: hidden;
}

.profile-container::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at 20% 30%,
      rgba(99, 102, 241, 0.15) 0%,
      transparent 25%),
    radial-gradient(circle at 80% 70%,
      rgba(99, 102, 241, 0.15) 0%,
      transparent 25%);
  z-index: 0;
  opacity: 0.6;
}

.profile-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(90deg, #6366f1, #4f46e5);
  color: #ffffff;
  padding: 1.25rem 2.5rem;
  border-radius: 1.25rem;
  box-shadow: 0 8px 24px rgba(79, 70, 229, 0.3);
  margin-bottom: 2.5rem;
  width: 100%;
  max-width: 1200px;
  z-index: 1;
  transition: transform 0.3s ease;
}

.app-name {
  font-size: 1.75rem;
  font-weight: 700;
  letter-spacing: 0.5px;
}

.user-actions {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.user-actions span {
  font-size: 1.1rem;
  font-weight: 600;
  background: rgba(255, 255, 255, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
}

.home-btn,
.logout-btn {
  background: #ffffff;
  color: #4f46e5;
  font-weight: 600;
  padding: 0.6rem 1.5rem;
  border-radius: 50px;
  border: none;
  cursor: pointer;
  box-shadow: 0 4px 12px rgba(79, 70, 229, 0.2);
  transition: all 0.3s ease;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.home-btn:hover,
.logout-btn:hover {
  background: linear-gradient(90deg, #e0e7ff, #c7d2fe);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(79, 70, 229, 0.3);
}

.profile-card-wrapper {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
}

.profile-card {
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

.profile-image img {
  width: 7rem;
  height: 7rem;
  border-radius: 50%;
  border: 5px solid #ffffff;
  object-fit: cover;
  margin-top: -4rem;
  margin-bottom: -1.5rem;
  background: #ffffff;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
  transition: transform 0.3s ease;
  position: relative;
  z-index: 1;
}

.profile-image img:hover {
  transform: scale(1.05);
}

.username {
  font-size: 1.1rem;
  color: #6b7280;
  margin: 1rem 0 0.5rem 0;
  font-weight: 500;
  position: relative;
  z-index: 2;
}

/* Follow button styles */
.follow-btn {
  background: #6366f1;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 2rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 0.5rem;
  font-size: 0.9rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.follow-btn:hover:not(:disabled) {
  background: #4f46e5;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}

.follow-btn.following {
  background: #ef4444;
}

.follow-btn.following:hover {
  background: #dc2626;
}

.follow-btn.pending {
  background: #f59e0b;
  cursor: not-allowed;
  opacity: 0.8;
}

.follow-btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.profile-details h2 {
  margin: 1.25rem 0 0.75rem;
  font-size: 2rem;
  font-weight: 700;
  color: #1f2937;
  letter-spacing: -0.5px;
}

.dob {
  color: #6b7280;
  font-size: 0.9rem;
  margin-bottom: 0.5rem;
}

.bio {
  font-size: 0.95rem;
  color: #4b5563;
  white-space: pre-line;
  margin-bottom: 2rem;
  background: rgba(243, 244, 246, 0.8);
  padding: 0.75rem;
  border-radius: 0.75rem;
  line-height: 1.6;
}

.profile-stats {
  display: flex;
  justify-content: space-around;
  margin: 2rem 0;
  gap: 1.5rem;
}

.stat {
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  gap: 0.5rem;
  text-align: center;
  transition: transform 0.3s ease;
}

.stat strong {
  display: block;
  font-size: 1.75rem;
  font-weight: 700;
  color: #1f2937;
}

.stat span {
  font-size: 0.9rem;
  color: #6b7280;
  font-weight: 500;
}

.stat strong.clickable {
  cursor: pointer;
  color: #6366f1;
  transition: color 0.3s ease;
}

.stat strong.clickable:hover {
  color: #4f46e5;
}

.followers-list,
.following-list {
  flex: 1 1 100%;
  margin-top: 25px;
  background: #ffffff;
  padding: 20px;
  border-radius: 15px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
  animation: fadeIn 0.5s ease;
}

.followers-list h3,
.following-list h3 {
  font-size: 20px;
  font-weight: 600;
  color: #222;
  margin-bottom: 15px;
}

.followers-list ul,
.following-list ul {
  list-style: none;
  padding: 0;
}

.followers-list li,
.following-list li {
  padding: 12px 0;
  font-size: 15px;
  color: #444;
  border-bottom: 1px solid #f1f1f1;
  transition: background 0.3s ease;
}

.followers-list li:last-child,
.following-list li:last-child {
  border-bottom: none;
}

.settings-icon-btn {
  position: absolute;
  top: 1.25rem;
  background: rgba(255, 255, 255, 0.95);
  border: none;
  border-radius: 50%;
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 3px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  z-index: 2;
}

.settings-icon-btn:hover {
  background: #ffffff;
  transform: rotate(45deg);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.2);
}

.settings-icon-btn svg {
  color: #4f46e5;
  width: 24px;
  height: 24px;
  transition: transform 0.3s ease;
}

.settings-icon-btn:hover svg {
  transform: scale(1.1);
}

.user-posts {
  width: 100%;
  max-width: 800px;
  margin: 2rem auto;
  padding: 2rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(16px);
  border-radius: 1.5rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  text-align: left;
  z-index: 1;
  transition: transform 0.3s ease;
}

.user-posts h3 {
  font-size: 1.75rem;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 2rem;
  position: relative;
  padding-bottom: 0.75rem;
}

.user-posts h3::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 60px;
  height: 4px;
  background: linear-gradient(90deg, #4f46e5, #6366f1);
  border-radius: 2px;
}

.user-posts .post {
  padding: 1.75rem;
  border: 1px solid #e5e7eb;
  border-radius: 1.25rem;
  margin-bottom: 1.5rem;
  background: #ffffff;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05), 0 2px 4px -1px rgba(0, 0, 0, 0.03);
}

.user-posts .post:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.08), 0 4px 6px -2px rgba(0, 0, 0, 0.04);
  border-color: #d1d5db;
}

.user-posts .post-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1.25rem;
}

.user-posts .post-author {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-posts .post-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e5e7eb;
  transition: transform 0.3s ease;
}

.user-posts .post-avatar:hover {
  transform: scale(1.05);
}

.user-posts .post-author-info {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.user-posts .post-author-info h4 {
  font-size: 1.35rem;
  font-weight: 700;
  color: #1f2937;
  margin: 0;
  line-height: 1.4;
}

.user-posts .post-author-name {
  font-size: 0.95rem;
  color: #6b7280;
  font-weight: 500;
}

.user-posts .post-date {
  font-size: 0.875rem;
  color: #6b7280;
  font-weight: 500;
}

.user-posts .post-content {
  margin-bottom: 1.5rem;
}

.user-posts .post-content p {
  font-size: 1.05rem;
  color: #374151;
  line-height: 1.7;
  white-space: pre-wrap;
  word-break: break-word;
  margin-bottom: 1.25rem;
}

.user-posts .post-image {
  width: 100%;
  height: auto;
  border-radius: 1rem;
  margin: 1rem 0;
  border: 1px solid #e5e7eb;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.user-posts .post-image:hover {
  transform: scale(1.01);
}

.user-posts .post-actions {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}

.user-posts .post-actions button {
  background: linear-gradient(90deg, #4f46e5, #6366f1);
  color: #fff;
  border: none;
  border-radius: 0.75rem;
  padding: 0.6rem 1.25rem;
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(79, 70, 229, 0.2);
}

.user-posts .post-actions button:hover {
  background: linear-gradient(90deg, #4338ca, #4f46e5);
  transform: translateY(-1px);
  box-shadow: 0 6px 8px -1px rgba(79, 70, 229, 0.25);
}

.user-posts .post-actions button:active {
  transform: translateY(0);
}

.user-posts .post-actions button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.user-posts .post-actions button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.user-posts .post-actions button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.user-posts .comments-count {
  font-size: 0.95rem;
  color: #6b7280;
  font-weight: 500;
}


.comments-section {
  margin-top: 1.25rem;
  padding: 1rem 0 0 0;
  border-top: 1px solid #e5e7eb;
  background: #f9fafb;
  border-radius: 0 0 1rem 1rem;
  animation: fadeIn 0.3s;
}

.comment-item {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 1rem;
  border-bottom: 1px solid #f1f1f1;
}

.comment-item:last-child {
  border-bottom: none;
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.comment-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e5e7eb;
  background: #fff;
}

.comment-info {
  display: flex;
  flex-direction: column;
  gap: 0.1rem;
}

.comment-author {
  font-weight: 600;
  color: #4f46e5;
  font-size: 1rem;
}

.comment-date {
  font-size: 0.85rem;
  color: #6b7280;
}

.comment-text {
  font-size: 0.98rem;
  color: #374151;
  margin: 0.25rem 0 0.5rem 0;
  white-space: pre-line;
}

.comment-image {
  max-width: 100%;
  max-height: 300px;
  width: auto;
  height: auto;
  object-fit: contain;
  border-radius: 0.5rem;
  margin-top: 0.7rem;
}

.private-message {
  width: 100%;
  max-width: 800px;
  margin: 2rem auto;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(16px);
  border-radius: 1.5rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  text-align: center;
  z-index: 1;
}

.lock-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.private-message h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 1rem;
}

.private-message p {
  font-size: 1rem;
  color: #6b7280;
  line-height: 1.6;
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

/* Responsive */
@media (max-width: 1200px) {
  .profile-main-content {
    padding: 1rem 0.5rem;
  }
}

@media (max-width: 900px) {
  .profile-layout {
    flex-direction: column;
  }

  .sidebar {
    flex-direction: row;
    width: 100vw;
    height: 60px;
    border-radius: 0 0 1.5rem 1.5rem;
    padding: 0 1rem;
    box-shadow: 0 2px 16px rgba(35, 38, 58, 0.08);
    position: static;
  }

  .sidebar-icons {
    flex-direction: row;
    gap: 2rem;
    width: 100%;
    justify-content: center;
  }

  .profile-main-content {
    padding: 1rem 0.5rem;
    align-items: center;
  }

  .profile-card-wrapper,
  .profile-card,
  .user-posts,
  .private-message {
    max-width: 100%;
    margin: 1rem auto;
  }

  .profile-card {
    padding: 1.5rem;
  }

  .profile-image img {
    width: 6rem;
    height: 6rem;
    margin-top: -3rem;
  }

  .profile-details h2 {
    font-size: 1.5rem;
  }

  .user-posts {
    padding: 1rem;
  }

  .user-posts .post {
    padding: 1.25rem;
  }

  .user-posts .post-author-info h4 {
    font-size: 1.2rem;
  }

  .user-posts .post-content p {
    font-size: 0.95rem;
  }

  .private-message {
    margin-left: 0;
    padding: 2rem 1rem;
  }
}

@media (max-width: 600px) {
  .profile-main-content {
    padding: 0.5rem 0.2rem;
    align-items: center;
  }

  .profile-card-wrapper,
  .profile-card,
  .user-posts,
  .private-message {
    width: 100%;
    margin: 0.5rem auto;
    padding: 1rem;
  }

  .profile-card {
    padding: 1.25rem;
  }

  .profile-stats {
    gap: 1rem;
  }

  .stat strong {
    font-size: 1.5rem;
  }

  .stat span {
    font-size: 0.85rem;
  }

  .user-posts .post-header {
    flex-direction: column;
    gap: 0.5rem;
  }

  .user-posts .post-author {
    width: 100%;
  }

  .user-posts .post-date {
    align-self: flex-start;
  }

  .user-posts .post-actions {
    flex-direction: column;
    gap: 0.75rem;
  }

  .user-posts .post-actions button {
    width: 100%;
  }

  .comments-section {
    margin-top: 1rem;
    padding-top: 1rem;
  }

  .comment-item {
    padding: 0.875rem;
  }

  .comment-input-row {
    flex-direction: column;
    align-items: stretch;
    gap: 0.5rem;
  }

  .comment-form .file-upload-label {
    width: 100%;
    justify-content: center;
  }

  .comment-submit-btn {
    width: 100%;
    padding: 0.7rem;
  }
}

@media (max-width: 480px) {
  .profile-main-content {
    padding: 0.5rem;
  }

  .profile-card {
    padding: 1rem;
    border-radius: 1rem;
  }

  .profile-image img {
    width: 5rem;
    height: 5rem;
    margin-top: -2.5rem;
  }

  .profile-details h2 {
    font-size: 1.25rem;
  }

  .bio {
    font-size: 0.9rem;
    padding: 0.5rem;
  }

  .profile-stats {
    gap: 0.75rem;
  }

  .stat strong {
    font-size: 1.25rem;
  }

  .stat span {
    font-size: 0.8rem;
  }

  .user-posts {
    padding: 0.75rem;
  }

  .user-posts .post {
    padding: 1rem;
    border-radius: 1rem;
  }

  .user-posts .post-author-info h4 {
    font-size: 1.1rem;
  }

  .user-posts .post-content p {
    font-size: 0.9rem;
  }

  .user-posts .post-avatar {
    width: 40px;
    height: 40px;
  }

  .comment-item {
    padding: 0.75rem;
  }

  .comment-avatar {
    width: 28px;
    height: 28px;
  }

  .comment-text {
    font-size: 0.85rem;
  }

  .private-message {
    padding: 1.5rem 1rem;
  }

  .private-message h3 {
    font-size: 1.25rem;
  }

  .private-message p {
    font-size: 0.9rem;
  }
}

@media (max-width: 360px) {
  .profile-card {
    padding: 0.875rem;
  }

  .profile-image img {
    width: 4.5rem;
    height: 4.5rem;
    margin-top: -2.25rem;
  }

  .profile-details h2 {
    font-size: 1.1rem;
  }

  .bio {
    font-size: 0.85rem;
  }

  .profile-stats {
    gap: 0.5rem;
  }

  .stat strong {
    font-size: 1.1rem;
  }

  .stat span {
    font-size: 0.75rem;
  }

  .user-posts .post {
    padding: 0.875rem;
  }

  .user-posts .post-author-info h4 {
    font-size: 1rem;
  }

  .user-posts .post-content p {
    font-size: 0.85rem;
  }

  .sidebar-icons {
    gap: 0.5rem;
  }

  .sidebar-icon {
    width: 32px;
    height: 32px;
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

.clickable-image {
  cursor: pointer;
  transition: box-shadow 0.2s;
}
.clickable-image:hover {
  box-shadow: 0 0 0 3px #6366f1;
}
.image-modal {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(30, 32, 48, 0.85);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}
.modal-image {
  max-width: 90vw;
  max-height: 90vh;
  border-radius: 1rem;
  box-shadow: 0 8px 32px rgba(35,38,58,0.25);
  background: #fff;
}
.close-modal-btn {
  position: absolute;
  top: 2rem;
  right: 2rem;
  font-size: 2.5rem;
  background: none;
  border: none;
  color: #fff;
  cursor: pointer;
  z-index: 10000;
  transition: color 0.2s;
}
.close-modal-btn:hover {
  color: #ef4444;
}
</style>