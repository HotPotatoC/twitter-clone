export interface AuthStatus {
  statusCode?: number
  message?: string
  isLoggedIn: boolean
}

export interface UserData {
  id: number
  name: string
  email: string
}

export interface State {
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
    email: '',
    name: '',
  },
}
