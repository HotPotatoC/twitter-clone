import {
  createStore,
  createLogger,
  useStore as useVuexStore,
  ModuleTree,
} from 'vuex'
import { authModule, AuthModule } from './auth'
import { tweetsModule, TweetsModule } from './tweets'

type StoreModules = {
  auth: AuthModule
  tweets: TweetsModule
}

type Store = AuthModule<Pick<StoreModules, 'auth'>> &
  TweetsModule<Pick<StoreModules, 'tweets'>>

const modules: ModuleTree<any> = {
  authModule,
  tweetsModule,
}

export const store = createStore({
  plugins: process.env.NODE_ENV === 'production' ? [] : [createLogger()],
  modules,
})

export function useStore() {
  return useVuexStore() as Store
}
