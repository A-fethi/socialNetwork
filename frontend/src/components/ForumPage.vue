<template>
  <div class="forum-layout">
    <!-- Left Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-icons">
        <!-- Home Icon -->
        <div class="sidebar-icon active" title="Home">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M3 12L12 3l9 9" />
            <path d="M9 21V9h6v12" />
          </svg>
        </div>
        <!-- Profile Icon -->
        <div class="sidebar-icon" title="Profile" @click="goToProfile">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <circle cx="12" cy="8" r="4" />
            <path d="M4 21v-2a4 4 0 014-4h8a4 4 0 014 4v2" />
          </svg>
        </div>
        <!-- Groups Icon -->
        <div class="sidebar-icon" title="Groups" @click="goToMyGroups">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M17 21v-2a4 4 0 00-4-4H7a4 4 0 00-4 4v2" />
            <circle cx="9" cy="7" r="4" />
            <path d="M23 21v-2a4 4 0 00-3-3.87" />
            <path d="M16 3.13a4 4 0 010 7.75" />
          </svg>
        </div>
        <!-- Settings Icon -->
        <div class="sidebar-icon" title="Settings" @click="$router.push('/settings')">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <circle cx="12" cy="12" r="3" />
            <path
              d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 01-2.83 2.83l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-4 0v-.09a1.65 1.65 0 00-1-1.51 1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83-2.83l.06-.06a1.65 1.65 0 00.33-1.82 1.65 1.65 0 00-1.51-1H3a2 2 0 010-4h.09a1.65 1.65 0 001.51-1 1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 012.83-2.83l.06.06a1.65 1.65 0 001.82.33h.09A1.65 1.65 0 0011 3.09V3a2 2 0 014 0v.09a1.65 1.65 0 001 1.51h.09a1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 2.83l-.06.06a1.65 1.65 0 00-.33 1.82v.09A1.65 1.65 0 0020.91 11H21a2 2 0 010 4h-.09a1.65 1.65 0 00-1.51 1z" />
          </svg>
        </div>
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
            <li v-if="!notifications || notifications.length === 0" class="notif-item empty">No notifications yet.</li>
            <li v-for="notif in notifications || []" :key="notif.time + notif.sender" class="notif-item"
              :class="{ 'unread': !notif.is_read }" @click="markNotificationAsRead(notif.id)">
              
              <span class="notif-text">
                <b>{{ notif.sender }}</b>
                <span v-if="notif.type === 'message'">sent you a message: {{ notif.text && notif.text.length > 40 ?
                  notif.text.slice(0, 40) + '…' : notif.text }}</span>
                <span v-else>{{ notif.text }}</span>
              </span>
              <span class="notif-time">{{ formatTimestamp(notif.time) }}</span>
            </li>
          </ul>
        </div>
        <div class="sidebar-icon" @click="logout">
          <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M17 16l4-4m0 0l-4-4m4 4H7" />
            <path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4" />
          </svg>
        </div>
      </div>
    </aside>

    <!-- Main Content Area -->
    <div class="main-area">
      <!-- Top Avatar Bar -->
      <div class="avatar-bar">
        <div class="avatar-list">
          <div v-for="other in otherUsers" :key="other.name" class="avatar-item">
            <img :src="other.avatar" :alt="other.name" @click="gotoxprofile(other.username)" />
            <span class="avatar-name">{{ other.name.split(' ')[0] }}</span>
          </div>
        </div>
      </div>

      <!-- Create Post -->
      <div class="create-post-card">
        <img class="profile-pic" :src="user.avatar" alt="Profile" />
        <form class="create-post-form" @submit.prevent="submitPost">
          <input type="text" v-model="newPost.title" placeholder="What's on your mind?" required />
          <textarea v-model="newPost.content" placeholder="Share your thoughts..." required></textarea>
          <label class="file-upload-label">
            <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
              stroke-linejoin="round" viewBox="0 0 24 24">
              <path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4" />
              <polyline points="17 8 12 3 7 8" />
              <line x1="12" y1="3" x2="12" y2="15" />
            </svg>
            Upload Image
            <input type="file" @change="handleImageUpload" accept="image/*" />
          </label>
          <img v-if="imagePreview" :src="imagePreview" class="image-preview" alt="Preview" @click="showFullImage(imagePreview)" />
          <div class="create-post-row">
            <div class="status-select">
              <select v-model="newPost.status" required>
                <option value="public">Public</option>
                <option value="private">Private</option>
                <option value="semi-private">Semi-Private</option>
              </select>
            </div>
          </div>
          <div v-if="newPost.status === 'semi-private'" class="allowed-users-select">
            <div v-if="allowedUsers.length === 0">No User is following you!</div>
            <label v-else v-for="user in allowedUsers" :key="user.username" class="allowed-user-pill">
              <input type="checkbox" v-model="selectedAllowedUsers" :value="user.username" />
              <span class="custom-checkbox">
                <svg v-if="selectedAllowedUsers.includes(user.username)" width="18" height="18" viewBox="0 0 18 18"
                  fill="none">
                  <circle cx="9" cy="9" r="9" fill="#4f46e5" />
                  <path d="M5 9l3 3 5-5" stroke="#fff" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" />
                </svg>
                <svg v-else width="18" height="18" viewBox="0 0 18 18" fill="none">
                  <circle cx="9" cy="9" r="9" fill="#e5e7eb" />
                </svg>
              </span>
              <img :src="user.avatar" :alt="user.username" class="allowed-user-avatar" />
              <span class="allowed-user-name">{{ user.username }}</span>
            </label>
          </div>
          <button class="post-btn" type="submit">Post</button>
          <div v-if="message" class="message">{{ message }}</div>
        </form>
      </div>

      <!-- Posts Grid -->
      <div class="posts-grid">
        <div v-if="posts.length === 0" class="no-posts">No posts available. Be the first to post!</div>
        <div v-else class="post-card" v-for="post in posts" :key="post.ID">
          <div class="post-image" v-if="post.imageUrl">
            <img :src="post.imageUrl" alt="Post image" @click="showFullImage(post.imageUrl)" />
          </div>
          <div class="post-header">
            <img class="post-author-pic" :src="post.authorAvatar" alt="Author" />
            <div>
              <h4>{{ post.Author }}</h4>
              <p class="post-timestamp">{{ formatTimestamp(post.Creation_date) }}</p>
            </div>
          </div>
          <h3 class="post-title">{{ post.Title }}</h3>
          <p class="post-content">{{ post.Content }}</p>
          <div class="post-image" v-if="post.Image" :class="post.Image ? 'post-image' : ''" @click="showFullImage(post.Image)">
            <img :src="post.Image" alt="Post Image" />
          </div>
          <div class="post-actions">
            <span class="icon-btn" @click="toggleComments(post)">
              <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <path d="M21 15a2 2 0 01-2 2H7l-4 4V5a2 2 0 012-2h14a2 2 0 012 2z" />
              </svg>
            </span>
          </div>
          <!-- Comments Section -->
          <div v-if="post.showComments" class="comments-section">
            <div v-if="post.comments.length === 0" class="no-comments">No comments yet. Be the first to comment!</div>
            <div v-else class="comment" v-for="comment in post.comments" :key="comment.id">
              <div class="comment-header">
                <img class="comment-author-pic" :src="comment.avatar" alt="Comment Author" />
                <div>
                  <h5>{{ comment.author }}</h5>
                  <p class="comment-timestamp">{{ formatTimestamp(comment.creation_date) }}</p>
                </div>
              </div>
              <p class="comment-content">{{ comment.comment }}</p>
              <div class="comment-image" v-if="comment.image" :class="comment.image ? 'comment-image' : ''" @click="showFullImage(comment.image)">
                <img class="comment-img" :src="comment.image" alt="Comment Image" />
              </div>
            </div>
            <form @submit.prevent="addComment(post)" class="comment-form">
              <div class="comment-input-row">
                <textarea v-model="post.newComment" placeholder="Add a comment..." required></textarea>
                <label class="file-upload-label">
                  <svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" viewBox="0 0 24 24">
                    <path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4" />
                    <polyline points="17 8 12 3 7 8" />
                    <line x1="12" y1="3" x2="12" y2="15" />
                  </svg>
                  <input type="file" @change="handleImageUpload($event, post.Id)" accept="image/*" />
                </label>
                <button type="submit" class="comment-submit-btn" title="Add Comment">
                  <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" viewBox="0 0 24 24">
                    <path d="M21 15a2 2 0 01-2 2H7l-4 4V5a2 2 0 012-2h14a2 2 0 012 2z" />
                  </svg>
                </button>
              </div>
              <div v-if="commentImages[post.Id] && commentImages[post.Id].fileName" class="file-name-display">
                Selected file: {{ commentImages[post.Id].fileName }}
              </div>
              <img v-if="commentImages[post.Id] && commentImages[post.Id].preview" :src="commentImages[post.Id].preview" class="image-preview" alt="Preview" @click="showFullImage(commentImages[post.Id].preview)" />
              <div v-if="post.commentError" class="comment-error">{{ post.commentError }}</div>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Sidebar -->
    <aside class="rightbar">
      <div class="groups-section">
        <div class="groups-header">
          <span>Groups</span>
          <button class="create-group-btn" @click="showCreateGroup = true" title="Create Group">
            <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
              stroke-linejoin="round" viewBox="0 0 24 24">
              <circle cx="12" cy="12" r="10" />
              <path d="M12 8v8M8 12h8" />
            </svg>
          </button>
          <!-- <button class="invite-group-btn" @click="inviteToGroup" title="Invite to Group">
            <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24"><path d="M21 10v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6"/><path d="M3 6l9 6 9-6"/></svg>
          </button> -->
        </div>
        <ul class="groups-list">
          <li v-for="group in groups" :key="group.id" class="group-list-item group-clickable"
            @click="goToGroup(group.id)">
            <div class="group-info">
              <span class="group-avatar">
                <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                  stroke-linejoin="round" viewBox="0 0 24 24">
                  <path d="M17 21v-2a4 4 0 00-4-4H7a4 4 0 00-4 4v2" />
                  <circle cx="9" cy="7" r="4" />
                  <path d="M23 21v-2a4 4 0 00-3-3.87" />
                  <path d="M16 3.13a4 4 0 010 7.75" />
                </svg>
              </span>
              <div class="group-details">
                <span class="group-title">{{ group.name }}</span>
                <span class="group-desc">{{ group.description }}</span>
                <span v-if="group.is_owner" class="owner-badge">Owner</span>
              </div>
            </div>
            <button v-if="group.member_status == 'accepted' && !group.is_owner" class="followed-btn"
              @click.stop="toggleJoin(group)">
              Leave
            </button>
            <button v-else-if="group.member_status == 'pending' && !group.is_owner" class="pending-btn"
              @click.stop="toggleJoin(group)">
              Cancel Request
            </button>
            <button v-else-if="group.member_status == 'invited' && !group.is_owner" class="pending-btn"
              @click.stop="toggleJoin(group)">
              Accept Invitation
            </button>
            <button v-else-if="!group.is_owner" class="follow-btn" @click.stop="toggleJoin(group)">
              Join
            </button>
          </li>
        </ul>
      </div>
      <div class="users-section">
        <div class="users-header">
          <span>Followed Users: </span>
        </div>
        <ul class="users-list">
          <li v-for="followed in followedUsers" :key="followed.name" class="user-list-item">
            <img class="mini-profile-pic" :src="followed.avatar" alt="User" @click="gotoxprofile(followed.username)" />
            <span class="user-name" @click="gotoxprofile(followed.username)">{{ followed.name }}</span>
          </li>
        </ul>
      </div>
    </aside>

    <!-- Create Group Modal -->
    <div v-if="showCreateGroup" class="modal-overlay">
      <div class="modal-content create-group-modal">
        <h3>Create New Group</h3>
        <form @submit.prevent="createGroup">
          <div class="form-group">
            <label for="groupTitle">Group Title</label>
            <input id="groupTitle" v-model="newGroup.title" type="text" required placeholder="Enter group title" />
          </div>
          <div class="form-group">
            <label for="groupDescription">Description</label>
            <textarea id="groupDescription" v-model="newGroup.description" required
              placeholder="Enter group description"></textarea>
          </div>
          <div class="modal-buttons">
            <button type="submit" class="submit-btn">Create Group</button>
            <button type="button" class="cancel-btn" @click="showCreateGroup = false">
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Chat Floating Button -->
    <button class="chat-fab" @click="handleChatClick" title="Open Chat">
      <svg width="28" height="28" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"
        stroke-linejoin="round" viewBox="0 0 24 24">
        <path d="M21 15a2 2 0 01-2 2H7l-4 4V5a2 2 0 012-2h14a2 2 0 012 2z" />
      </svg>
    </button>

    <!-- Chat Modal -->
    <div v-if="isChatExpanded" class="chat-modal-overlay" @click.self="isChatExpanded = false">
      <div class="chat-modal">
        <div class="chat-modal-header">
          <span>Chat</span>
          <button class="chat-modal-close" @click="isChatExpanded = false">&times;</button>
        </div>
        <div class="chat-users-list">
          <div class="chat-users-title">Chat Users</div>
          <ul>
            <li v-for="user in chatUsers" :key="user.name" @click="selectChatUser(user)"
              :class="{ selected: selectedChatUser && selectedChatUser.name === user.name }">
              <img :src="user.avatar" :alt="user.name" class="chat-user-avatar" />
              <span>{{ user.fullname }}</span>
            </li>
          </ul>
        </div>
        <div class="chat-messages-area" v-if="selectedChatUser">
          <div class="chat-messages-title">Chat with {{ selectedChatUser.name }}</div>
          <div class="chat-messages-list">
            <div v-for="msg in selectedChatUser.messages" :key="msg.id"
              :class="['chat-msg', msg.sender === 'self' ? 'self' : 'other']">
              <span>{{ msg.text }}</span>
              <span class="chat-msg-time">{{ formatTimestamp(msg.timestamp) }}</span>
            </div>
          </div>
          <form class="chat-send-form" @submit.prevent="sendMessage">
            <input v-model="newMessage" placeholder="Type a message..." required />
            <button type="submit" class="chat-send-btn" title="Send">
              <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <line x1="22" y1="2" x2="11" y2="13" />
                <polygon points="22 2 15 22 11 13 2 9 22 2" />
              </svg>
            </button>
          </form>
        </div>
        <div v-else class="chat-no-user">Select a user to start chatting</div>
      </div>
    </div>

    <!-- Add Image Preview Modal -->
    <div v-if="showImageModal" class="image-modal-overlay" @click="closeImageModal">
      <div class="image-modal-content" @click.stop>
        <img :src="selectedImage" alt="Full size image" />
        <button class="close-modal-btn" @click="closeImageModal">&times;</button>
      </div>
    </div>
  </div>
