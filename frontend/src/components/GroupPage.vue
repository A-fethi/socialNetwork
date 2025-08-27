<template>
  <div class="group-layout">
    <!-- Notifications -->
    <!-- Add notification component at the top level -->
    <div v-if="notification.visible" :class="['notification', `notification-${notification.type}`]">
      {{ notification.message }}
    </div>

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
        <div class="sidebar-icon" title="Profile" @click="goToProfile">
          <svg width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <circle cx="12" cy="8" r="4" />
            <path d="M4 21v-2a4 4 0 014-4h8a4 4 0 014 4v2" />
          </svg>
        </div>
        <!-- Groups Icon -->
        <div class="sidebar-icon active" title="Groups" @click="goToMyGroups">
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
              d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 01-2.83 2.83l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-4 0v-.09a1.65 1.65 0 00-1-1.51 1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83-2.83l.06-.06a1.65 1.65 0 00.33-1.82 1.65 1.65 0 00-1.51-1H3a2 2 0 010-4h.09a1.65 1.65 0 001.51-1 1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 012.83-2.83l.06.06a1.65 1.65 0 001.82.33h.09A1.65 1.65 0 008.91 3H9a2 2 0 014 0v.09a1.65 1.65 0 001 1.51h.09a1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 2.83l-.06.06a1.65 1.65 0 00-.33 1.82v.09a1.65 1.65 0 001.51 1H21a2 2 0 010 4h-.09a1.65 1.65 0 00-1.51 1z" />
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
            <li v-if="!notifications || notifications.length === 0" class="notif-item empty">
              No notifications yet.
            </li>
            <li v-for="notif in notifications || []" :key="notif.time + notif.sender" class="notif-item"
              :class="{ unread: !notif.is_read }" @click="markNotificationAsRead(notif.id)">
              <span class="notif-text">
                <b>{{ notif.sender }}</b>
                <span v-if="notif.type === 'message'">sent you a message:
                  {{
                    notif.text && notif.text.length > 40
                      ? notif.text.slice(0, 40) + "…"
                      : notif.text
                  }}</span>
                <span v-else>{{ notif.text }}</span>
              </span>
              <span class="notif-time">{{ formatTimestamp(notif.time) }}</span>
            </li>
          </ul>
        </div>
        <!-- Logout Icon -->
        <div class="sidebar-icon" title="Logout" @click="logout">
          <svg width="22" height="22" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" viewBox="0 0 24 24">
            <path d="M17 16l4-4m0 0l-4-4m4 4H7" />
            <path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4" />
          </svg>
        </div>
      </div>
    </aside>
    <!-- Main Content -->
    <div class="group-main-content">
      <div class="group-container">
        <!-- Group Header -->
        <header class="group-header">
          <div class="group-info">
            <h2>{{ group.title }}</h2>
            <p>{{ group.description }}</p>
            <div v-if="
              !isMember &&
              membershipStatus !== 'pending' &&
              membershipStatus !== 'invited'
            " class="membership-status not-member">
              You are not a member of this group
            </div>
            <div v-else-if="membershipStatus === 'pending'" class="membership-status pending">
              Your membership is pending approval
            </div>
            <div v-else-if="membershipStatus === 'invited'" class="membership-status invited">
              You have been invited to this group
              <div class="invitation-actions">
                <button class="accept-btn" @click="handleInvitation('accept')">
                  Accept Invitation
                </button>
                <button class="decline-btn" @click="handleInvitation('decline')">
                  Decline Invitation
                </button>
              </div>
            </div>
            <div v-else class="membership-status accepted">
              <span v-if="group.is_owner">You are the owner of this group</span>
              <span v-else>You are a member of this group</span>
            </div>
          </div>
          <div class="group-actions">
            <button class="back-btn" @click="goBack">Back to Forum</button>
            <button v-if="group.is_owner" class="member-requests-btn" @click="showMemberRequestsModal = true">
              Member Requests
              <span v-if="pendingRequests.length > 0" class="badge">{{
                pendingRequests.length
                }}</span>
            </button>
            <button v-if="membershipStatus === 'accepted'" class="invite-btn" @click="openInviteModal">
              Invite Members
            </button>
            <button v-if="membershipStatus === 'accepted'" class="event-btn" @click="showEventModal = true">
              Create Event
            </button>
            <button v-if="!group.is_owner && membershipStatus === 'accepted'" class="leave-btn" @click="leaveGroup">
              Leave Group
            </button>
          </div>
        </header>

        <!-- Show content only if user is an accepted member -->
        <div v-if="membershipStatus === 'accepted'" class="group-content">
          <!-- Create Post Section -->
          <div class="create-post-section">
            <h3>Create a Post</h3>
            <form @submit.prevent="createPost">
              <input v-model="newPost.title" type="text" placeholder="Post title" required />
              <textarea v-model="newPost.content" placeholder="What's on your mind?" required></textarea>
              <label class="file-upload-label">
                <svg width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                  stroke-linejoin="round" viewBox="0 0 24 24">
                  <path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4" />
                  <polyline points="17 8 12 3 7 8" />
                  <line x1="12" y1="3" x2="12" y2="15" />
                </svg>
                Upload Image
                <input type="file" accept="image/*" @change="handleImageUpload" />
              </label>
              <img v-if="image" :src="image" class="image-preview" alt="Preview" />
              <button type="submit">Post</button>
            </form>
          </div>

          <!-- Posts Section -->
          <div class="posts-section">
            <h3>Group Posts</h3>
            <div v-if="posts.length === 0" class="no-posts">
              No posts yet. Be the first to post!
            </div>
            <div v-else class="post-card" v-for="post in posts" :key="post.id">
              <div class="post-header">
                <img :src="post.authorAvatar" alt="Author" class="post-author-pic" />
                <div>
                  <h4>{{ post.author }}</h4>
                  <p class="post-timestamp">
                    {{ formatTimestamp(post.created_at) }}
                  </p>
                </div>
              </div>
              <h3 class="post-title">{{ post.title }}</h3>
              <p class="post-content">{{ post.content }}</p>
              <div v-if="post.image" class="post-image">
                <img :src="post.image" alt="Post Image" id="groupostimage" style="cursor:pointer"
                  @click="openImageModal(post.image)" />
              </div>
              <div class="post-actions">
                <button @click="toggleComments(post)">
                  {{ post.showComments ? "Hide Comments" : "Show Comments" }}
                </button>
              </div>
              <div v-if="post.showComments" class="comments-section">
                <div v-if="post.comments.length === 0" class="no-comments">
                  No comments yet
                </div>
                <div v-else class="comment" v-for="comment in post.comments" :key="comment.id">
                  <div class="comment-header">
                    <img :src="comment.authorAvatar" alt="Comment Author" class="comment-author-pic" />
                    <div>
                      <h5>{{ comment.author }}</h5>
                      <p class="comment-timestamp">
                        {{ formatTimestamp(comment.created_at) }}
                      </p>
                    </div>
                  </div>
                  <p class="comment-content">{{ comment.content }}</p>
                  <div v-if="comment.image" class="comment-image">
                    <img :src="`http://localhost:8080/uploads/${comment.image}`" alt="Comment Image"
                      style="max-width: 200px; max-height: 200px; margin-top: 8px; cursor:pointer"
                      @click="openImageModal(`http://localhost:8080/uploads/${comment.image}`)" />
                  </div>
                </div>
                <form @submit.prevent="addComment(post)" class="comment-form" enctype="multipart/form-data">
                  <textarea v-model="post.newComment" placeholder="Add a comment..." required></textarea>
                  <label class="file-upload-label">
                    <svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2"
                      stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                      <path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4" />
                      <polyline points="17 8 12 3 7 8" />
                      <line x1="12" y1="3" x2="12" y2="15" />
                    </svg>
                    <input type="file" accept="image/*" @change="onCommentImageChange($event, post)" />
                  </label>
                  <img v-if="post.commentImagePreview" :src="post.commentImagePreview" class="image-preview"
                    alt="Preview" />
                  <button type="submit">Comment</button>
                </form>
              </div>
            </div>
          </div>
          <!-- Group Chat Section -->
          <div class="group-chat-section">
            <h3>Group Chat</h3>
            <div class="chat-messages">
              <div v-if="chatMessages.length === 0" class="no-messages">
                No messages yet
              </div>
              <div v-else class="message" v-for="msg in chatMessages" :key="msg.id">
                <img :src="msg.authorAvatar" alt="Author" class="message-author-pic" />
                <div class="message-header">
                  <div class="message-meta">
                    <span class="message-author">{{ msg.username }}</span>
                    <span class="message-time">{{
                      formatMessageTime(msg.created_at)
                      }}</span>
                  </div>
                  <p class="message-content">{{ msg.content }}</p>
                </div>
              </div>
            </div>
            <form @submit.prevent="sendMessage" class="chat-input">
              <input v-model="newMessage" type="text" placeholder="Type a message..." required />
              <button type="submit">Send</button>
            </form>
          </div>
        </div>

        <!-- Show message if user is not a member -->
        <div v-else class="non-member-content">
          <div class="non-member-message">
            <h3 v-if="membershipStatus === 'pending'">Your membership request is pending approval</h3>
            <h3 v-else>Join this group to see its content</h3>
            <p v-if="membershipStatus === 'pending'">Please wait for an admin to approve your request.</p>
            <p v-else>This is a private group. You need to be a member to see posts, events, and participate in
              discussions.</p>
          </div>
        </div>

        <!-- Event Modal -->
        <div v-if="showEventModal" class="modal-overlay">
          <div class="modal-content">
            <div class="modal-header">
              <h3>Group Events</h3>
              <button class="close-btn" @click="showEventModal = false">
                &times;
              </button>
            </div>

            <!-- Create Event Form -->
            <div class="create-event-form">
              <h4>Create New Event</h4>
              <form @submit.prevent="createEvent">
                <div class="form-group">
                  <label>Event Title</label>
                  <input v-model="newEvent.title" type="text" required placeholder="Enter event title" />
                </div>
                <div class="form-group">
                  <label>Date and Time</label>
                  <input v-model="newEvent.date" type="datetime-local" required />
                </div>
                <div class="form-group">
                  <label>Description</label>
                  <textarea v-model="newEvent.description" required placeholder="Enter event description"></textarea>
                </div>
                <div class="form-group">
                  <label>Location</label>
                  <input v-model="newEvent.location" type="text" required placeholder="Enter event location" />
                </div>
                <button type="submit" class="submit-btn">Create Event</button>
              </form>
            </div>

            <!-- Events Lists -->

            <div class="events-container">
              <!-- Upcoming Events -->
              <div class="events-section">
                <h4>Upcoming Events</h4>
                <div v-if="upcomingEvents.length === 0" class="no-events">
                  No upcoming events
                </div>
                <div v-else class="event-list">
                  <div v-for="event in upcomingEvents" :key="event.id" class="event-item">
                    <div class="event-info">
                      <h5>{{ event.title }}</h5>
                      <p class="event-date">📅 {{ formatDate(event.date) }}</p>
                      <p class="event-location">📍 {{ event.location }}</p>
                      <p class="event-description">{{ event.description }}</p>
                      <div class="event-response-stats">
                        <span class="response-stat going">
                          <span class="response-icon">👥</span>
                          <span class="response-count">{{
                            event.going_count || 0
                            }}</span>
                          <span class="response-label">Going</span>
                        </span>
                        <span class="response-stat not-going">
                          <span class="response-icon">✕</span>
                          <span class="response-count">{{
                            event.not_going_count || 0
                            }}</span>
                          <span class="response-label">Not Going</span>
                        </span>
                      </div>
                    </div>
                    <div class="event-actions">
                      <button v-if="isMember" @click="respondToEvent(event.id, 1)" :class="{
                        'response-btn': true,
                        going: event.user_response === 1,
                        active: event.user_response === 1,
                      }" :disabled="event.user_response === 1">
                        <span class="btn-icon">{{
                          event.user_response === 1 ? "✓" : "👤"
                          }}</span>
                        {{
                          event.user_response === 1 ? "You're Going" : "Going"
                        }}
                      </button>
                      <button v-if="isMember" @click="respondToEvent(event.id, -1)" :class="{
                        'response-btn': true,
                        'not-going': event.user_response === -1,
                        active: event.user_response === -1,
                      }" :disabled="event.user_response === -1">
                        <span class="btn-icon">{{
                          event.user_response === -1 ? "✓" : "✕"
                          }}</span>
                        {{
                          event.user_response === -1 ? "Not Going" : "Not Going"
                        }}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
              <!-- Past Events -->
              <div class="events-section">
                <h4>Past Events</h4>
                <div v-if="pastEvents.length === 0" class="no-events">
                  No past events
                </div>
                <div v-else class="event-list">
                  <div v-for="event in pastEvents" :key="event.id" class="event-item past">
                    <div class="event-info">
                      <h5>{{ event.title }}</h5>
                      <p class="event-date">📅 {{ formatDate(event.date) }}</p>
                      <p class="event-location">📍 {{ event.location }}</p>
                      <p class="event-description">{{ event.description }}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Member Requests Modal -->
        <div v-if="showMemberRequestsModal" class="modal-overlay">
          <div class="modal-content">
            <div class="modal-header">
              <h3>Member Requests</h3>
              <button class="close-btn" @click="showMemberRequestsModal = false">
                &times;
              </button>
            </div>
            <div class="member-requests-container">
              <div v-if="pendingRequests.length === 0" class="no-requests">
                No pending member requests
              </div>
              <div v-else class="requests-list">
                <div v-for="request in pendingRequests" :key="request.id" class="request-item">
                  <div class="request-info">
                    <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${request.username}`" alt="User Avatar"
                      class="request-avatar" />
                    <div class="request-details">
                      <h4>{{ request.full_name }}</h4>
                      <p>{{ request.email }}</p>
                    </div>
                  </div>
                  <div class="request-actions">
                    <button class="accept-btn" @click="handleMemberRequest(request.id, 'accept')">
                      Accept
                    </button>
                    <button class="decline-btn" @click="handleMemberRequest(request.id, 'decline')">
                      Decline
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Invite Members Modal -->
        <div v-if="showInviteMembersModal" class="modal-overlay">
          <div class="modal-content invite-members-modal">
            <h3>Invite Members</h3>
            <div class="search-box">
              <input type="text" v-model="searchQuery" placeholder="Search users..." @input="searchUsers" />
            </div>
            <div class="users-list">
              <div v-if="!filteredUsers || filteredUsers.length === 0" class="no-users">
                {{
                  searchQuery
                    ? "No users found matching your search"
                    : "No users available to invite"
                }}
              </div>
              <div v-else class="users-list-container">
                <div v-for="user in filteredUsers" :key="user.id" class="user-item">
                  <div class="user-info">
                    <img :src="user.avatar" :alt="user.username" class="user-avatar" />
                    <span class="user-name">{{ user.fullname }}</span>
                  </div>
                  <button class="invite-user-btn" :disabled="user.status === 'invited' ||
                    user.status === 'pending' ||
                    user.status === 'accepted'
                    " @click="inviteUser(user)">
                    {{
                      user.status === "invited"
                        ? "Invited"
                        : user.status === "pending"
                          ? "Request Pending"
                          : user.status === "accepted"
                            ? "Member"
                            : "Invite"
                    }}
                  </button>
                </div>
              </div>
            </div>
            <div class="modal-buttons">
              <button class="close-btn" @click="showInviteMembersModal = false">Close</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showImageModal" class="image-modal-overlay" @click.self="closeImageModal">
      <div class="image-modal-content">
        <img :src="modalImageSrc" alt="Large Image" />
        <button class="close-image-modal" @click="closeImageModal">&times;</button>
      </div>
    </div>
  </div>
</template>

<script>
import { notificationWebSocket } from '../services/notificationWebSocket.js';

export default {
  props: ["showNotifications"],
  name: "GroupPage",
  data() {
    return {
      isLoading: true,
      user: {
        id: "",
        username: "",
      },
      group: {
        id: null,
        title: "",
        description: "",
        is_owner: false,
      },
      isMember: false,
      isAdmin: false,
      membershipStatus: "not_member",
      posts: [],
      chatMessages: [],
      newPost: {
        title: "",
        content: "",
        image: "",
      },
      newMessage: "",
      socket: null,
      showEventModal: false,
      showMemberRequestsModal: false,
      pendingRequests: [],
      newEvent: {
        title: "",
        date: "",
        description: "",
        location: "",
      },
      upcomingEvents: [],
      pastEvents: [],
      notification: {
        visible: false,
        message: "",
        type: "success",
      },
      image: "",
      notifications: [],
      showInviteMembersModal: false,
      searchQuery: "",
      filteredUsers: [],
      allUsers: [],
      invitedUsers: new Set(),
      groupMembers: [],
      showNotifications: false,
      showImageModal: false,
      modalImageSrc: "",
    };
  },
  beforeRouteEnter(to, from, next) {
    fetch("http://localhost:8080/api/info", {
      method: "GET",
      credentials: "include",
    })
      .then(async res => {
        if (res.ok) {
          const userData = await res.json();
          next(vm => {
            vm.user.username = userData.Username;
          });
        } else {
          next('/login');
        }
      })
      .catch(() => {
        next('/login');
      });
  },
  async created() {
    console.log("GroupPage component created");
    const groupId = this.$route.params.id;
    console.log("Group ID:", groupId);
    await this.checkMembership(groupId);
    await this.fetchGroupDetails(groupId);
    await this.fetchEvents(this.$route.params.id);
    await this.fetchPosts(groupId);
    await this.fetchNotifications();
    if (this.group.is_owner) {
      await this.fetchPendingRequests();
    }
    this.initializeWebSocket(groupId);
    this.scrollToBottom(); 
  },
  methods: {
    openImageModal(src) {
      this.modalImageSrc = src;
      this.showImageModal = true;
    },
    closeImageModal() {
      this.showImageModal = false;
      this.modalImageSrc = "";
    },
    scrollToBottom() {
      this.$nextTick(() => {
        const chatContainer = document.querySelector('.chat-messages');
        if (chatContainer) {
          chatContainer.scrollTop = chatContainer.scrollHeight;
        }
      });
    },
    async fetchGroupDetails(groupId) {
      try {
        const response = await fetch(
          `http://localhost:8080/api/getgroups?id=${groupId}`,
          {
            method: "GET",
            credentials: "include",
          }
        );
        if (response.ok) {
          const data = await response.json();
          for (const group of data) {
            if (group.id === groupId) {
              this.group = group;
              break;
            }
          }
        }
      } catch (error) {
        console.error("Error fetching group details:", error);
      }
    },
    handleNotifClick() {
      this.showNotifications = !this.showNotifications;
      if (this.showNotifications) {
        //mark ga3 notif bli read
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
            ? data.map((notif) => ({
              id: notif.id,
              sender: notif.sender_username || "Unknown",
              text: notif.content || "",
              time: notif.time || new Date().toISOString(),
              type: notif.type || "notification",
              is_read: notif.is_read || false,
            }))
            : [];

          this.unreadNotificationCount = this.notifications.filter(n => !n.is_read).length;
        } else {
          this.notifications = [];
        }
      } catch (e) {
        console.error('Error fetching notifications:', e);
        this.notifications = [];
      }
    },
    async markNotificationAsRead(notificationId) {
      try {
        const notification = this.notifications.find(n => n.id === notificationId);
        if (!notification) return;

        
        setTimeout(async () => {
          try {
            const res = await fetch(`http://localhost:8080/api/markasread`, {
              method: 'POST',
              credentials: 'include',
              headers: {
                'Content-Type': 'application/json'
              },
              body: JSON.stringify({
                notificationId,
                is_read: 1
              })
            });
            if (res.ok) {
              notification.is_read = true;
            }
          } catch (error) {
            console.error('Error marking notification as read:', error);
          }
        }, 2000);
      } catch (error) {
        console.error('Error marking notification as read:', error);
      }
    },
    async checkMembership(groupId) {
      try {
        const response = await fetch(
          `http://localhost:8080/api/checkmem?group_id=${groupId}`,
          {
            method: "GET",
            credentials: "include",
          }
        );

        if (response.ok) {
          const data = await response.json();
          console.log("Membership data:", data);

          this.isMember = data.is_member;
          this.membershipStatus = data.status || "not_member";
          this.isAdmin = data.is_admin || false;
          this.group.is_owner = data.is_owner || false;

          
          if (!this.isMember || this.membershipStatus === 'pending') {
            this.disableFeatures();
          }
        }
      } catch (error) {
        console.error("Error checking membership:", error);
      }
    },
    initializeWebSocket(groupId) {
      this.socket = new WebSocket(
        `ws://localhost:8080/ws/group/${groupId}`
      );

      this.socket.onopen = () => {
        console.log('WebSocket connected to group chat');
      };

      this.socket.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);

          
          if (Array.isArray(data)) {
            this.chatMessages = data
              .map(msg => ({
                ...msg,
                authorAvatar: msg.avatar ? `http://localhost:8080/uploads/${msg.avatar}` : `https://api.dicebear.com/7.x/avataaars/svg?seed=${msg.username}`,
                created_at: new Date(msg.created_at)
              }))
              .sort((a, b) => a.created_at - b.created_at);
            return;
          }

          if (data && data.error) {
            console.error('WebSocket error:', data.error);
            this.showNotification(data.error, 'error');
            return;
          }

          if (data && data.username && data.content) {
            this.chatMessages.push({
              ...data,
              authorAvatar: data.avatar ? `http://localhost:8080/uploads/${data.avatar}` : `https://api.dicebear.com/7.x/avataaars/svg?seed=${data.username}`,
              created_at: new Date(data.created_at)
            });

            this.chatMessages.sort((a, b) => a.created_at - b.created_at);

            this.$nextTick(() => {
              const chatContainer = document.querySelector('.chat-messages');
              if (chatContainer) {
                chatContainer.scrollTop = chatContainer.scrollHeight;
              }
            });
          }
        } catch (error) {
          console.error('Error processing WebSocket message:', error);
          this.showNotification('Error processing message', 'error');
        }
      };

      this.socket.onerror = (error) => {
        console.error('WebSocket error:', error);
        this.showNotification('Error connecting to group chat', 'error');
      };

      this.socket.onclose = () => {
        console.log('WebSocket disconnected from group chat');
      };
    },
    async sendMessage() {
      if (this.socket && this.socket.readyState === WebSocket.OPEN && this.newMessage.trim()) {
        try {
          this.socket.send(
            JSON.stringify({
              content: this.newMessage.trim(),
              group_id: this.group.id
            })
          );
          this.newMessage = "";
        } catch (error) {
          console.error('Error sending message:', error);
          this.showNotification('Failed to send message', 'error');
        }
      } else {
        this.showNotification('Cannot send message - chat disconnected', 'error');
      }
    },
    handleImageUpload(event) {
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

        const reader = new FileReader();
        reader.onload = (e) => {
          this.image = e.target.result; 
        };
        reader.readAsDataURL(file);
      }
    },
    async leaveGroup() {
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

        const response = await fetch(
          "http://localhost:8080/api/removememberfromgroup",
          {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              group_id: this.group.id,
              user_id: userData.id, 
            }),
          },
        );
        if (response.ok) {
          this.showNotification("Successfully left the group", "success");
          this.$router.push("/home");
        } else {
          const errorText = await response.text();
          console.error("Failed to leave group:", errorText);
          this.showNotification("Failed to leave group", "error");
        }
      } catch (error) {
        console.error("Error leaving group:", error);
        this.showNotification("Error leaving group", "error");
      }
    },
    goBack() {
      this.$router.push("/home");
    },
    formatTimestamp(timestamp) {
      return new Date(timestamp).toLocaleString();
    },
    validateEventDate(date) {
      const eventDate = new Date(date);
      const now = new Date();
      const maxDate = new Date();
      maxDate.setFullYear(now.getFullYear() + 2); 

      
      if (eventDate < now) {
        return { valid: false, message: "Event date cannot be in the past" };
      }

      
      if (eventDate > maxDate) {
        return { valid: false, message: "Event date cannot be more than 2 years in the future" };
      }

      return { valid: true };
    },
    async createEvent() {
      try {
        
        const dateValidation = this.validateEventDate(this.newEvent.date);
        if (!dateValidation.valid) {
          this.showNotification(dateValidation.message, "error");
          return;
        }

        const response = await fetch(
          `http://localhost:8080/api/events/add?group_id=${this.$route.params.id}`,
          {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              group_id: this.group.id,
              ...this.newEvent,
            }),
          }
        );

        if (response.ok) {
          this.showNotification("Event created successfully!", "success");
          this.newEvent = {
            title: "",
            date: "",
            description: "",
            location: "",
          };
          await this.fetchEvents(this.$route.params.id);
          this.showEventModal = false;
        } else {
          throw new Error("Failed to create event");
        }
      } catch (error) {
        console.error("Error creating event:", error);
        this.showNotification("Failed to create event", "error");
      }
    },

    async respondToEvent(eventId, response) {
      try {
        const res = await fetch(
          `http://localhost:8080/api/events/join?event_id=${eventId}&response=${response}`,
          {
            method: "POST",
            credentials: "include",
          }
        );

        if (!res.ok) {
          throw new Error("Failed to update response");
        }

        const data = await res.json();

        
        const eventToUpdate = this.upcomingEvents.find(e => e.id === eventId);
        if (eventToUpdate) {
          eventToUpdate.going_count = data.going_count;
          eventToUpdate.not_going_count = data.not_going_count;
          eventToUpdate.user_response = response;
        }

        const message = response === 1 ? "You're going to the event!" : "You've marked as not going";
        this.showNotification(message, "success");
      } catch (error) {
        console.error("Error responding to event:", error);
        this.showNotification("Failed to update response", "error");
      }
    },

    showNotification(message, type = 'success') {
      this.notification = {
        visible: true,
        message,
        type,
      };
      setTimeout(() => {
        this.notification.visible = false;
      }, 3000);
    },
    async fetchEvents(group_id) {
      try {
        const response = await fetch(
          `http://localhost:8080/api/events?id=${group_id}`,
          {
            credentials: "include",
          }
        );

        if (!response.ok) {
          const errorText = await response.text();
          throw new Error(errorText);
        }

        const events = await response.json();
        const now = new Date();

        if (events) {
          const normalizedEvents = events.map((event) => ({
            ...event,
            date: event.date || event.event_datetime || event.eventDateTime,
            going_count: event.going_count || 0,
            not_going_count: event.not_going_count || 0,
            user_response: event.response || null
          }));

          this.upcomingEvents = normalizedEvents
            .filter((event) => new Date(event.date) > now)
            .sort((a, b) => new Date(a.date) - new Date(b.date));

          this.pastEvents = normalizedEvents
            .filter((event) => new Date(event.date) <= now)
            .sort((a, b) => new Date(b.date) - new Date(a.date));

          console.log("Fetched events with response counts:", this.upcomingEvents);
        }
      } catch (error) {
        console.error("Error fetching events:", error);
        this.showNotification(error.message || "Error fetching events", "error");
        this.upcomingEvents = [];
        this.pastEvents = [];
      }
    },

    async fetchPosts(group_id) {
      try {
        const response = await fetch(
          `http://localhost:8080/api/groupposts?group_id=${group_id}`,
          {
            credentials: "include",
          }
        );

        if (response.ok) {
          const posts = await response.json();
          console.log("Fetched posts:", posts);

          if (posts) {
            this.posts = posts.map((post) => ({
              ...post,
              id: post.id,
              author: post.author,
              authorAvatar: post.avatar ? `http://localhost:8080/uploads/${post.avatar}` : `https://api.dicebear.com/7.x/avataaars/svg?seed=${post.author}`,
              image: post.image ? post.image : null,
              created_at: post.creation_date,
              comments: [],
              showComments: false,
              newComment: "",
              commentError: "",
              commentImage: null,
            }));
          }
        } else {
          console.error("Failed to fetch posts");
          this.showNotification("Failed to fetch posts", "error");
        }
      } catch (error) {
        console.error("Error fetching posts:", error);
        this.showNotification("Error fetching posts", "error");
      }
    },

    async createPost() {
      const groupId = this.$route.params.id;

      try {
        const response = await fetch(`http://localhost:8080/api/groupposts/add?group_id=${groupId}`, {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            title: this.newPost.title,
            content: this.newPost.content,
            image: this.image,
          }),
        });

        if (response.ok) {
          this.newPost = { title: "", content: "", image: "" };
          this.image = "";
          
          const fileInput = document.querySelector('input[type="file"]');
          if (fileInput) {
            fileInput.value = '';
          }
          await this.fetchPosts(groupId);
        } else {
          const errorText = await response.text();
          console.error("Post failed:", errorText);
        }
      } catch (error) {
        console.error("Error creating post:", error);
      }
    },
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        
        const maxSize = 2 * 1024 * 1024; 
        if (file.size > maxSize) {
          alert('Image file size must be less than 2MB');
          event.target.value = ''; 
          return;
        }

        const reader = new FileReader();
        reader.onload = (e) => {
          this.image = e.target.result; 
        };
        reader.readAsDataURL(file);
      }
    },
    async addComment(post) {
      if (!post.newComment.trim() && !post.commentImage) {
        post.commentError = "Comment cannot be empty.";
        return;
      }
      const formData = new FormData();
      formData.append("group_post_id", post.id);
      formData.append("content", post.newComment.trim());
      if (post.commentImage) {
        formData.append("image", post.commentImage);
      }
      try {
        const res = await fetch("http://localhost:8080/api/groupcomments/add", {
          method: "POST",
          credentials: "include",
          body: formData,
        });
        if (res.ok) {
          post.commentError = "";
          post.newComment = "";
          post.commentImage = null;
          post.commentImagePreview = null;
          
          const fileInput = document.querySelector(`input[type="file"][data-post-id="${post.id}"]`);
          if (fileInput) {
            fileInput.value = '';
          }
          this.showNotification("Comment added successfully", "success");
          await this.fetchComments(post);
        } else {
          post.commentError = "Failed to add comment.";
          this.showNotification("Failed to add comment", "error");
        }
      } catch (error) {
        post.commentError = "Error adding comment.";
        this.showNotification("Error adding comment", "error");
      }
    },
    onCommentImageChange(event, post) {
      const file = event.target.files[0];
      if (file) {
        
        const maxSize = 2 * 1024 * 1024; 
        if (file.size > maxSize) {
          this.showNotification('Image file size must be less than 2MB', 'error');
          event.target.value = ''; 
          return;
        }

        post.commentImage = file;
        
        post.commentImagePreview = URL.createObjectURL(file);
      } else {
        post.commentImage = null;
        post.commentImagePreview = null;
      }
    },

    async fetchComments(post) {
      try {
        const res = await fetch(
          `http://localhost:8080/api/groupcomments?group_post_id=${post.id}`,
          {
            method: "GET",
            credentials: "include",
          }
        );
        if (res.ok) {
          const data = await res.json();
          post.comments = data.map((comment) => ({
            ...comment,
            authorAvatar: comment.avatar
              ? `http://localhost:8080/uploads/${comment.avatar}`
              : `https://api.dicebear.com/7.x/avataaars/svg?seed=${comment.author}`,
            created_at: comment.creation_date,
            image: comment.image
          }));
        } else {
          post.commentError = "Failed to load comments.";
        }
      } catch (error) {
        post.commentError = "Error loading comments.";
      }
    },

    toggleComments(post) {
      console.log(post);

      post.showComments = !post.showComments;
      if (post.showComments && post.comments.length === 0) {
        this.fetchComments(post);
      }
    },
    formatDate(date) {
      return new Date(date).toLocaleString();
    },
    goToHome() {
      this.$router.push('/home');
    },
    goToProfile() {
      this.$router.push(`/profile/${this.user.username}`);
    },
    goToMyGroups() {
      this.$router.push('/mygroups');
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
            this.$router.push('/login');
          }
        })
        .catch(() => {
          this.$router.push('/login');
        });
    },
    disableFeatures() {
      
      this.posts = [];
      this.chatMessages = [];
      this.upcomingEvents = [];
      this.pastEvents = [];
    },
    async fetchPendingRequests() {
      if (!this.group.id) return;

      try {
        const response = await fetch(
          `http://localhost:8080/api/GetInvitations?group_id=${this.group.id}`,
          {
            credentials: "include",
          }
        );
        if (response.ok) {
          const data = await response.json();
          console.log("Pending requests data:", data);

          
          
          this.pendingRequests = Array.isArray(data) ? data : [];
        }
      } catch (error) {
        console.error("Error fetching pending requests:", error);
        this.pendingRequests = [];
      }
    },

    async handleMemberRequest(userId, action) {
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
            group_id: this.group.id,
            user_id: userId
          })
        });

        if (response.ok) {
          await this.fetchPendingRequests();
          this.showNotification(
            `Member request ${action === "accept" ? "accepted" : "declined"
            } successfully`,
            "success"
          );
        } else {
          throw new Error('Failed to handle member request');
        }
      } catch (error) {
        console.error('Error handling member request:', error);
        this.showNotification(
          `Failed to ${action} member request`,
          'error'
        );
      }
    },
    async handleInvitation(action) {
      try {
        const endpoint = action === 'accept'
          ? 'http://localhost:8080/api/acceptgroupinvite'
          : 'http://localhost:8080/api/declinegroupinvite';

        const response = await fetch(endpoint, {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            group_id: this.group.id,
            action: action
          })
        });
        console.log(`Invitation ${action} request sent for group ID: ${this.group.id}`);

        
        if (action === 'accept') {
          this.membershipStatus = 'accepted';
          this.isMember = true;
          this.showNotification('Successfully joined the group', 'success');
        } else {
          this.membershipStatus = null;
          this.isMember = false;
          this.showNotification('Successfully declined the invitation', 'success');
        }

        await this.fetchGroupDetails(this.group.id);
        await this.fetchPosts(this.group.id);
        await this.fetchEvents(this.group.id);
        await this.fetchPendingRequests();
      } catch (error) {
        console.error(`Error ${action}ing invitation:`, error);
        this.showNotification(
          `Failed to ${action} invitation. Please try again.`,
          'error'
        );
      }
    },
    async fetchAllUsers() {
      try {
        const response = await fetch("http://localhost:8080/api/allusers", {
          credentials: "include",
        });
        if (response.ok) {
          const users = await response.json();
          console.log("Fetched users:", users);
          this.allUsers = users.map(user => ({
            fullname: user.fullname,
            id: user.id,
            username: user.username,
            avatar: user.avatar
              ? `http://localhost:8080/uploads/${user.avatar}`
              : `https://api.dicebear.com/7.x/avataaars/svg?seed=${user.username}`,
            invited: false
          }));
          console.log("Processed users:", this.allUsers);
          
          this.filteredUsers = [...this.allUsers];
        }
      } catch (error) {
        console.error("Error fetching users:", error);
        this.showNotification("Failed to fetch users", "error");
      }
    },

    async searchUsers() {
      console.log("Search query:", this.searchQuery);
      console.log("All users:", this.allUsers);

      
      try {
        const response = await fetch(`http://localhost:8080/api/groupmembers/status?group_id=${this.group.id}`, {
          credentials: "include",
        });
        if (response.ok) {
          const memberStatuses = await response.json();
          console.log("Member statuses:", memberStatuses);

          
          this.allUsers = this.allUsers.map(user => ({
            ...user,
            status: memberStatuses[user.id] || null
          }));
        }
      } catch (error) {
        console.error("Error fetching member statuses:", error);
      }

      if (!this.searchQuery.trim()) {
        
        this.filteredUsers = [...this.allUsers];
      } else {
        const query = this.searchQuery.toLowerCase();
        this.filteredUsers = this.allUsers.filter(user =>
          user.username.toLowerCase().includes(query)
        );
      }

      console.log("Filtered users:", this.filteredUsers);
    },

    async inviteUser(user) {
      try {
        const response = await fetch("http://localhost:8080/api/addmembertogroup", {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            group_id: this.group.id,
            user_id: user.id,
            status: "invited"
          }),
        });
        console.log("Invitation sent to:", user.id, this.group.id, "status:", "invited");

        if (response.ok) {
          
          user.status = 'invited';
          this.invitedUsers.add(user.id);
          this.showNotification(`Invitation sent to ${user.username}`, "success");

          
          try {
            await fetch("http://localhost:8080/api/notifications", {
              method: "POST",
              credentials: "include",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                user_id: user.id,
                type: "group_invite",
                content: `You have been invited to join ${this.group.title}`,
              }),
            });
          } catch (notifError) {
            console.error("Error creating notification:", notifError);
            
          }
        } else {
          const errorText = await response.text();
          console.error("Failed to send invitation:", errorText);
          this.showNotification("Failed to send invitation", "error");
        }
      } catch (error) {
        console.error("Error inviting user:", error);
        this.showNotification("Failed to send invitation", "error");
      }
    },

    async fetchGroupMembers() {
      try {
        const response = await fetch(`http://localhost:8080/api/groupmembers?group_id=${this.group.id}`, {
          credentials: "include",
        });
        if (response.ok) {
          const members = await response.json();
          console.log("Fetched group members:", members);
          this.groupMembers = members;
        }
      } catch (error) {
        console.error("Error fetching group members:", error);
      }
    },

    async openInviteModal() {
      this.showInviteMembersModal = true;
      await this.fetchAllUsers();
      await this.fetchGroupMembers();
      this.searchUsers(); 
    },
    formatMessageTime(date) {
      const now = new Date();
      const messageDate = new Date(date);
      const diffInHours = Math.abs(now - messageDate) / 36e5;

      if (diffInHours < 24) {
        
        return messageDate.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
      } else if (diffInHours < 48) {
        
        return 'Yesterday ' + messageDate.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
      } else {
        
        return messageDate.toLocaleDateString([], {
          month: 'short',
          day: 'numeric',
          hour: '2-digit',
          minute: '2-digit'
        });
      }
    },
    setupNotificationWebSocket() {
      
      notificationWebSocket.onNotification('group-page', (notification) => {
        console.log('Received real-time notification:', notification);

        
        this.fetchNotifications();
      });
    },
  },
  mounted() {
    
    this.setupNotificationWebSocket();

    this.fetchGroupDetails(this.$route.params.id);
    this.fetchPosts(this.$route.params.id);
    this.fetchNotifications();
    document.addEventListener('click', this.handleNotifClose);
    this.scrollToBottom(); 
  },
  beforeUnmount() {
    
    notificationWebSocket.removeNotificationHandler('group-page');

    document.removeEventListener('click', this.handleNotifClose);
    if (this.socket) {
      this.socket.close();
    }
  },
  computed: {
    unreadNotificationCount() {
      return this.notifications.filter(notif => !notif.is_read).length;
    }
  },
  watch: {
    chatMessages: {
      handler() {
        this.scrollToBottom();
      },
      deep: true
    }
  },
};
</script>

