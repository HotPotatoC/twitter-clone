import {
  createStore,
  createLogger,
  useStore as useVuexStore,
  ModuleTree,
  Store as VuexStore,
  ActionContext,
  CommitOptions,
  DispatchOptions,
} from 'vuex'
import { ThemeModule } from './modules/theme/store'
import { authModule, AuthModule } from './modules/auth/store'
import { tweetsModule, TweetsModule } from './modules/tweets/store'
import { profileModule, ProfileModule } from './modules/user/store'

export type AnyRecord = Record<any, any>

export type AugmentedActionContext<M extends AnyRecord, S> = {
  commit<K extends keyof M>(
    key: K,
    payload: Parameters<M[K]>[1]
  ): ReturnType<M[K]>
} & Omit<ActionContext<S, S>, 'commit'>

export type AugmentedModule<
  S = any,
  M extends AnyRecord = any,
  G extends AnyRecord = any,
  A extends AnyRecord = any
> = Omit<VuexStore<S>, 'commit' | 'getters' | 'dispatch'> & {
  commit<K extends keyof M, P extends Parameters<M[K]>[1]>(
    key: K,
    payload?: P,
    options?: CommitOptions
  ): ReturnType<M[K]>
} & {
  getters: {
    [K in keyof G]: ReturnType<G[K]>
  }
} & {
  dispatch<K extends keyof A>(
    key: K,
    payload?: Parameters<A[K]>[1],
    options?: DispatchOptions
  ): ReturnType<A[K]>
}

type StoreModules = {
  theme: ThemeModule
  auth: AuthModule
  tweets: TweetsModule
  profile: ProfileModule
}

export type Store = ThemeModule<Pick<StoreModules, 'theme'>> &
  AuthModule<Pick<StoreModules, 'auth'>> &
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
