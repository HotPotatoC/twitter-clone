import { ActionTypes } from '../modules/auth/store/actions'
import { GuardContext } from '../types'

export default async function (ctx: GuardContext) {
  await ctx.store.dispatch(ActionTypes.REFRESH_AUTH_TOKEN)
  await ctx.store.dispatch(ActionTypes.GET_USER_DATA)

  if (ctx.store.getters['isLoggedIn']) ctx.next('/home')
  else ctx.next()
}