<style scoped>
.group-layout {
  display: flex;
  min-height: 100vh;
  background: linear-gradient(135deg, #f3f4f6, #e5e7eb);
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
  position: relative;
}

.sidebar-icon.active,
.sidebar-icon:hover {
  background: #35395a;
  opacity: 1;
}

#groupostimage {
  width: 100%;
}

.group-main-content {
  flex: 1;
  margin-left: 70px;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  min-width: 0;
}

.group-container {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
  min-height: 100vh;
  background: linear-gradient(135deg, #f3f4f6, #e5e7eb);
}

.group-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem 2rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 1rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.group-info h2 {
  margin: 0;
  color: #1f2937;
  font-size: 1.75rem;
  font-weight: 600;
}

.group-info p {
  margin: 0.75rem 0 0;
  color: #6b7280;
  font-size: 1.1rem;
  line-height: 1.5;
}

.group-actions {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.back-btn,
.leave-btn,
.event-btn {
  padding: 0.75rem 1.5rem;
  border-radius: 0.75rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.back-btn {
  background: #e5e7eb;
  color: #4b5563;
  border: none;
}

.back-btn:hover {
  background: #d1d5db;
}

.leave-btn {
  background: #ef4444;
  color: white;
  border: none;
}

.leave-btn:hover {
  background: #dc2626;
}

.event-btn {
  background: #4f46e5;
  color: white;
  border: none;
}

.event-btn:hover {
  background: #4338ca;
}

.group-content {
  display: grid;
  grid-template-columns: 350px 1fr 350px;
  gap: 2rem;
}

.create-post-section,
.posts-section,
.group-chat-section {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 1rem;
  padding: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.2);
  height: fit-content;
}

.create-post-section h3,
.posts-section h3,
.group-chat-section h3 {
  color: #1f2937;
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  padding-bottom: 0.75rem;
  border-bottom: 2px solid #e5e7eb;
}

.create-post-section form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.create-post-section input,
.create-post-section textarea {
  padding: 0.75rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 1rem;
  background: rgba(255, 255, 255, 0.8);
  transition: all 0.3s ease;
}

.create-post-section input:focus,
.create-post-section textarea:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.create-post-section textarea {
  min-height: 150px;
  resize: vertical;
}

.create-post-section button {
  background: #4f46e5;
  color: white;
  padding: 0.75rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.3s ease;
}

.create-post-section button:hover {
  background: #4338ca;
}

.post-card {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 1rem;
  padding: 1.5rem;
  margin-bottom: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.2);
  transition: transform 0.3s ease;
}

.post-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
}

