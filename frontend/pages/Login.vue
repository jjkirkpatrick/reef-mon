<template>
  <div id="LoginPage">
    <div class="flex flex-wrap w-full">
      <div class="flex flex-col w-full md:w-1/2">
          <div class="flex justify-center pt-12 md:justify-start md:pl-12 md:-mb-24">
              <a href="#" class="p-4 text-xl font-bold text-white bg-black">
                  Reef-Mon.
              </a>
          </div>
          <div class="flex flex-col justify-center px-8 pt-8 my-auto md:justify-start md:pt-0 md:px-24 lg:px-32">
              <p class="text-3xl text-center">
                  Welcome.
              </p>

                          
              <button @click="discordLogin" type="button" class="py-2 px-4 flex justify-center items-center  bg-gray-600 hover:bg-gray-700 focus:ring-gray-500 focus:ring-offset-gray-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2  rounded-lg ">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-discord" viewBox="0 0 16 16">
                <path d="M13.545 2.907a13.227 13.227 0 0 0-3.257-1.011.05.05 0 0 0-.052.025c-.141.25-.297.577-.406.833a12.19 12.19 0 0 0-3.658 0 8.258 8.258 0 0 0-.412-.833.051.051 0 0 0-.052-.025c-1.125.194-2.22.534-3.257 1.011a.041.041 0 0 0-.021.018C.356 6.024-.213 9.047.066 12.032c.001.014.01.028.021.037a13.276 13.276 0 0 0 3.995 2.02.05.05 0 0 0 .056-.019c.308-.42.582-.863.818-1.329a.05.05 0 0 0-.01-.059.051.051 0 0 0-.018-.011 8.875 8.875 0 0 1-1.248-.595.05.05 0 0 1-.02-.066.051.051 0 0 1 .015-.019c.084-.063.168-.129.248-.195a.05.05 0 0 1 .051-.007c2.619 1.196 5.454 1.196 8.041 0a.052.052 0 0 1 .053.007c.08.066.164.132.248.195a.051.051 0 0 1-.004.085 8.254 8.254 0 0 1-1.249.594.05.05 0 0 0-.03.03.052.052 0 0 0 .003.041c.24.465.515.909.817 1.329a.05.05 0 0 0 .056.019 13.235 13.235 0 0 0 4.001-2.02.049.049 0 0 0 .021-.037c.334-3.451-.559-6.449-2.366-9.106a.034.034 0 0 0-.02-.019Zm-8.198 7.307c-.789 0-1.438-.724-1.438-1.612 0-.889.637-1.613 1.438-1.613.807 0 1.45.73 1.438 1.613 0 .888-.637 1.612-1.438 1.612Zm5.316 0c-.788 0-1.438-.724-1.438-1.612 0-.889.637-1.613 1.438-1.613.807 0 1.451.73 1.438 1.613 0 .888-.631 1.612-1.438 1.612Z"/>
              </svg>
                  Sign in with Discord
              </button>

                      <div class="pt-12 pb-12 text-center">
                          <p>
                              <a>
                              Login currently only supported via Discord
                              </a>
                          </p>
                      </div>
                  </div>
              </div>
              <div class="w-1/2 shadow-2xl">
                  <img class="hidden object-cover w-full h-screen md:block" src="https://www.atinorthamerica.com/wp-content/uploads/2020/04/how-to-clean-a-reef-tank.jpg"/>
              </div>
          </div>  
  </div>
</template>


<script>
export default {
    name: 'LoginPage',
    created () {
        if (this.$auth.loggedIn) {
            this.$router.push('logout')
        }
    },

    methods: {
    discordLogin () { 
        if (!this.$auth.loggedIn) {
            this.$auth.login('discord')}
        }
      
    },
}
</script>