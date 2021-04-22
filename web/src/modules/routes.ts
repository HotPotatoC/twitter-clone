import authRoutes from './auth/routes'
import homeRoutes from './home/routes'
import userRoutes from './user/routes'
import searchRoutes from './search/routes'
import { RouteRecordRaw } from 'vue-router'
import Root from './root/index.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: Root,
  },
  ...authRoutes,
  homeRoutes,
  searchRoutes,
  userRoutes,
]

export default routes