.post-author-pic,
.comment-author-pic,
.message-author-pic {
  width: 3rem;
  height: 3rem;
  border-radius: 50%;
  border: 2px solid #e5e7eb;
  transition: transform 0.3s ease;
}

.post-author-pic:hover,
.comment-author-pic:hover,
.message-author-pic:hover {
  transform: scale(1.1);
}

.post-title {
  margin: 0.75rem 0;
  color: #1f2937;
  font-size: 1.25rem;
  font-weight: 600;
}

.post-content {
  color: #4b5563;
  margin-bottom: 1rem;
  line-height: 1.6;
  font-size: 1.1rem;
}

.comments-section {
  margin-top: 1.5rem;
  padding-top: 1.5rem;
  border-top: 2px solid #e5e7eb;
}

.comment {
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: rgba(249, 250, 251, 0.8);
  border-radius: 0.75rem;
  transition: transform 0.3s ease;
}



.comment:hover {
  transform: translateX(5px);
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.5rem;
}

.comment-content {
  color: #4b5563;
  margin-left: 3.5rem;
  line-height: 1.5;
}

.chat-messages {
  height: 500px;
  overflow-y: auto;
  margin-bottom: 1.5rem;
  padding-right: 0.5rem;
}

.chat-messages::-webkit-scrollbar {
  width: 6px;
}

