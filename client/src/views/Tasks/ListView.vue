<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Мои задачи</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <ProfileInfo />
        <div class="mt-7 mx-3 md:mx-0">
          <div v-for="task in tasks" v-bind:key="task.id" :id="'task_' + task.id" class="bg-white text-sm mt-3 shadow rounded-xl hover:shadow-xl transition duration-300 flex flex-col">
            <div class="px-4 py-3 bg-blue-400 text-white rounded-t-xl" :class="{
              'bg-slate-400': task.status === 0 && !task.deleted_at && moment(task.delivery_date).utc(0).add(1, 'days') > moment().utc(0),
              'bg-blue-400': task.status === 1 && !task.deleted_at && moment(task.delivery_date).utc(0).add(1, 'days') > moment().utc(0),
              'bg-red-400': task.deleted_at,
              'bg-yellow-400': !task.deleted_at && moment(task.delivery_date).utc(0).add(1, 'days') <= moment().utc(0)
            }">
              Задача <span class="font-medium">#{{ task.id }}</span>
              <span v-if="!task.deleted_at" class="ml-1">&#x2022; Осталось: <strong class="font-medium">{{ task.amount }} шт.</strong></span>
              &#x2022; <span v-if="task.deleted_at" class="font-medium">удалено</span>
              <span v-else-if="moment(task.delivery_date).utc(0).add(1, 'days') <= moment().utc(0)" class="font-medium">просрочено</span>
              <span v-else-if="task.status === 0" class="font-medium">приостановлено</span>
              <span v-else-if="task.status === 1" class="font-medium">активно</span>
              <div v-if="!task.deleted_at" class="font-medium sm:float-right mt-4 text-center md:text-left md:mt-0">на сумму {{ $filters.currencyFormat(task.structed.overallPrice) }}</div>
            </div>
            <div class="text-base px-4 py-2 pb-4">
              <div class="text-center text-sm mt-1 mb-2 text-slate-500">
                Дата сдачи: {{ moment(task.delivery_date).utc(0).format('dddd, Do MMMM YYYY') }}
              </div>
              <h1 class="font-medium text-lg">
                {{ task.structed.main.title }}
              </h1>
              <p class="text-slate-600 text-sm mb-2">{{ task.description }}</p>
              <a :href="task.link" class="text-sm font-medium text-blue-500" target="_blank">{{ task.link }}</a>
              <div v-if="task.status != 0" class="text-sm mt-2 mb-1 border-t-2 pt-3 pb-1 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-2">
                <button @click="setTaskStatus(0, task.id)" class="bg-slate-200 text-slate-700 font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-slate-300 transition duration-300">Остановить задачу</button>
                <button v-if="showFull !== task.id" @click="showFull = task.id" class="bg-blue-100 text-blue-500 font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-blue-200 transition duration-300">
                  Полное описание
                </button>
                <button v-if="showFull === task.id" @click="showFull = 0" class="bg-blue-100 text-blue-500 font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-blue-200 transition duration-300">
                  Закрыть
                </button>
              </div>
              <div v-else-if="task.status == 0" class="text-sm mt-2 border-t-2 pt-3 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-2">
                <button v-if="!task.deleted_at" @click="setTaskStatus(1, task.id)" class="bg-blue-200 text-blue-700 font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-blue-300 transition duration-300">Запустить</button>
                <RouterLink v-if="!task.deleted_at" :to="{ name: 'edit-task', params: { id: task.id }}" class="bg-green-200 text-center py-3 text-green-700 font-medium py-1 px-2 rounded-lg shadow hover:shadow-md hover:bg-green-300 transition duration-300">Редактировать</RouterLink>
                <button v-if="!task.deleted_at" @click="deleteTask(task.id)" class="bg-red-200 text-red-700 font-medium py-1 px-2 rounded-lg shadow hover:shadow-md hover:bg-red-300 transition duration-300">
                  Удалить
                  <br>
                  <p class="text-xs -mt-1 text-red-500">сумма задачи вернется на баланс</p>
                </button>
                <button v-if="showFull !== task.id" @click="showFull = task.id" class="bg-blue-100 text-blue-500 font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-blue-200 transition duration-300">
                  Полное описание
                </button>
                <button v-if="showFull === task.id" @click="showFull = 0" class="bg-blue-100 text-blue-500 font-medium py-3 px-2 rounded-lg shadow hover:shadow-md hover:bg-blue-200 transition duration-300">
                  Закрыть
                </button>
              </div>
            </div>
            <transition name="slide-fade">
              <div v-if="showFull === task.id">
                <div class="py-2 px-4">
                  Создано: <span class="font-medium">{{ moment(task.created_at).utc(0).format('dddd, Do MMMM YYYY, в HH:mm') }}</span>
                </div>
                <div class="pb-2 px-4" v-if="task.deleted_at">
                  Удалено: <span class="font-medium">{{ moment(task.deleted_at).utc(0).format('dddd, Do MMMM YYYY, в HH:mm') }}</span>
                </div>
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-2 px-3 py-3">
                  <div class="text-sm py-2 px-4 rounded-full bg-slate-500 text-white shadow-lg">
                    Категория: <span class="font-medium">{{ task.structed.main.title }}</span> <span class="text-white">{{ $filters.currencyFormat(task.structed.main.price) }}</span>
                  </div>
                  <div v-for="option in task.structed.options" class="text-sm py-2 px-4 rounded-full bg-blue-500 text-white shadow-lg">
                    <span class="font-medium">{{ option.title }}</span> <span class="text-white">{{ $filters.currencyFormat(option.price) }}</span>
                  </div>
                </div>
              </div>
            </transition>
          </div>
        </div>
        <div v-if="tasks.length === 0" class="text-gray-500 mt-7 text-sm text-center">
          У Вас пока нет задач
        </div>
        <div v-if="tasks.length > 0" class="mt-7 text-sm grid justify-items-center">
          <div class="text-sm text-gray-500 dark:text-gray-400">
            Показано <span class="font-semibold text-gray-700 dark:text-white">{{ offset + 1 }}-{{ offsetEnd }}</span> из <span class="font-semibold text-gray-700 dark:text-white">{{ count }}</span> {{ $filters.declOfNum(count, ['задачи', 'задач', 'задач']) }}
          </div>
          <div>
            <nav class="mt-8">
              <ul class="inline-flex -space-x-px">
                <li v-for="page in pages">
                  <RouterLink v-if="page === 1" :to="{ 'name': 'tasks-list' }" @click="currentPage = page" class="rounded-l-xl" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
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
      showFull: 0,
      perPage: 10,
      offset: 0,
      pages: 0,
      count: 0,
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
    document.title = 'Мои задачи'

    if (this.$route.query.page) {
      this.currentPage = parseInt(this.$route.query.page)
    }

    this.getTasks()
  },
  methods: {
    getTaskOptionStructured(task) {
      let mainOption = {}
      let options = []
      let overallPrice = 0

      for(let i = 0; i < task.options.length; i++) {
        if (!task.options[i].parent_id) {
          mainOption = task.options[i]
        } else {
          options.push(task.options[i])
        }
        overallPrice += task.options[i].price
      }

      return {
        'main': mainOption,
        'overallPrice': overallPrice * task.amount,
        'options': options
      }
    },
    setTaskStatus(status, taskId) {
      NProgress.start()

      console.log(window.scrollY)
      window.scrollTo(0, 500)

      axios.put(import.meta.env.VITE_API_URL + 'tasks/' + taskId, {
        status: status
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getTasks(false)
        this.e.emit('updateUser', false)

        NProgress.done()
      })
    },
    deleteTask(taskId) {
      NProgress.start()

      axios.delete(import.meta.env.VITE_API_URL + 'tasks/' + taskId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getTasks(false)
        this.e.emit('updateUser', false)

        NProgress.done()
      })
    },
    getTasks(scroll = true) {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'tasks/user/' + this.user.id + '?page=' + this.currentPage + '&per_page=' + this.perPage, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
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
  },
  watch: {
    currentPage() {
      this.getTasks()
    }
  }
}
</script>