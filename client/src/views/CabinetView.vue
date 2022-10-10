<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 v-if="user.type === 1" class="text-3xl font-bold tracking-tight text-gray-900">Кабинет исполнителя</h1>
        <h1 v-if="user.type === 2" class="text-3xl font-bold tracking-tight text-gray-900">Кабинет заказчика</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <div class="mx-3 md:mx-0 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
            <div class="float-left ml-2 mt-1">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-gray-800">
                <path fill-rule="evenodd" d="M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" clip-rule="evenodd" />
              </svg>
            </div>
            <div class="ml-20 py-3">
              {{ user.username }}
              <br>
              <span class="text-gray-500 font-medium">id:{{ user.id }}</span>
            </div>
          </div>
          <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
            <div class="float-left ml-2 mt-1">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-lime-500">
                <path d="M12 7.5a2.25 2.25 0 100 4.5 2.25 2.25 0 000-4.5z" />
                <path fill-rule="evenodd" d="M1.5 4.875C1.5 3.839 2.34 3 3.375 3h17.25c1.035 0 1.875.84 1.875 1.875v9.75c0 1.036-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 011.5 14.625v-9.75zM8.25 9.75a3.75 3.75 0 117.5 0 3.75 3.75 0 01-7.5 0zM18.75 9a.75.75 0 00-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 00.75-.75V9.75a.75.75 0 00-.75-.75h-.008zM4.5 9.75A.75.75 0 015.25 9h.008a.75.75 0 01.75.75v.008a.75.75 0 01-.75.75H5.25a.75.75 0 01-.75-.75V9.75z" clip-rule="evenodd" />
                <path d="M2.25 18a.75.75 0 000 1.5c5.4 0 10.63.722 15.6 2.075 1.19.324 2.4-.558 2.4-1.82V18.75a.75.75 0 00-.75-.75H2.25z" />
              </svg>
            </div>
            <div class="ml-20 py-3">
              Баланс
              <br>
              <span class="text-gray-500 font-medium">{{ $filters.currencyFormat(user.balance) }}</span>
            </div>
          </div>
          <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
            <div class="float-left ml-2 mt-1">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-blue-500">
                <path fill-rule="evenodd" d="M9.315 7.584C12.195 3.883 16.695 1.5 21.75 1.5a.75.75 0 01.75.75c0 5.056-2.383 9.555-6.084 12.436A6.75 6.75 0 019.75 22.5a.75.75 0 01-.75-.75v-4.131A15.838 15.838 0 016.382 15H2.25a.75.75 0 01-.75-.75 6.75 6.75 0 017.815-6.666zM15 6.75a2.25 2.25 0 100 4.5 2.25 2.25 0 000-4.5z" clip-rule="evenodd" />
                <path d="M5.26 17.242a.75.75 0 10-.897-1.203 5.243 5.243 0 00-2.05 5.022.75.75 0 00.625.627 5.243 5.243 0 005.022-2.051.75.75 0 10-1.202-.897 3.744 3.744 0 01-3.008 1.51c0-1.23.592-2.323 1.51-3.008z" />
              </svg>
            </div>
            <div class="ml-20 py-3">
              Задачи
              <br>
              <span class="text-gray-500 font-medium">0 {{ $filters.declOfNum(0, ['активная', 'активных', 'активных']) }}</span>
            </div>
          </div>
          <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
            <div class="float-left ml-2 mt-1">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-red-500">
                <path fill-rule="evenodd" d="M12.963 2.286a.75.75 0 00-1.071-.136 9.742 9.742 0 00-3.539 6.177A7.547 7.547 0 016.648 6.61a.75.75 0 00-1.152-.082A9 9 0 1015.68 4.534a7.46 7.46 0 01-2.717-2.248zM15.75 14.25a3.75 3.75 0 11-7.313-1.172c.628.465 1.35.81 2.133 1a5.99 5.99 0 011.925-3.545 3.75 3.75 0 013.255 3.717z" clip-rule="evenodd" />
              </svg>

            </div>
            <div class="ml-20 py-3">
              Активность на сайте
              <br>
              <span class="text-gray-500 font-medium">{{ activity }} {{ $filters.declOfNum(activity, ['день', 'дня', 'дней']) }}</span>
            </div>
          </div>
        </div>
        <div class="mx-3 mt-7 md:mx-0 grid grid-cols-1 lg:grid-cols-4">
          <div class="bg-white shadow rounded-lg text-center text-sm lg:mr-3 py-4 hover:shadow-lg transition duration-300">
            Здравствуйте, <span class="font-medium">{{ user.username }}</span>

            <div class="mt-4 flex flex-wrap">
              <RouterLink :to="{ name: 'home' }" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Мои задачи</RouterLink>
              <RouterLink :to="{ name: 'home' }" v-if="user.type === 1" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Найти задачу</RouterLink>
              <RouterLink :to="{ name: 'home' }" v-if="user.type === 2" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Добавить задачу</RouterLink>
              <RouterLink :to="{ name: 'home' }" v-if="user.type === 2" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Мои заказы</RouterLink>
              <a href="#" @click="logout" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Выйти</a>

              <RouterLink :to="{ name: 'home' }" class="bg-yellow-500 text-white mt-5 py-1 text-center w-full font-medium hover:underline hover:bg-yellow-700 transition duration-300">Написать в тех. поддержку</RouterLink><br>
            </div>
          </div>
          <div class="col-span-3">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="md:col-span-2 bg-white shadow rounded-lg text-sm mt-3 lg:mt-0 py-3 px-3 hover:shadow-lg transition duration-300">
                <div v-if="user.type === 1">
                  У Вас <span class="font-medium">{{ user.points }} {{ $filters.declOfNum(user.points, ['балл', 'балла', 'баллов'])}}</span>, Ваш уровень исполнителя — <span class="font-medium">Новичок</span>
                </div>
                <div v-if="user.type === 2">
                  У Вас <span class="font-medium">{{ user.points }} {{ $filters.declOfNum(user.points, ['балл', 'балла', 'баллов'])}}</span>, Ваш уровень заказчика — <span class="font-medium">Новичок</span>
                </div>
                <div class="grid grid-cols-6 mt-2">
                  <div class="h-16 bg-gray-300 rounded-l-lg border-x border-gray-400"></div>
                  <div class="h-16 bg-gray-300 border-r border-gray-400"></div>
                  <div class="h-16 bg-gray-300 border-r border-gray-400"></div>
                  <div class="h-16 bg-gray-300 border-r border-gray-400"></div>
                  <div class="h-16 bg-gray-300 border-r border-gray-400"></div>
                  <div class="h-16 bg-gray-300 rounded-r-lg border-r border-gray-400"></div>
                </div>
                <div class="mt-3">
                  <RouterLink :to="{ name: 'home' }" class="underline">Как это работает?</RouterLink>
                </div>
              </div>
              <div class="bg-white shadow hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="float-left ml-2 mt-1">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14 text-gray-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6.633 10.5c.806 0 1.533-.446 2.031-1.08a9.041 9.041 0 012.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 00.322-1.672V3a.75.75 0 01.75-.75A2.25 2.25 0 0116.5 4.5c0 1.152-.26 2.243-.723 3.218-.266.558.107 1.282.725 1.282h3.126c1.026 0 1.945.694 2.054 1.715.045.422.068.85.068 1.285a11.95 11.95 0 01-2.649 7.521c-.388.482-.987.729-1.605.729H13.48c-.483 0-.964-.078-1.423-.23l-3.114-1.04a4.501 4.501 0 00-1.423-.23H5.904M14.25 9h2.25M5.904 18.75c.083.205.173.405.27.602.197.4-.078.898-.523.898h-.908c-.889 0-1.713-.518-1.972-1.368a12 12 0 01-.521-3.507c0-1.553.295-3.036.831-4.398C3.387 10.203 4.167 9.75 5 9.75h1.053c.472 0 .745.556.5.96a8.958 8.958 0 00-1.302 4.665c0 1.194.232 2.333.654 3.375z" />
                  </svg>
                </div>
                <div class="ml-20 py-3">
                  Приглашайте друзей и получайте
                  <br>
                  <span class="text-gray-500 font-medium">25% комиссии</span>
                </div>
              </div>
              <div class="bg-white shadow border-2 border-red-200 hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="float-left ml-2 mt-1">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14 text-red-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.362 5.214A8.252 8.252 0 0112 21 8.25 8.25 0 016.038 7.048 8.287 8.287 0 009 9.6a8.983 8.983 0 013.361-6.867 8.21 8.21 0 003 2.48z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 18a3.75 3.75 0 00.495-7.467 5.99 5.99 0 00-1.925 3.546 5.974 5.974 0 01-2.133-1A3.75 3.75 0 0012 18z" />
                  </svg>

                </div>
                <div class="ml-20 py-3">
                  Получайте особый
                  <br>
                  <span class="text-gray-500 font-medium">PRO-аккаунт</span>
                </div>
              </div>
              <div v-if="user.type === 1" class="bg-white shadow hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="float-left ml-2 mt-1">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14 text-yellow-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 13.5l10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75z" />
                  </svg>
                </div>
                <div class="ml-20 py-3">
                  Активных задач
                  <br>
                  <span class="text-gray-500 font-medium">0 в работе</span>
                </div>
              </div>
              <div v-if="user.type === 2" class="bg-white shadow hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="float-left ml-2 mt-1">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14 text-yellow-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 13.5l10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75z" />
                  </svg>
                </div>
                <div class="ml-20 py-3">
                  Активных заказов
                  <br>
                  <span class="text-gray-500 font-medium">0 в работе</span>
                </div>
              </div>
              <div v-if="user.type === 1" class="bg-white shadow hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="float-left ml-2 mt-1">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14 text-blue-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M20.25 8.511c.884.284 1.5 1.128 1.5 2.097v4.286c0 1.136-.847 2.1-1.98 2.193-.34.027-.68.052-1.02.072v3.091l-3-3c-1.354 0-2.694-.055-4.02-.163a2.115 2.115 0 01-.825-.242m9.345-8.334a2.126 2.126 0 00-.476-.095 48.64 48.64 0 00-8.048 0c-1.131.094-1.976 1.057-1.976 2.192v4.286c0 .837.46 1.58 1.155 1.951m9.345-8.334V6.637c0-1.621-1.152-3.026-2.76-3.235A48.455 48.455 0 0011.25 3c-2.115 0-4.198.137-6.24.402-1.608.209-2.76 1.614-2.76 3.235v6.226c0 1.621 1.152 3.026 2.76 3.235.577.075 1.157.14 1.74.194V21l4.155-4.155" />
                  </svg>
                </div>
                <div class="ml-20 py-3">
                  Диалоги с заказчиками
                  <br>
                  <span class="text-gray-500 font-medium">15 новых сообщений</span>
                </div>
              </div>
              <div v-if="user.type === 2" class="bg-white shadow hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="float-left ml-2 mt-1">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14 text-blue-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M20.25 8.511c.884.284 1.5 1.128 1.5 2.097v4.286c0 1.136-.847 2.1-1.98 2.193-.34.027-.68.052-1.02.072v3.091l-3-3c-1.354 0-2.694-.055-4.02-.163a2.115 2.115 0 01-.825-.242m9.345-8.334a2.126 2.126 0 00-.476-.095 48.64 48.64 0 00-8.048 0c-1.131.094-1.976 1.057-1.976 2.192v4.286c0 .837.46 1.58 1.155 1.951m9.345-8.334V6.637c0-1.621-1.152-3.026-2.76-3.235A48.455 48.455 0 0011.25 3c-2.115 0-4.198.137-6.24.402-1.608.209-2.76 1.614-2.76 3.235v6.226c0 1.621 1.152 3.026 2.76 3.235.577.075 1.157.14 1.74.194V21l4.155-4.155" />
                  </svg>
                </div>
                <div class="ml-20 py-3">
                  Диалоги с исполнителями
                  <br>
                  <span class="text-gray-500 font-medium">15 новых сообщений</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import {mapActions, mapState} from "pinia";
import {useUserStore} from "@/stores/user";

export default {
  computed: {
    ...mapState(useUserStore, ['user']),
    activity() {
      let secondsFromUserCreating = (new Date().getTime() - new Date(new Date(this.user.created_at).toDateString()).getTime()) / 1000

      return Math.round(secondsFromUserCreating / 60 / 60 / 24)
    }
  },
  mounted() {
    if(this.user.type === 1) document.title = 'Кабинет исполнителя'
    else if(this.user.type === 2) document.title = 'Кабинет заказчика'
  },
  methods: {
    ...mapActions(useUserStore, ['setUser', 'setToken']),
    logout(e) {
      e.preventDefault()

      this.setUser({type: 0})
      this.setToken('')

      this.showProfileMenu = false
      this.showMobileMenu = false
      this.showNotifications = false

      this.$router.push({ name: 'sign-in' })
    }
  }
}
</script>