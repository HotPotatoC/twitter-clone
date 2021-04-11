import {
  ActionContext,
  Store as VuexStore,
  CommitOptions,
  DispatchOptions,
} from 'vuex'

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
