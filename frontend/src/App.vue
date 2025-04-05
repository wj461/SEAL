<script setup lang="ts">
import { provide } from 'vue';
var animationRecord: Record<string, string[]> = {};

// 這裡使用 import.meta.glob 來動態導入圖片 and store in animationRecord
(function () { 
    const images = import.meta.glob("/src/image/pet/**/*.png", { eager: true });
    const sortedPaths = Object.keys(images).sort((a, b) =>
        a.localeCompare(b, undefined, { numeric: true })
        );

    sortedPaths.forEach((key) => {
        // / key = src/image/pet/move/1.png
        const path = key.split('/');
        const animation_name = path[path.length - 2];

        if (animationRecord.hasOwnProperty(animation_name)) {
            animationRecord[animation_name].push(key);
        } else {
            animationRecord[animation_name] = [key];
        }
    })
  }
 )();

 provide('animations', animationRecord);
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
