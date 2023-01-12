<template>
  
  <div class=" flex w-screen items-center justify-center py-4 bg-gray-00">
    <form method="GET">
      <div class="relative text-blue-200 focus-within:text-gray-400 " >
        <span class="absolute inset-y-0 left-0 flex items-center pl-2">
          <button type="submit" class="p-1 focus:outline-none focus:shadow-outline">
            <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" class="w-6 h-6"><path d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
          </button>
        </span>
        <input type="search" name="q" class="py-2 px-6 text-md text-white bg-gray-800 rounded-md pl-10 focus:outline-none focus:bg-gray-800 focus:text-gray-00" placeholder="Search email..." autocomplete="off">
      </div>
    </form>
  </div>
 
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  props: {
    msg: String
  }
})
export default class Searcher extends Vue {
  msg!: string
}
</script>


