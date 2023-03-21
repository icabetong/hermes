<script setup lang="ts">
import { reactive } from 'vue'
import { PlusIcon, SearchIcon, ChevronDownIcon } from 'vue-tabler-icons'
import { Popover, PopoverButton, PopoverPanel } from '@headlessui/vue'
import DataTable from '@/components/data-table/DataTable.vue'
import DataEditor from '@/components/data-editor/DataEditor.vue'

const state = reactive<{
  open: boolean
  medicine: Medicine | null
  filter: { others: boolean; expired: boolean; near: boolean }
}>({
  open: false,
  medicine: null,
  filter: {
    others: true,
    expired: true,
    near: true
  }
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

const triggerEditorCreate = () => {
  state.open = !state.open
  state.medicine = null
}
const triggerEditorUpdate = (medicine: Medicine) => {
  state.medicine = medicine
  state.open = true
}
const triggerEditorSubmit = (medicine: Medicine) => {
  console.log(medicine)
  state.open = false
}
</script>

<template>
  <main class="flex h-full flex-col p-4">
    <div class="flex h-full flex-1 flex-col items-center rounded-lg border bg-white p-8">
      <div class="flex w-full items-center justify-between">
        <h1 class="text-2xl font-bold">Medicines</h1>
        <button type="button" class="button-primary flex items-center" @click="triggerEditorCreate">
          <plus-icon class="mr-2 h-4 w-4" />
          Add Medicine
        </button>
      </div>
      <div class="flex w-full items-center justify-between bg-white py-4 dark:bg-gray-800">
        <div>
          <popover as="div" class="relative inline-block text-left text-sm">
            <popover-button
              type="button"
              class="inline-flex items-center rounded-lg border border-gray-300 bg-white px-6 py-1.5 text-sm font-medium text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:border-gray-600 dark:bg-gray-800 dark:text-gray-400 dark:hover:border-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-700">
              <span class="sr-only">Filter button</span>
              Filter
              <chevron-down-icon class="ml-2 h-3 w-3" />
            </popover-button>
            <popover-panel
              class="absolute left-0 z-10 mt-2 w-56 origin-top-left rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
              <ul class="space-y-3 p-3 text-sm text-gray-700 dark:text-gray-200">
                <li class="flex items-center">
                  <input type="checkbox" id="expired" v-model="state.filter.expired" />
                  <label
                    for="expired"
                    class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">
                    Hide expired
                  </label>
                </li>
                <li class="flex items-center">
                  <input type="checkbox" id="near" v-model="state.filter.near" />
                  <label
                    for="near"
                    class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">
                    Hide nearing expiration
                  </label>
                </li>
                <li class="flex items-center">
                  <input type="checkbox" id="others" v-model="state.filter.others" />
                  <label
                    for="other"
                    class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">
                    Hide others
                  </label>
                </li>
              </ul>
            </popover-panel>
          </popover>
        </div>
        <label for="table-search" class="sr-only">Search</label>
        <div class="relative">
          <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
            <search-icon class="h-4 w-4 text-slate-500" />
          </div>
          <input
            type="text"
            id="table-search-medicines"
            class="form-input pl-10"
            placeholder="Search for medicines" />
        </div>
      </div>
      <data-table :items="medicines" @select="triggerEditorUpdate" />
    </div>
  </main>
  <data-editor
    :medicine="state.medicine"
    :open="state.open"
    :title="state.medicine ? 'Edit Medicine' : 'Create Medicine'"
    @dismiss="triggerEditorCreate"
    @submit="triggerEditorSubmit" />
</template>
