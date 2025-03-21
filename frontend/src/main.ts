import { createApp } from 'vue'
import App from './App.vue'

import { createMemoryHistory, createRouter } from 'vue-router'

import HomeView from './components/Home.vue'
import Pet from './components/Pet.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/Pet', component: Pet },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

createApp(App)
  .use(router)
  .mount('#app')