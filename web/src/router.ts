import { createMemoryHistory, createRouter } from "vue-router";

import HomeView from "./views/Menu.vue";
import TaskView from "./views/Task.vue";

const routes = [
  { path: "/", component: HomeView },
  { path: "/task/:id", component: TaskView, props: true },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;
