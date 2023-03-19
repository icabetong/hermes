<script setup lang="ts">
import { reactive } from 'vue'
import { PlusIcon } from 'vue-tabler-icons'
import DataTable from '@/components/data-table/DataTable.vue'
import DataEditor from '@/components/data-editor/DataEditor.vue'

const state = reactive({
  open: false
})

const medicines: Medicine[] = [
  {
    id: '1',
    description: 'Biogesic',
    quantity: 100,
    unit: 'Box',
    batch: 'Kebab',
    expiry: Date.now(),
    price: 6.25,
    createdAt: Date.now(),
    updatedAt: Date.now()
  },
  {
    id: '2',
    description: 'Biogesic',
    quantity: 100,
    unit: 'Box',
    batch: 'Kebab',
    expiry: Date.now(),
    price: 6.25,
    createdAt: Date.now(),
    updatedAt: Date.now()
  }
]

const triggerEditor = () => (state.open = !state.open)
</script>

<template>
  <main class="flex h-full flex-col p-4">
    <div class="flex h-full flex-1 flex-col items-center rounded-lg border bg-white p-2">
      <div class="flex w-full items-center justify-between p-4">
        <h1 class="text-2xl font-bold">Medicines</h1>
        <button type="button" class="button-primary flex items-center" @click="triggerEditor">
          <plus-icon class="mr-2 h-4 w-4" />
          Add Medicine
        </button>
      </div>
      <data-table :items="medicines" />
      <data-editor :open="state.open" title="Create Medicine" @dismiss="triggerEditor" />
    </div>
  </main>
</template>
