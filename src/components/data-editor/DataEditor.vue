<script setup lang="ts">
import SideBar from '@/components/side-bar/SideBar.vue'
import { DeviceFloppyIcon } from 'vue-tabler-icons'

const props = defineProps<{
  open: boolean
  title: string
}>()
const emits = defineEmits(['dismiss'])

const dismiss = () => emits('dismiss')
</script>

<template>
  <side-bar :open="props.open" :title="props.title" @dismiss="dismiss">
    <div class="flex flex-1 flex-col">
      <div class="flex-1 space-y-4">
        <div>
          <label class="form-label" for="description">Description</label>
          <input type="text" class="form-input" />
        </div>
        <div>
          <label class="form-label" for="quantity">Quantity</label>
          <input type="number" id="quantity" class="form-input" />
        </div>
        <div>
          <label class="form-label" for="unit">Unit</label>
          <input type="text" id="unit" class="form-input" />
        </div>
        <div>
          <label class="form-label" for="batch">Batch</label>
          <input type="text" id="batch" class="form-input" />
        </div>
        <div>
          <label class="form-label" for="expiry">Expiry</label>
          <input type="text" id="expiry" class="form-input" />
        </div>
        <div>
          <label class="form-label" for="price">Price</label>
          <input type="number" id="price" class="form-input" />
        </div>
      </div>

      <button
        type="submit"
        class="button-primary flex w-full items-center justify-center"
        @click="dismiss"
      >
        <device-floppy-icon class="mr-2 h-4 w-4" />
        Save
      </button>
    </div>
  </side-bar>
</template>
