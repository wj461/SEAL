<script setup lang="ts">
import { Application, Events, Window } from '@wailsio/runtime';
import { PetService } from '../../bindings/github.com/wj461/SEAL/Pet';
import { ref, onMounted, onUnmounted, inject, watch, computed } from "vue";

const animations = inject('animations') as Record<string, string[]>;

// 當前狀態與動畫
let id = 0;
let bound = {X: 0, Y: 0, Width: 0, Height: 0};
let state = ref<string>();
let currentFrame = ref<string>();
let isGround = ref<boolean>(true);

async function getIsGround() {
  const {x:x ,y:y}  = await Window.Position();
  // FIX ME
  return y+600 >= bound.Height
}

// 控制動畫播放
const isPlaying = ref<boolean>(true);
let frameIndex = 0;
let animationId: number | null = null;
const fps = 8;
const frameInterval = 1000 / fps;
let lastTime = 0;

// 切換動畫幀
const animate = (time: number) => {
  if (!isPlaying.value || !state.value) return;
  // TEST
  RunAroundInScreen();
  
  if (time - lastTime > frameInterval) {
    lastTime = time;
    const frames = animations[state.value]; // 取得當前狀態的動畫幀
    frameIndex = (frameIndex + 1) % frames.length;
    currentFrame.value = frames[frameIndex];
  }
  PetService.Update();
  animationId = requestAnimationFrame(animate);
};

// 切換狀態
// Have different animations for different action and the action is folder name in src
const changeState = (newState: keyof typeof animations) => {
  PetService.SetAction(id, newState).then(() => {
    PetService.GetState(id).then((m) => {
      state.value = m as keyof typeof animations;

      frameIndex = 0; // 重置幀索引
      currentFrame.value = animations[state.value][0]; // 顯示第一幀
      ResizeWindowByImage (currentFrame.value);
    });
  });
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

  PetService.NewPetForFrontend().then(([genId, b]) => {
    PetService.GetState(genId).then((m) => {
      id = genId;
      bound = b;
      state.value = m as keyof typeof animations;
      currentFrame.value = animations[state.value][0];

      changeState("idle");
      animationId = requestAnimationFrame(animate);
      window.addEventListener('keydown', handleKeyDown)
    });
  });

  watch(isGround, async (newVal, oldVal) => {
    const {x:x ,y:y}  = await Window.Position();
    console.log("isGround: " + newVal, x, y);
    const resolvedNewVal = newVal;
    if (resolvedNewVal != oldVal && resolvedNewVal) {
      changeState("right_ground_walk");
      Window.SetPosition(x,bound.Height-600);
    } else {
      changeState("failing");
    }
  })

});

// 組件卸載時清除動畫
onUnmounted(() => {
  if (animationId) cancelAnimationFrame(animationId);
  window.removeEventListener('keydown', handleKeyDown);
});

// Listen keyboard events
function handleKeyDown(event: KeyboardEvent) {
  if (event.key === "ArrowLeft") {
    changeState("left_ground_walk");
  } else if (event.key === "ArrowRight") {
    changeState("right_ground_walk");
  } else if (event.key === "ArrowDown") {
    changeState("failing");
    isGround;
  // } else if (event.key === "ArrowLeft"){
  //   changeState("left_walk");
  }
}

function ResizeWindowByImage(image_path: string) {
    const img = new Image();
    img.src = image_path;
    img.onload = function() {
        Window.SetSize(img.width, img.height);
    }
}


// some custom actions group
async function RunAroundInScreen() {
  // get the windows position use wails3 API
  const {x:x ,y:y}  = await Window.Position();
  // console.log("x: " + x + " y: " + y);
  // console.log("width: " + bound.Width + " height: " + bound.Height);
  isGround.value = await getIsGround();

  if (bound.X > x) {
    console.log("x < 0");
    changeState("right_ground_walk");
    Window.SetPosition(bound.X,y);
  } else if (bound.Width-200 < x) {
    changeState("left_ground_walk");
    Window.SetPosition(bound.Width-200, y);
  }

  //  else if (bound.Y > y) {
  //   console.log("y < 0");
  //   changeState("up_walk");
  //   Window.SetPosition(x,bound.Y);
  // } else if (bound.Width-200 < x) {
  //   changeState("right_walk");
  //   Window.SetPosition(bound.Width-200, y);
  // } else if (bound.Height-100 < y) {
  //   console.log("y > height");
  //   changeState("walk");
  //   Window.SetPosition(x,bound.Height-100);
  // }

}


</script>

<template>
        <img :src="currentFrame" alt="Pet image">
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
