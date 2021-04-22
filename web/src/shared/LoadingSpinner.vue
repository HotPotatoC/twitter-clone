<script lang="ts">
import { defineComponent, h } from 'vue'

export default defineComponent({
  name: 'LoadingSpinner',
  setup() {
    return () =>
      h('div', {
        class: ['loading-spinner', 'inline-block'],
      })
  },
})
</script>

<style scoped>
.loading-spinner {
  width: 80px;
  height: 80px;
}
.loading-spinner:after {
  content: ' ';
  display: block;
  width: 64px;
  height: 64px;
  margin: 8px;
  border-radius: 50%;
  border: 6px solid #1da1f2;
  border-color: #1da1f2 rgba(39, 149, 217, 0.2) #1da1f2 rgba(39, 149, 217, 0.2);
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
