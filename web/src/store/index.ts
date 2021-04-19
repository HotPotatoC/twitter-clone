import {
  createStore,
  createLogger,
  useStore as useVuexStore,
  ModuleTree,
} from 'vuex'
import { authModule, AuthModule } from '../modules/auth/store'
import { tweetsModule, TweetsModule } from '../modules/tweets/store'
import { profileModule, ProfileModule } from '../modules/user/store'

type StoreModules = {
  auth: AuthModule
  tweets: TweetsModule
  profile: ProfileModule
}

type Store = AuthModule<Pick<StoreModules, 'auth'>> &
  TweetsModule<Pick<StoreModules, 'tweets'>> &
  ProfileModule<Pick<StoreModules, 'profile'>>

const modules: ModuleTree<any> = {
  authModule,
  tweetsModule,
  profileModule,
}

export const store = createStore({
  plugins: process.env.NODE_ENV === 'production' ? [] : [createLogger()],
  modules,
})

export function useStore() {
  return useVuexStore() as Store
}
