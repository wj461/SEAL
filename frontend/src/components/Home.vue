<script setup lang="ts">
import HelloWorld from './HelloWorld.vue'
</script>

<template>
  <div class="container">
     <p>
    <strong>Current route path:</strong> {{ $route.fullPath }}
  </p>
    <div>
      <a wml-openURL="https://wails.io">
        <img src="/wails.png" class="logo" alt="Wails logo"/>
      </a>
      <a wml-openURL="https://vuejs.org/">
        <img src="/vue.svg" class="logo vue" alt="Vue logo"/>
      </a>
    </div>
    <HelloWorld msg="Wails + Vue" />
    <RouterLink to="/Pet">Go to Home</RouterLink>
  </div>
  <main>
    <RouterView />
  </main>
</template>

<style scoped>
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
