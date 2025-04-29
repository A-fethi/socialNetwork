import { createRouter, createWebHistory } from 'vue-router';
import UserLogin from './components/UserLogin.vue';
import UserRegister from './components/UserRegister.vue';
import AuthMiddleware from './components/AuthMiddleware.vue';
import ChechInfo from './components/ChechInfo.vue';
import ForumPage from './components/ForumPage.vue';
import ProfilePage from './components/ProfilePage.vue';
const routes = [
  { path: '/login', component: UserLogin },
  { path: '/register', component: UserRegister },
  { path: '/', component: AuthMiddleware },
  { path :"/info", component: ChechInfo},
  { path: "/home", component: ForumPage },
  { path : '/profile',component: ProfilePage },
];
const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
