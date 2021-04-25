import { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { Store } from './store'

export type GuardContext = {
  to: RouteLocationNormalized
  from: RouteLocationNormalized
  next: NavigationGuardNext
  store: Store
}

export type ResponseItems<T> = {
  items: T[]
  total_items: number
}

export type Month =
  | 'January'
  | 'February'
  | 'March'
  | 'April'
  | 'May'
  | 'June'
  | 'July'
  | 'August'
  | 'September'
  | 'October'
  | 'November'
  | 'December'

export type Birthdate = `${number}-${number}-${number}`
