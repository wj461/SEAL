<script setup lang="ts">
</script>

<template>
    <div class="container">
    </div>
    <main>
      <RouterView />
    </main>
</template>

<style scoped>

.Pet{
  display: block;
}

.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #e80000aa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
