<template >

    <div class="overflow-hidden">
    
        <Contact></Contact>

    </div>

        

</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

import Contact from '@/components/Contact.vue';
    @Options({
    components: {
    
        Contact: Contact,
    },
    })
export default class AboutView extends Vue {}
</script>