<script setup lang="ts">
import { Application, Events, Window } from '@wailsio/runtime';
import { PetService } from '../../bindings/github.com/wj461/SEAL/Service';
import { ref, onMounted, onUnmounted, inject, watch, computed } from "vue";
import { useWindowStatus } from './UseWindowStatus';
import { FrontendPetService } from '../services/PetService';


const frontendPetService = ref<FrontendPetService | null>(null);
const animations = inject('animations') as Record<string, string[]>;

let id = 0;
let bound = {X: 0, Y: 0, Width: 0, Height: 0};
let state = ref<string>();
let currentFrame = ref<string>();
let isGround = ref<boolean>(true);
let imageHight = 0;
let imageWidth = 0;
const { isFocused, isMoving, dragDirection, isDragging, unfocus } = useWindowStatus()
let _isDragging = ref<boolean>(false);

async function getIsGround() {
  // FIX ME 考慮靠牆跟天花板的情況
  const {x:x ,y:y}  = await Window.Position();
  return y >= bound.Height-imageHight;
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
  // PetService.Update();
  frontendPetService.value?.update();
  animationId = requestAnimationFrame(animate);
};

// 切換狀態
// Have different animations for different action and the action is folder name in src
const changeState = (newState: keyof typeof animations) => {
  PetService.SetAction(id, newState).then(() => {
    PetService.GetState(id).then((m) => {
      state.value = m as keyof typeof animations;
      frontendPetService.value?.setAction(state.value);

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
onMounted(async () => {
  isPlaying.value = true;
  const windowName = await Window.Name();

  PetService.NewPetForFrontend(windowName).then(([genId, b]) => {
    PetService.GetState(genId).then((m) => {
      id = genId;
      bound = b;
      state.value = m as keyof typeof animations;
      currentFrame.value = animations[state.value][0];

      // Init state
      changeState("failing");
      console.log(state.value)
      frontendPetService.value = new FrontendPetService();
      animationId = requestAnimationFrame(animate);
      window.addEventListener('keydown', handleKeyDown)
    });
  });

  watch(isGround, async (newVal, oldVal) => {
    const {x:x ,y:y}  = await Window.Position();
    const resolvedNewVal = newVal;
    if (resolvedNewVal != oldVal && resolvedNewVal) {
      changeState("right_ground_walk");
      Window.SetPosition(x,bound.Height-imageHight);
    }
  })
  watch(isFocused, async (newVal, oldVal) => {
    const resolvedNewVal = newVal;
    if (resolvedNewVal != oldVal && resolvedNewVal && isGround.value) {
      changeState("focused_ground");
    }else if (resolvedNewVal != oldVal && !resolvedNewVal && isGround.value) {
      changeState("left_ground_walk");
    }
  })

  watch(dragDirection, async (newVal, oldVal) => {
    const resolvedNewVal = newVal;
    if (resolvedNewVal != oldVal && resolvedNewVal == "left" && _isDragging.value) {
      changeState("left_dragging");
    } else if (resolvedNewVal != oldVal && resolvedNewVal == "right" && _isDragging.value) {
      changeState("right_dragging");
    } else if (resolvedNewVal != oldVal && resolvedNewVal == "mid" && _isDragging.value) {
      changeState("mid_dragging");
    }
  })

  watch(_isDragging, async (newVal, oldVal) => {
    const resolvedNewVal = newVal;
    if (resolvedNewVal != oldVal && !resolvedNewVal) {
      changeState("failing");
      unfocus()
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
}

function ResizeWindowByImage(image_path: string) {
    const img = new Image();
    img.src = image_path;
    img.onload = function() {
      imageHight = img.height;
      imageWidth = img.width;
      Window.SetSize(img.width, img.height);
    }
}


// some custom actions group
async function RunAroundInScreen() {
  const {x:x ,y:y}  = await Window.Position();
  isGround.value = await getIsGround();
  _isDragging.value = isDragging.value && !isGround.value;
  console.log(isDragging.value)

  if (bound.X > x && isGround.value) {
    changeState("right_ground_walk");
    Window.SetPosition(bound.X,y);
  } else if (bound.Width-imageWidth < x && isGround.value) {
    changeState("left_ground_walk");
    Window.SetPosition(bound.Width-imageWidth, y);
  }
}



// DLC 可以跟其他pet互動

// import { Events } from '@wailsio/runtime';

// // 發布此 pet 的事件
// function publishMyAction(action: string) {
//   Events.Emit(`pet:${id}:action`, action);
// }

// // 監聽所有其他 pet 的事件
// function listenToOtherPets() {
//   Events.On('pet:*:action', (eventName: string, action: string) => {
//     const otherPetId = eventName.split(':')[1];
//     if (otherPetId !== id.toString()) {
//       console.log(`Pet ${otherPetId} is doing action: ${action}`);
//       // 你可以據此做出反應
//     }
//   });
// }

// onMounted(() => {
//   // ... 其他初始化代碼
//   listenToOtherPets();
// });

// onUnmounted(() => {
//   // 清理事件監聽
//   Events.Off('pet:*:action');
// });

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
