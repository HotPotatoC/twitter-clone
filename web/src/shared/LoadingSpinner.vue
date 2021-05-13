<script lang="ts">
import { defineComponent, h, ref, toRefs } from 'vue'

export default defineComponent({
  name: 'LoadingSpinner',
  props: {
    size: {
      type: [String, Number],
      default: '64px',
    },
    color: {
      type: String as () => 'primary' | 'white',
      default: 'primary',
    },
  },
  setup(props) {
    const { size, color } = toRefs(props)
    const borderColor = ref('')
    const primaryTheme =
      '#1da1f2 rgba(39, 149, 217, 0.2) #1da1f2 rgba(39, 149, 217, 0.2)'
    const whiteTheme =
      '#F5F8FA rgba(245, 248, 250, 0.2) #F5F8FA rgba(245, 248, 250, 0.2)'

    if (color.value === 'primary') {
      borderColor.value = primaryTheme
    }

    if (color.value === 'white') {
      borderColor.value = whiteTheme
    }

    return () =>
      h('div', {
        class: ['loading-spinner', 'inline-block'],
        style: {
          '--size': size.value,
          '--borderColor': borderColor.value,
        },
      })
  },
})
</script>

<style scoped>
.loading-spinner {
  width: var(--size);
  height: var(--size);
}
.loading-spinner:after {
  content: ' ';
  display: block;
  width: var(--size);
  height: var(--size);
  margin: 8px;
  border-radius: 50%;
  border: 6px solid;
  border-color: var(--borderColor);
  animation: spin 1.2s linear infinite;
}
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
