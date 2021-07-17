import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '../components/TopPage.vue'
import DiariesPage from '../components/DiariesPage.vue'
import NewPage from '../components/NewPage.vue'

const routes = [
  { path: "/", name: "TopPage", component: TopPage },
  { path: "/diaries", name: "Diaries", component: DiariesPage },
  { path: "/diaries/new", name: "New", component: NewPage },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
