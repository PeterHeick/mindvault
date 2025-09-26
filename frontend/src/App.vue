<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    </header>

  <main>
    <RouterView />
  </main>
</template>

<style scoped>
main {
  padding: 2rem;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  max-width: 900px;
  margin: 0 auto;
}
</style>