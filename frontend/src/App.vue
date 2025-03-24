<script setup lang="ts">
import { provide } from 'vue';
const img_path = new Map<string, string[]>();
var animationRecord: Record<string, string[]> = {};


// Initialize the window size based on the image size and set the default image
(function () { 
    const images = import.meta.glob("/src/image/pet/**/*.png", { eager: true });
    const sortedPaths = Object.keys(images).sort((a, b) =>
        a.localeCompare(b, undefined, { numeric: true })
        );

    // do the same thing but use Record<string, string[]> instead of Map<string, string[]>
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

    // sortedPaths.forEach((key) => {
    //     // / key = src/image/pet/move/1.png
    //     const path = key.split('/');
    //     const animation_name = path[path.length - 2];

    //     if (img_path.has(animation_name)) {
    //         img_path.get(animation_name)?.push(key);
    //     } else {
    //         img_path.set(animation_name, [key]);
    //     }
    // })
  }
 )();

//  const animations: Record<string, string[]> = Object.fromEntries(animationsMap);
 provide('img_path', animationRecord);

// provide('im', sharedData);

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
