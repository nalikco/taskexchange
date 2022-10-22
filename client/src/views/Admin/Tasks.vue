<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Управление задачами</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <AdminMenu />
        <div class="overflow-x-auto relative shadow-md sm:rounded-lg">
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" class="py-3 px-6">
                ID
              </th>
              <th scope="col" class="py-3 px-6">
                Заказчик
              </th>
              <th scope="col" class="py-3 px-6">
                Статус
              </th>
              <th scope="col" class="py-3 px-6">
                Категория
              </th>
              <th scope="col" class="py-3 px-6">
                Количество
              </th>
              <th scope="col" class="py-3 px-6">
                Дата сдачи
              </th>
              <th scope="col" class="py-3 px-6">
                Ссылка
              </th>
              <th scope="col" class="py-3 px-6">
                Комментарий
              </th>
              <th scope="col" class="py-3 px-6">
                Создано
              </th>
              <th scope="col" class="py-3 px-6">
                Действия
              </th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="task in tasks" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
              <td class="py-4 px-6">
                {{ task.id }}
              </td>
              <td class="py-4 px-6 font-medium text-slate-700">
                {{ task.customer.username }}
              </td>
              <td v-if="task.deleted_at" class="py-4 px-6 font-medium text-red-500">
                Удалено
              </td>
              <td v-else-if="task.status === 0" class="py-4 px-6 font-medium text-slate-500">
                Приостановлено
              </td>
              <td v-else-if="task.status === 1" class="py-4 px-6 font-medium text-blue-500">
                Активно
              </td>
              <td class="py-4 px-6 font-medium">
                {{ getTaskOptionStructured(task).main.title }}
              </td>
              <td class="py-4 px-6 text-slate-700">
                <span class="font-medium">{{ task.amount }}</span> шт.<br>
                <span class="font-medium">{{ $filters.currencyFormat(getTaskOptionStructured(task).price) }}</span> за шт.
              </td>
              <td class="py-4 px-6 text-slate-700">
                {{ moment(task.delivery_date).utc(0).format('DD.MM.YYYY') }}
              </td>
              <td class="py-4 px-6 text-slate-700" :title="task.link">
                <a :href="task.link" class="text-blue-500 font-medium" target="_blank">{{ task.link.substring(0, 15) }}...</a>
              </td>
              <td class="py-4 px-6 text-slate-700" :title="task.description">
                {{ task.description.substring(0, 15) }}...
              </td>
              <td class="py-4 px-6 text-slate-700">
                {{ moment(task.created_at).utc(0).format('DD.MM.YYYY в HH:mm') }}
              </td>
              <td class="py-4 px-6">
                <button v-if="!task.deleted_at" @click="deleteTask(task.id)" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</button>
                <button v-else @click="deleteTask(task.id)" class="font-medium text-yellow-600 dark:text-yellow-500 hover:underline">Восстановить</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import axios from 'axios';
import NProgress from "nprogress";
import {useUserStore} from "@/stores/user";
import {mapState} from "pinia";

export default {
  data() {
    return {
      tasks: [],
      currentPage: 1,
      perPage: 10,
      offset: 0,
      pages: 0,
      count: 0,
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
    document.title = 'Управление задачами'

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
        'price': overallPrice,
        'options': options
      }
    },
    getTasks() {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'tasks/admin?page=' + this.currentPage + '&per_page=' + this.perPage, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.tasks = res.data.data

        }

        NProgress.done();
      })
    },
    deleteTask(taskId) {
      NProgress.start()

      axios.delete(import.meta.env.VITE_API_URL + 'tasks/' + taskId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getTasks()
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