.chat-messages::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 3px;
}

.chat-messages::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.message {
  margin-bottom: 1rem;
  padding: 0.75rem;
  background: rgba(249, 250, 251, 0.8);
  border-radius: 0.75rem;
  transition: transform 0.3s ease;
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
  max-width: 85%;
}

.message:hover {
  transform: translateX(5px);
}

.message-header {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  min-width: 0;
  flex: 1;
}

.message-meta {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.25rem;
}

.message-author {
  font-weight: 600;
  color: #1f2937;
  font-size: 0.9rem;
}

.message-time {
  color: #6b7280;
  font-size: 0.8rem;
}

.message-content {
  color: #4b5563;
  line-height: 1.5;
  word-break: break-word;
  margin-top: 0.25rem;
}

.chat-input {
  display: flex;
  gap: 1rem;
  padding-top: 1rem;
  border-top: 2px solid #e5e7eb;
  margin-top: auto;
  justify-content: center;
  align-items: center;
}

.chat-input input {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 1rem;
  background: rgba(255, 255, 255, 0.8);
  transition: all 0.3s ease;
}

.chat-input input:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.chat-input button {
  background: #4f46e5;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.3s ease;
}

.chat-input button:hover {
  background: #4338ca;
}

.no-posts,
.no-comments,
.no-messages {
  text-align: center;
  color: #6b7280;
  padding: 3rem;
  font-size: 1.1rem;
  background: rgba(249, 250, 251, 0.8);
  border-radius: 0.75rem;
}

