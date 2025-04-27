import { ref, onMounted, onBeforeUnmount } from 'vue'
import { Window } from '@wailsio/runtime'

export function useWindowStatus() {
  // const isFocused = ref(document.hasFocus())
  const isFocused = ref(false)
  const isMoving = ref(false)
  const isDragging = ref(false)
  const dragDirection = ref<'left' | 'right' | 'mid' | null>(null)

  let lastX = window.screenX
  let lastY = window.screenY
  let interval: number | null = null
  const unfocus = () => {
    isFocused.value = false
    // if (document.activeElement instanceof HTMLElement) {
    //   document.activeElement.blur()
    //   isDragging.value = false
    // }
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
    console.log('blur')
  }

  const checkWindowMove = async () => {
    const currentX = window.screenX
    const currentY = window.screenY

    if (currentX !== lastX || currentY !== lastY) {
      isMoving.value = true
    } else {
      isMoving.value = false
    }

    if (currentX > lastX + 10) {
      dragDirection.value = 'right'
    } else if (currentX < lastX - 10) {
      dragDirection.value = 'left'
    } else{
      dragDirection.value = 'mid'
    }

    lastX = currentX
    lastY = window.screenY
    // isFocused.value = await Window.IsFocused();
  }

  onMounted(() => {
    window.addEventListener('mousedown', onMouseDown)
    window.addEventListener('mouseup', onMouseUp)
    window.addEventListener('mousemove', onMouseMove)
    window.addEventListener('blur', onBlur)
    interval = window.setInterval(checkWindowMove, 100)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('mousedown', onMouseDown)
    window.removeEventListener('mouseup', onMouseUp)
    window.removeEventListener('mousemove', onMouseMove)
    window.removeEventListener('blur', onBlur)
    if (interval) {
      clearInterval(interval)
    }
  })

  return {
    isFocused,
    isMoving,
    dragDirection,
    isDragging,
    unfocus
  }
}