import { createRouter, createWebHistory } from "vue-router";
import UserLogin from "./components/UserLogin.vue";
import UserRegister from "./components/UserRegister.vue";
import AuthMiddleware from "./components/AuthMiddleware.vue";
import ChechInfo from "./components/ChechInfo.vue";
import ForumPage from "./components/ForumPage.vue";

import WebSocket from "./components/WebSocket.vue";
import MyGroups from "./components/MyGroups.vue";
import GroupPage from "./components/GroupPage.vue";
import ProfileInfo from "./components/ProfileInfo.vue";
import SettingsPage from "./components/SettingsPage.vue";
import ErrorPage from "./components/ErrorPage.vue";
import { CheckLogin } from "./api/middle";

const routes = [
  {
    path: "/login",
    component: UserLogin,
    beforeEnter: async (to, from, next) => {
      const isLoggedIn = await CheckLogin();
      if (isLoggedIn) {
        next("/home");
      } else {
        next();
      }
    },
    meta: { requiresAuth: false },
  },
  {
    path: "/register",
    component: UserRegister,
    beforeEnter: async (to, from, next) => {
      const isLoggedIn = await CheckLogin();
      if (isLoggedIn) {
        next("/home");
      } else {
        next();
      }
    },
  },
  { path: "/", component: AuthMiddleware },
  { path: "/info", component: ChechInfo },
  { path: "/home", component: ForumPage },

  { path: "/messages", component: WebSocket },
  { path: "/mygroups", component: MyGroups },
  { path: "/group/:id", component: GroupPage },
  { path: "/profile/:id", component: ProfileInfo },
  { path: "/settings", component: SettingsPage },

  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: ErrorPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
