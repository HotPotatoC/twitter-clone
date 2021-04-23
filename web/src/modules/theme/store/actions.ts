import { ActionTree } from 'vuex'
import { Mutations, MutationTypes } from './mutations'
import { State } from './state'
import { AugmentedActionContext } from '../../../store'
import { Theme } from '../types'

export enum ActionTypes {
  TOGGLE_THEME = 'TOGGLE_THEME',
}

export type Actions = {
  [ActionTypes.TOGGLE_THEME](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: Theme
  ): Promise<void>
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.TOGGLE_THEME]({ commit }, theme): Promise<void> {
    commit(MutationTypes.SET_THEME, theme)
  },
}
