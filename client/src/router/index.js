import { createRouter, createWebHistory } from "vue-router";

const Home = () => import("../views/Home.vue");

const routes = [
  { path: "/", component: Home },
  { path: "/cart", component: () => import("../views/Cart.vue") },
  { path: "/login", component: () => import("../views/Login.vue") },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
