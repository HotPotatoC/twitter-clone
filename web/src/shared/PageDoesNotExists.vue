<script lang="ts">
import { defineComponent, h } from 'vue'

export default defineComponent({
  name: 'PageDoesNotExists',
  setup() {
    return () =>
      h(
        'h1',
        {
          class: ['font-bold', 'text-3xl', 'dark:text-lighter', 'text-center'],
        },
        "Sorry, That page doesn't exist!"
      )
  },
})
</script>
