import { Theme } from '../types'

export type State = {
  theme: Theme
}

export const state: State = {
  theme: localStorage.getItem('theme'),
}
