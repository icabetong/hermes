<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavigationBar from './components/navigation-bar/NavigationBar.vue'
</script>

<template>
  <div class="flex min-h-screen bg-slate-50">
    <header>
      <navigation-bar />
    </header>

    <div class="flex-1"><router-view /></div>
  </div>
</template>
