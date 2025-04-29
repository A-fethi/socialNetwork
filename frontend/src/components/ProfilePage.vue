<template>
  <div class="profile-container">
    <!-- Header Bar -->
    <header class="profile-header">
      <div class="app-name">MySocialApp</div>
      <div class="user-actions">
        <button @click="goToHome" class="home-btn">Back to Home</button>
        <span>{{ user.firstname }} {{ user.lastname }}</span>
        <button @click="logout" class="logout-btn">Logout</button>
      </div>
    </header>

    <!-- Profile Card -->
    <div class="profile-card-wrapper">
      <div class="profile-card">
        <div class="profile-image">
          <img :src="user.avatar" alt="Profile Picture" />
        </div>
        <div class="profile-details">
          <h2>{{ user.firstname }} {{ user.lastname }}</h2>
          <p class="username">@{{ user.username }}</p>
          <p class="email">{{ user.email }}</p>
          <p class="bio">{{ user.bio || "No bio provided" }}</p>
          
          <!-- Privacy Toggle (only visible to profile owner) -->
          <div v-if="isOwnProfile" class="privacy-toggle">
            <label class="switch">
              <input type="checkbox" v-model="user.isPrivate" @change="updatePrivacy">
              <span class="slider round"></span>
            </label>
            <span>{{ user.isPrivate ? 'Private Profile' : 'Public Profile' }}</span>
          </div>
        </div>
        
        <!-- Profile Stats -->
        <div v-if="canViewProfile" class="profile-stats">
          <div class="stat">
            <strong>{{ followers.length }}</strong>
            <span>Followers</span>
          </div>
          <div class="stat">
            <strong>{{ following.length }}</strong>
            <span>Following</span>
          </div>
          <div class="stat">
            <strong>{{ userPosts.length }}</strong>
            <span>Posts</span>
          </div>
        </div>
        
        <!-- Follow/Unfollow Button (only visible to other users) -->
        <div v-if="!isOwnProfile" class="follow-action">
          <button v-if="!isFollowing" @click="followUser" class="follow-btn">Follow</button>
          <button v-else @click="unfollowUser" class="unfollow-btn">Unfollow</button>
        </div>
      </div>
    </div>
    
    <!-- Content Sections (only visible if can view profile) -->
    <div v-if="canViewProfile" class="profile-content">
      <!-- User Posts -->
      <div class="profile-section">
        <h3>Posts</h3>
        <div v-if="userPosts.length === 0" class="no-content">
          No posts yet.
        </div>
        <div v-else class="posts-grid">
          <div v-for="post in userPosts" :key="post.id" class="post-card">
            <h4>{{ post.title }}</h4>
            <p>{{ post.content }}</p>
            <span class="post-date">{{ formatTimestamp(post.creation_date) }}</span>
          </div>
        </div>
      </div>
      
      <!-- Followers Section -->
      <div class="profile-section">
        <h3>Followers</h3>
        <div v-if="followers.length === 0" class="no-content">
          No followers yet.
        </div>
        <div v-else class="users-grid">
          <div v-for="follower in followers" :key="follower.id" class="user-card" @click="viewProfile(follower.id)">
            <img :src="follower.avatar" alt="Follower" class="user-avatar">
            <span>{{ follower.firstname }} {{ follower.lastname }}</span>
          </div>
        </div>
      </div>
      
      <!-- Following Section -->
      <div class="profile-section">
        <h3>Following</h3>
        <div v-if="following.length === 0" class="no-content">
          Not following anyone yet.
        </div>
        <div v-else class="users-grid">
          <div v-for="followedUser in following" :key="followedUser.id" class="user-card" @click="viewProfile(followedUser.id)">
            <img :src="followedUser.avatar" alt="Following" class="user-avatar">
            <span>{{ followedUser.firstname }} {{ followedUser.lastname }}</span>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Private Profile Message -->
    <div v-else-if="!canViewProfile && !isOwnProfile" class="private-profile-message">
      <div class="lock-icon">🔒</div>
      <h3>This profile is private</h3>
      <p>Follow this user to see their profile content.</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProfilePage",
  data() {
    return {
      user: {
        id: null,
        avatar: "",
        firstname: "",
        lastname: "",
        username: "",
        email: "",
        bio: "",
        isPrivate: false
      },
      profileUserId: null,
      isOwnProfile: true,
      isFollowing: false,
      followers: [],
      following: [],
      userPosts: []
    };
  },
  computed: {
    canViewProfile() {
      // Can view if: it's own profile, profile is public, or user is following a private profile
      return this.isOwnProfile || !this.user.isPrivate || this.isFollowing;
    }
  },
  methods: {
    logout() {
      // ... existing code ...
    },
    goToHome() {
      // ... existing code ...
    },
    async fetchUserProfile() {
      try {
        // Determine if viewing own profile or someone else's
        const urlParams = new URLSearchParams(window.location.search);
        this.profileUserId = urlParams.get('id');
        
        if (this.profileUserId) {
          // Viewing someone else's profile
          this.isOwnProfile = false;
          const res = await fetch(`http://localhost:8080/api/profile/${this.profileUserId}`, {
            method: 'GET',
            credentials: 'include'
          });
          
          if (!res.ok) throw new Error('Failed to fetch profile');
          
          const data = await res.json();
          this.user = {
            id: data.id,
            firstname: data.firstname,
            lastname: data.lastname,
            username: data.username,
            email: data.email,
            bio: data.bio,
            isPrivate: data.is_private,
            avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${data.username}`
          };
          
          // Check if current user is following this profile
          this.checkFollowStatus();
        } else {
          // Viewing own profile
          this.isOwnProfile = true;
          const res = await fetch('http://localhost:8080/api/info', {
            method: 'GET',
            credentials: 'include'
          });
          
          if (!res.ok) throw new Error('Failed to fetch profile');
          
          const data = await res.json();
          this.user = {
            id: data.ID,
            firstname: data.Firstname,
            lastname: data.Lastname,
            username: data.Username,
            email: data.Email,
            bio: data.Bio,
            isPrivate: data.IsPrivate,
            avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${data.Username}`
          };
        }
        
        // Fetch additional profile data if can view
        if (this.canViewProfile) {
          await Promise.all([
            this.fetchFollowers(),
            this.fetchFollowing(),
            this.fetchUserPosts()
          ]);
        }
      } catch (err) {
        console.error('Profile fetch failed:', err);
        this.$router.push('/login');
      }
    },
    async checkFollowStatus() {
      try {
        const res = await fetch(`http://localhost:8080/api/follow/status/${this.profileUserId}`, {
          method: 'GET',
          credentials: 'include'
        });
        
        if (!res.ok) throw new Error('Failed to check follow status');
        
        const data = await res.json();
        this.isFollowing = data.is_following;
      } catch (err) {
        console.error('Failed to check follow status:', err);
      }
    },
    async fetchFollowers() {
      try {
        const userId = this.profileUserId || 'me';
        const res = await fetch(`http://localhost:8080/api/followers/${userId}`, {
          method: 'GET',
          credentials: 'include'
        });
        
        if (!res.ok) throw new Error('Failed to fetch followers');
        
        const data = await res.json();
        this.followers = data.map(follower => ({
          ...follower,
          avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${follower.username}`
        }));
      } catch (err) {
        console.error('Failed to fetch followers:', err);
      }
    },
    async fetchFollowing() {
      try {
        const userId = this.profileUserId || 'me';
        const res = await fetch(`http://localhost:8080/api/following/${userId}`, {
          method: 'GET',
          credentials: 'include'
        });
        
        if (!res.ok) throw new Error('Failed to fetch following');
        
        const data = await res.json();
        this.following = data.map(followed => ({
          ...followed,
          avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${followed.username}`
        }));
      } catch (err) {
        console.error('Failed to fetch following:', err);
      }
    },
    async fetchUserPosts() {
      try {
        const userId = this.profileUserId || 'me';
        const res = await fetch(`http://localhost:8080/api/posts/user/${userId}`, {
          method: 'GET',
          credentials: 'include'
        });
        
        if (!res.ok) throw new Error('Failed to fetch user posts');
        
        this.userPosts = await res.json();
      } catch (err) {
        console.error('Failed to fetch user posts:', err);
      }
    },
    async followUser() {
      try {
        const res = await fetch(`http://localhost:8080/api/follow/${this.profileUserId}`, {
          method: 'POST',
          credentials: 'include'
        });
        
        if (!res.ok) throw new Error('Failed to follow user');
        
        this.isFollowing = true;
        // Refresh followers count
        await this.fetchFollowers();
      } catch (err) {
        console.error('Failed to follow user:', err);
      }
    },
    async unfollowUser() {
      try {
        const res = await fetch(`http://localhost:8080/api/unfollow/${this.profileUserId}`, {
          method: 'POST',
          credentials: 'include'
        });
        
        if (!res.ok) throw new Error('Failed to unfollow user');
        
        this.isFollowing = false;
        // Refresh followers count
        await this.fetchFollowers();
      } catch (err) {
        console.error('Failed to unfollow user:', err);
      }
    },
    async updatePrivacy() {
      try {
        const res = await fetch('http://localhost:8080/api/profile/privacy', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
          body: JSON.stringify({ is_private: this.user.isPrivate })
        });
        
        if (!res.ok) throw new Error('Failed to update privacy settings');
        
        // Show success message
        alert(this.user.isPrivate ? 'Profile set to private' : 'Profile set to public');
      } catch (err) {
        console.error('Failed to update privacy settings:', err);
        // Revert the toggle if update fails
        this.user.isPrivate = !this.user.isPrivate;
      }
    },
    formatTimestamp(timestamp) {
      return new Date(timestamp).toLocaleString();
    },
    viewProfile(userId) {
      this.$router.push(`/profile?id=${userId}`);
    }
  },
  created() {
    this.fetchUserProfile();
  }
};
</script>

