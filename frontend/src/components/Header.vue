<!-- eslint-disable vue/multi-word-component-names -->
<template>   
<header >
    <nav class=" fixed w-screen bg-white border-gray-200 px-4 lg:px-6 py-1 dark:bg-gray-900 ">
        <div class="flex  items-center mx-auto max-w-screen-xl">
            <a href="#" class="flex items-center">
                <img src="../assets/email-icon.png" class="mr-3 h-6 sm:h-9" alt="Logo" />
                <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Enron Mail</span>
            </a>

            <div class=" flex w-screen items-center justify-center py-4 bg-gray-00">
                <form method="GET">
                  <div class="relative text-blue-200 focus-within:text-gray-400 " >
                    <span class="absolute inset-y-0 left-0 flex items-center pl-2">
                      <button type="submit" class="p-1 focus:outline-none focus:shadow-outline">
                        <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" class="w-6 h-6"><path d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </span>
                    <input type="search" name="q" class="py-2 px-6 text-md text-white bg-blue-900 rounded-md pl-10 focus:outline-none focus:bg-white focus:text-gray-800" placeholder="Search email..." autocomplete="off">
                  </div>
                </form>
              </div>



            <div class="flex items-center lg:order-2">
                <a href="#" class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800">About</a>
                <a href="#" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">Contact</a>
        
            </div>
        </div>
    </nav>
</header>

</template>

<script lang="ts">
export default {
    methods: {
        reload() {
            location.reload();
        }
    },
};
</script>

