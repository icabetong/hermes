<script setup lang="ts">
import { reactive } from 'vue'
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
  }
]

const triggerEditor = () => (state.open = !state.open)
</script>

<template>
  <main>
    <div class="flex items-center justify-between">
      <h1>Medicines</h1>
      <button type="button" @click="triggerEditor">Create</button>
    </div>
    <DataTable :items="medicines" />
    <DataEditor :open="state.open" title="Create Medicine" @dismiss="triggerEditor" />
  </main>
</template>