<style scoped>
.profile-container {
  font-family: 'Inter', sans-serif;
  min-height: 100vh;
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(135deg, #f5f7fa, #e2e8f0);
  position: relative;
  overflow: hidden;
}

/* Google-inspired background wave effect */
.profile-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at 10% 20%, rgba(79, 70, 229, 0.1) 0%, transparent 20%),
              radial-gradient(circle at 90% 80%, rgba(79, 70, 229, 0.1) 0%, transparent 20%);
  z-index: 0;
  opacity: 0.5;
}

.profile-header {
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
  width: 100%;
  max-width: 1200px;
  transition: transform 0.3s ease;
  z-index: 1;
}

.profile-header:hover {
  transform: translateY(-4px);
}

.app-name {
  font-size: 1.5rem;
  font-weight: 700;
}

.user-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-actions span {
  font-size: 1rem;
  font-weight: 500;
}

.home-btn, .logout-btn {
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

.home-btn:hover, .logout-btn:hover {
  background: #e0e7ff;
  transform: translateY(-2px);
}

.profile-card-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  width: 100%;
  z-index: 1;
}

.profile-card {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border-radius: 1rem;
  padding: 2rem;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 100%;
  max-width: 600px;
  transition: transform 0.3s ease;
}

.profile-card:hover {
  transform: translateY(-4px);
}

