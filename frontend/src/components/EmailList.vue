<!-- eslint-disable vue/multi-word-component-names -->
<template>

  <div class=" bg-slate-700 h-screen grid grid-cols-3 gap-4 p-4 pt-24" >

    <div class="col-span-2 bg-slate-200 rounded-xl shadow">
      <div class="px-4 ">
        <div class="-mx-4 sm:-mx-8 px-4 sm:px-8 py-4 ">
          <div class="inline-block min-w-full shadow rounded-lg ">
            <table class="min-w-full leading-normal overflow-auto ">
              <thead class="rounded-xl shadow border">
                <tr class="bg-gray-700 text-white rounded-xl shadow border ">
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 text-left text-md font-semibold tracking-wider ">
                    Subject
                  </th>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 text-left text-md font-semibold tracking-wider">
                    From
                  </th>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 text-left text-md font-semibold tracking-wider">
                    To
                  </th>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 text-left text-md font-semibold tracking-wider">
                    Content
                  </th>
                  <th
                  class="px-5 py-3 border-b-2 border-gray-200 text-left text-md font-semibold tracking-wider">
                  
                </th>
                 
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                      <p class="text-gray-900 whitespace-no-wrap">
                            Invitation
                      </p>
                       
                  </td>
                  <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                    <p class="text-gray-900 whitespace-no-wrap">@HemmaRsko</p>
                  </td>
                  <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                    <p class="text-gray-900 whitespace-no-wrap">
                      @demonst
                    </p>
                  </td>
                  <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                    <p class="text-gray-900 whitespace-no-wrap truncate">
                      Hola...
                    </p>
                  </td>

                  <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                    <span
                                          class="relative inline-block px-3 py-1 text-right font-semibold text-green-900 ">
                                          <span aria-hidden
                                              class="absolute inset-0 bg-green-200 opacity-50 rounded-full"></span>
                    <span class="relative">View Full</span>
                    </span>
                  </td>
                  
                </tr>
   

                
              </tbody>
            </table>
            <div
              class="px-5 py-5 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between          ">
              <span class="text-xs xs:text-sm text-gray-900">
                              Showing 1 to 4 of 50 Entries
                          </span>
              <div class="inline-flex mt-2 xs:mt-0">
                <button
                                  class="text-sm text-indigo-50 transition duration-150 hover:bg-indigo-500 bg-indigo-600 font-semibold py-2 px-4 rounded-l">
                                  Prev
                              </button>
                &nbsp; &nbsp;
                <button
                                  class="text-sm text-indigo-50 transition duration-150 hover:bg-indigo-500 bg-indigo-600 font-semibold py-2 px-4 rounded-r">
                                  Next
                              </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
 
    <div class="bg-slate-200 rounded-xl shadow border py-4 px-5">
        <div class="  borderborder-gray-500" >
          <div class=" text-gray-800 p-6   px-6 py-3 text-xl font-bold  text-left ">
            <h4 class="text-center">Content Mail</h4>
          </div>
          <br>
          <div class="mx-4">
            <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
              Date:
            </span>
             mail.date 
          </div>
          <div class="mx-4 ">
            <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
              Subject:
            </span>
             mail.subject 
          </div>
          <div class="mx-4">
            <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
              From:
            </span>
             mail.from 
          </div>
          <div class="mx-4">
            <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
              To:
            </span>
             mail.to.join('; ') 
          </div>  <div class="mx-4">
            <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
              file:
            </span>
             mail.path 
          </div>
  
          <div class="mx-4 my-5">
            <p class="text-xs underline font-bold ">
              Message:
            </p>
             mail.content 
          </div>
        </div>
    </div>

  </div>

</template>


