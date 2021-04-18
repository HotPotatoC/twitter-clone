import { RouteRecordRaw } from 'vue-router'
import BaseLayout from '../../layouts/BaseLayout.vue'
import { makeRoutesWithLayout } from '../../services/routes'

const routes: RouteRecordRaw[] = [
  {
    path: '',
    component: () => import('./index.vue'),
    meta: {
      requiresAuth: true,
    },
  },
]

export default makeRoutesWithLayout('/search', BaseLayout, routes)
