import { RouteRecordRaw } from 'vue-router'
import BaseLayout from '../../modules/layouts/BaseLayout.vue'
import { makeRoutesWithLayout } from '../../utils/routes'
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
