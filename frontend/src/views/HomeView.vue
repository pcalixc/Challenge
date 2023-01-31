<template >
    <Home></Home>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

import Home from '@/components/Home.vue';

@Options({
  components: {
    
    Home: Home,
  },
})
export default class HomeView extends Vue {}
</script>