.post-actions {
  display: flex;
  gap: 0.75rem;
  margin-top: 1rem;
}

.post-image img {
  max-width: 100%;
  max-height: 300px;
  width: auto;
  height: auto;
  object-fit: contain;
  border-radius: 0.5rem;
  margin-top: 0.7rem;
}

.post-actions button {
  background: #e5e7eb;
  color: #4b5563;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.comment-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-top: 1rem;
  padding: 1rem;
  background: rgba(249, 250, 251, 0.8);
  border-radius: 0.75rem;
}

.comment-form textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 1rem;
  background: rgba(255, 255, 255, 0.8);
  transition: all 0.3s ease;
  min-height: 80px;
  resize: vertical;
}

.comment-form textarea:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.comment-form .file-upload-label {
  align-self: flex-start;
  margin-top: 0.5rem;
}

.comment-form .image-preview {
  max-width: 200px;
  max-height: 200px;
  border-radius: 0.5rem;
  margin-top: 0.5rem;
  object-fit: contain;
}

.comment-form button {
  align-self: flex-end;
  background: #4f46e5;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.3s ease;
}

.comment-form button:hover {
  background: #4338ca;
}

/* Responsive Design */
@media (max-width: 1200px) {
  .group-content {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .group-chat-section {
    grid-column: 1 / -1;
  }

  .group-container {
    padding: 1.5rem;
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

  .group-main-content {
    margin-left: 0;
    margin-top: 70px;
  }

  .group-container {
    padding-top: 1rem;
  }

  .group-header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
    padding: 1rem;
  }

  .group-actions {
    width: 100%;
    justify-content: center;
    flex-wrap: wrap;
    gap: 0.75rem;
  }

  .group-info h2 {
    font-size: 1.5rem;
  }

  .group-info p {
    font-size: 1rem;
  }

  .create-post-section,
  .posts-section,
  .group-chat-section {
    padding: 1rem;
  }

  .post-card {
    padding: 1rem;
  }

  .post-header {
    gap: 0.75rem;
  }

  .post-author-pic,
  .comment-author-pic,
  .message-author-pic {
    width: 2.5rem;
    height: 2.5rem;
  }

  .post-title {
    font-size: 1.1rem;
  }

  .post-content {
    font-size: 1rem;
  }

  .chat-messages {
    height: 400px;
  }

  .message {
    max-width: 95%;
  }

  .modal-content {
    width: 95%;
    max-height: 90vh;
  }

  .create-event-form,
  .events-container {
    padding: 1rem;
  }

  .event-item {
    flex-direction: column;
    gap: 1rem;
  }

  .event-actions {
    width: 100%;
    justify-content: flex-start;
  }

  .response-btn {
    padding: 0.5rem 1rem;
  }
}

@media (max-width: 480px) {
  .group-container {
    padding: 0.75rem;
  }

  .group-header {
    padding: 0.75rem;
  }

  .group-actions {
    flex-direction: column;
    width: 100%;
  }

  .back-btn,
  .leave-btn,
  .event-btn,
  .invite-btn,
  .member-requests-btn {
    width: 100%;
    margin: 0.25rem 0;
  }

  .create-post-section form {
    gap: 1rem;
  }

  .post-card {
    padding: 0.75rem;
  }

  .post-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }

  .post-actions {
    flex-direction: column;
    width: 100%;
  }

  .post-actions button {
    width: 100%;
    margin: 0.25rem 0;
  }

  .comment-form {
    padding: 0.75rem;
  }

  .comment-form textarea {
    min-height: 60px;
  }

  .comment-form button {
    width: 100%;
  }

  .chat-input {
    flex-direction: column;
    gap: 0.75rem;
  }

  .chat-input input {
    width: 100%;
  }

  .chat-input button {
    width: 100%;
  }

  .modal-content {
    width: 100%;
    max-height: 100vh;
    border-radius: 0;
  }

  .modal-header {
    padding: 1rem;
  }

  .create-event-form,
  .events-container {
    padding: 0.75rem;
  }

  .event-item {
    padding: 0.75rem;
  }

  .event-response-stats {
    flex-direction: column;
    gap: 0.75rem;
  }

  .response-stat {
    width: 100%;
    justify-content: center;
  }

  .event-actions {
    flex-direction: column;
    gap: 0.5rem;
  }

  .response-btn {
    width: 100%;
  }

  .request-item {
    flex-direction: column;
    gap: 0.75rem;
  }

  .request-info {
    width: 100%;
  }

  .request-actions {
    width: 100%;
    justify-content: space-between;
  }

  .accept-btn,
  .decline-btn {
    flex: 1;
  }

  .user-item {
    flex-direction: column;
    gap: 0.75rem;
  }

  .user-info {
    width: 100%;
  }

  .invite-user-btn {
    width: 100%;
  }
}

