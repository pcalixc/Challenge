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
                <svg class="fill-current h-6 w-6 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M12.9 14.32a8 8 0 1 1 1.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12z" /></svg>
              </div>
              <input id="searchInput"  v-model="term"  @keyup.enter="Search()" class="focus:ring w-full border pl-12 pr-4 py-2 rounded-md focus:shadow-outline outline-none" type="text" placeholder="Search email..." />
              <button @click="Search" type="submit" class="text-white absolute right-1 bottom-1 bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>
            </div>
          </div>
          <!-- ---- -->
            <div class="flex items-center lg:order-2">
                <router-link to="/contact" href="" class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800">Contact</router-link>
                <router-link to="/about" href="" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">Read Me</router-link>
        
            </div>
        </div>
    </nav>
</header>


  <div class=" bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-blue-700 via-blue-800 to-gray-900 h-screen grid grid-cols-5 gap-4 p-4 pt-24" >
    <div class="col-span-3 rounded-xl bg-white m-3 ">
      <div  class="flex-grow">
        <div class="flex py-0 w-full">
            <div class="shadow-md sm:rounded-lg h-96 relative w-full  max-h-full">
              <table v-if="emailsFounded" class="overflow-y-auto w-full  align-middle h-80 table-auto grow text-sm text-left text-gray-500 dark:text-gray-400">
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
                          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-gray-900 bg-blue-100 rounded-full">
                            {{ email.from }}
                          </span>
                        </td>
                        <td class="px-4 py-1 text-xs">
                          {{ email.to }}
                        </td>
                        <td class="px-4 py-1 text-xs font-semibold text-slate-600">
                          {{ email.date }}
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
            <div
              class="px-4 py-0 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between mt-20         ">
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
          <div class=" flex justify-center  text-gray-800 p-3 text-xl font-bold ">
            <h4> {{ emailSelectedValues.subject }}
            </h4>
          </div>
          <div class="mx-4 font-bold text-blue-900">
            👤 {{  emailSelectedValues.from }}
          </div>
          <div class="mx-4">
            <span class="inline-flex  text-lt font-semibold ml-3">
              To: {{ emailSelectedValues.to }}
            </span>
          </div>

          <div class="mx-4 my-5 text-xs font-ligth  overflow-auto h-72 custom-scrollbar">
            
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
  emailsFounded
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

