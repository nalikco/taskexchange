<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
</script>
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
        <ProfileInfo />
        <div class="mx-3 mt-7 md:mx-0 grid grid-cols-1 lg:grid-cols-4">
          <div class="bg-white shadow rounded-lg text-center text-sm lg:mr-3 py-4 hover:shadow-lg transition duration-300">
            Здравствуйте, <span class="font-medium">{{ user.first_name }} {{ user.last_name }}</span>

            <div class="mt-4 flex flex-wrap">
              <RouterLink :to="{ name: 'tasks-my' }" v-if="user.type === 2" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Мои задачи</RouterLink>
              <RouterLink :to="{ name: 'orders-performer' }" v-if="user.type === 1" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Мои задачи</RouterLink>
              <RouterLink :to="{ name: 'tasks' }" v-if="user.type === 1" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Найти задачу</RouterLink>
              <RouterLink :to="{ name: 'create-task' }" v-if="user.type === 2" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Добавить задачу</RouterLink>
              <RouterLink :to="{ name: 'orders-customer' }" v-if="user.type === 2" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Мои заказы</RouterLink>
              <a href="#" @click="logout" class="text-gray-800 text-white mt-1 py-1 text-center w-full font-medium hover:underline">Выйти</a>

              <RouterLink :to="{ name: 'home' }" class="text-white mt-5 py-1 text-center w-full font-medium flex items-center justify-center hover:underline h-14 transition duration-300 mx-4 rounded" style="background: #242447;">
                <svg width="27" height="27" class="text-white" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path fill-rule="evenodd" clip-rule="evenodd" d="M18.6562 20.897L20.8733 18.6798L20.0925 15.843L20.4327 15.0305L23 13.5818V10.4464L20.44 8.99173L20.1055 8.18067L20.8961 5.34235L18.6774 3.12683L15.8403 3.90748L15.0296 3.56758L13.5808 1H10.4454L8.99072 3.56004L8.17985 3.89446L5.34198 3.10281L3.1267 5.31809L3.90748 8.15567L3.56758 8.96634L1 10.4151V13.5496L3.55774 15.0076L3.89252 15.8193L3.10197 18.6572L5.31809 20.8733L8.15567 20.0925L8.96644 20.4325L10.4153 22.999H13.5498L15.0067 20.4412L15.8183 20.1065L18.6562 20.897ZM18.8527 13.6256L17.9809 15.7078L18.6362 18.0886L18.0678 18.657L15.692 17.9951L13.609 18.8542L12.3873 20.999H11.5829L10.3714 18.8529L8.29155 17.9808L5.90947 18.6362L5.34203 18.0688L6.00385 15.693L5.14482 13.6101L3 12.3876V11.583L5.1471 10.3715L6.0192 8.29155L5.36375 5.90947L5.93001 5.34321L8.30576 6.00595L10.3895 5.14655L11.6093 3H12.4129L13.6245 5.1471L15.7044 6.0192L18.087 5.36362L18.6558 5.93166L17.9941 8.30696L18.8534 10.3906L21 11.6103V12.4139L18.8527 13.6256ZM12 16C9.79086 16 8 14.2091 8 12C8 9.79086 9.79086 8 12 8C14.2091 8 16 9.79086 16 12C16 14.2091 14.2091 16 12 16ZM14 12C14 13.1046 13.1046 14 12 14C10.8954 14 10 13.1046 10 12C10 10.8954 10.8954 10 12 10C13.1046 10 14 10.8954 14 12Z" fill="#ffffff"/>
                </svg>
                <div class="ml-3">
                  Написать в тех. поддержку
                </div>
              </RouterLink>
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
              <div class="bg-white shadow hover:shadow-lg flex flex-row items-center pl-5 text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="bg-yellow-500 rounded-md p-3">
                  <svg width="30" height="30" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M17 17V19.382L12 16.882L7 19.382V4H17V5H19V4C19 2.89543 18.1046 2 17 2H7C5.89543 2 5 2.89543 5 4V22.618L12 19.118L19 22.618V17H17ZM18 8V10H16V12H18V14H20V12H22V10H20V8H18Z" fill="#ffffff"/>
                  </svg>
                </div>
                <div class="py-3 ml-5">
                  Приглашайте друзей и получайте
                  <br>
                  <span class="text-gray-500 font-medium">25% комиссии</span>
                </div>
              </div>
              <div class="bg-white shadow border-2 border-red-200 hover:shadow-lg flex flex-row items-center pl-5 text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="bg-red-600 rounded-md p-3">
                  <svg width="30" height="30" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M11.1758 11.53C11.7222 10.763 11.6613 9.9553 11.2824 8.50235C10.533 5.62882 10.8314 4.05383 13.4147 2.18916L14.6756 1.27905L14.9806 2.80389C15.3047 4.42441 15.8369 5.42671 17.2036 7.35464C17.2579 7.43122 17.2579 7.43122 17.3124 7.50815C19.281 10.2854 20 11.9596 20 15C20 18.6884 16.2714 22 12 22C7.72841 22 4 18.6888 4 15C4 14.9311 4.00007 14.9331 3.98839 14.6285C3.89803 12.2718 4.33381 10.4274 6.09706 8.43586C6.46961 8.01509 6.89308 7.61068 7.36962 7.22371L8.42162 6.36946L8.92766 7.62658C9.30158 8.55547 9.73969 9.28567 10.2346 9.82151C10.6538 10.2754 10.9647 10.8461 11.1758 11.53ZM7.59448 9.76166C6.23712 11.2947 5.91441 12.6606 5.98692 14.5518C6.00042 14.9039 6 14.8915 6 15C6 17.5279 8.7836 20 12 20C15.2161 20 18 17.5275 18 15C18 12.4582 17.4317 11.135 15.6807 8.6647C15.6265 8.58818 15.6265 8.58818 15.5719 8.51125C14.5085 7.01111 13.8747 5.96759 13.4553 4.80052C12.7705 5.62118 12.8107 6.43739 13.2176 7.99766C13.967 10.8712 13.6686 12.4462 11.0853 14.3108L9.61227 15.3741L9.50185 13.5608C9.43129 12.4023 9.16906 11.6156 8.76539 11.1785C8.36656 10.7467 8.00647 10.2411 7.68355 9.66279C7.65343 9.69566 7.62374 9.72862 7.59448 9.76166Z" fill="#ffffff"/>
                  </svg>
                </div>
                <div class="ml-5 py-3">
                  Получайте особый
                  <br>
                  <span class="text-gray-500 font-medium">PRO-аккаунт</span>
                </div>
              </div>
              <div v-if="user.type === 1" class="bg-white shadow flex flex-row items-center pl-5 hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="bg-indigo-500 rounded-md p-3">
                  <svg width="30" height="30" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M3.91937 15H9V22.8507L20.0806 8.99995H15V1.14917L3.91937 15ZM11 13H8.08062L13 6.85073V11H15.9194L11 17.1492V13Z" fill="#ffffff"/>
                  </svg>
                </div>
                <div class="ml-5 py-3">
                  Активных задач
                  <br>
                  <span class="text-gray-500 font-medium">{{ performerActiveOrdersCount }} в работе</span>
                </div>
              </div>
              <div v-if="user.type === 2" class="bg-white shadow flex flex-row items-center pl-5 hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="bg-indigo-500 rounded-md p-3">
                  <svg width="30" height="30" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M3.91937 15H9V22.8507L20.0806 8.99995H15V1.14917L3.91937 15ZM11 13H8.08062L13 6.85073V11H15.9194L11 17.1492V13Z" fill="#ffffff"/>
                  </svg>
                </div>
                <div class="ml-5 py-3">
                  Активных заказов
                  <br>
                  <span class="text-gray-500 font-medium">{{ customerActiveOrdersCount }} в работе</span>
                </div>
              </div>
              <RouterLink v-if="user.type === 1" :to="{name:'messages'}" class="bg-white shadow flex flex-row items-center pl-5 hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="bg-green-500 rounded-md p-3">
                  <svg width="30" height="30" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M6 21.8042L12.0868 18H20C21.1046 18 22 17.1046 22 16V4C22 2.89543 21.1046 2 20 2H4C2.89543 2 2 2.89543 2 4V16C2 17.1046 2.89543 18 4 18H6V21.8042ZM11.5132 16L8 18.1958V16H4V4H20V16H11.5132ZM7 13V11H14V13H7ZM7 7V9H16V7H7Z" fill="#ffffff"/>
                  </svg>
                </div>
                <div class="ml-5 py-3">
                  Диалоги с заказчиками
                  <br>
                  <span class="text-gray-500 font-medium">{{ unViewedMessagesCount }} {{ $filters.declOfNum(unViewedMessagesCount, ['новое сообщение', 'новых сообщения', 'новых сообщений'])}}</span>
                </div>
              </RouterLink>
              <RouterLink v-if="user.type === 2" :to="{name:'messages'}" class="bg-white shadow flex flex-row pl-5 items-center hover:shadow-lg text-sm py-2 px-2 rounded-lg text-gray-800 transition duration-300">
                <div class="bg-green-500 rounded-md p-3">
                  <svg width="30" height="30" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M6 21.8042L12.0868 18H20C21.1046 18 22 17.1046 22 16V4C22 2.89543 21.1046 2 20 2H4C2.89543 2 2 2.89543 2 4V16C2 17.1046 2.89543 18 4 18H6V21.8042ZM11.5132 16L8 18.1958V16H4V4H20V16H11.5132ZM7 13V11H14V13H7ZM7 7V9H16V7H7Z" fill="#ffffff"/>
                  </svg>
                </div>
                <div class="ml-5 py-3">
                  Диалоги с исполнителями
                  <br>
                  <span class="text-gray-500 font-medium">{{ unViewedMessagesCount }} {{ $filters.declOfNum(unViewedMessagesCount, ['новое сообщение', 'новых сообщения', 'новых сообщений'])}}</span>
                </div>
              </RouterLink>
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
import NProgress from "nprogress";
import axios from "axios"

export default {
  data() {
    return {
      unViewedMessagesCount: 0,
      customerActiveOrdersCount: 0,
      performerActiveOrdersCount: 0
    }
  },
  computed: {
    ...mapState(useUserStore, ['user', 'token']),
  },
  mounted() {
    if(this.user.type === 1) {
      document.title = 'Кабинет исполнителя'
      this.getActivePerformerOrders()
    }
    else if(this.user.type === 2) {
      document.title = 'Кабинет заказчика'
      this.getActiveCustomerOrders()
    }

    NProgress.done()
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
    },
    getUnViewedMessagesCount() {
      axios.get(import.meta.env.VITE_API_URL + 'messages/count-un-viewed', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if (res.data.data) {
          this.unViewedMessagesCount = res.data.data
        }
      })
    },
    getActiveCustomerOrders() {
      axios.get(import.meta.env.VITE_API_URL + 'orders/customer-active', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if (res.data.data) {
          this.customerActiveOrdersCount = res.data.data.length
        }
      })
    },
    getActivePerformerOrders() {
      axios.get(import.meta.env.VITE_API_URL + 'orders/performer-active', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if (res.data.data) {
          this.performerActiveOrdersCount = res.data.data.length
        }
      })
    }
  }
}
</script>