import { ref, Ref, onMounted, onUnmounted, watchEffect } from 'vue'

type ScrollHook = [Ref<HTMLElement | undefined>, Ref<boolean>]

export function useScroll(): ScrollHook {
  const elementRef = ref<HTMLElement>()
  const isBottom = ref(false)

  onMounted(() => {
    const element = elementRef.value
    watchEffect(() => {
      if (element) {
        element.addEventListener('scroll', () => {
          if (
            element.scrollTop + element.clientHeight + 1 >=
            element.scrollHeight
          ) {
            isBottom.value = true
          }
        })
      }
    })
  })

  onUnmounted(() => {
    const element = elementRef.value
    if (element) {
      element.removeEventListener('scroll', () => {})
    }
  })

  return [elementRef, isBottom]
}
