import { Action } from '../modules/storeActionTypes'
import { GuardContext } from '../types'

export default async function (ctx: GuardContext) {
  await ctx.store.dispatch(Action.AuthActionTypes.REFRESH_AUTH_TOKEN)
  await ctx.store.dispatch(Action.AuthActionTypes.GET_USER_DATA)

  if (!ctx.store.getters['isLoggedIn'])
    ctx.next({ path: '/login', query: { redirect: ctx.to.fullPath } })
  else ctx.next()
}
