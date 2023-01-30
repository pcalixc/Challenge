<!-- eslint-disable vue/multi-word-component-names -->
<template>
    <div v-if="show">
        <div class="modal-overlay"></div>
        <div class="modal-content">
        <button @click="closeModal">Cerrar</button>
        <slot></slot>
        </div>
    </div>  
</template>


<script lang="ts">

import { ref } from 'vue'

export default {
  setup() {
    const showModal = ref(false)
    const closeModal = () => {
      showModal.value = false
    }
    return {
      showModal,
      closeModal
    }
  }
}


</script>