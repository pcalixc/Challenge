<!-- eslint-disable vue/valid-v-on -->
<!-- eslint-disable vue/multi-word-component-names -->
<template>
  
  <header >
    <nav class=" fixed w-screen border-gray-200 px-4 lg:px-6 py-1 bg-gray-800 ">
        <div class="flex  justify-between mx-auto max-w-screen-xl">
            <a href="http://localhost:8080/" class="flex items-center">
                <img src="../assets/email-icon.png" class="mr-3 h-6 sm:h-9" alt="Logo" />
                <span class="self-center text-xl font-bold whitespace-nowrap dark:text-white">Enron Mail</span>
            </a>
              <!-- Search bar -->
            <div class="m-2 px-4 w-1/4">
              <div class="relative">
                <div class="absolute left-0 inset-y-0 pl-3 flex items-center">
                  <svg class="fill-current h-6 w-6 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"></svg>
                </div>
                <input id="searchInput"  v-model="term"  @keyup.enter="Search()" class="focus:ring w-full border pl-12 pr-4 py-2 rounded-md focus:shadow-outline outline-none" type="text" placeholder="Search email..." />
                <button @click="Search" type="submit" class="text-white absolute right-1 bottom-1 bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>
              </div>
            </div>
          <!-- ---- -->
            <div class="flex items-center lg:order-2">
                <button @click="showModal = true" class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800">Contact</button>
                <router-link to="/about" href="" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">About</router-link>
            </div>
        </div>
    </nav>
  </header>


  <div class=" bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-blue-700 via-blue-800 to-gray-900 h-screen grid grid-cols-5 gap-4 p-4 pt-24" >
    <div class="col-span-3 rounded-xl bg-white m-3 ">
      <div  class="flex-grow">
        <div class="flex py-0 w-full">
            <div class="shadow-md sm:rounded-lg h-96 relative w-full  max-h-full">
              <table v-if="emailsFounded" class="overflow-y-auto w-full  align-middle h-80 h- table-auto grow text-sm text-left text-gray-500 dark:text-gray-400">
                  <thead class="sticky top-0 text-md text-gray-700  bg-gray-800 dark:text-white w-full">
                    <tr>
                        <th scope="col" class="px-4 py-3">Subject</th>
                        <th scope="col" class="px-4 py-3">From</th>
                        <th scope="col" class="px-4 py-3">To</th>
                        <th scope="col" class="px-4 py-3">Date</th>
                    </tr>
                  </thead>
                  <tbody  class="overflow-y-auto " v-for="email in emailsByPage" :key="email.message_id">
                    <tr @click="viewEmail(email)" class=" border-b  text-gray-900 bg-white dark:border-gray-700 hover:bg-gray-200  cursor-pointer rounded-full">
                        <td class="px-4 py-1 text-xs font-semibold  h-14">
                          {{ email.subject }}
                        </td>
                        <td class="px-4 py-1 text-xs">
                            {{ email.from }}
                        </td>
                        <td class="px-4 py-1 text-xs">
                          {{ email.to }}
                        </td>
                        <td class="px-4 py-1 text-xs font-semibold text-slate-600">
                          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-gray-900 bg-green-100 rounded-full">
                            {{ email.date }}
                          </span>
                        </td>         
                    </tr>
                  </tbody>
              </table>
              <div v-if="!emailsFounded" class="flex items-center justify-center">
                <div class="flex  items-center justify-center m-10   w-60 bg-yellow-100 rounded-lg p-4 mb-4 text-sm text-yellow-700" role="alert">
                  <svg class="w-5 h-5 inline mr-3" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
                  <div>
                    <span class="font-medium">No emails founded</span> 
                  </div>
                </div>
              </div>
            </div>
        </div>
      </div>
            <!-- pagination -->
      <div class="px-4 py-0 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between mt-20         ">
        <span class="text-xs xs:text-sm font-bold text-gray-900">
          Showing {{ ini }} to {{ fin }} of {{ totalResults }} Entries
        </span>
        <div class="inline-flex mt-2 xs:mt-0">
          <button v-if="(page==1)"
            class="text-sm text-indigo-50 transition duration-150 bg-gray-500 font-semibold  px-4 rounded-l">
              Prev
          </button>
          <button 
            v-if="(page>1)"
            @click="pagination(totalResults, page -1)" 
            class="text-sm text-indigo-50 transition duration-150 hover:bg-blue-700 bg-blue-900 font-semibold  px-4 rounded-l">
                Prev
          </button>
                &nbsp; &nbsp;
          <button 
            v-if="(page==totalPages)"
            class="text-sm text-indigo-50 transition duration-150 bg-gray-500 font-semibold py-1 px-4 rounded-r">
              Next
          </button>
          <button 
            @click="pagination(totalResults, page + 1)"
            class="text-sm text-indigo-50 transition duration-150 hover:bg-blue-700 bg-blue-900 font-semibold py-1 px-4 rounded-r">
              Next
          </button>
        </div>
      </div>  
      </div>
      <!-- mail content -->
      <div class="col-span-2  rounded-xl shadow-lg border py-4 px-5 m-3 bg-white h-90">
        <div v-if="emailSelected" class="  borderborder-gray-500 scroll-y-auto" >
          <div class=" flex justify-center items-center text-gray-800 p-3 text-xl font-bold ">
            <h4> {{ emailSelectedValues.subject }}
            </h4>
          </div>
          <div class="mx-4 font-bold text-blue-900">
            {{  emailSelectedValues.from }}
          </div>
          <div class="mx-4">
            <span class="inline-flex justify-center items-center text-lt font-semibold ">
              <img class="w-4 h-4 mr-4" src="https://www.flaticon.com/svg/vstatic/svg/9252/9252742.svg?token=exp=1674883372~hmac=93cb570e03deca7818768cd849b3cdc2" alt="">
              {{emailSelectedValues.to }}
            </span>
          </div>
          <div class="mx-4 my-5 text-sm font-ligth  overflow-auto h-72 custom-scrollbar">
            {{ emailSelectedValues.content }}
          </div>
        </div>

        <div v-if="!emailSelected" class="font-bold text-lg flex items-center justify-center m-6">
          <div class="flex bg-blue-100 rounded-lg p-4 mb-4 text-sm text-blue-700" role="alert">
            <svg class="w-5 h-5 inline mr-3" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
            <div>
              <span class="font-medium">No email Selected</span> 
            </div>
          </div>
        </div>
    </div>
  </div>

  <div>
    
    <div v-if="showModal" class=" transition ease-linear fixed top-0 left-0 h-screen w-screen flex items-center justify-center" @click.away="showModal = false">
      <div class="transition ease-linear  bg-black bg-opacity-60 fixed top-0 left-0 h-screen w-screen"></div>
      
        <div class="relative w-full group max-w-md min-w-0 mx-auto mt-6 mb-6 break-words bg-white border shadow-3xl dark:bg-gray-800 dark:border-gray-700 md:max-w-sm rounded-xl">
            <div class="pb-6  bg-gray-900 rounded-lg">
                <div class="flex flex-wrap justify-center">
                    <div class="flex justify-center w-full">
                        <div class="ml-14">
                            <img src="../assets/email-icon.png" class="dark:shadow-xl border-white dark:border-gray-800  rounded-full align-middle border-8  -m-18 -ml-18 lg:-ml-16 max-w-[150px]" />
                        </div>
                    </div>
                </div>
            <div class="mt-2  text-center ">
                <h3 class="mb-1 text-2xl font-bold leading-normal dark:text-gray-300">Paola Cálix</h3>
                <div class="flex flex-row justify-center w-full mx-auto space-x-2 text-center">
                    <svg  class="w-4 h-4  text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd"></path>
                    </svg>
                    <div class="font-bold tracking-wide text-gray-600 dark:text-gray-300 font-mono text-xl">Systems Ingineering Student</div>
                </div>
            <div class="w-full text-center">
                <div class="flex justify-center pt-8 pb-0 lg:pt-4">
                    <div class="flex space-x-2">
                    

                        <a class="p-1 -m-1 text-gray-400 hover:text-pink-500 focus:outline-none focus-visible:ring-2 ring-primary" href="https://www.github.com/Smilesharks" rel="noopener" aria-label="Ariel Cerda on Github" target="_blank">
                            <svg class="w-6 h-6 overflow-visible fill-current" alt="" aria-hidden="true" viewBox="0 0 140 140">
                                <path
                                    d="M70 1.633a70 70 0 00-21.7 136.559h1.692a5.833 5.833 0 006.183-6.184v-1.225-6.358a2.917 2.917 0 00-1.167-1.925 2.917 2.917 0 00-2.45-.583C36.925 125.3 33.6 115.5 33.367 114.858a27.067 27.067 0 00-10.034-12.775 6.767 6.767 0 00-.875-.641 3.675 3.675 0 012.217-.409 8.575 8.575 0 016.708 5.134 17.5 17.5 0 0023.334 6.766 2.917 2.917 0 001.691-2.1 11.667 11.667 0 013.267-7.175 2.917 2.917 0 00-1.575-5.075c-13.825-1.575-27.942-6.416-27.942-30.275a23.333 23.333 0 016.125-16.216A2.917 2.917 0 0036.808 49a20.533 20.533 0 01.059-14 32.317 32.317 0 0114.7 6.708 2.858 2.858 0 002.45.409A61.892 61.892 0 0170 39.958a61.075 61.075 0 0116.042 2.159 2.858 2.858 0 002.391-.409 32.608 32.608 0 0114.7-6.708 20.825 20.825 0 010 13.883 2.917 2.917 0 00.525 3.092 23.333 23.333 0 016.125 16.042c0 23.858-14.175 28.641-28.058 30.216a2.917 2.917 0 00-1.575 5.134 12.833 12.833 0 013.558 10.15v18.55a6.183 6.183 0 002.1 4.841 7 7 0 006.184 1.109A70 70 0 0070 1.633z"
                                ></path>
                            </svg>
                        </a>

                        <a
                            class="p-1 -m-1 text-gray-400 hover:text-pink-500 focus:outline-none focus-visible:ring-2 ring-primary"
                            href="https://www.linkedin.com/in/"
                            rel="noopener"
                            aria-label=""
                            target="_blank"
                        >
                            <svg class="w-6 h-6 overflow-visible fill-current" alt="" aria-hidden="true" viewBox="0 0 140 140">
                                <path
                                    d="M23.4 44.59h-4.75a12.76 12.76 0 00-9.73 2.19 9.44 9.44 0 00-2.35 7.12V131a9.08 9.08 0 002.3 7 9.24 9.24 0 006.82 2c2.22 0 4.15-.21 8.24-.06a12 12 0 009.25-2 9.1 9.1 0 002.29-7V53.9a9.44 9.44 0 00-2.34-7.12 12.68 12.68 0 00-9.73-2.19zM21 0A16.19 16.19 0 005.09 15.6 16.52 16.52 0 0021 31.86 16.12 16.12 0 0037 15.6 16.18 16.18 0 0021 0zM99.74 43.63a31.09 31.09 0 00-20.93 6.3A7.25 7.25 0 0077 46.34a6.08 6.08 0 00-4.52-1.78 119.08 119.08 0 00-15 .3c-4.16.84-6.18 3.79-6.18 9V131a9.14 9.14 0 002.28 7 12.06 12.06 0 009.26 2c4.47-.17 5.74.06 8.22.06a9.26 9.26 0 006.83-2 9.12 9.12 0 002.3-7V89.88A12.48 12.48 0 0192.93 76 12.44 12.44 0 01106 89.88V131a9.1 9.1 0 002.29 7 12 12 0 009.24 2c1.83-.07 4-.07 5.8 0a12.09 12.09 0 009.26-2 9.14 9.14 0 002.28-7V78.32a33.07 33.07 0 00-35.13-34.69z"
                                ></path>
                            </svg>
                        </a>

                        <a
                            class="p-1 -m-1 text-gray-400 hover:text-pink-500 focus:outline-none focus-visible:ring-2 ring-primary"
                            href="https://www.instagram.com/"
                            rel="noopener"
                            aria-label=""
                            target="_blank"
                        >
                            <svg class="w-6 h-6 overflow-visible fill-current" alt="" aria-hidden="true" viewBox="0 0 140 140" xmlns="http://www.w3.org/2000/svg">
                                <path
                                    fill-rule="evenodd"
                                    clip-rule="evenodd"
                                    d="M73.1735 2C89.9405 2 92.3831 2.0897 99.4487 2.414C106.79 2.7521 111.807 3.9182 116.195 5.6225C120.801 7.35353 124.973 10.0686 128.422 13.5782C131.931 17.0266 134.646 21.1993 136.378 25.805C138.082 30.1934 139.248 35.2097 139.586 42.5513C139.917 49.9136 140 52.2596 140 71V71.552C140 89.7887 139.917 92.1623 139.586 99.4487C139.248 106.79 138.082 111.807 136.378 116.195C134.646 120.801 131.931 124.973 128.422 128.422C124.973 131.931 120.801 134.646 116.195 136.378C111.807 138.082 106.79 139.248 99.4487 139.586C92.0864 139.917 89.7404 140 71 140H70.448C52.2113 140 49.8377 139.917 42.5513 139.586C35.2097 139.248 30.1934 138.082 25.805 136.378C21.1993 134.646 17.0266 131.931 13.5782 128.422C10.0686 124.973 7.35353 120.801 5.6225 116.195C3.9182 111.807 2.7521 106.79 2.414 99.4487C2.0897 92.3831 2 89.9336 2 73.1735V68.8265C2 52.0595 2.0897 49.6169 2.414 42.5513C2.7521 35.2097 3.9182 30.1934 5.6225 25.805C7.35353 21.1993 10.0686 17.0266 13.5782 13.5782C17.0266 10.0686 21.1993 7.35353 25.805 5.6225C30.1934 3.9182 35.2097 2.7521 42.5513 2.414C49.6169 2.0897 52.0664 2 68.8265 2H73.1735V2ZM72.6146 14.4338H69.3854C52.439 14.4338 50.1758 14.5097 43.1171 14.834C36.3896 15.1445 32.7395 16.2623 30.3038 17.2076C27.0815 18.4634 24.7838 19.9538 22.3688 22.3688C19.9538 24.7838 18.4634 27.0815 17.2076 30.3038C16.2623 32.7395 15.1376 36.3896 14.834 43.1171C14.5097 50.1758 14.4338 52.439 14.4338 69.3854V72.6146C14.4338 89.561 14.5097 91.8242 14.834 98.8829C15.1445 105.61 16.2623 109.261 17.2076 111.696C18.4634 114.912 19.9607 117.216 22.3688 119.631C24.7838 122.046 27.0815 123.537 30.3038 124.792C32.7395 125.738 36.3896 126.862 43.1171 127.166C50.3897 127.497 52.5701 127.566 71 127.566H71.552C89.4713 127.566 91.6793 127.497 98.876 127.166C105.61 126.855 109.261 125.738 111.696 124.792C114.912 123.537 117.216 122.046 119.631 119.631C122.046 117.216 123.537 114.919 124.792 111.696C125.738 109.261 126.862 105.61 127.166 98.8829C127.497 91.6034 127.566 89.4299 127.566 71V70.448C127.566 52.5287 127.497 50.3207 127.166 43.124C126.855 36.3896 125.738 32.7395 124.792 30.3038C123.687 27.3062 121.923 24.5947 119.631 22.3688C117.405 20.077 114.694 18.3134 111.696 17.2076C109.261 16.2623 105.61 15.1376 98.8829 14.834C91.8242 14.5097 89.561 14.4338 72.6146 14.4338ZM71 35.5685C75.6529 35.5685 80.2603 36.485 84.559 38.2656C88.8578 40.0462 92.7637 42.656 96.0539 45.9461C99.344 49.2363 101.954 53.1422 103.734 57.441C105.515 61.7397 106.432 66.3471 106.432 71C106.432 75.6529 105.515 80.2603 103.734 84.559C101.954 88.8578 99.344 92.7637 96.0539 96.0539C92.7637 99.344 88.8578 101.954 84.559 103.734C80.2603 105.515 75.6529 106.432 71 106.432C61.603 106.432 52.5908 102.699 45.9461 96.0539C39.3015 89.4092 35.5685 80.397 35.5685 71C35.5685 61.603 39.3015 52.5908 45.9461 45.9461C52.5908 39.3015 61.603 35.5685 71 35.5685V35.5685ZM71 48.0023C64.9006 48.0023 59.0511 50.4253 54.7382 54.7382C50.4253 59.0511 48.0023 64.9006 48.0023 71C48.0023 77.0994 50.4253 82.9489 54.7382 87.2618C59.0511 91.5747 64.9006 93.9977 71 93.9977C77.0994 93.9977 82.9489 91.5747 87.2618 87.2618C91.5747 82.9489 93.9977 77.0994 93.9977 71C93.9977 64.9006 91.5747 59.0511 87.2618 54.7382C82.9489 50.4253 77.0994 48.0023 71 48.0023V48.0023ZM107.832 25.8878C110.028 25.8878 112.134 26.7602 113.687 28.313C115.24 29.8658 116.112 31.9718 116.112 34.1678C116.112 36.3638 115.24 38.4698 113.687 40.0226C112.134 41.5754 110.028 42.4478 107.832 42.4478C105.636 42.4478 103.53 41.5754 101.977 40.0226C100.425 38.4698 99.5522 36.3638 99.5522 34.1678C99.5522 31.9718 100.425 29.8658 101.977 28.313C103.53 26.7602 105.636 25.8878 107.832 25.8878Z"
                                ></path>
                            </svg>
                        </a>  
                    </div>
                </div>
            </div>
        </div>
        <div class="pt-6 mx-6 mt-6 text-center border-t border-gray-200 dark:border-gray-700/50">
            <div class="flex flex-wrap justify-center">
                <div class="w-full px-6">
                    <p class="mb-4 font-light leading-relaxed text-gray-600 dark:text-gray-400">
                        Committed to learning and developing skills related to web development, databases, and related fields. 
                    </p>
                </div>
            </div>
        </div>
        <div class="relative h-6 overflow-hidden translate-y-6 rounded-b-lg">
            <div class="absolute flex -space-x-12 rounded-b-2xl">
                <div class="w-36 h-8 transition-colors duration-200 delay-75 transform skew-x-[35deg] bg-pink-400/90 group-hover:bg-pink-600/90 z-10"></div>
                <div class="w-28 h-8 transition-colors duration-200 delay-100 transform skew-x-[35deg] bg-pink-300/90 group-hover:bg-pink-500/90 z-20"></div>
                <div class="w-28 h-8 transition-colors duration-200 delay-150 transform skew-x-[35deg] bg-pink-200/90 group-hover:bg-pink-400/90 z-30"></div>
                <div class="w-28 h-8 transition-colors duration-200 delay-200 transform skew-x-[35deg] bg-pink-100/90 group-hover:bg-pink-300/90 z-40"></div>
                <div class="w-28 h-8 transition-colors duration-200 delay-300 transform skew-x-[35deg] bg-pink-50/90 group-hover:bg-pink-200/90 z-50"></div>
            </div>
        </div>
            </div>
        </div>
      
    </div>
  </div>

