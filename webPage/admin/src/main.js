import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import axios from 'axios'

import './assets/main.css'

// antui plugin
import { Button, Form, Input } from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.use(Button)
app.use(Form)
app.use(Input)

app.provide('$axios', axios)

app.mount('#app')
