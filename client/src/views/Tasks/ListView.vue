<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Биржа задач</h1>
        <p class="text-slate-500">
          Активных задач на бирже: <strong>{{ count }}</strong>
        </p>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <ProfileInfo v-if="user.type !== 0" />
        <div class="mt-7 mx-3 md:mx-0">
          <div v-for="task in tasks" v-bind:key="task.id" :id="'task_' + task.id" @click="toggleShowFull(task.id)" class="shadow-md cursor-pointer rounded-md mb-5">
            <div class="hidden md:grid grid-cols-12 bg-slate-500 text-white py-2 rounded-t-md">
              <div class="col-span-2 text-center">
                Название
              </div>
              <div class="col-start-10 col-span-2">
                Оплата
              </div>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-12 bg-white pt-8 rounded-b-md" :class="{
              'pb-8': showFull !== task.id
            }">
              <div class="col-span-2 text-center">
                <div class="bg-sky-800 text-white mx-5 rounded-md h-36 flex items-center justify-center">
                  <h2 class="text-4xl font-semibold">{{ task.structed.main.short }}</h2>
                </div>
              </div>
              <div class="col-span-7">
                <transition name="slide-fade">
                  <div v-if="showFull !== task.id" class="md:absolute text-center mt-5 md:mt-0">
                    <span class="font-semibold">Задача:</span> 0000{{ task.id }}
                  </div>
                </transition>
                <div class="md:mt-10">
                  <h1 class="text-4xl font-semibold text-center md:text-left">{{ task.structed.main.title }}</h1>
                </div>
                <div class="mt-5 flex flex-col md:flex-row px-5 md:px-0 gap-1 md:gap-4 text-sm">
                  <span v-for="option in task.structed.options" class="bg-gray-200 text-gray-700 py-1 px-7 rounded-md">
                    {{ option.title }}
                  </span>
                </div>
              </div>
              <div class="col-span-2">
                <div class="mt-7 md:mt-10">
                  <h1 class="text-4xl font-semibold text-center md:text-left">{{ $filters.currencyFormat(task.structed.main.price) }}</h1>
                </div>
              </div>
            </div>
            <transition name="slide-fade">
              <div v-if="showFull === task.id">
                <div class="grid pt-7 md:pt-0 px-5 md:px-0 md:grid-cols-12 bg-white">
                  <div class="col-start-3 col-span-10 text-sm md:pr-12 pb-8">
                    <div class="bg-gray-200 px-4 py-5 rounded-md">
                      {{ task.description }}
                    </div>
                    <div class="mt-4">
                      <RouterLink :to="{ name: 'profile', params: { user_id: task.customer.id }}">
                        <span class="font-semibold">Заказчик:</span> {{ task.customer.first_name }} {{ task.customer.last_name }}
                      </RouterLink>
                    </div>
                    <div class="mt-4">
                      <span class="font-semibold">Дата сдачи:</span> {{ moment(task.delivery_date).utc(0).format('dddd, DD MMMM YYYY') }}
                    </div>
                    <div class="mt-4">
                      <span class="font-semibold">Время для работы:</span> {{ workTime(task.delivery_date) }}
                    </div>
                    <div class="mt-4">
                      <span class="font-semibold">Предоставить в отчёте:</span> {{ task.report }}
                    </div>
                  </div>
                </div>
                <div v-if="task.customer_id === user.id" class="text-sm flex justify-center text-center">
                  <RouterLink :to="{ name: 'tasks-my' }" class="bg-gray-200 text-gray-700 w-full font-semibold py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-gray-300 transition duration-300">Ваша задача</RouterLink>
                </div>
                <div v-if="user.type === 1" class="text-sm">
                  <button v-if="!checkIfOfferSendForTask(task.id) && !checkIfExistsActiveOrderForTask(task.id)" @click="sendOffer(task.id)" class="bg-green-200 text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-green-300 transition duration-300">Отправить предложение</button>
                  <button v-else-if="checkIfOfferSendForTask(task.id)" disabled class="bg-slate-200 text-slate-700 w-full font-semibold py-3 px-2 rounded-lg shadow">Предложение отправлено</button>
                  <button v-else-if="checkIfExistsActiveOrderForTask(task.id)" disabled class="bg-slate-200 text-slate-700 w-full font-semibold py-3 px-2 rounded-lg shadow">В работе</button>
                </div>
              </div>
            </transition>
          </div>
        </div>
        <div v-if="tasks.length === 0" class="text-gray-500 mt-7 text-sm text-center">
          Пока нет задач
        </div>
        <div v-if="tasks.length > 0" class="mt-7 text-sm flex flex-row justify-between">
          <div>
            <nav>
              <ul class="inline-flex -space-x-px">
                <li v-for="page in pages">
                  <RouterLink v-if="page === 1" :to="{ 'name': 'tasks-my' }" @click="currentPage = page" class="rounded-l-xl" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                  <RouterLink v-else-if="page === pages" :to="'?page=' + page" @click="currentPage = page" class="rounded-r-xl" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                  <RouterLink v-else :to="'?page=' + page" @click="currentPage = page" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                </li>
              </ul>
            </nav>
          </div>
          <div class="text-sm text-gray-500 dark:text-gray-400">
            Показано <span class="font-semibold text-gray-700 dark:text-white">{{ offset + 1 }}-{{ offsetEnd }}</span> из <span class="font-semibold text-gray-700 dark:text-white">{{ count }}</span> {{ $filters.declOfNum(count, ['задачи', 'задач', 'задач']) }}
          </div>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import NProgress from "nprogress";
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";
import axios from "axios";
import {emitter} from "@/emitter";

