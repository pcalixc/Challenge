<template >
    <Contact></Contact>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

import Contact from '@/components/Contact.vue';
    @Options({
    components: {
    
        Contact: Contact,
    },
    })
export default class ContactView extends Vue {}
</script>