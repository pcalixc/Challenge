<template >
    <EmailList></EmailList>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

import EmailList from '@/components/EmailList.vue';

@Options({
  components: {
    
    EmailList: EmailList,
  },
})
export default class HomeView extends Vue {}
</script>


