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
          <div v-for="task in tasks" v-bind:key="task.id" :id="'task_' + task.id" class="bg-white text-sm mt-3 shadow rounded-xl hover:shadow-xl transition duration-300 flex flex-col">
            <div class="px-4 py-3 bg-slate-500 text-white rounded-t-xl">
              Задача <span class="font-medium">#{{ task.id }}</span>
              &#x2022; Количество: <span class="font-medium">{{ task.amount }} шт.</span>
              <div v-if="!task.deleted_at" class="font-medium sm:float-right mt-4 text-center md:text-left md:mt-0">{{ $filters.currencyFormat(task.structed.price) }} за задачу</div>
            </div>
            <div class="text-base px-4 py-2 pb-4">
              <div class="text-center text-sm mt-1 mb-2 text-slate-500">
                Дата сдачи: {{ moment(task.delivery_date).utc(0).format('dddd, Do MMMM YYYY') }}
              </div>
              <h1 class="font-medium text-lg">
                {{ task.structed.main.title }}, {{ $filters.currencyFormat(task.structed.main.price) }}
              </h1>
              <p class="text-slate-600 text-sm mb-2">{{ task.description }}</p>
              <a :href="task.link" class="text-sm font-medium text-blue-500" target="_blank">{{ task.link }}</a>
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-2 py-3">
                <div v-for="option in task.structed.options" class="text-sm py-2 px-4 rounded-full bg-blue-500 text-white shadow-lg">
                  <span class="font-medium">+ {{ option.title }}</span> <span class="text-white">{{ $filters.currencyFormat(option.price) }}</span>
                </div>
              </div>
              <div v-if="user.type === 1" class="text-sm mt-2 mb-1 border-t-2 pt-3">
                <button v-if="!checkIfOfferSendForTask(task.id) && !checkIfExistsActiveOrderForTask(task.id)" @click="sendOffer(task.id)" class="bg-green-200 text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-green-300 transition duration-300">Отправить предложение</button>
                <button v-else-if="checkIfOfferSendForTask(task.id)" disabled class="bg-slate-200 text-slate-700 w-full font-medium py-3 px-2 rounded-lg shadow">Предложение отправлено</button>
                <button v-else-if="checkIfExistsActiveOrderForTask(task.id)" disabled class="bg-slate-200 text-slate-700 w-full font-medium py-3 px-2 rounded-lg shadow">В работе</button>
              </div>
            </div>
          </div>
        </div>
        <div v-if="tasks.length === 0" class="text-gray-500 mt-7 text-sm text-center">
          Пока нет задач
        </div>
        <div v-if="tasks.length > 0" class="mt-7 text-sm grid justify-items-center">
          <div class="text-sm text-gray-500 dark:text-gray-400">
            Показано <span class="font-semibold text-gray-700 dark:text-white">{{ offset + 1 }}-{{ offsetEnd }}</span> из <span class="font-semibold text-gray-700 dark:text-white">{{ count }}</span> {{ $filters.declOfNum(count, ['задачи', 'задач', 'задач']) }}
          </div>
          <div>
            <nav class="mt-8">
              <ul class="inline-flex -space-x-px">
                <li v-for="page in pages">
                  <RouterLink v-if="page === 1" :to="{ 'name': 'tasks-my' }" @click="currentPage = page" class="rounded-l-xl" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                  <RouterLink v-else-if="page === pages" :to="'?page=' + page" @click="currentPage = page" class="rounded-r-xl" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                  <RouterLink v-else :to="'?page=' + page" @click="currentPage = page" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                </li>
              </ul>
            </nav>
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

      axios.get(import.meta.env.VITE_API_URL + 'tasks/?page=' + this.currentPage + '&per_page=' + this.perPage).then(res => {
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