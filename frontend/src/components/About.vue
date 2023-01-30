<!-- eslint-disable vue/multi-word-component-names -->

<template>
    <div class='flex flex-col items-center justify-center min-h-screen bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-blue-700 via-blue-800 to-gray-900'>
        <router-link to="/" href="" class=" hover:bg-blue-100 bg-white rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 focus:outline-none">
            <img class="w-8 h-8" src="https://cdn-icons-png.flaticon.com/512/9283/9283458.png" alt="">
        </router-link>
        <div class="p-2 mx-24 ">
            <div class="bg-gray-900 text-white py-6 rounded-lg shadow-xl ">
                <div class="flex flex-col gap-1 xl:gap-12 px-2 xl:px-12 m-1 ">
                    <div class="flex space-x-8 mt-8">
                        <div>
                            <img class="w-12 h-12" src="https://cdn-icons-png.flaticon.com/512/8298/8298031.png" alt="">
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
                            <div>
                                <img class="w-10 h-10" src="https://cdn-icons-png.flaticon.com/512/8298/8298031.png" alt="">
                            </div>
                        </div>
                        <div>
                            <h4 class="text-xl font-bold mb-4">Database</h4>
                            <ul>
                                <li class="font-semibold" >Enron Corp's email database: <a class="text-blue-200 hover:text-blue-500 hover:underline font-light" href="http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz">http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz</a> </li>
                            </ul>
                        </div>
                    </div>
                    <div class="flex space-x-8 mt-8 m-b-8">
                        <div>
                            <div>
                                <img class="w-10 h-10" src="https://cdn-icons-png.flaticon.com/512/8298/8298031.png" alt="">
                            </div>
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