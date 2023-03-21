<script setup lang="ts">
import { XIcon } from 'vue-tabler-icons'
const props = defineProps<{
  open: boolean
  title: string
}>()
const emits = defineEmits(['dismiss'])
</script>

<template>
  <div
    v-if="props.open"
    class="fixed top-0 right-0 flex h-full w-80 flex-col overflow-y-auto border-l bg-white p-4"
  >
    <div class="mb-8 flex items-center justify-between">
      <h6 class="text-lg font-bold" v-text="props.title" />
      <button
        type="button"
        class="rounded-lg p-2 transition-all hover:bg-gray-200"
        @click="emits('dismiss')"
      >
        <x-icon class="h-4 w-4" />
      </button>
    </div>
    <slot />
  </div>
</template>
