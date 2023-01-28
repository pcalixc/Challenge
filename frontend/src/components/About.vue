v<!-- eslint-disable vue/multi-word-component-names -->

<template>

    <!-- component -->

<div class='flex flex-col items-center justify-center min-h-screen bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-blue-700 via-blue-800 to-gray-900'>
    <router-link to="/" href="" class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800">Back</router-link>

    <div class="p-2 mx-24 ">
        <div class="bg-gray-900 text-white py-6 rounded-lg shadow-xl ">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-1 xl:gap-12 px-2 xl:px-12 m-4">
                <div class="flex space-x-8 mt-8">
                    <div>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path d="M12 14l9-5-9-5-9 5 9 5z"></path>
                            <path d="M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z"></path>
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14l9-5-9-5-9 5 9 5zm0 0l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14zm-4 6v-7.5l4-2.222"></path>
                        </svg>
                    </div>
                    <div>
                        <h4 class="text-xl font-bold font mb-4">Technologies</h4>
                    
                            <ul>
                            <li class="font-semibold" >Backend: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="https://go.dev">[Go](https://go.dev/)</a> </li>
                            <li class="font-semibold">API Router: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="https://github.com/go-chi/chi">[Chi](https://github.com/go-chi/chi)</a> </li>
                            <li class="font-semibold">Search Engine: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="https://zincsearch.com/">[ZincSearch](https://zincsearch.com/)</a> </li>
                            <li class="font-semibold">Frontend: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="https://vuejs.org/"> [Vue](https://vuejs.org/)</a> + 
                                <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="https://developer.mozilla.org/es/docs/Web/JavaScript"> [JavaScript](https://developer.mozilla.org/es/docs/Web/JavaScript)</a> </li>
                            </ul>
                    
                    </div>
                </div>
                
                <div class="flex space-x-8 mt-8">
                    <div>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                        </svg>
                    </div>
                    <div>
                        <h4 class="text-xl font-bold mb-4">Database</h4>
                        <ul>
                            <li class="font-semibold" >Enron Corp's email database: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz">http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz</a> </li>
                        </ul>
                    </div>
                </div>

                <div class="flex space-x-8 mt-8">
                    <div>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z"></path>
                            </svg>
                    </div>
                    <div>
                        <h4 class="text-xl font-bold mb-4">Repository</h4>
                        <ul>
                            <li class="font-semibold" >Challenge: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href=" https://github.com/pcalixc/Challenge.git"> https://github.com/pcalixc/Challenge.git</a> </li>
                        </ul>
                       
                        
                    </div>
                </div>

            
            </div>
        </div>
    </div>

</div>
</template>