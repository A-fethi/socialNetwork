<template>
  <div class="forum-container">
    <!-- Header -->
    <header class="forum-header">
      <div class="user-info">
        <img class="profile-pic" :src="user.avatar" alt="Profile" />
        <div>
          <h2>{{ user.name }}</h2>
          <p>{{ user.email }}</p>
        </div>
      </div>
      <div class="header-buttons">
        <button class="profile-btn" @click="goToProfile">Profile</button>
        <button class="logout-btn" @click="logout">Logout</button>
      </div>
    </header>

    <!-- Main Grid -->
    <div class="forum-grid">
      <!-- Left Sidebar: Groups -->
      <aside class="forum-sidebar">
        <h3>🔥 Groups</h3>
        <ul class="group-list">
          <li v-for="group in groups" :key="group.name">
            {{ group.name }}
            <button @click="toggleJoin(group)">
              {{ group.joined ? 'Leave' : 'Join' }}
            </button>
          </li>
        </ul>
      </aside>

      <!-- Center Content: Create Post & Posts List -->
      <main class="forum-main">
        <div class="create-post">
          <h3>Create a Post</h3>
          <form @submit.prevent="submitPost">
            <input
              type="text"
              v-model="newPost.title"
              placeholder="What's on your mind?"
              required
            />
            <textarea
              v-model="newPost.content"
              placeholder="Share your thoughts..."
              required
            ></textarea>

            <!-- Categories inside post -->
            <div class="categories-inside">
              <h4>Categories:</h4>
              <label v-for="cat in categories" :key="cat">
                <input type="checkbox" v-model="selectedCategories" :value="cat" />
                {{ cat }}
              </label>
            </div>

            <button type="submit">Post</button>
          </form>
          <div v-if="message" class="message">{{ message }}</div>
        </div>

        <!-- Posts List -->
        <div class="posts-list">
          <h3>Recent Posts</h3>
          <div v-if="posts.length === 0" class="no-posts">
            No posts available. Be the first to post!
          </div>
          <div v-else class="post-card" v-for="post in posts" :key="post.ID">
            <div class="post-header">
              <img class="post-author-pic" :src="post.authorAvatar" alt="Author" />
              <div>
                <h4>{{ post.Author }}</h4>
                <p class="post-timestamp">{{ formatTimestamp(post.Creation_date) }}</p>
              </div>
            </div>
            <h3 class="post-title">{{ post.Title }}</h3>
            <p class="post-content">{{ post.Content }}</p>
            <div class="post-categories">
              <span v-for="cat in post.Categories.split(',')" :key="cat" class="category-tag">
                {{ cat.trim() }}
              </span>
            </div>
            <button class="comments-toggle" @click="toggleComments(post)">
              {{ post.showComments ? 'Hide Comments' : 'Show Comments' }} ({{ post.comments.length }})
            </button>
            <div v-if="post.showComments" class="comments-section">
              <div v-if="post.comments.length === 0" class="no-comments">
                No comments yet. Be the first to comment!
              </div>
              <div v-else class="comment" v-for="comment in post.comments" :key="comment.id">
                <div class="comment-header">
                  <img class="comment-author-pic" :src="comment.authorAvatar" alt="Comment Author" />
                  <div>
                    <h5>{{ comment.Author }}</h5>
                    <p class="comment-timestamp">{{ formatTimestamp(comment.Creation_date
) }}</p>
                  </div>
                </div>
                <p class="comment-content">{{ comment.Comment }}</p>
              </div>
              <form @submit.prevent="addComment(post)" class="comment-form">
                <textarea
                  v-model="post.newComment"
                  placeholder="Add a comment..."
                  required
                ></textarea>
                <button type="submit">Comment</button>
                <div v-if="post.commentError" class="comment-error">{{ post.commentError }}</div>
              </form>
            </div>
          </div>
        </div>
      </main>

      <!-- Right Sidebar: Other Users -->
      <aside class="forum-info">
        <h3>👥 People to Follow</h3>
        <ul class="user-list">
          <li v-for="other in otherUsers" :key="other.name">
            <img class="mini-profile-pic" :src="other.avatar" alt="User" />
            <span>{{ other.name }}</span>
            <button @click="toggleFollow(other)">
              {{ other.following ? 'Unfollow' : 'Follow' }}
            </button>
          </li>
        </ul>
      </aside>
    </div>

    <!-- Chatbox -->
    <div class="chatbox-container" :class="{ expanded: isChatExpanded }">
      <div class="chatbox-header" @click="toggleChat">
        <span>💬 Chat</span>
        <span class="toggle-icon">{{ isChatExpanded ? '▼' : '▲' }}</span>
      </div>
      <div class="chatbox-content" v-if="isChatExpanded">
        <div class="chat-users">
          <h4>Online Users</h4>
          <ul>
            <li v-for="user in chatUsers" :key="user.name" @click="selectChatUser(user)">
              <span :class="{ 'active-user': selectedChatUser === user }">{{ user.name }}</span>
            </li>
          </ul>
        </div>
        <div class="chat-messages">
          <div v-if="selectedChatUser">
            <h4>Chat with {{ selectedChatUser.name }}</h4>
            <div class="messages">
              <div v-for="msg in selectedChatUser.messages" :key="msg.id" :class="msg.sender === 'self' ? 'self' : 'other'">
                <p>{{ msg.text }}</p>
              </div>
            </div>
            <form @submit.prevent="sendMessage">
              <input
                v-model="newMessage"
                placeholder="Type a message..."
                required
              />
              <button type="submit">Send</button>
            </form>
          </div>
          <div v-else class="no-chat-selected">
            Select a user to start chatting
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ForumPage",
  data() {
    return {
      user: {
        avatar: '',
        name: '',
        email: '',
        username: 'guest'
      },
      newPost: {
        author: '',
        title: '',
        content: ''
      },
      posts: [],
      message: '',
      categories: ['General', 'Help', 'News', 'Programming', 'Design'],
      selectedCategories: [],
      groups: [
        { name: 'Vue.js Devs', joined: false },
        { name: 'Design Lovers', joined: false },
        { name: 'Tech News', joined: true }
      ],
      otherUsers: [
        { name: 'Alice', avatar: 'https://i.pravatar.cc/40?img=1', following: false },
        { name: 'Bob', avatar: 'https://i.pravatar.cc/40?img=2', following: true },
        { name: 'Charlie', avatar: 'https://i.pravatar.cc/40?img=3', following: false }
      ],
      // Chatbox data
      isChatExpanded: false,
      newMessage: '',
      selectedChatUser: null,
      chatUsers: [
        { name: 'Alice', messages: [] },
        { name: 'Bob', messages: [] },
        { name: 'Charlie', messages: [] }
      ]
    };
  },
  async created() {
    try {
      // Fetch user info
      const userRes = await fetch('http://localhost:8080/api/info', {
        method: 'GET',
        credentials: 'include'
      });
      const userData = await userRes.json();
      this.user.name = userData.Username;
      this.user.email = userData.Email;
      this.user.username = userData.Username.toLowerCase();
      this.user.avatar = 'https://api.dicebear.com/7.x/avataaars/svg?seed=' + userData.Username;

      // Fetch posts
      await this.fetchPosts();
    } catch (err) {
      console.error('User fetch failed:', err);
      this.$router.push('/login');
    }
  },
  methods: {
    async fetchPosts() {
      try {
        const res = await fetch('http://localhost:8080/api/getposts', {
          method: 'GET',
          credentials: 'include'
        });
        if (res.ok) {
          const data = await res.json();
          // Initialize posts with comments-related fields
          this.posts = data.map(post => ({
            ...post,
            authorAvatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${post.Author}`,
            comments: [],
            newComment: '',
            showComments: false,
            commentError: '' // Add error field for user feedback
          }));
        } else {
          console.error('Failed to fetch posts');
        }
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    },
    async fetchComments(post) {
      try {
        const res = await fetch(`http://localhost:8080/api/getcomments?postId=${post.Id}`, {
          method: 'GET',
          credentials: 'include'
        });
        if (res.ok) {
          const data = await res.json();
          console.log(data);
        
          // Assuming API returns comments with id, content, author, createdAt
          post.comments = data.map(comment => ({
            
            
    
            ...comment,
            authorAvatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${comment.Author}`
          }));
        } else {
          console.error('Failed to fetch comments for post', post.ID);
          post.commentError = 'Failed to load comments.';
        }
      } catch (error) {
        console.error('Error fetching comments:', error);
        post.commentError = 'Error loading comments.';
      }
    },
    async addComment(post) {
      if (!post.newComment.trim()) {
        post.commentError = 'Comment cannot be empty.';
        return;
      }

      const payload = {
        postId: post.Id.toString(),
        comment: post.newComment.trim(),
        author: this.user.username
      };
      console.log('Sending comment payload:', payload); // Debug log

      try {
        const res = await fetch('http://localhost:8080/api/addcomments', {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(payload)
        });
        if (res.ok) {
          post.commentError = '';
          post.newComment = ''; // Clear only after success
          await this.fetchComments(post); // Refresh comments
        } else {
          console.error('Failed to add comment');
          post.commentError = 'Failed to add comment.';
        }
      } catch (error) {
        console.error('Error adding comment:', error);
        post.commentError = 'Error adding comment.';
      }
    },
    async submitPost() {
      try {
        const res = await fetch('http://localhost:8080/api/posts', {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            ...this.newPost,
            author: this.user.username,
            categories: this.selectedCategories.join(',')
          })
        });

        if (res.ok) {
          this.message = 'Post created successfully!';
          this.newPost.title = '';
          this.newPost.content = '';
          this.selectedCategories = [];
          // Refresh posts after submitting
          await this.fetchPosts();
        } else {
          this.message = 'Failed to create post.';
        }
      } catch (error) {
        console.error('Post creation failed:', error);
        this.message = 'Error submitting post.';
      }
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
        .catch(err => {
          console.error('Logout error:', err);
          this.$router.push('/login');
        });
    },
    toggleFollow(user) {
      user.following = !user.following;
    },
    toggleJoin(group) {
      group.joined = !group.joined;
    },
    goToProfile() {
      this.$router.push('/profile');
    },
    toggleChat() {
      this.isChatExpanded = !this.isChatExpanded;
    },
    selectChatUser(user) {
      this.selectedChatUser = user;
    },
    sendMessage() {
      if (this.newMessage.trim() && this.selectedChatUser) {
        this.selectedChatUser.messages.push({
          id: Date.now(),
          text: this.newMessage,
          sender: 'self'
        });
        this.newMessage = '';
        // Simulate a reply after a short delay
        setTimeout(() => {
          this.selectedChatUser.messages.push({
            id: Date.now(),
            text: `Hi! Thanks for your message.`,
            sender: 'other'
          });
        }, 1000);
      }
    },
    formatTimestamp(timestamp) {
      return new Date(timestamp).toLocaleString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      });
    },
    toggleComments(post) {
      post.showComments = !post.showComments;
      if (post.showComments && post.comments.length === 0) {
        this.fetchComments(post);
      }
    }
  }
};
</script>

<style scoped>
.forum-container {
  font-family: 'Inter', sans-serif;
  background: linear-gradient(135deg, #e0e7ff, #c7d2fe);
  min-height: 100vh;
  padding: 1.5rem;
  position: relative;
}

.forum-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: rgba(79, 70, 229, 0.9);
  backdrop-filter: blur(12px);
  color: white;
  padding: 1rem 2rem;
  border-radius: 1rem;
  box-shadow: 0 6px 24px rgba(79, 70, 229, 0.3);
  margin-bottom: 2rem;
  transition: transform 0.3s ease;
}

.forum-header:hover {
  transform: translateY(-4px);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.profile-pic, .mini-profile-pic, .post-author-pic, .comment-author-pic {
  width: 3.5rem;
  height: 3.5rem;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffffff;
  transition: transform 0.3s ease;
}

.profile-pic:hover, .post-author-pic:hover, .comment-author-pic:hover {
  transform: scale(1.1);
}

.mini-profile-pic, .post-author-pic, .comment-author-pic {
  width: 2.5rem;
  height: 2.5rem;
}

.header-buttons {
  display: flex;
  gap: 0.75rem;
}

.profile-btn, .logout-btn, .comments-toggle {
  background: #ffffff;
  color: #4f46e5;
  font-weight: 600;
  padding: 0.5rem 1.25rem;
  border-radius: 0.75rem;
  border: none;
  cursor: pointer;
  box-shadow: 0 3px 12px rgba(79, 70, 229, 0.2);
  transition: all 0.3s ease;
}

.profile-btn:hover, .logout-btn:hover, .comments-toggle:hover {
  background: #e0e7ff;
  transform: translateY(-2px);
}

.forum-grid {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  gap: 1.5rem;
}

@media (max-width: 1024px) {
  .forum-grid {
    grid-template-columns: 1fr 2fr;
  }
  .forum-info {
    display: none;
  }
}

@media (max-width: 768px) {
  .forum-grid {
    grid-template-columns: 1fr;
  }
  .forum-sidebar {
    display: none;
  }
}

.forum-sidebar, .forum-main, .forum-info {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border-radius: 1rem;
  padding: 1.5rem;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.forum-sidebar:hover, .forum-main:hover, .forum-info:hover {
  transform: translateY(-4px);
}

h3 {
  color: #4f46e5;
  margin-bottom: 1.25rem;
  font-size: 1.5rem;
}

.group-list, .user-list {
  list-style: none;
  padding: 0;
}

.group-list li, .user-list li {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 1rem;
  font-size: 0.95rem;
}

.group-list button, .user-list button {
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: white;
  border: none;
  padding: 0.4rem 0.9rem;
  border-radius: 0.5rem;
  font-size: 0.85rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.group-list button:hover, .user-list button:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: scale(1.05);
}

.forum-main h3 {
  color: #4f46e5;
}

.create-post input, .create-post textarea {
  width: 100%;
  margin-bottom: 1rem;
  padding: 0.75rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 1rem;
  background: #f8fafc;
  transition: all 0.3s ease;
}

.create-post input:focus, .create-post textarea:focus {
  border-color: #4f46e5;
  background: #ffffff;
  outline: none;
}

.create-post textarea {
  min-height: 120px;
  resize: vertical;
}

.create-post button {
  padding: 0.75rem 1.25rem;
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: white;
  border: none;
  border-radius: 0.75rem;
  font-size: 1rem;
  cursor: pointer;
  width: 100%;
  font-weight: 600;
  transition: all 0.3s ease;
}

.create-post button:hover {
  transform: translateY(-2px);
  background: linear-gradient(135deg, #6366f1, #4f46e5);
}

.categories-inside {
  margin-bottom: 1.25rem;
}

.categories-inside h4 {
  margin-bottom: 0.75rem;
  font-weight: 600;
  color: #4f46e5;
}

.categories-inside label {
  display: inline-flex;
  align-items: center;
  margin-right: 0.75rem;
  font-size: 0.9rem;
}

.categories-inside input[type="checkbox"] {
  margin-right: 0.4rem;
  accent-color: #4f46e5;
}

.message {
  margin-top: 1rem;
  font-weight: 600;
  color: #10b981;
}

/* Posts List Styles */
.posts-list {
  margin-top: 2rem;
}

.no-posts {
  text-align: center;
  color: #6b7280;
  font-size: 1rem;
  padding: 1rem;
}

.post-card {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border-radius: 1rem;
  padding: 1.25rem;
  margin: 1.5rem 0;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.post-card:hover {
  transform: translateY(-4px);
}

.post-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
}

.post-header h4 {
  color: #1f2937;
  font-size: 1rem;
  margin: 0;
}

.post-timestamp {
  color: #6b7280;
  font-size: 0.85rem;
  margin: 0;
}

.post-title {
  color: #4f46e5;
  font-size: 1.25rem;
  margin: 0.5rem 0;
}

.post-content {
  color: #4b5563;
  font-size: 0.95rem;
  margin-bottom: 0.75rem;
}

.post-categories {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
}

.category-tag {
  background: #e0e7ff;
  color: #4f46e5;
  padding: 0.25rem 0.5rem;
  border-radius: 0.5rem;
  font-size: 0.85rem;
  font-weight: 500;
}

/* Comments Styles */
.comments-section {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}

.no-comments {
  text-align: center;
  color: #6b7280;
  font-size: 0.9rem;
  padding: 0.5rem;
}

.comment {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 0.75rem;
  padding: 0.75rem;
  margin-bottom: 0.75rem;
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.comment-header h5 {
  color: #1f2937;
  font-size: 0.9rem;
  margin: 0;
}

.comment-timestamp {
  color: #6b7280;
  font-size: 0.8rem;
  margin: 0;
}

.comment-content {
  color: #4b5563;
  font-size: 0.9rem;
}

.comment-form {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-top: 1rem;
}

.comment-form textarea {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  font-size: 0.9rem;
  background: #f8fafc;
  min-height: 80px;
  resize: vertical;
}

.comment-form textarea:focus {
  border-color: #4f46e5;
  background: #ffffff;
  outline: none;
}

.comment-form button {
  padding: 0.5rem 1rem;
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: white;
  border: none;
  border-radius: 0.75rem;
  font-size: 0.9rem;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s ease;
}

.comment-form button:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: translateY(-2px);
}

.comment-error {
  color: #ef4444;
  font-size: 0.85rem;
  margin-top: 0.5rem;
}

/* Chatbox Styles */
.chatbox-container {
  position: fixed;
  bottom: 1.5rem;
  right: 1.5rem;
  width: 360px;
  max-height: 500px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(12px);
  border-radius: 1rem;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.15);
  overflow: hidden;
  transition: all 0.3s ease;
  z-index: 1000;
}

.chatbox-container.expanded {
  transform: translateY(0);
}

.chatbox-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem 1.25rem;
  background: #4f46e5;
  color: white;
  cursor: pointer;
  font-weight: 600;
}

.toggle-icon {
  font-size: 0.9rem;
}

.chatbox-content {
  display: flex;
  height: 400px;
}

.chat-users {
  width: 120px;
  padding: 1rem;
  border-right: 1px solid #e5e7eb;
  overflow-y: auto;
}

.chat-users h4 {
  font-size: 0.9rem;
  color: #4f46e5;
  margin-bottom: 0.75rem;
}

.chat-users ul {
  list-style: none;
  padding: 0;
}

.chat-users li {
  padding: 0.5rem 0;
  cursor: pointer;
  font-size: 0.9rem;
  color: #1f2937;
}

.chat-users li:hover {
  background: #f1f5f9;
  border-radius: 0.5rem;
}

.active-user {
  font-weight: 600;
  color: #4f46e5;
}

.chat-messages {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 1rem;
}

.chat-messages h4 {
  font-size: 0.9rem;
  color: #4f46e5;
  margin-bottom: 0.75rem;
}

.messages {
  flex: 1;
  overflow-y: auto;
  margin-bottom: 1rem;
}

.messages .self, .messages .other {
  margin-bottom: 0.75rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.75rem;
  max-width: 80%;
}

.messages .self {
  background: #4f46e5;
  color: white;
  align-self: flex-end;
  margin-left: auto;
}

.messages .other {
  background: #e5e7eb;
  color: #1f2937;
  align-self: flex-start;
}

.chat-messages form {
  display: flex;
  gap: 0.5rem;
}

.chat-messages input {
  flex: 1;
  padding: 0.5rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  font-size: 0.9rem;
}

.chat-messages button {
  padding: 0.5rem 1rem;
  background: #4f46e5;
  color: white;
  border: none;
  border-radius: 0.5rem;
  font-size: 0.9rem;
  cursor: pointer;
}

.chat-messages button:hover {
  background: #6366f1;
}

.no-chat-selected {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
  font-size: 0.9rem;
}

@media (max-width: 480px) {
  .chatbox-container {
    width: 100%;
    bottom: 0;
    right: 0;
    border-radius: 1rem 1rem 0 0;
  }

  .post-card {
    padding: 1rem;
  }

  .post-title {
    font-size: 1.1rem;
  }

  .post-content {
    font-size: 0.9rem;
  }

  .post-author-pic, .comment-author-pic {
    width: 2rem;
    height: 2rem;
  }

  .comment {
    padding: 0.5rem;
  }

  .comment-header h5 {
    font-size: 0.85rem;
  }

  .comment-content {
    font-size: 0.85rem;
  }

  .comment-form textarea {
    font-size: 0.85rem;
    min-height: 60px;
  }

  .comment-form button {
    font-size: 0.85rem;
    padding: 0.4rem 0.8rem;
  }
}
</style>