<script setup lang="ts">
import SideBar from '../side-bar/SideBar.vue'

const props = defineProps<{
  open: boolean
  title: string
}>()
const emits = defineEmits(['dismiss'])

const dismiss = () => emits('dismiss')
</script>

<template>
  <SideBar :open="props.open" :title="props.title" @dismiss="dismiss">
    <div>
      <div>
        <label for="description">Description</label>
        <input type="text" id="description" />
      </div>
      <div>
        <label for="quantity">Quantity</label>
        <input type="number" id="quantity" />
      </div>
      <div>
        <label for="unit">Unit</label>
        <input type="text" id="unit" />
      </div>
      <div>
        <label for="batch">Batch</label>
        <input type="text" id="batch" />
      </div>
      <div>
        <label for="expiry">Expiry</label>
        <input type="text" id="expiry" />
      </div>
      <div>
        <label for="price">Price</label>
        <input type="number" id="price" />
      </div>

      <button type="submit" @click="dismiss">Save</button>
    </div>
  </SideBar>
</template>
