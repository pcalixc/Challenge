<!-- eslint-disable vue/multi-word-component-names -->
<template>

  <div class=" bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-blue-700 via-blue-800 to-gray-900 h-screen grid grid-cols-5 gap-4 p-4 pt-24" >

    <div class="col-span-3 rounded-xl bg-white m-3">
            <!-- table emails-->
            <div class="rounded-xl shadow">
                <table class=" text-sm text-left w-full ">
                    <!--header table with from, subject and date -->
                    <thead class="text-sm email-head  bg-gray-800 text-white  ">
                        <tr>
                            <th  class="py-3 px-6  ">
                                From
                            </th>
                            <th  class="py-3 px-6 ">
                                To
                            </th>
                            <th  class="py-3 px-6 ">
                                Subject
                            </th>
                            <th  class="py-3 px-6 ">
                                Date
                            </th>
                        </tr>
                    </thead>
                    <!-- end header table with from, subject and date -->
                    <tbody>
                        <tr 
                            
                            class=" border-b bg-white dark:border-gray-700 cursor-pointer email-body text-black hover:bg-gray-100">
                            <td class=" py-4 px-6 r font-medium"> brent.hendry@enron.com
                            </td>
                            <td class="py-4 px-6 r" >
                              <span class="inline-flex px-2 text-lt font-semibold leading-5 text-blue-800 bg-green-100 rounded-full">
                                araspe@elementre.com
            </span>
                            </td>
                            <td class="py-4 px-6 text-sm" >RE: XL Trading Partners Ltd
                            </td>
                            <td class="py-4 px-6 text-gray-600 font-semibold text-xs"  >Mon, 8 Oct 2001
                            </td>
                        </tr>
           
                        
                    
                    
                    </tbody>
                </table>
                <!-- end table emails-->
            </div>
        
        <!-- email content preview-->
     
        <!-- end email content preview-->
    
            <div
              class="px-5 py-5 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between          ">
              <span class="text-xs xs:text-sm text-gray-900">
                              Showing 1 to 4 of 50 Entries
              </span>
              <div class="inline-flex mt-2 xs:mt-0">
                <button
                                  class="text-sm text-indigo-50 transition duration-150 hover:bg-blue-700 bg-blue-900 font-semibold  px-4 rounded-l">
                                  Prev
                              </button>
                &nbsp; &nbsp;
                <button
                                  class="text-sm text-indigo-50 transition duration-150 hover:bg-blue-700 bg-blue-900 font-semibold py-1 px-4 rounded-r">
                                  Next
                              </button>
              </div>
            </div>
          
        
      
    </div>
 
    <div class="col-span-2 bg-white rounded-xl shadow-lg border py-4 px-5 m-3">
        <div class="  borderborder-gray-500" >
          <div class=" flex justify-center  text-gray-800 p-3 text-xl font-bold ">
            <h4 >RE: XL Trading Partners Ltd</h4>
            <img src="../assets/mail_5.png" alt="" class="h-11 w-11 ml-4">
          </div>
          <div class="mx-4 font-bold text-blue-900">
           
            brent.hendry@enron.com
          </div>

          <div class="mx-4">
            <span class="inline-flex  text-lt font-semibold ">
              To:
            </span>
            araspe@elementre.com
          </div>
          

         
  
          <div class="mx-4 my-5 text-sm font-ligth ">
            
            Adele,
            I reviewed the opinion and there were several changes to the form from the prior draft.  I would like to discuss two of the additions with your counsel if that would be ok with you.  The first addition at issue is subpart k at the end of the first paragraph on page two.  It looks like we may need a representation from you that you are not and will not carry on investment business in or from within Bermuda and that you will not perform your obligations under the Agreement in or from within Bermuda in order for the opinion to have value.  The second addition at issue is subpart (e) under section 8 of the opinion.  As written it looks like a large loop hole that can be taken advantage of just by finding evidence that was not submitted at the trial.  Our General Counsel is concerned with what looks like an easy way to circumvent a foreign judgment.  I would appreciate your thoughts or any clarification I can get. 
          </div>
        </div>
    </div>

  </div>

</template>


<!-- <script>

import axios from 'axios';
import moment from 'moment';
export default {
  name: 'MyIndexer',
  data() {
    return {
      text: '',
      user: '',
      users: [],
      fields: ['subject', 'from', 'to'],
      emails: [],
      email: {},
      showTable: false,
      showContent: false,
      chargeEmails: false,
    }
  },

  created() {
    this.getUsers();
  },

  methods: {
    async searchInformation() {
      let response = await axios.get(`http://localhost:8000/search/${this.text}`);
      this.showTable = true;
      this.emails = response.data.emails.map(email => {
        return {
          subject: email.subject,
          content: email.content,
          to: email.to,
          date: moment(email.date).format("MMM Do YYYY"),
          from: email.from,
          id: email.message_id
        }
      });
    },

    async getUsers() {
      let response = await axios.get('http://localhost:8000/enron/users/');
      this.users = response.data.users;
    },

    async getUserInformation() {
      this.chargeEmails = true;
      let response = await axios.get(`http://localhost:8000/enron/users/${this.user}`);
      if (response) {
        alert('Se insertaron ' + response.data.emails.record_count + ' registros');
      }
      this.chargeEmails = false;
    },

    setEmailInformation(email) {
      this.email = email;
    }
  }
}
</script> -->