</template>


<script lang="ts">
import axios from 'axios';
import {ref} from 'vue'
import moment from 'moment'

export default {
  setup() {
    type Email = {
    message_id: string
    subject: string
    date: string
    from: string
    to: string
    content: string
  }
    
    const showModal = ref(false)

    const emails = ref<Email[]>([])
    var ini = ref(0)
    var fin = ref(0)
    var totalResults= ref(0)
    var emailSelected= ref(false)
    var emailSelectedValues= ref()
    var emailsByPage= ref<Email[]>([])
    var totalPages= ref(0)
    var page= ref(1)
    var searchString= ref("")
    var term = ref("")
    var emailsFounded= ref(true)


  getEmails('http://localhost:8004/search/')

  async function getEmails(url:string){
    emails.value=([])
    term.value=""
    emailSelected.value= false
    const response = await getData(url)
    totalResults.value = response.hits.Total.value

    if(totalResults.value>0){
      let email = response.hits.hits
      for (let i = 0; i < 100; i++) {
        let Email = {
          message_id: email[i]._source.message_id,
          subject: email[i]._source.subject,
          date: moment(email[i]._source.date).format("L"),
          from: email[i]._source.from,
          to: email[i]._source.to,
          content: email[i]._source.content,
        };
      emails.value.push(Email);
    }
    pagination(totalResults.value, page.value)
    }
    else{
      emailsFounded.value=false
    }
  }

  async function getData(url: string) {
      try {
        const response = await axios.get(url);
        return response.data;
      } catch (error) {
        console.error(error);
        throw error;
      }
    }

  const viewEmail = (email:Email) =>{
    emailSelected.value= true
    emailSelectedValues.value=email
    return emailSelectedValues
  }

  const pagination = ( totalR:number, page_:number)=>{
    emailsByPage.value = ([])
    const pageElements=7
    totalPages.value = Math.ceil(totalR / pageElements) 
    console.log("totalPages",totalPages.value)
    ini.value = (page_* pageElements) - pageElements
    fin.value = (page_ * pageElements)
    for (let index:number = ini.value; index< fin.value; index++){
      emailsByPage.value.push(emails.value[index])
      console.log(emails.value[1].toString)
  }
  page.value = page_
}

function Search() {
    searchString.value=(  'http://localhost:8004/search/'+ term.value)
    console.log(searchString.value)
    getEmails(searchString.value)
    
  }

return {
  getData,
  getEmails,
  emails,
  ini,
  fin,
  totalResults,
  emailSelected,
  viewEmail,
  emailSelectedValues,
  pagination,
  emailsByPage,
  totalPages,
  page,
  searchString,
  term,
  Search,
  emailsFounded,
  showModal
} 
  },
}

</script>

<style>
.custom-scrollbar::-webkit-scrollbar {
  width: 10px; 
  background-color: gray;
}


/* Track */
.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1; 
}
 
/* Handle */
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #888; 
}

/* Handle on hover */
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #555; 
}

</style>

