<script setup lang="ts">
import { ref, watchEffect } from 'vue'
import SideBar from '@/components/side-bar/SideBar.vue'
import { DeviceFloppyIcon } from 'vue-tabler-icons'

const props = defineProps<{
  open: boolean
  title: string
  medicine: Medicine | null
}>()
const state = ref<Medicine>({
  id: '',
  description: '',
  quantity: 0,
  unit: '',
  batch: '',
  expiry: Date.now(),
  price: 0,
  createdAt: Date.now(),
  updatedAt: Date.now()
})
const emits = defineEmits(['dismiss', 'submit'])

const dismiss = () => emits('dismiss')
const submit = () => {
  emits('submit', state.value)
}

watchEffect(() => {
  if (!props.medicine) {
    state.value = {
      id: '',
      description: '',
      quantity: 0,
      unit: '',
      batch: '',
      expiry: Date.now(),
      price: 0,
      createdAt: Date.now(),
      updatedAt: Date.now()
    }
    return
  }

  state.value = props.medicine
})
</script>

<template>
  <side-bar :open="props.open" :title="props.title" @dismiss="dismiss">
    <form class="flex flex-1 flex-col" @submit.prevent="submit">
      <div class="flex-1 space-y-4">
        <div>
          <label class="form-label" for="description">Description</label>
          <input type="text" class="form-input" v-model="state.description" />
        </div>
        <div>
          <label class="form-label" for="quantity">Quantity</label>
          <input type="number" id="quantity" class="form-input" v-model="state.quantity" />
        </div>
        <div>
          <label class="form-label" for="unit">Unit</label>
          <input type="text" id="unit" class="form-input" v-model="state.unit" />
        </div>
        <div>
          <label class="form-label" for="batch">Batch</label>
          <input type="text" id="batch" class="form-input" v-model="state.batch" />
        </div>
        <div>
          <label class="form-label" for="expiry">Expiry</label>
          <input type="text" id="expiry" class="form-input" v-model="state.expiry" />
        </div>
        <div>
          <label class="form-label" for="price">Price</label>
          <input type="number" id="price" class="form-input" v-model="state.price" />
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
    </form>
  </side-bar>
</template>
