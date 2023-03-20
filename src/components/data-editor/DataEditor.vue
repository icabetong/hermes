<script setup lang="ts">
import { ref, watchEffect, reactive } from 'vue'
import SideBar from '@/components/side-bar/SideBar.vue'
import { DeviceFloppyIcon, CurrencyPesoIcon } from 'vue-tabler-icons'

const props = defineProps<{
  open: boolean
  title: string
  medicine: Medicine | null
}>()
const errors = reactive<Record<keyof Omit<Medicine, 'createdAt' | 'updatedAt' | 'id'>, string>>({
  description: '',
  quantity: '',
  unit: '',
  batch: '',
  expiry: '',
  price: ''
})
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
  const { description, quantity, unit, batch, expiry, price } = state.value

  if (!description || (description && description.trim().length <= 0)) {
    errors.description = 'A medicine name is required'
    return
  }
  if (quantity <= 0) {
    errors.quantity = 'Quantity needs to be above 0'
    return
  }
  if (!unit || (unit && unit.trim().length <= 0)) {
    errors.unit = 'A unit of measurement needs to be defined'
    return
  }
  if (!batch || (batch && batch.trim().length <= 0)) {
    errors.batch = 'Batch data needs to be defined'
    return
  }
  if (expiry <= 0) {
    errors.expiry = 'Expiration Date is required'
    return
  }
  if (price <= 0) {
    errors.price = 'Price is required'
    return
  }

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

const date = new Date(state.value.expiry).toISOString()
const expiryDate = date.substring(0, date.indexOf('T')).split('/').join('-')
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
          <input type="date" id="expiry" class="form-input" v-model="expiryDate" />
        </div>
        <div>
          <label class="form-label" for="price">Price</label>
          <div class="relative">
            <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
              <currency-peso-icon class="h-4 w-4 text-gray-500 dark:text-gray-400" />
            </div>
            <input type="number" id="price" class="form-input pl-10" v-model="state.price" />
          </div>
        </div>
      </div>

      <button type="submit" class="button-primary flex w-full items-center justify-center">
        <device-floppy-icon class="mr-2 h-4 w-4" />
        Save
      </button>
    </form>
  </side-bar>
</template>
