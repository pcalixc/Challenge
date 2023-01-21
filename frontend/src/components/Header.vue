<!-- eslint-disable vue/multi-word-component-names -->
<template>   
<header >
    <nav class=" fixed w-screen border-gray-200 px-4 lg:px-6 py-1 bg-gray-800 ">
        <div class="flex  justify-between mx-auto max-w-screen-xl">
          <!-- ------ -->
            <a href="#" class="flex items-center">
                <img src="../assets/email-icon.png" class="mr-3 h-6 sm:h-9" alt="Logo" />
                <span class="self-center text-xl font-bold whitespace-nowrap dark:text-white">Enron Mail</span>
            </a>


            <!-- ------ -->
            <div class="m-2 px-4 w-1/4">
              <div class="relative">
              <div class="absolute left-0 inset-y-0 pl-3 flex items-center">
                <svg class="fill-current h-6 w-6 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M12.9 14.32a8 8 0 1 1 1.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12z" /></svg>
              </div>

              <input class="w-full border pl-12 pr-4 py-2 rounded-md focus:shadow-outline outline-none" type="text" placeholder="Search email..." />
            </div>
          </div>


            


          <!-- ---- -->
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

