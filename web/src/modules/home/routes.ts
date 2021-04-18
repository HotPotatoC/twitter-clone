import { RouteRecordRaw } from 'vue-router'
import BaseLayout from '../../layouts/BaseLayout.vue'
import { makeRoutesWithLayout } from '../../services/routes'
import Home from './index.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '',
    component: Home,
    meta: {
      requiresAuth: true,
    },
  },
]

export default makeRoutesWithLayout('/home', BaseLayout, routes)
