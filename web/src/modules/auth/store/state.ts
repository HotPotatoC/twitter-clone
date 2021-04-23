import { UserData } from '../types'

export type AuthStatus = {
  statusCode?: number
  message?: string
  isLoggedIn: boolean
}

export type State = {
  accessToken: string
  authStatus: AuthStatus
  user: UserData
}

export const state: State = {
  authStatus: {
    isLoggedIn: false,
  },
  accessToken: '',
  user: {
    id: 0,
    name: '',
    handle: '',
    email: '',
  },
}
