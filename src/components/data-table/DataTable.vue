<script setup lang="ts">
const props = defineProps<{ items: Medicine[] }>()
</script>

<template>
  <table>
    <thead>
      <th>Description</th>
      <th>Quantity</th>
      <th>Unit</th>
      <th>Batch</th>
      <th>Expiry</th>
      <th>Price</th>
      <th><span class="sr-only">Actions</span></th>
    </thead>
    <tbody>
      <tr v-bind:key="medicine.id" v-for="medicine in props.items">
        <th>{{ medicine.description }}</th>
        <td>{{ medicine.quantity }}</td>
        <td>{{ medicine.unit }}</td>
        <td>{{ medicine.batch }}</td>
        <td>{{ medicine.expiry }}</td>
        <td>{{ medicine.price }}</td>
        <td><button type="button">Remove</button></td>
      </tr>
    </tbody>
  </table>
</template>
