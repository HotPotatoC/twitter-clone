import { RouteComponent, RouteRecordRaw } from 'vue-router'

export function makeRoutesWithLayout(
  rootPath: string,
  layout: RouteComponent,
  routes: RouteRecordRaw[]
): RouteRecordRaw {
  return {
    path: rootPath,
    component: layout,
    children: routes,
  }
}