.membership-status {
  margin-top: 0.5rem;
  padding: 0.5rem;
  border-radius: 0.5rem;
  font-size: 0.9rem;
  font-weight: 500;
}

.membership-status.not-member {
  background-color: #fee2e2;
  color: #991b1b;
}

.membership-status.pending {
  background-color: #fef3c7;
  color: #92400e;
}

.membership-status.accepted {
  background-color: #dcfce7;
  color: #166534;
}

.membership-status.invited {
  background: #fef3c7;
  color: #92400e;
  padding: 1rem;
  border-radius: 0.5rem;
  margin-top: 1rem;
}

.create-post-section.disabled,
.group-chat-section.disabled {
  opacity: 0.5;
  pointer-events: none;
}

input:disabled,
textarea:disabled,
button:disabled {
  cursor: not-allowed;
  opacity: 0.7;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(5px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

.modal-content {
  background: white;
  border-radius: 1.5rem;
  width: 90%;
  max-width: 800px;
  max-height: 85vh;
  overflow-y: auto;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1),
    0 10px 10px -5px rgba(0, 0, 0, 0.04);
  animation: slideIn 0.3s ease;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem 2rem;
  border-bottom: 1px solid #e5e7eb;
  background: #f8fafc;
  border-radius: 1.5rem 1.5rem 0 0;
}

.modal-header h3 {
  margin: 0;
  color: #1f2937;
  font-size: 1.5rem;
  font-weight: 600;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.75rem;
  color: #6b7280;
  cursor: pointer;
  padding: 0.5rem;
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.close-btn:hover {
  background: #f1f5f9;
  color: #1f2937;
}

.create-event-form {
  padding: 2rem;
  border-bottom: 1px solid #e5e7eb;
  background: white;
}

.create-event-form h4 {
  margin: 0 0 1.5rem;
  color: #1f2937;
  font-size: 1.25rem;
  font-weight: 600;
}

.form-group {
  margin-bottom: 1.25rem;
}

.form-group label {
  display: block;
  color: #4b5563;
  margin-bottom: 0.5rem;
  font-weight: 500;
  font-size: 0.95rem;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.875rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 1rem;
  transition: all 0.2s ease;
  background: #f8fafc;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
  background: white;
}

.form-group textarea {
  min-height: 120px;
  resize: vertical;
}

.events-container {
  padding: 2rem;
  background: #f8fafc;
}

.events-section {
  margin-bottom: 2.5rem;
}

.events-section h4 {
  color: #1f2937;
  margin-bottom: 1.25rem;
  font-size: 1.25rem;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.events-section h4::before {
  content: "";
  display: block;
  width: 4px;
  height: 1.25rem;
  background: #4f46e5;
  border-radius: 2px;
}

.event-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.event-item {
  background: white;
  border-radius: 1rem;
  padding: 1.25rem;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 1.5rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;
}

.event-item:hover {
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.event-info {
  flex: 1;
}

.event-info h5 {
  margin: 0 0 0.75rem;
  color: #1f2937;
  font-size: 1.1rem;
  font-weight: 600;
}

.event-date,
.event-location {
  color: #6b7280;
  font-size: 0.9rem;
  margin-bottom: 0.5rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.event-description {
  color: #4b5563;
  font-size: 0.95rem;
  margin-top: 0.75rem;
  line-height: 1.5;
}

.event-response-stats {
  display: flex;
  gap: 1.5rem;
  margin-top: 1rem;
  padding: 0.75rem;
  background: rgba(243, 244, 246, 0.5);
  border-radius: 0.75rem;
}

.response-stat {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  background: white;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.response-stat.going {
  color: #059669;
}

.response-stat.not-going {
  color: #dc2626;
}

.response-icon {
  font-size: 1.25rem;
}

.response-count {
  font-weight: 600;
  font-size: 1.1rem;
}

.response-label {
  font-size: 0.9rem;
  color: #6b7280;
}

.response-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.25rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 0.95rem;
  font-weight: 500;
  transition: all 0.2s ease;
  background: #f3f4f6;
  color: #374151;
}

.response-btn .btn-icon {
  font-size: 1.1rem;
}

.response-btn.going:not(:disabled) {
  background: #059669;
  color: white;
}

.response-btn.not-going:not(:disabled) {
  background: #dc2626;
  color: white;
}

.response-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  opacity: 0.9;
}

.response-btn:disabled {
  cursor: not-allowed;
  opacity: 0.8;
}

.response-btn.active {
  position: relative;
  overflow: hidden;
}

.response-btn.active::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.1);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% {
    opacity: 0;
  }

  50% {
    opacity: 0.3;
  }

  100% {
    opacity: 0;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

@keyframes slideIn {
  from {
    transform: translateY(20px);
    opacity: 0;
  }

  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* Scrollbar Styling */
.modal-content::-webkit-scrollbar {
  width: 8px;
}

.modal-content::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 4px;
}

.modal-content::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 4px;
}

.modal-content::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

.join-btn:disabled {
  background: #10b981;
  cursor: not-allowed;
  opacity: 0.8;
}

.join-btn:disabled:hover {
  transform: none;
  background: #10b981;
}

.accept-btn,
.decline-btn {
  background: #4f46e5;
  color: white;
  padding: 0.625rem 1.25rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 0.9rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.accept-btn:hover,
.decline-btn:hover {
  background: #4338ca;
  transform: translateY(-1px);
}

.accept-btn.selected {
  background-color: #28a745;
  color: white;
}

.decline-btn.selected {
  background-color: #dc3545;
  color: white;
}

.non-member-content {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
  padding: 2rem;
}

.non-member-message {
  text-align: center;
  background: rgba(255, 255, 255, 0.9);
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  max-width: 600px;
  width: 100%;
}

.non-member-message h3 {
  color: #1f2937;
  font-size: 1.5rem;
  margin-bottom: 1rem;
}

.non-member-message p {
  color: #6b7280;
  font-size: 1.1rem;
  line-height: 1.6;
}

/* Notification Styles */
.notification {
  position: fixed;
  top: 20px;
  right: 20px;
  padding: 1rem 1.5rem;
  border-radius: 0.5rem;
  z-index: 2000;
  animation: slideIn 0.3s ease-out;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  font-weight: 500;
}

.notification-success {
  background-color: #10b981;
  color: white;
  border-left: 4px solid #059669;
}

.notification-error {
  background-color: #ef4444;
  color: white;
  border-left: 4px solid #dc2626;
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }

  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.member-requests-btn {
  background: #4f46e5;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.3s ease;
  margin-right: 1rem;
}

.member-requests-btn:hover {
  background: #4338ca;
}

.member-requests-container {
  padding: 1.5rem;
}

.no-requests {
  text-align: center;
  padding: 2rem;
  color: #6b7280;
  font-size: 1.1rem;
}

.requests-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.request-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: #f9fafb;
  border-radius: 0.75rem;
  transition: transform 0.2s ease;
}

.request-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.request-avatar {
  width: 3rem;
  height: 3rem;
  border-radius: 50%;
  border: 2px solid #e5e7eb;
}

.request-details h4 {
  margin: 0;
  color: #1f2937;
  font-size: 1.1rem;
}

.request-details p {
  margin: 0.25rem 0 0;
  color: #6b7280;
  font-size: 0.9rem;
}

.request-actions {
  display: flex;
  gap: 0.75rem;
}

.accept-btn,
.decline-btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  font-size: 0.9rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.accept-btn {
  background: #10b981;
  color: white;
}

.accept-btn:hover {
  background: #059669;
}

.decline-btn {
  background: #ef4444;
  color: white;
}

.decline-btn:hover {
  background: #dc2626;
}

.invitation-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

.invite-btn {
  background: #4f46e5;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.3s ease;
}

.invite-btn:hover {
  background: #4338ca;
}

.invite-members-modal {
  padding: 2rem;
  border-bottom: 1px solid #e5e7eb;
  background: white;
}

.invite-members-modal h3 {
  margin: 0 0 1.5rem;
  color: #1f2937;
  font-size: 1.25rem;
  font-weight: 600;
}

.search-box {
  margin-bottom: 1.5rem;
}

.search-box input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 1rem;
  background: #f8fafc;
}

.search-box input:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.users-list {
  margin-bottom: 1.5rem;
}

.users-list-container {
  max-height: 300px;
  overflow-y: auto;
  padding: 0.5rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: #f9fafb;
  border-radius: 0.5rem;
  margin-bottom: 0.5rem;
  transition: all 0.2s ease;
}

.user-item:hover {
  background: #f3f4f6;
  transform: translateY(-1px);
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.user-name {
  font-size: 0.95rem;
  font-weight: 500;
  color: #374151;
}

.invite-user-btn {
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  font-weight: 500;
  transition: all 0.2s ease;
  background: #4f46e5;
  color: white;
  border: none;
  cursor: pointer;
}

.invite-user-btn:disabled {
  background: #e5e7eb;
  color: #9ca3af;
  cursor: not-allowed;
}

.invite-user-btn:not(:disabled):hover {
  background: #4338ca;
}

.no-users {
  text-align: center;
  padding: 2rem;
  color: #6b7280;
  font-style: italic;
}

.modal-buttons {
  display: flex;
  justify-content: flex-end;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}

.modal-buttons button {
  background: #4f46e5;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: all 0.3s ease;
}

.modal-buttons button:hover {
  background: #4338ca;
}

#groupostimage {
  width: 100%;
}

.file-upload-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.9rem;
  color: #4b5563;
  width: fit-content;
}

.file-upload-label:hover {
  background: #e5e7eb;
  border-color: #d1d5db;
}

.file-upload-label svg {
  color: #4f46e5;
}

.file-upload-label input[type="file"] {
  display: none;
}

.image-preview {
  max-width: 100%;
  max-height: 300px;
  border-radius: 0.5rem;
  margin-top: 1rem;
  object-fit: contain;
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

.post-title,
.post-content,
.comment-content {
  word-break: break-word;
  overflow-wrap: break-word;
  white-space: pre-line;
  /* preserves line breaks and wraps text */
}

.image-modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(30, 32, 48, 0.85);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.image-modal-content {
  position: relative;
  max-width: 90vw;
  max-height: 90vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.image-modal-content img {
  max-width: 90vw;
  max-height: 90vh;
  border-radius: 1rem;
  box-shadow: 0 8px 32px rgba(35,38,58,0.25);
  background: #fff;
}

.close-image-modal {
  position: absolute;
  top: 10px;
  right: 10px;
  background: rgba(0, 0, 0, 0.5);
  color: #fff;
  border: none;
  font-size: 2rem;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-image-modal:hover {
  background: #ef4444;
}

/* Add responsive styles for sidebar */
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
  }

  .sidebar-icons {
    flex-direction: row;
    justify-content: space-around;
    margin: 0;
    padding: 0 1rem;
  }

  .group-main-content {
    margin-left: 0;
    margin-top: 70px;
  }

  .group-container {
    padding-top: 1rem;
  }
}

@media (max-width: 480px) {
  .sidebar {
    height: 60px;
  }

  .sidebar-icons {
    padding: 0 0.5rem;
  }

  .group-main-content {
    margin-top: 60px;
  }

  .sidebar-icon {
    width: 40px;
    height: 40px;
  }
}
</style>