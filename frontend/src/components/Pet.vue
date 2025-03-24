<script setup lang="ts">
import { Events, Window } from '@wailsio/runtime';
import { PetService } from '../../bindings/changeme';
import { ref, onMounted, onUnmounted } from "vue";

var img_path = new Map<string, Array<string>>();
var current_status = "Normal";
const current_display_path = ref('');

// Initialize the window size based on the image size and set the default image
(function () { 
    const images = import.meta.glob("/src/image/pet/**/*.png", { eager: true });
    const sortedPaths = Object.keys(images).sort((a, b) =>
        a.localeCompare(b, undefined, { numeric: true })
        );

    sortedPaths.forEach((key) => {
        // / key = src/image/pet/move/1.png
        const path = key.split('/');
        const animation_name = path[path.length - 2];
        current_display_path.value = key;

        if (img_path.has(animation_name)) {
            img_path.get(animation_name)?.push(key);
        } else {
            img_path.set(animation_name, [key]);
        }

    })

    // set the default image
  }
 )();

 function ResizeWindowByImage(image_path: string) {
    const img = new Image();
    img.src = image_path;
    img.onload = function() {
        Window.SetSize(img.width, img.height);
    }
}

</script>

<template>
        <img :src="current_display_path"  class="Pet" alt="Pet logo"/>
</template>

<style scoped>
</style>

