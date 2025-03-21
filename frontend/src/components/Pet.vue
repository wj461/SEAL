<script setup lang="ts">
import { Events, Window } from '@wailsio/runtime';
import { PetService } from '../../bindings/changeme';
import { ref, onMounted, onUnmounted } from "vue";

const img_path = new Map<string, string>();
var current_display_key = "Normal";
const current_display_path = ref('');
const display_image = (image: string) => {
    current_display_path.value = img_path.get(image) ?? '';
}

// Initialize the window size based on the image size and set the default image
(function () { 
    // import all images in the folder
    const images = import.meta.glob('/public/PetImage/*.{png,jpg,jpeg,svg}')
    Object.keys(images).forEach((key) => {
        console.log(key);
        // use regex to get the image name
        const regex = /\/public\/PetImage\/(.*)\.(png|jpg|jpeg|svg)/;
        const match = key.match(regex);
        if (match) {
            img_path.set(match[1], "/PetImage/" + match[1] + "." + match[2]);
        }
    })

    // set the default image
    SetImage(current_display_key);
  }
 )();

//  Set display image based on the key and resize the window based on the image size
 function SetImage(key) {
    current_display_path.value = img_path.get(key)?.toString() ?? '';
    current_display_key = key;

    const img = new Image();
    img.src = current_display_path.value;
    img.onload = function() {
        Window.SetSize(img.width, img.height);
    }
}

let index = 0;
let intervalId: number | null = null;
const walk_animation = ["Normal", "Move"];

const PlayWalkAnimation = () => {
  index = (index + 1) % walk_animation.length; // 依序循環
  SetImage(walk_animation[index]);
  Window.Position().then((currentPosition) => {
    Window.SetPosition(currentPosition.x - 5, currentPosition.y);
  });
};

onMounted(() => {
  intervalId = setInterval(PlayWalkAnimation, 150);
});

onUnmounted(() => {
  if (intervalId) clearInterval(intervalId);
});

</script>

<template>
        <img :src="current_display_path"  class="Pet" alt="Pet logo"/>
</template>

<style scoped>
</style>

