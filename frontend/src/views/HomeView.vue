<template >

  <div class="overflow-hidden">
    <Header></Header>
      
      
        <EmailList></EmailList>

  </div>

        
 

</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import Header from '@/components/Header.vue';
import EmailList from '@/components/EmailList.vue';

import Footer from '@/components/Footer.vue'

import EmailSelected from '@/components/EmailSelected.vue';



@Options({
  components: {
    
    Header: Header,
    EmailList: EmailList,
    Footer: Footer,
    EmailSelected: EmailSelected
  },
})
export default class HomeView extends Vue {}
</script>


