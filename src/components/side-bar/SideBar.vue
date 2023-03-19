<script setup lang="ts">
const props = defineProps<{
  open: boolean
  title: string
}>()
const emits = defineEmits(['dismiss'])
</script>

<template>
  <div v-if="props.open" class="fixed top-0 right-0 flex h-full w-80 flex-col overflow-y-auto">
    <div>
      <h6 v-text="props.title" />
      <button type="button" @click="emits('dismiss')">X</button>
    </div>
    <slot />
  </div>
</template>
