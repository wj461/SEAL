import { ref, onMounted, onBeforeUnmount } from 'vue'
import { Window } from '@wailsio/runtime'

// Singleton state
const isFocused = ref(false)
const isMoving = ref(false)
const isDragging = ref(false)
const dragDirection = ref<'left' | 'right' | 'mid' | null>(null)
const currentX = ref(window.screenX)
const currentY = ref(window.screenY)

let lastX = window.screenX
let lastY = window.screenY
let interval: number | null = null
let isInitialized = false

// Singleton instance
const unfocus = () => {
  isFocused.value = false
}

const onMouseDown = () => {
  isDragging.value = true
  isFocused.value = true
}

const onMouseUp = () => {
  isDragging.value = false
}

const onMouseMove = (event: MouseEvent) => {
  if (event.buttons === 0) {
    isDragging.value = false
  }
}
const onBlur = () => {
  isFocused.value = false
}

const checkWindowMove = async () => {
  currentX.value = window.screenX
  currentY.value = window.screenY

  if (currentX.value !== lastX || currentY.value !== lastY) {
    isMoving.value = true
  } else {
    isMoving.value = false
  }

  if (currentX.value > lastX + 10) {
    dragDirection.value = 'right'
  } else if (currentX.value < lastX - 10) {
    dragDirection.value = 'left'
  } else{
    dragDirection.value = 'mid'
  }

  lastX = currentX.value
  lastY = currentY.value
}

// Initialize once
function initialize() {
  if (isInitialized) return
  
  window.addEventListener('mousedown', onMouseDown)
  window.addEventListener('mouseup', onMouseUp)
  window.addEventListener('mousemove', onMouseMove)
  window.addEventListener('blur', onBlur)
  interval = window.setInterval(checkWindowMove, 100)
  
  isInitialized = true
}

// Cleanup function
function cleanup() {
  if (!isInitialized) return
  
  window.removeEventListener('mousedown', onMouseDown)
  window.removeEventListener('mouseup', onMouseUp)
  window.removeEventListener('mousemove', onMouseMove)
  window.removeEventListener('blur', onBlur)
  if (interval) {
    clearInterval(interval)
    interval = null
  }
  
  isInitialized = false
}

export function useWindowStatus() {
  onMounted(() => {
    initialize()
  })

  onBeforeUnmount(() => {
    cleanup()
  })

  return {
    isFocused,
    isMoving,
    dragDirection,
    isDragging,
    unfocus,
    currentX,
    currentY
  }
}