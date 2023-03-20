<script setup lang="ts">
import { TrashIcon } from 'vue-tabler-icons'
const props = defineProps<{ items: Medicine[] }>()
const emits = defineEmits(['select', 'remove'])
</script>

<template>
  <div class="relative w-full overflow-x-auto">
    <table class="w-full text-left text-sm text-gray-500 dark:text-gray-400">
      <thead class="bg-gray-50 text-xs uppercase text-gray-700 dark:bg-gray-700 dark:text-gray-400">
        <th scope="col" class="px-6 py-3">Description</th>
        <th scope="col" class="px-6 py-3">Quantity</th>
        <th scope="col" class="px-6 py-3">Unit</th>
        <th scope="col" class="px-6 py-3">Batch</th>
        <th scope="col" class="px-6 py-3">Expiry</th>
        <th scope="col" class="px-6 py-3">Price</th>
        <th scope="col" class="px-6 py-3">Actions</th>
      </thead>
      <tbody class="divide-y">
        <tr
          class="border-b bg-white hover:cursor-pointer hover:bg-gray-100 dark:border-gray-600 dark:bg-gray-800 dark:hover:bg-gray-700"
          v-bind:key="medicine.id"
          v-for="medicine in props.items"
          @click.stop="emits('select', medicine)"
        >
          <th
            scope="row"
            class="whitespace-nowrap px-6 py-4 font-medium text-gray-900 dark:text-white"
          >
            {{ medicine.description }}
          </th>
          <td class="px-6 py-4">{{ medicine.quantity }}</td>
          <td class="px-6 py-4">{{ medicine.unit }}</td>
          <td class="px-6 py-4">{{ medicine.batch }}</td>
          <td class="px-6 py-4">{{ medicine.expiry }}</td>
          <td class="px-6 py-4">{{ medicine.price }}</td>
          <td class="px-6 py-4">
            <button type="button" @click.stop="emits('remove', medicine)">
              <span class="sr-only">Remove</span>
              <trash-icon class="h-4 w-4" />
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