</template>

<script>
import { notificationWebSocket } from '../services/notificationWebSocket.js';

export default {
  props: ["showNotification"],
  name: "ForumPage",
  data() {
    return {
      user: {
        avatar: "",
        name: "",
        email: "",
        username: "guest",
      },
      socket: null,
      newPost: {
        author: "",
        title: "",
        content: "",
        image: "",
        status: "public", 
      },
      image: null,
      imagePreview: null,
      imageFileName: "", 
      posts: [],
      message: "",
      allowedUsers: [],
      selectedAllowedUsers: [],
      groups: [],
      otherUsers: [],
      followedUsers: [], 
      
      isChatExpanded: false,
      newMessage: "",
      selectedChatUser: null,
      chatUsers: [],
      currentMessage: {
        type: "",
        message: "",
        receiver: "",
        username: "",
      },
      showCreateGroup: false,
      newGroup: {
        title: "",
        description: "",
        creator_id: null,
      },
      groupMembership: {},
      postsPrvMap: {},
      showNotifications: false,
      notifications: [],
      commentImages: {},
      unreadNotificationCount: 0,
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
  async created() {
    this.$router.push('/home');
    try {
      const userRes = await fetch("http://localhost:8080/api/info", {
        method: "GET",
        credentials: "include",
      });
      const userData = await userRes.json();
      console.log("userData", userData);

      this.user.name = userData.Username;
      this.user.email = userData.Email;
      this.user.username = userData.Username.toLowerCase();

      if (userData.Avatar) {
        this.user.avatar = `http://localhost:8080/uploads/${userData.Avatar}`;
      } else {
        this.user.avatar = `https://api.dicebear.com/7.x/avataaars/svg?seed=${userData.Username}`;
      }
      
      this.initializeWebSocket();
      await this.fetchPostsPrv();
      await this.fetchPosts();
      await this.fetchAllUsers();
      await this.fetchAllowedUsers();
      await this.fetchGroups();
      await this.fetchNotifications();
    } catch (err) {
      console.error("Initialization failed:", err);
      this.$router.push("/login");
    }
  },
  methods: {
    async fetchAllowedUsers() {
      try {
        const res = await fetch("http://localhost:8080/api/postsprivacy", {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          const data = await res.json();
          if (!data.followers) {
            return;
          }
          console.log("Allowed users data:", data);

          this.allowedUsers = data.followers.map((user) => ({
            username: user,
            avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${user}`,
          }));
        } else {
          this.showNotification("Failed to fetch allowed users", "error");
          console.error("Failed to fetch allowed users");
        }
      } catch (error) {
        this.showNotification("Failed to fetch allowed users", "error");
        console.error("Error fetching allowed users:", error);
      }
      console.log("Allowed users:", this.allowedUsers);

    },
    async fetchAllUsers() {
      try {
        
        const allUsersRes = await fetch("http://localhost:8080/api/allusers", {
          method: "GET",
          credentials: "include",
        });
        if (allUsersRes.ok) {
          const allUsersData = await allUsersRes.json();
          if ((!allUsersData) || allUsersData.length === 0) {
            return;
          }
          console.log("All users data:", allUsersData);
          this.otherUsers = allUsersData.map((user) => ({
            name: user.fullname,
            username: user.username,
            avatar: user.avatar
              ? `http://localhost:8080/uploads/${user.avatar}`
              : `https://api.dicebear.com/7.x/avataaars/svg?seed=${user.username}`,
            followed: user.followed || false, 
          }));

          
          this.followedUsers = this.otherUsers.filter(
            (user) => user.followed
          );
          console.log("Followed users:", this.followedUsers);

        }

        
        const chatUsersRes = await fetch("http://localhost:8080/api/openchat", {
          method: "GET",
          credentials: "include",
        });
        if (chatUsersRes.ok) {
          const chatUsersData = await chatUsersRes.json();
          console.log("Chat users data:", chatUsersData);
          if (!chatUsersData || chatUsersData.length === 0) {
            return;
          } else {
            this.chatUsers = chatUsersData.map((user) => ({
              id: user.id,
              fullname: user.full_name,
              name: user.username,
              messages: [],
              avatar:   user.avatar 
                ? `http://localhost:8080/uploads/${user.avatar}`
                : `https://api.dicebear.com/7.x/avataaars/svg?seed=${user.username}`,
            }));
            
            
          }
        } else {
          this.showNotification("Failed to fetch chat users", "error");
          console.error("Failed to fetch chat users");
        }
      } catch (error) {
        this.showNotification("Failed to fetch chat users", "error");
        console.error("Error fetching users:", error);
      }
    },
    initializeWebSocket() {
      this.socket = new WebSocket("ws://localhost:8080/ws");

      this.socket.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.type == "message" ){
          this.showNotification("New message from " + data.username, "success");
        }
        console.log("Received message:", data);

        
        const user = this.chatUsers.find((u) => u.name === data.username);
        if (user) {
          this.showNotification("New message from " + data.username, "success");
          user.messages.push({
            id: Date.now(),
            text: data.message,
            sender: "other",
            timestamp: new Date().toISOString() 
          });
          
          
          if (this.selectedChatUser && this.selectedChatUser.name === data.username) {
            this.$nextTick(() => {
              this.scrollToBottom();
            });
          }
        }
      };
      this.socket.onopen = () => {
        console.log("WebSocket connected");
      };

      this.socket.onerror = (error) => {
        console.error("WebSocket error:", error);
      };

      this.socket.onclose = () => {
        console.log("WebSocket disconnected");
      };
    },
    async fetchPosts() {
      try {
        const res = await fetch("http://localhost:8080/api/getposts", {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          const data = await res.json();

          console.log("data", data);

          
          if (!data) {
            return;
          }
          this.posts = data.map((post) => ({
            ...post,
            authorAvatar: post.Avatar
              ? `http://localhost:8080/uploads/${post.Avatar}`
              : `https://api.dicebear.com/7.x/avataaars/svg?seed=${post.Author}`,
            Image: post.Image, 
            comments: [],
            newComment: "",
            showComments: false,
            commentError: "",
          }));
        } else {
          const errorText = await res.text();
          this.showNotification(errorText || "Failed to fetch posts", "error");
          console.error("Failed to fetch posts:", errorText);
        }
      } catch (error) {
        this.showNotification(error.message || "Failed to fetch posts", "error");
        console.error("Error fetching posts:", error);
      }
    },
    async fetchComments(post) {
      try {
        const res = await fetch(
          `http://localhost:8080/api/getcomments?post_id=${post.Id}`,
          {
            method: "GET",
            credentials: "include",
          }
        );
        if (res.ok) {
          const data = await res.json();
          console.log(data);
          if (!data) {
            return;
          }
          post.comments = data.map((comment) => ({
            ...comment,
            avatar: comment.avatar
              ? `http://localhost:8080/uploads/${comment.avatar}`
              : `https://api.dicebear.com/7.x/avataaars/svg?seed=${comment.author}`,
            image: comment.image ? `http://localhost:8080/uploads/${comment.image}` : null,
          }));

        } else {
          this.showNotification("Failed to fetch comments", "error");
          console.error("Failed to fetch comments for post", post.ID);
          post.commentError = "Failed to load comments.";
        }
      } catch (error) {
        this.showNotification("Failed to fetch comments", "error");
        console.error("Error fetching comments:", error);
      }
    },
    async addComment(post) {
      if (!post.newComment.trim()) {
        post.commentError = "Comment cannot be empty.";
        return;
      }

      
      if (post.newComment.trim().length < 1) {
        post.commentError = "Comment must be at least 1 character long.";
        return;
      }

      if (post.newComment.trim().length > 500) {
        post.commentError = "Comment must not exceed 500 characters.";
        return;
      }

      const formData = new FormData();
      formData.append("post_id", post.Id.toString());
      formData.append("comment", post.newComment.trim());

      
      if (this.commentImages[post.Id] && this.commentImages[post.Id].file) {
        formData.append("image", this.commentImages[post.Id].file);
      }

      try {
        const res = await fetch("http://localhost:8080/api/addcomments", {
          method: "POST",
          credentials: "include",
          body: formData
        });

        if (res.ok) {
          post.commentError = "";
          
          post.newComment = "";
          
          this.commentImages[post.Id] = null;
          this.showNotification("Comment added successfully", "success");
          await this.fetchComments(post);
        } else {
          this.showNotification("Failed to add comment", "error");
          console.error("Failed to add comment");
          post.commentError = "Failed to add comment.";
        }
      } catch (error) {
        this.showNotification("Failed to add comment", "error");
        console.error("Error adding comment:", error);
        post.commentError = "Error adding comment.";
      }
    },
    async submitPost() {
      if (this.newPost.title.length < 1) {
        this.message = "Title must be at least 1 characters long.";
        this.showNotification("Title must be at least 1 characters long.", "error");
        return;
      }

      if (this.newPost.title.length > 50 ) {
        this.message = "Title must not exceed 50 characters.";
        this.showNotification("Title must not exceed 100 characters.", "error");
        return;
      }

      if (this.newPost.content.length < 1) {
        this.message = "Content must be at least 1 characters long.";
        this.showNotification("Content must be at least 10 characters long.", "error");
        return;
      }

      if (this.newPost.content.length > 1000) {
        this.message = "Content must not exceed 1000 characters.";
        this.showNotification("Content must not exceed 1000 characters.", "error");
        return;
      }

      try {
        const formData = new FormData();
        formData.append("title", this.newPost.title);
        formData.append("content", this.newPost.content);
        formData.append("status", this.newPost.status);
        formData.append("allowed_users", this.selectedAllowedUsers.join(","));
        if (this.image) {
          formData.append("image", this.image);
        }

        const res = await fetch("http://localhost:8080/api/posts", {
          method: "POST",
          credentials: "include",
          body: formData
        });
        console.log(this.selectedAllowedUsers);

        if (res.ok) {
          this.newPost = {
            title: "",
            content: "",
            status: "public"
          };
          this.selectedAllowedUsers = [];
          this.image = null;
          this.imagePreview = null;
          this.imageFileName = "";
          this.showNotification("Post created successfully!", "success");
          await this.fetchPosts();
        } else {
          const errorText = await res.text();
          this.showNotification("Failed to create post: "+ errorText, "error");
        }
      } catch (error) {
        this.showNotification("Failed to create post", "error");
        console.error("Post creation failed:", error);
        this.message = "Error submitting post.";
      }
    },
    logout() {
      fetch("http://localhost:8080/api/auth/logout", {
        method: "POST",
        credentials: "include",
      })
        .then((res) => {
          if (res.ok) {
            this.showNotification("Logged out successfully", "success");
            this.$router.push("/login");
          } else {
            this.showNotification("Failed to logout", "error");
            this.$router.push("/login");
          }
        })
        .catch((err) => {
          this.showNotification("wax nta hbit rak mamlogich", "error");
          console.error("Logout error:", err);
          this.$router.push("/login");
        });
    },
    gotoxprofile(username) {
      this.$router.push(`/profile/${username}`);
    },
    async toggleJoin(group) {
      try {
        if (!group || !group.id) {
          console.warn("toggleJoin called with invalid group");
          return;
        }

        const userRes = await fetch("http://localhost:8080/api/info", {
          method: "GET",
          credentials: "include",
        });
        
        if (!userRes.ok) {
          this.showNotification("Failed to get user information", "error");
          return;
        }
        
        const userData = await userRes.json();

        let endpoint;
        let action;
        let requestBody = {};

        if (group.member_status === 'accepted') {
          endpoint = "http://localhost:8080/api/removememberfromgroup";
          action = "leave";
          requestBody = {
            group_id: group.id,
            user_id: userData.id  
          };
        } else if (group.member_status === 'pending') {
          endpoint = "http://localhost:8080/api/cancelgrouprequest";
          action = "cancel request from";
          requestBody = {
            group_id: group.id
          };
        } else if (group.member_status === 'invited') {
          endpoint = "http://localhost:8080/api/acceptgroupinvite";
          action = "accept invitation to";
          requestBody = {
            group_id: group.id,
            action: 'accept'
          };
        } else {
          endpoint = "http://localhost:8080/api/requesttojoingroup";
          action = "join";
          requestBody = {
            group_id: group.id
          };
        }

        const response = await fetch(endpoint, {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(requestBody),
        });

        if (response.ok) {
          if (group.member_status === 'accepted') {
            group.member_status = null;
            this.showNotification("Left group successfully", "success");
          } else if (group.member_status === 'pending') {
            group.member_status = null;
            this.showNotification("Request cancelled successfully", "success");
          } else if (group.member_status === 'invited') {
            group.member_status = 'accepted';
            this.showNotification("Invitation accepted successfully", "success");
          } else {
            group.member_status = 'pending';
            this.showNotification("Join request sent successfully", "success");
          }

          await this.fetchGroups();
        } else {
          const errorText = await response.text();
          console.error(`Failed to ${action} group: ${errorText}`);
          this.showNotification(`Failed to ${action} group`, "error");
        }
      } catch (error) {
        console.error("Error updating group membership:", error);
        this.showNotification("Failed to update group membership", "error");
      }
    },
    goToProfile() {
      this.$router.push(`/profile/${this.user.username}`);
    },
    goToMyGroups() {
      this.$router.push('/mygroups');
    },
    goToGroup(groupId) {
      this.$router.push(`/group/${groupId}`);
    },
    toggleChat() {
      this.isChatExpanded = !this.isChatExpanded;
    },
    selectChatUser(user) {
      console.log("user", user.name);
      this.selectedChatUser = user.name;
      console.log("selectedChatUser", this.selectedChatUser);
      this.fetchMessages(this.selectedChatUser);
    },
    async fetchMessages(user) {
      console.log("Fetching messages for use11:", user);
      
      try {
        const res = await fetch(
          `http://localhost:8080/api/getmessages?sender=${this.user.username}&receiver=${user}`,
          {
            method: "GET",
            credentials: "include",
          }
        );

        if (res.ok) {
          const data = await res.json();
          console.log("Fetched messages:", data);

          const chatUser = this.chatUsers.find(u => u.name === user);
          if (!chatUser) {
            console.error("Chat user not found");
            return;
          }

          
          chatUser.messages = [];

          
          if (data && Array.isArray(data)) {
            chatUser.messages = data.map(msg => ({
              id: Date.now() + Math.random(), 
              text: msg.message,
              sender: msg.username === this.user.username ? "self" : "other",
              timestamp: msg.time
            }));

            
            chatUser.messages.sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp));
          }

          
          this.selectedChatUser = chatUser;
          
          
          this.$nextTick(() => {
            this.scrollToBottom();
          });
        } else {
          console.error("Failed to fetch messages");
          this.showNotification("Failed to load messages", "error");
        }
      } catch (error) {
        console.error("Error fetching messages:", error);
        this.showNotification("Error loading messages", "error");
      }
    },
    scrollToBottom() {
      const messagesList = document.querySelector('.chat-messages-list');
      if (messagesList) {
        messagesList.scrollTop = messagesList.scrollHeight;
      }
    },
    sendMessage() {
      if (this.newMessage.trim() && this.selectedChatUser) {
        this.currentMessage.type = "message";
        this.currentMessage.receiver = this.selectedChatUser.name;
        this.currentMessage.username = this.user.username;
        this.currentMessage.message = this.newMessage;

        
        if (this.socket) {
          this.socket.send(JSON.stringify(this.currentMessage));
        }

        
        this.selectedChatUser.messages.push({
          text: this.newMessage,
          sender: "self",
          timestamp: Date.now()
        });

        this.newMessage = "";
        
        
        this.$nextTick(() => {
          this.scrollToBottom();
        });
      }
    },
    formatTimestamp(timestamp) {
      return new Date(timestamp).toLocaleString("en-US", {
        month: "short",
        day: "numeric",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    },
    toggleComments(post) {
      post.showComments = !post.showComments;
      if (post.showComments && post.comments.length === 0) {
        this.fetchComments(post);
      }
    },
    getMessageClass(msg) {
      return msg.sender === "self" ? "self" : "other";
    },
    async fetchGroups() {
      try {
        const response = await fetch("http://localhost:8080/api/getgroups", {
          method: "GET",
          credentials: "include",
        });

        if (response.ok) {
          const data = await response.json();
          if (!data) {
            return;
          }
          console.log("Fetched groups data:", data);
          this.groups = data.map((group) => ({
            id: group.id,
            owner: group.creator_id,
            is_owner: group.is_owner,
            name: group.title,
            description: group.description,
            member_status: group.member_status
          }));
          console.log("Fetched groups:", this.groups);

          
          for (const group of this.groups) {
            await this.checkGroupMembership(group.id);
          }
        } else {
          this.showNotification("Failed to fetch groups", "error");
          console.error("Failed to fetch groups");
        }
      } catch (error) {
        this.showNotification("Failed to fetch groups", "error");
        console.error("Error fetching groups:", error);
      }
    },
    async checkGroupMembership(groupId) {
      try {
        if (!groupId) {
          console.warn("checkGroupMembership called with invalid groupId");
          return;
        }
        const response = await fetch(
          `http://localhost:8080/api/ismember?group_id=${groupId}`,
          {
            method: "GET",
            credentials: "include",
          }
        );


        if (!response.ok) {
          const errorText = await response.text();
          console.error(`Failed to check group membership: ${errorText}`);
          this.showNotification("Failed to check group membership", "error");
          return;
        }

        const data = await response.json();
        if (data && typeof data.is_member !== 'undefined') {
          this.groupMembership[groupId] = data.is_member;
        } else {
          console.warn("Received invalid data format from server");
          this.groupMembership[groupId] = false; 
        }
      } catch (error) {
        console.error("Error checking group membership:", error);
        this.showNotification("Failed to check group membership", "error");
      }
    },
    async createGroup() {
      try {
        const userRes = await fetch("http://localhost:8080/api/info", {
          method: "GET",
          credentials: "include",
        });
        
        if (!userRes.ok) {
          this.showNotification("Failed to get user information", "error");
          return;
        }
        
        const userData = await userRes.json();
        this.newGroup.creator_id = userData.id; 

        const response = await fetch("http://localhost:8080/api/creategroups", {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(this.newGroup),
        });

        if (response.ok) {
          this.showNotification("Group created successfully", "success");
          const data = await response.json();
          this.groups.push({
            id: data.id,
            name: data.title,
            owner: this.user.username,
            is_owner: true,
            member_status: 'accepted',
            description: data.description,
          });
          this.showCreateGroup = false;
          this.newGroup = {
            title: "",
            description: "",
            creator_id: null,
          };
        } else {
          this.showNotification("Failed to create group", "error");
          console.error("Failed to create group");
        }
      } catch (error) {
        this.showNotification("Failed to create group", "error");
        console.error("Error creating group:", error);
      }
    },
    async fetchPostsPrv() {
      try {
        const res = await fetch("http://localhost:8080/api/postsprv", {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          const data = await res.json();
          if (!data || data.length === 0) {
            console.warn("No posts privacy data found");
            return;
          }
          const map = {};
          data.forEach(item => {
            const postId = item.post_id.toString();
            if (!map[postId]) {
              map[postId] = [];
            }
            map[postId].push(item.username);
          });
          this.postsPrvMap = map;
          console.log("postsPrvMap", this.postsPrvMap);
        }
      } catch (error) {
        console.error("Error fetching postsprv:", error);
      }
    },
    handleNotifClick() {
      this.showNotifications = !this.showNotifications;
      if (this.showNotifications) {
        this.notifications.forEach(notif => {
          if (!notif.is_read) {
            this.markNotificationAsRead(notif.id);
          }
        });
      }
    },
    handleNotifClose(e) {
      if (!this.showNotifications) return;
      const popup = this.$refs.notifPopup;
      const btn = this.$refs.notifBtn;
      if (
        (!popup || !btn) ||
        (!popup.contains(e.target) && !btn.contains(e.target))
      ) {
        this.showNotifications = false;
      }
    },
    async fetchNotifications() {
      try {
        const res = await fetch("http://localhost:8080/api/notifications", {
          method: "GET",
          credentials: "include",
        });
        if (res.ok) {
          const data = await res.json();
          console.log("sassasaassa", data);

          this.notifications = Array.isArray(data)
            ? data.map(notif => ({
              id: notif.id,
              sender: notif.sender_username || 'Unknown',
              text: notif.content || '',
              time: notif.time || new Date().toISOString(),
              type: notif.type || 'notification',
              is_read: notif.is_read || false
            }))
            : [];
          
          
          this.unreadNotificationCount = this.notifications.filter(n => !n.is_read).length;
        } else {
          this.notifications = [];
          this.unreadNotificationCount = 0;
        }
      } catch (e) {
        console.error('Error fetching notifications:', e);
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
    handleImageUpload(event, postId = null) {
      const file = event.target.files[0];
      if (file) {
      const allowedTypes = ["image/jpeg", "image/png", "image/gif"];
      if (!allowedTypes.includes(file.type)) {
          this.showNotification('Only JPEG, PNG, and GIF images are allowed', 'error');
          event.target.value = ''; 
          return;
        }
        
        const maxSize = 2 * 1024 * 1024; 
        if (file.size > maxSize) {
          this.showNotification('Image file size must be less than 2MB', 'error');
          event.target.value = ''; 
          return;
        }

        if (postId) {
          
          this.commentImages[postId] = {
            file: file,
            preview: URL.createObjectURL(file),
            fileName: file.name
          };
        } else {
          
          this.image = file;
          this.imageFileName = file.name;
          this.imagePreview = URL.createObjectURL(file);
        }
      } else {
        if (postId) {
          
          this.commentImages[postId] = null;
        } else {
          
          this.image = null;
          this.imagePreview = null;
          this.imageFileName = "";
        }
      }
    },
    formatTime(timestamp) {
      if (!timestamp) return '';
      const date = new Date(timestamp);
      const now = new Date();
      const diff = now - date;

      
      if (diff < 60000) {
        return 'just now';
      }
      
      if (diff < 3600000) {
        const minutes = Math.floor(diff / 60000);
        return `${minutes}m ago`;
      }
      
      if (diff < 86400000) {
        const hours = Math.floor(diff / 3600000);
        return `${hours}h ago`;
      }
      
      if (diff < 604800000) {
        const days = Math.floor(diff / 86400000);
        return `${days}d ago`;
      }
      
      return date.toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric'
      });
    },
    async fetchChatUsers() {
      try {
        const chatUsersRes = await fetch("http://localhost:8080/api/openchat", {
          method: "GET",
          credentials: "include",
        });
        if (chatUsersRes.ok) {
          const chatUsersData = await chatUsersRes.json();
          console.log("Chat users data:", chatUsersData);
          if (!chatUsersData || chatUsersData.length === 0) {
            return;
          } else {
            this.chatUsers = chatUsersData.map((user) => ({
              id: user.id,
              fullname: user.full_name,
              name: user.username,
              messages: [],
              avatar: user.avatar 
                ? `http://localhost:8080/uploads/${user.avatar}`
                : `https://api.dicebear.com/7.x/avataaars/svg?seed=${user.username}`,
            }));
          }
        } else {
          this.showNotification("Failed to fetch chat users", "error");
          console.error("Failed to fetch chat users");
        }
      } catch (error) {
        console.error("Error fetching chat users:", error);
        this.showNotification("Error fetching chat users", "error");
      }
    },

    async handleChatClick() {
      if (!this.isChatExpanded) {
        await this.fetchChatUsers();
        this.selectedChatUser = null; 
      }
      this.isChatExpanded = !this.isChatExpanded;
    },
    setupNotificationWebSocket() {
      
      notificationWebSocket.onNotification('forum-page', (notification) => {
        console.log('Received real-time notification:', notification);
        
        
        this.fetchNotifications();
      });
    },
    async fetchNotificationCount() {
      try {
        const res = await fetch('http://localhost:8080/api/notifications', {
          method: 'GET',
          credentials: 'include'
        });
        if (res.ok) {
          const data = await res.json();
          
          this.unreadNotificationCount = (data || []).filter(n => !n.is_read).length;
        }
      } catch (err) {
        console.error('Failed to fetch notification count:', err);
      }
    },
    showFullImage(imageUrl) {
      this.selectedImage = imageUrl;
      this.showImageModal = true;
    },
    closeImageModal() {
      this.showImageModal = false;
      this.selectedImage = null;
    },
  },
  mounted() {
    
    this.setupNotificationWebSocket();
    
    this.fetchPosts();
    this.fetchNotifications();
    this.fetchPostsPrv();
    document.addEventListener('click', this.handleNotifClose);
  },
  beforeUnmount() {
    
    notificationWebSocket.removeNotificationHandler('forum-page');
    document.removeEventListener('click', this.handleNotifClose);
  },
  computed: {
  },
};
</script>

<style scoped>
/* Layout */
.forum-layout {
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
}

.sidebar-icon.active,
.sidebar-icon:hover {
  background: #35395a;
  opacity: 1;
}

/* Main Area */
.main-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 2rem 2rem 2rem 1.5rem;
  min-width: 0;
  margin-left: 70px;
}

/* Top Avatar Bar */
.avatar-bar {
  width: 100%;
  margin-bottom: 1.5rem;
  overflow-x: auto;
  padding-bottom: 0.5rem;
}

.avatar-list {
  display: flex;
  gap: 1.5rem;
  align-items: center;
  padding: 0.5rem 0;
  min-width: min-content;
}

.avatar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
  min-width: 56px;
}

.avatar-item img {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  border: 2px solid #fff;
  box-shadow: 0 2px 8px rgba(35, 38, 58, 0.08);
  margin-bottom: 0.25rem;
  object-fit: cover;
  transition: transform 0.2s;
}

.avatar-item img:hover {
  transform: scale(1.08);
}

.avatar-name {
  font-size: 0.8rem;
  color: #23263a;
  opacity: 0.7;
  text-align: center;
  max-width: 60px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.avatar-item.see-all .see-all-btn {
  width: 48px;
  height: 48px;
  background: #f3f4f8;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #23263a;
  font-size: 0.9rem;
  font-weight: 600;
  border: 1px solid #e5e7eb;
  margin-bottom: 0.25rem;
  cursor: pointer;
}

/* Custom scrollbar for avatar bar */
.avatar-bar::-webkit-scrollbar {
  height: 6px;
}

.avatar-bar::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.avatar-bar::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.avatar-bar::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* Create Post Card */
.create-post-card {
  display: flex;
  align-items: flex-start;
  background: #fff;
  border-radius: 1.25rem;
  box-shadow: 0 4px 24px rgba(35, 38, 58, 0.08);
  padding: 1.5rem 2rem;
  margin-bottom: 2rem;
  gap: 1.5rem;
}

.create-post-card .profile-pic {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  object-fit: cover;
  margin-top: 0.5rem;
}

.create-post-form {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.create-post-form input,
.create-post-form textarea {

  border: none;
  background: #f3f4f8;
  border-radius: 0.75rem;
  padding: 0.9rem 1.2rem;
  font-size: 1rem;
  margin-bottom: 0.2rem;
  resize: none;
}

.create-post-form textarea {
  min-height: 70px;
  resize: vertical;
}

.create-post-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
}


.status-select select {
  background: #f3f4f8;
  border: none;
  border-radius: 0.5rem;
  padding: 0.3rem 1rem;
  font-size: 0.95rem;
  color: #23263a;
}

.allowed-users-select {
  display: flex;
  gap: 0.7rem;
  flex-wrap: wrap;
  margin-bottom: 0.5rem;
}

.allowed-user-pill {
  display: flex;
  align-items: center;
  background: #f3f4f8;
  border-radius: 2rem;
  padding: 0.25rem 1rem 0.25rem 0.4rem;
  font-size: 0.97rem;
  color: #23263a;
  cursor: pointer;
  gap: 0.5rem;
  box-shadow: 0 1px 4px rgba(35, 38, 58, 0.06);
  border: 1.5px solid #e5e7eb;
  transition: border 0.18s, box-shadow 0.18s, background 0.18s;
  position: relative;
}

.allowed-user-pill:hover,
.allowed-user-pill:focus-within {
  border: 1.5px solid #4f46e5;
  background: #e0e7ff;
  box-shadow: 0 2px 8px rgba(79, 70, 229, 0.10);
}

.allowed-user-pill input[type="checkbox"] {
  display: none;
}

.custom-checkbox {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 22px;
  height: 22px;
  margin-right: 0.1rem;
}

.allowed-user-avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  object-fit: cover;
  border: 1.5px solid #e5e7eb;
  background: #fff;
}

.allowed-user-name {
  font-weight: 500;
  color: #23263a;
  font-size: 0.97rem;
  margin-left: 0.1rem;
}

.post-btn {
  background: linear-gradient(90deg, #23263a 60%, #4f46e5 100%);
  color: #fff;
  border: none;
  border-radius: 0.7rem;
  padding: 0.7rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  margin-top: 0.5rem;
  align-self: flex-end;
  transition: background 0.2s;
}

.post-btn:hover {
  background: linear-gradient(90deg, #4f46e5 60%, #23263a 100%);
}

/* Posts Grid */
.posts-grid {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  margin-top: 1.5rem;
  width: 100%;
  max-width: 900px;
  margin-left: auto;
  margin-right: auto;
}

.post-card {
  background: #fff;
  border-radius: 1.25rem;
  box-shadow: 0 4px 24px rgba(35, 38, 58, 0.08);
  padding: 1.5rem 1.8rem 1.2rem 1.8rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 100%;
  overflow-wrap: break-word;
  word-wrap: break-word;
  word-break: break-word;
  height: auto;
}

.post-image {
  width: 100%;
  max-height: 400px;
  overflow: hidden;
  border-radius: 0.8rem;
  margin-bottom: 0.7rem;
  display: flex;
  border-radius: 0.5rem;
  align-items: center;
  justify-content: center;
}

.post-image img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  border-radius: 0.8rem;
}

.post-header {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.post-author-pic {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
}

.post-title {
  font-size: 1.3rem;
  color: #23263a;
  font-weight: 600;
  margin: 0.2rem 0;
}

.post-content {
  color: #4b5563;
  font-size: 1.1rem;
  margin-bottom: 0.2rem;
  word-wrap: break-word;
  white-space: pre-wrap;
  overflow-wrap: break-word;
  max-width: 100%;
  line-height: 1.6;
  height: auto;
}

.post-actions {
  display: flex;
  gap: 1.2rem;
  margin-top: 0.2rem;
}

.icon-btn {
  font-size: 1.2rem;
  color: #b0b3c6;
  cursor: pointer;
  transition: color 0.2s;
}

.icon-btn:hover {
  color: #4f46e5;
}

.no-posts {
  color: #b0b3c6;
  text-align: center;
  font-size: 1.1rem;
  grid-column: 1/-1;
}

/* Right Sidebar */
.rightbar {
  width: 400px;
  background: #fff;
  border-radius: 0 1.5rem 1.5rem 0;
  box-shadow: -2px 0 16px rgba(35, 38, 58, 0.08);
  padding: 2rem 1.5rem 2rem 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}


.invite-group-btn {
  background: #f3f4f8;
  color: #23263a;
  border: none;
  border-radius: 0.5rem;
  padding: 0.3rem 0.6rem;
  margin-right: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: background 0.2s, color 0.2s;
  position: relative;
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

.notif-btn:hover,
.invite-group-btn:hover {
  background: #e0e7ff;
  color: #4f46e5;
}

.logout-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: #f3f4f8;
  color: #23263a;
  border: none;
  border-radius: 0.5rem;
  padding: 0.4rem 1rem;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}

.logout-btn:hover {
  background: #e5e7eb;
}

.create-group-btn {
  background: #4f46e5;
  color: #fff;
  border: none;
  border-radius: 0.5rem;
  padding: 0.2rem 0.5rem;
  margin-right: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: background 0.2s;
}

.create-group-btn:hover {
  background: #23263a;
}

.groups-header {
  display: flex;
  align-items: center;
  gap: 0.3rem;
  justify-content: space-between;
}

.groups-header span {
  flex: 1;
}

.groups-header .see-all-link {
  flex: unset;
}

.groups-section {
  margin-bottom: 2rem;
  height: 300px;
  overflow-y: auto;
  padding-right: 10px;
}

.groups-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.group-list-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #f3f4f8;
  border-radius: 0.7rem;
  padding: 0.7rem 1rem;
  gap: 0.7rem;
}

.group-info {
  display: flex;
  align-items: center;
  gap: 0.7rem;
}

.group-avatar {
  font-size: 1.5rem;
  color: #4f46e5;
  background: #e0e7ff;
  border-radius: 50%;
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.group-details {
  display: flex;
  flex-direction: column;
}

.group-title {
  font-size: 1rem;
  font-weight: 600;
  color: #23263a;
}

.group-desc {
  font-size: 0.85rem;
  color: #6b7280;
  max-width: 120px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.pending-btn {
  background: #ad8f0a;
}

.follow-btn,
.pending-btn {
  color: #fff;
  border: none;
  border-radius: 0.5rem;
  padding: 0.4rem 1.1rem;
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.follow-btn {
  background: #4f46e5;
}

.follow-btn:hover {
  background: #23263a;
}

.followed-btn {
  background: #e5e7eb;
  color: #4f46e5;
  border: none;
  border-radius: 0.5rem;
  padding: 0.4rem 1.1rem;
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
}

.users-section {
  margin-top: 2rem;
  height: 300px;
  overflow-y: auto;
  padding-right: 10px;
}

.users-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
  font-size: 1.1rem;
  font-weight: 600;
  color: #23263a;
}

.users-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.user-list-item {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  background: #f3f4f8;
  border-radius: 0.7rem;
  padding: 0.7rem 1rem;
}

.mini-profile-pic {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
  cursor: pointer;
}

.user-name {
  font-size: 1rem;
  color: #23263a;
  font-weight: 500;
  cursor: pointer;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(35, 38, 58, 0.45);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  transition: background 0.3s;
}

.create-group-modal {
  background: #fff;
  border-radius: 1.2rem;
  box-shadow: 0 8px 32px rgba(35, 38, 58, 0.18);
  padding: 2.5rem 2.5rem 2rem 2.5rem;
  min-width: 340px;
  max-width: 95vw;
  animation: modalPopIn 0.25s cubic-bezier(.4, 2, .6, 1) both;
}

@keyframes modalPopIn {
  0% {
    transform: scale(0.95) translateY(30px);
    opacity: 0;
  }

  100% {
    transform: scale(1) translateY(0);
    opacity: 1;
  }
}

.create-group-modal h3 {
  color: #4f46e5;
  margin-bottom: 1.5rem;
  text-align: center;
  font-size: 1.3rem;
  font-weight: 700;
}

.create-group-modal .form-group label {
  color: #23263a;
  font-weight: 600;
}

.create-group-modal .form-group input,
.create-group-modal .form-group textarea {
  background: #f3f4f8;
  border: none;
  border-radius: 0.7rem;
  padding: 0.8rem 1.1rem;
  font-size: 1rem;
  margin-bottom: 0.7rem;
  width: 100%;
  box-sizing: border-box;
}

.create-group-modal .form-group textarea {
  min-height: 80px;
  resize: vertical;
}

.create-group-modal .modal-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.2rem;
}

.create-group-modal .submit-btn {
  background: #4f46e5;
  color: #fff;
  padding: 0.7rem 1.5rem;
  border: none;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.create-group-modal .submit-btn:hover {
  background: #23263a;
}

.create-group-modal .cancel-btn {
  background: #e5e7eb;
  color: #23263a;
  padding: 0.7rem 1.5rem;
  border: none;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.create-group-modal .cancel-btn:hover {
  background: #cbd5e1;
}

/* Comments Section */
.comments-section {
  margin-top: 1.1rem;
  padding: 1.1rem;
  background: #f8fafc;
  border-radius: 0.7rem;
  border: 1px solid #e5e7eb;
}

.no-comments {
  text-align: center;
  color: #6b7280;
  font-size: 0.95rem;
  padding: 1rem 0;
  background: #fff;
  border-radius: 0.5rem;
  border: 1px dashed #e5e7eb;
}

.comment {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  background: #fff;
  border-radius: 0.7rem;
  padding: 1rem;
  margin-bottom: 0.7rem;
  border: 1px solid #e5e7eb;
  transition: transform 0.2s, box-shadow 0.2s;
}

.comment:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(35, 38, 58, 0.06);
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  margin-bottom: 0.5rem;
}

.comment-author-pic {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e5e7eb;
}

.comment-header h5 {
  color: #23263a;
  font-size: 0.95rem;
  margin: 0;
  font-weight: 600;
}

.comment-timestamp {
  color: #6b7280;
  font-size: 0.85rem;
  margin: 0;
}

.comment-content {
  color: #4b5563;
  font-size: 0.95rem;
  line-height: 1.5;
  margin: 0.5rem 0;
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

.comment-form {
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
  margin-top: 1rem;
  background: #fff;
  padding: 1rem;
  border-radius: 0.7rem;
  border: 1px solid #e5e7eb;
}

.comment-input-row {
  display: flex;
  align-items: flex-end;
  gap: 0.7rem;
}

.comment-form textarea {
  flex: 1;
  padding: 0.8rem 1rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.7rem;
  font-size: 0.95rem;
  background: #f8fafc;
  min-height: 42px;
  resize: vertical;
  transition: border-color 0.2s, background 0.2s;
}

.comment-form textarea:focus {
  border-color: #4f46e5;
  background: #fff;
  outline: none;
}

.comment-form input[type="file"] {
  display: none;
}

.comment-form .file-upload-label {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: #f3f4f8;
  border-radius: 0.5rem;
  padding: 0.5rem 0.8rem;
  font-size: 0.9rem;
  color: #23263a;
  cursor: pointer;
  transition: background 0.2s;
}

.comment-form .file-upload-label:hover {
  background: #e0e7ff;
}

.comment-form .file-upload-label svg {
  width: 18px;
  height: 18px;
  color: #4f46e5;
}

.comment-submit-btn {
  background: #4f46e5;
  color: #fff;
  border: none;
  border-radius: 0.5rem;
  padding: 0.5rem 0.8rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.2s;
}

.comment-submit-btn:hover {
  background: #23263a;
}

.comment-error {
  color: #ef4444;
  font-size: 0.9rem;
  margin-top: 0.3rem;
  padding: 0.5rem;
  background: #fee2e2;
  border-radius: 0.5rem;
  border: 1px solid #fecaca;
}

.comment-img {
  max-width: 100%;
  max-height: 300px;
  width: auto;
  height: auto;
  object-fit: contain;
  border-radius: 0.5rem;
  margin-top: 0.7rem;
  border: 1px solid #e5e7eb;
}

.group-clickable {
  cursor: pointer;
  transition: background 0.18s;
}

.group-clickable:hover {
  background: #e0e7ff;
}

/* Chat Floating Action Button */
.chat-fab {
  position: fixed;
  bottom: 2.2rem;
  right: 2.2rem;
  z-index: 3000;
  background: #4f46e5;
  color: #fff;
  border: none;
  border-radius: 50%;
  width: 60px;
  height: 60px;
  box-shadow: 0 4px 16px rgba(35, 38, 58, 0.18);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  cursor: pointer;
  transition: background 0.2s, box-shadow 0.2s;
}

.chat-fab:hover {
  background: #23263a;
  box-shadow: 0 8px 32px rgba(35, 38, 58, 0.22);
}

/* Chat Modal */
.chat-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(35, 38, 58, 0.25);
  z-index: 4000;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chat-modal {
  background: #fff;
  border-radius: 1.1rem;
  box-shadow: 0 8px 32px rgba(35, 38, 58, 0.18);
  width: 370px;
  max-width: 95vw;
  padding: 1.2rem 1.2rem 1.2rem 1.2rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  animation: modalPopIn 0.22s cubic-bezier(.4, 2, .6, 1) both;
}

.chat-modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 1.2rem;
  font-weight: 600;
  color: #4f46e5;
  margin-bottom: 0.2rem;
}

.chat-modal-close {
  background: none;
  border: none;
  font-size: 2rem;
  color: #b0b3c6;
  cursor: pointer;
  transition: color 0.2s;
}

.chat-modal-close:hover {
  color: #ef4444;
}

.chat-users-list {
  background: #f8fafc;
  border-radius: 0.7rem;
  padding: 0.7rem 0.5rem;
  margin-bottom: 0.2rem;
}

.chat-users-title {
  font-size: 1rem;
  color: #23263a;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.chat-users-list ul {
  list-style: none;
  padding: 0;
  margin: 0;
  max-height: 120px;
  overflow-y: auto;
}

.chat-users-list li {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  padding: 0.4rem 0.2rem;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: background 0.18s;
}

.chat-users-list li.selected,
.chat-users-list li:hover {
  background: #e0e7ff;
}

.chat-user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
}

.chat-messages-area {
  background: #f8fafc;
  border-radius: 0.7rem;
  padding: 0.7rem 0.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.chat-messages-title {
  font-size: 1rem;
  color: #4f46e5;
  font-weight: 600;
  margin-bottom: 0.3rem;
}

.chat-messages-list {
  flex: 1;
  max-height: 180px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
  scroll-behavior: smooth;
  padding-right: 0.5rem;
}

.chat-messages-list::-webkit-scrollbar {
  width: 6px;
}

.chat-messages-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.chat-messages-list::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.chat-messages-list::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

.chat-msg {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  background: #fff;
  border-radius: 0.5rem;
  padding: 0.4rem 0.7rem;
  font-size: 0.97rem;
  box-shadow: 0 1px 4px rgba(35, 38, 58, 0.06);
  max-width: 80%;
}

.chat-msg.self {
  align-self: flex-end;
  background: #4f46e5;
  color: #fff;
}

.chat-msg.other {
  align-self: flex-start;
  background: #e5e7eb;
  color: #23263a;
}

.chat-msg-time {
  font-size: 0.75rem;
  color: #b0b3c6;
  margin-top: 0.1rem;
  align-self: flex-end;
}

.chat-send-form {
  display: flex;
  gap: 0.4rem;
  margin-top: 0.3rem;
}

.chat-send-form input {
  flex: 1;
  padding: 0.5rem 0.8rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  font-size: 0.97rem;
  background: #fff;
}

.chat-send-btn {
  background: #4f46e5;
  color: #fff;
  border: none;
  border-radius: 0.5rem;
  padding: 0.4rem 0.7rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.2s;
}

.chat-send-btn:hover {
  background: #23263a;
}

.chat-no-user {
  color: #b0b3c6;
  text-align: center;
  font-size: 1rem;
  padding: 1.2rem 0;
}

/* Custom scrollbar */
.groups-section::-webkit-scrollbar {
  width: 6px;
}

.groups-section::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.groups-section::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.groups-section::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* Custom scrollbar for users section */
.users-section::-webkit-scrollbar {
  width: 6px;
}

.users-section::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.users-section::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.users-section::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
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

.notif-avatar img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e5e7eb;
  flex-shrink: 0;
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

.owner-badge {
  font-size: 0.75rem;
  color: #4f46e5;
  background: #e0e7ff;
  padding: 0.1rem 0.4rem;
  border-radius: 0.3rem;
  margin-top: 0.2rem;
  display: inline-block;
}

.create-post-form input[type="file"] {
  display: none;
}

.create-post-form .file-upload-label {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: #f3f4f8;
  border-radius: 0.75rem;
  padding: 0.7rem 1.2rem;
  font-size: 0.95rem;
  color: #23263a;
  cursor: pointer;
  transition: background 0.2s;
  margin-bottom: 0.5rem;
}

.create-post-form .file-upload-label:hover {
  background: #e0e7ff;
}

.create-post-form .file-upload-label svg {
  width: 20px;
  height: 20px;
  color: #4f46e5;
}

.create-post-form .image-preview {
  max-width: 300px;
  max-height: 300px;
  width: auto;
  height: auto;
  object-fit: contain;
  border-radius: 0.75rem;
  margin-top: 0.5rem;
  display: block;
  background: #f3f4f8;
}

.file-name-display {
  font-size: 0.9rem;
  color: #4f46e5;
  background: #e0e7ff;
  padding: 0.5rem 0.8rem;
  border-radius: 0.5rem;
  margin-top: 0.5rem;
  display: inline-block;
}

.post-image {
  object-fit: cover;
  width: fit-content;
  height: 300px;
}

@media (max-width: 768px) {
  .post-card {
    width: 100%;
  }
  .post-image img {
    height: 250px;
  }
}

@media (max-width: 900px) {
  .posts-grid {
    max-width: 100%;
    padding: 0 1rem;
  }
}

/* Remove the grid layout for larger screens */
@media (min-width: 768px) {
  .posts-grid {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }
}

/* Responsive Design */
@media (max-width: 1024px) {
  .forum-layout {
    flex-direction: column;
  }

  .sidebar {
    width: 100%;
    flex-direction: row;
    padding: 1rem;
    border-radius: 0;
    justify-content: center;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: auto;
    height: auto;
    z-index: 1000;
  }

  .sidebar-icons {
    flex-direction: row;
    justify-content: center;
    gap: 1.5rem;
    position: relative;
  }

  .main-area {
    margin-left: 0;
    margin-top: 80px;
    padding: 1rem;
  }

  .rightbar {
    width: 100%;
    border-radius: 0;
    margin-top: 1rem;
  }

  .posts-grid {
    max-width: 100%;
  }

  .chat-fab {
    bottom: 1.5rem;
    right: 1.5rem;
  }

  .notif-popup {
    left: 50%;
    transform: translateX(-50%);
    top: 80px;
  }
}

@media (max-width: 768px) {
  .sidebar-icons {
    gap: 1rem;
  }

  .sidebar-icon {
    width: 40px;
    height: 40px;
  }

  .create-post-card {
    padding: 1rem;
  }

  .post-card {
    padding: 1rem;
  }

  .post-image img {
    height: 200px;
  }

  .post-title {
    font-size: 1.2rem;
  }

  .post-content {
    font-size: 1rem;
  }

  .chat-modal {
    width: 95vw;
    height: 90vh;
  }

  .groups-section,
  .users-section {
    height: 250px;
  }
}

@media (max-width: 480px) {
  .sidebar-icons {
    gap: 0.8rem;
  }

  .sidebar-icon {
    width: 36px;
    height: 36px;
  }

  .post-header {
    gap: 0.7rem;
  }

  .post-author-pic {
    width: 40px;
    height: 40px;
  }

  .post-image img {
    height: 180px;
  }

  .create-post-form input,
  .create-post-form textarea {
    font-size: 0.95rem;
  }

  .chat-modal {
    padding: 1rem;
  }

  .chat-messages-list {
    max-height: 150px;
  }
}

/* Ensure content is properly centered on mobile */
@media (max-width: 1024px) {
  .main-area {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .avatar-bar {
    width: 100%;
    max-width: 900px;
  }

  .create-post-card {
    width: 100%;
    max-width: 900px;
  }
}

/* Image Modal Styles */
.image-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.85);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 5000;
  animation: fadeIn 0.2s ease;
  padding: 2rem;
}

.image-modal-content {
  position: relative;
  max-width: 90vw;
  max-height: 85vh;
  animation: zoomIn 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f8;
  border-radius: 0.5rem;
  padding: 1rem;
}

.image-modal-content img {
  max-width: 100%;
  max-height: 85vh;
  width: auto;
  height: auto;
  object-fit: contain;
  border-radius: 0.5rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.close-modal-btn {
  position: absolute;
  top: -40px;
  right: 0;
  background: none;
  border: none;
  color: white;
  font-size: 2rem;
  cursor: pointer;
  padding: 0.5rem;
  transition: opacity 0.2s;
  z-index: 1;
}

.close-modal-btn:hover {
  opacity: 0.8;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes zoomIn {
  from { transform: scale(0.95); }
  to { transform: scale(1); }
}

/* Add cursor pointer to clickable images */
.post-image img,
.comment-image,
.create-post-form .image-preview {
  cursor: pointer;
  transition: opacity 0.2s;
}

.post-image img:hover,
.comment-image:hover,
.create-post-form .image-preview:hover {
  opacity: 0.9;
}
</style>