import { createRouter, createWebHistory } from "vue-router";
import HomePageView from "../views/HomePageView.vue";
import LoginView from "../views/LoginView.vue";
import DashboardView from "../views/DashboardView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomePageView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/dashboard",
    name: "dasboard",
    component: DashboardView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
