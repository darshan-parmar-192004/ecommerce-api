import { createApp } from 'vue'
import { devtools } from '@vue/devtools'
import './style.css'
import App from './App.vue'

const app = createApp(App)

if (import.meta.env.DEV) {
  app.use(devtools)
}

app.mount('#app')
