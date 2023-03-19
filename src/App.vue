<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navigation from './components/navigation-bar/NavigationBar.vue'
</script>

<template>
  <header>
    <Navigation />
  </header>

  <RouterView />
</template>
