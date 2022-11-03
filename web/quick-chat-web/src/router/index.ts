import { createRouter, createWebHashHistory } from "vue-router";

import LoginView from "@/views/LoginView/LoginView.vue";
import HomeView from "@/views/HomeView/HomeView.vue";
import HomeChatView from "@/views/HomeView/HomeChatView.vue";
import HomeContactView from "@/views/HomeView/HomeContactView.vue";
import HomeMomentView from "@/views/HomeView/HomeMomentView.vue";
import HomeMyView from "@/views/HomeView/HomeMyView.vue";

export const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/home",
    name: "home",
    redirect: "/home/chat",
    component: HomeView,
    children: [
      { path: "/home/chat", name: "chat", component: HomeChatView },
      { path: "/home/contact", name: "contact", component: HomeContactView },
      { path: "/home/moment", name: "moment", component: HomeMomentView },
      { path: "/home/my", name: "my", component: HomeMyView },
    ],
  },
];

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
