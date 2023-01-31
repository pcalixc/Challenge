<!-- eslint-disable vue/multi-word-component-names -->
<!-- eslint-disable vue/valid-v-on -->

<template>
  <!-- Inicio Header -->
  <header >
    <nav class=" fixed w-screen border-gray-200 px-4 lg:px-6 py-1 bg-gray-800 ">
        <div class="flex  justify-between mx-auto max-w-screen-xl">
            <a href="http://localhost:8080/" class="flex items-center">
                <img src="../assets/email-icon.png" class="mr-3 h-6 sm:h-9" alt="Logo" />
                <span class="self-center text-xl font-bold whitespace-nowrap dark:text-white">Enron Mail</span>
            </a>
            <!-- Search Bar -->
            <div class="m-2 px-4 w-1/4">
              <div class="relative">
                <div class="absolute left-0 inset-y-0 pl-3 flex items-center">
                  <img class="w-4 h-4" src="https://cdn-icons.flaticon.com/svg/3917/3917132.svg?token=exp=1675113557~hmac=3f74fbf732d9ee0c6cd1eb70db16c2db" alt="">
                </div>
                <input id="searchInput"  v-model="term"  @keyup.enter="Search()" class="focus:ring w-full border pl-12 pr-4 py-2 rounded-md focus:shadow-outline outline-none" type="text" placeholder="Search email..." />
                <button @click="Search" type="submit" class="text-white absolute right-1 bottom-1 bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>
              </div>
            </div>
            <!-- contact-about button -->
            <div class="flex items-center lg:order-2">
                <button @click="showModal = true" class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800">Contact</button>
                <router-link to="/about" href="" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">About</router-link>
            </div>
        </div>
    </nav>
  </header>
  <!-- Fin Header -->


  <!-- Inicio HomeScreen -->
  <div class=" bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-blue-700 via-blue-800 to-gray-900 h-screen grid grid-cols-5 gap-4 p-4 pt-24" >
    <div v-if="(termSearched=='')" class="absolute top-20 pl-4 text-white text-md font-semibold"> Showing all resulst</div>
    <div v-else class="absolute top-20 pl-4 text-white text-md font-semibold"> Showing results for <span class="bg-green-100 pl-2 pr-2 rounded-2xl text-gray-900 font-semibold"> {{ termSearched }} </span> </div>
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
                          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-gray-900 bg-blue-100 rounded-full">
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
      
      <!-- Pagination -->
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
      <!--Fin Pagination -->


      <!-- Mail content -->
      <div class="col-span-2  rounded-xl shadow-lg border py-4 px-5 m-3 bg-white h-90 ">
        <div v-if="emailSelected" class="  borderborder-gray-500 scroll-y-auto w-92 " >
          <div class=" flex justify-center items-center text-gray-800 p-3 text-xl font-bold ">
            <div v-for="word in emailSelectedValues.subject.split(' ')" v-bind:key="word">
              <span v-if="((word.toLowerCase()===termSearched.toLowerCase()) && termSearched != '')" class="bg-green-200 text-gray-900 rounded-2xl px-1"> {{word}}</span> 
              <span v-else class="px-1" > {{word}}</span> 
              </div>
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
          <div class="flex justify-center mx-4 my-5 text-sm font-ligth overflow-y-auto h-72 custom-scrollbar">
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
  <!-- Fin HomeScreen -->

  <!-- Contact modal -->
  <div>
    <div v-if="showModal" class="transition ease-in-out delay-150 fixed top-0 left-0 h-screen w-screen flex items-center justify-center" @click.away="showModal = false">
      <div class="transition ease-linear  bg-black bg-opacity-60 fixed top-0 left-0 h-screen w-screen"></div>
      
        <div class="relative w-full group max-w-md min-w-0 mx-auto mt-6 mb-6 break-words  shadow-3xl md:max-w-sm rounded-xl">
            <div class="pb-6  bg-gray-900 rounded-xl">
                <div class="flex flex-wrap justify-center">
                    <div class="flex justify-center w-full">
                        <div class="ml-14">
                            <img src="" class="relative dark:shadow-xl border-white dark:border-gray-800  rounded-full align-middle border-8  -m-18 -ml-18 lg:-ml-16 max-w-[150px]" />
                        </div>
                    </div>
                </div>
            <div class="mt-2  text-center ">
                <h3 class="mb-1 text-2xl font-bold leading-normal dark:text-gray-300">Paola Gisselle Cálix</h3>
                <div class="flex flex-row justify-center w-full mx-auto space-x-2 text-center">
                    <svg  class="w-4 h-4  text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd"></path>
                    </svg>
                    <div class="font-bold tracking-wide text-gray-600 dark:text-gray-300 font-mono text-xl">Systems Ingineering Student</div>
                </div>
            <div class="w-full text-center">
                <div class="flex justify-center pt-8 pb-0 lg:pt-4">
                    <div class="flex space-x-2">
                    

                        <a class="p-1 -m-1 text-gray-400 hover:text-pink-500 focus:outline-none focus-visible:ring-2 ring-primary" href="https://github.com/pcalixc" rel="noopener" aria-label="Ariel Cerda on Github" target="_blank">
                            <svg class="w-6 h-6 overflow-visible fill-current" alt="" aria-hidden="true" viewBox="0 0 140 140">
                                <path
                                    d="M70 1.633a70 70 0 00-21.7 136.559h1.692a5.833 5.833 0 006.183-6.184v-1.225-6.358a2.917 2.917 0 00-1.167-1.925 2.917 2.917 0 00-2.45-.583C36.925 125.3 33.6 115.5 33.367 114.858a27.067 27.067 0 00-10.034-12.775 6.767 6.767 0 00-.875-.641 3.675 3.675 0 012.217-.409 8.575 8.575 0 016.708 5.134 17.5 17.5 0 0023.334 6.766 2.917 2.917 0 001.691-2.1 11.667 11.667 0 013.267-7.175 2.917 2.917 0 00-1.575-5.075c-13.825-1.575-27.942-6.416-27.942-30.275a23.333 23.333 0 016.125-16.216A2.917 2.917 0 0036.808 49a20.533 20.533 0 01.059-14 32.317 32.317 0 0114.7 6.708 2.858 2.858 0 002.45.409A61.892 61.892 0 0170 39.958a61.075 61.075 0 0116.042 2.159 2.858 2.858 0 002.391-.409 32.608 32.608 0 0114.7-6.708 20.825 20.825 0 010 13.883 2.917 2.917 0 00.525 3.092 23.333 23.333 0 016.125 16.042c0 23.858-14.175 28.641-28.058 30.216a2.917 2.917 0 00-1.575 5.134 12.833 12.833 0 013.558 10.15v18.55a6.183 6.183 0 002.1 4.841 7 7 0 006.184 1.109A70 70 0 0070 1.633z"
                                ></path>
                            </svg>
                        </a>

                        <a
                            class="p-1 -m-1 text-gray-400 hover:text-pink-500 focus:outline-none focus-visible:ring-2 ring-primary"
                            href="https://www.linkedin.com/in/paola-c-9b5905226/"
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
    var termSearched= ref("")


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
    emailSelectedValues.value=('')
    emailSelectedValues.value=email
    console.log("em value",emailSelectedValues.value)
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
  }
  //const content= emailsByPage.value[1].content.split('')
  //console.log("here",emailsByPage.value[1].content)

  page.value = page_
}

function Search() {
    termSearched.value=term.value

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
  showModal,
  termSearched
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