export default {
  data() {
    return {
      tasks: [],
      currentPage: 1,
      perPage: 10,
      offset: 0,
      pages: 0,
      count: 0,
      showFull: 2,
      performerOffers: [],
      performerOrders: [],
      e: emitter
    }
  },
  computed: {
    offsetEnd() {
      let offsetEnd = this.offset + this.perPage
      if (offsetEnd > this.count) offsetEnd = this.count
      return offsetEnd
    },
    ...mapState(useUserStore, ['user', 'token'])
  },
  mounted() {
    document.title = 'Биржа задач'

    if (this.$route.query.page) {
      this.currentPage = parseInt(this.$route.query.page)
    }

    this.getTasks()
  },
  methods: {
    workTime(deliveryDate) {
      let currentTimestamp = moment().unix()
      let deliveryTimestamp = moment(deliveryDate).utc(0).unix()

      let minutes = Math.round((deliveryTimestamp - currentTimestamp) / 60)
      if (minutes < 60) return minutes + " " + this.$filters.declOfNum(minutes, ['минута', 'минуты', 'минут'])

      let hours = Math.round(minutes / 60)
      if (hours < 24) return hours + " " + this.$filters.declOfNum(hours, ['час', 'часа', 'часов'])

      let days = Math.round(hours / 24)

      return days + " " + this.$filters.declOfNum(days, ['день', 'дня', 'дней'])
    },
    toggleShowFull(taskId) {
      if(this.showFull === 0) {
        this.showFull = taskId
        return
      }

      this.showFull = 0
    },
    checkIfOfferSendForTask(taskId) {
      let send = false

      for (let i = 0; i < this.performerOffers.length; i++) {
        if (this.performerOffers[i].task_id === taskId) {
          send = true
        }
      }

      return send
    },
    checkIfExistsActiveOrderForTask(taskId) {
      let exists = false

      for (let i = 0; i < this.performerOrders.length; i++) {
        if (this.performerOrders[i].task_id === taskId) {
          exists = true
        }
      }

      return exists
    },
    getTaskOptionStructured(task) {
      let mainOption = {}
      let options = []
      let price = 0

      for(let i = 0; i < task.options.length; i++) {
        if (!task.options[i].parent_id) {
          mainOption = task.options[i]
        } else {
          options.push(task.options[i])
        }
        price += task.options[i].price
      }

      return {
        'main': mainOption,
        'price': price,
        'options': options
      }
    },
    sendOffer(taskId) {
      axios.post(import.meta.env.VITE_API_URL + 'offers/', {
        task_id: taskId
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getTasks(false)
      })
    },
    getTasks(scroll = true) {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'tasks/?page=' + this.currentPage + '&per_page=' + this.perPage, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          if(this.user.type === 1) this.getOffers()

          this.tasks = res.data.data
          for(let i = 0; i < this.tasks.length; i++) {
            this.tasks[i].structed = this.getTaskOptionStructured(this.tasks[i])
          }

          this.pages = res.data.pagination.pages
          this.count = res.data.pagination.count
          this.offset = res.data.pagination.offset

          if(scroll) window.scrollTo(0,0)
        }

        NProgress.done();
      })
    },
    getOffers() {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'offers/performer', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.performerOffers = res.data.data
        }

        this.getOrders()

        NProgress.done();
      })
    },
    getOrders() {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'orders/performer-active', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.performerOrders = res.data.data
        }

        NProgress.done();
      })
    }
  },
  watch: {
    currentPage() {
      this.getTasks()
    }
  }
}
</script>