<script setup lang="ts">
import { Events, Window } from '@wailsio/runtime';
import { PetService } from '../../bindings/changeme';
import { ref, onMounted, onUnmounted, inject } from "vue";

const img_path = inject('img_path');
var current_status = "Normal";
const current_display_path = ref('');


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

