import authRoutes from './auth/routes'
import homeRoutes from './home/routes'
import userRoutes from './user/routes'
import searchRoutes from './search/routes'

export default [...authRoutes, homeRoutes, searchRoutes, userRoutes]
