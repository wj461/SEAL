<script setup lang="ts">
import { Events, Window } from '@wailsio/runtime';
import { PetService } from '../../bindings/changeme';
import { ref, onMounted, onUnmounted, inject } from "vue";

const animations = inject('animations') as Record<string, string[]>;

// 當前狀態與動畫
const state = ref<keyof typeof animations>("idle");
const currentFrame = ref<string>(animations[state.value][0]);

// 控制動畫播放
const isPlaying = ref<boolean>(true);
let frameIndex = 0;
let animationId: number | null = null;
const fps = 10;
const frameInterval = 1000 / fps;
let lastTime = 0;

// 切換動畫幀
const animate = (time: number) => {
  if (!isPlaying.value) return;
  
  if (time - lastTime > frameInterval) {
    lastTime = time;
    const frames = animations[state.value]; // 取得當前狀態的動畫幀
    frameIndex = (frameIndex + 1) % frames.length;
    currentFrame.value = frames[frameIndex];
  }
  animationId = requestAnimationFrame(animate);
};

// 切換狀態
const changeState = (newState: keyof typeof animations) => {
  state.value = newState;
  frameIndex = 0; // 重置幀索引
  currentFrame.value = animations[newState][0]; // 顯示第一幀
};

// 控制播放/暫停動畫
const toggleAnimation = () => {
  isPlaying.value = !isPlaying.value;
  if (isPlaying.value) {
    animationId = requestAnimationFrame(animate);
  } else if (animationId) {
    cancelAnimationFrame(animationId);
  }
};

// 組件掛載時啟動動畫
onMounted(() => {
  isPlaying.value = true;
  animationId = requestAnimationFrame(animate);
});

// 組件卸載時清除動畫
onUnmounted(() => {
  if (animationId) cancelAnimationFrame(animationId);
});

function ResizeWindowByImage(image_path: string) {
    const img = new Image();
    img.src = image_path;
    img.onload = function() {
        Window.SetSize(img.width, img.height);
    }
}

</script>

<template>
        <img :src="currentFrame" alt="動畫">
    
        <div class="controls">
          <button @click="changeState('idle')">idle</button>
          <button @click="changeState('move')">move</button>
          <button @click="toggleAnimation">{{ isPlaying ? "暫停" : "播放" }}</button>
        </div>
</template>

<style scoped>
.animation-container {
  text-align: center;
}

.controls {
  margin-top: 10px;
}

button {
  margin: 5px;
  padding: 10px;
  border: none;
  background-color: #007bff;
  color: white;
  cursor: pointer;
  border-radius: 5px;
}

button:hover {
  background-color: #0056b3;
}
</style>