.profile-image img {
  width: 8rem;
  height: 8rem;
  border-radius: 50%;
  border: 4px solid #ffffff;
  object-fit: cover;
  margin-top: -4rem;
  background: #ffffff;
  transition: transform 0.3s ease;
}

.profile-image img:hover {
  transform: scale(1.1);
}

.profile-details h2 {
  margin: 1rem 0 0.5rem;
  font-size: 1.75rem;
  color: #1f2937;
}

.location {
  color: #6b7280;
  font-size: 0.95rem;
  margin-bottom: 0.75rem;
}

.bio {
  font-size: 0.9rem;
  color: #4b5563;
  white-space: pre-line;
  margin-bottom: 1.5rem;
}

.profile-stats {
  display: flex;
  justify-content: space-around;
  margin: 1.5rem 0;
}

.stat {
  text-align: center;
}

.stat strong {
  display: block;
  font-size: 1.5rem;
  color: #1f2937;
}

.stat span {
  font-size: 0.85rem;
  color: #6b7280;
}

.show-more {
  background: linear-gradient(135deg, #4f46e5, #6366f1);
  color: white;
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
}

.show-more:hover {
  background: linear-gradient(135deg, #6366f1, #4f46e5);
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  .profile-header {
    flex-direction: column;
    align-items: flex-start;
    padding: 1rem;
  }

  .user-actions {
    margin-top: 0.75rem;
    flex-wrap: wrap;
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
}

@media (max-width: 480px) {
  .profile-container {
    padding: 1rem;
  }

  .profile-card {
    max-width: 100%;
  }

  .profile-stats {
    flex-direction: column;
    gap: 1rem;
  }

  .user-actions {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
}

/* New styles for profile content */
.profile-content {
  width: 100%;
  max-width: 800px;
  margin-top: 2rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.profile-section {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border-radius: 1rem;
  padding: 1.5rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.profile-section h3 {
  font-size: 1.25rem;
  color: #4f46e5;
  margin-bottom: 1rem;
  border-bottom: 2px solid #e5e7eb;
  padding-bottom: 0.5rem;
}

.no-content {
  color: #6b7280;
  font-style: italic;
  text-align: center;
  padding: 1rem;
}

.posts-grid, .users-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}

.post-card {
  background: white;
  border-radius: 0.5rem;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s ease;
}

.post-card:hover {
  transform: translateY(-2px);
}

.post-card h4 {
  font-size: 1rem;
  margin-bottom: 0.5rem;
  color: #1f2937;
}

.post-card p {
  font-size: 0.875rem;
  color: #4b5563;
  margin-bottom: 0.5rem;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}

.post-date {
  font-size: 0.75rem;
  color: #6b7280;
}

.user-card {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  background: white;
  border-radius: 0.5rem;
  padding: 0.75rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.user-card:hover {
  background-color: #f3f4f6;
}

.user-avatar {
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 50%;
  object-fit: cover;
}

.username, .email {
  color: #6b7280;
  font-size: 0.9rem;
  margin-bottom: 0.5rem;
}

.privacy-toggle {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-top: 1rem;
}

/* Toggle Switch */
.switch {
  position: relative;
  display: inline-block;
  width: 50px;
  height: 24px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 16px;
  width: 16px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: .4s;
}

input:checked + .slider {
  background-color: #4f46e5;
}

input:focus + .slider {
  box-shadow: 0 0 1px #4f46e5;
}

input:checked + .slider:before {
  transform: translateX(26px);
}

.slider.round {
  border-radius: 24px;
}

.slider.round:before {
  border-radius: 50%;
}

.follow-btn, .unfollow-btn {
  background: #4f46e5;
  color: white;
  border: none;
  border-radius: 0.5rem;
  padding: 0.5rem 1.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s ease;
  margin-top: 1rem;
}

.follow-btn:hover {
  background: #4338ca;
}

.unfollow-btn {
  background: #e5e7eb;
  color: #1f2937;
}

.unfollow-btn:hover {
  background: #d1d5db;
}

.private-profile-message {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  border-radius: 1rem;
  padding: 2rem;
  text-align: center;
  margin-top: 2rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  max-width: 400px;
}

.lock-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.private-profile-message h3 {
  font-size: 1.25rem;
  color: #1f2937;
  margin-bottom: 0.5rem;
}

.private-profile-message p {
  color: #6b7280;
}
</style>