<script setup>
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Финансы</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <div class="flex flex-col md:flex-row bg-white shadow rounded-md py-7 px-6">
          <div class="basis-1/3 text-center flex flex-row justify-center items-center">
            <h2 class="text-2xl font-semibold">{{ $filters.currencyFormat(user.balance) }}</h2>
            <div class="ml-5">
              <p class="text-sm text-gray-900 font-semibold text-left">Баланс аккаунта</p>
              <p class="text-sm text-gray-500 text-left">
                Последнее списание:
                <span v-if="lastWriteOff">{{ $filters.currencyFormat(lastWriteOff.sum) }}</span>
                <span v-else>{{ $filters.currencyFormat(0) }}</span>
              </p>
            </div>
          </div>
          <div v-if="user.type === 2" class="basis-1/3 text-center flex flex-row justify-center items-center">
            <h2 class="text-2xl font-semibold">{{ $filters.currencyFormat(activeTasksSum) }}</h2>
            <div class="ml-0 md:ml-5">
              <p class="text-sm text-gray-900 font-semibold text-left">Заморожено</p>
              <p class="text-sm text-gray-500 text-left">{{ activeTasks.length }} {{ $filters.declOfNum(activeTasks.length, ['задание', 'задания', 'заданий']) }} в работе</p>
            </div>
          </div>
          <div v-else class="basis-1/3">

          </div>
          <div class="basis-1/3 text-center flex flex-row gap-2 justify-center items-center">
            <button class="font-semibold bg-green-500 text-white px-7 py-3 rounded-lg shadow hover:bg-green-600">Пополнить</button>
            <button class="font-semibold bg-indigo-500 text-white px-7 py-3 rounded-lg shadow hover:bg-indigo-600">Вывести</button>
          </div>
        </div>
        <header class="my-10">
          <h1 class="text-3xl font-bold tracking-tight text-gray-900">История операций</h1>
        </header>
        <div v-if="payments.length > 0" class="overflow-x-auto relative">
          <table class="w-full text-left dark:text-gray-400">
            <tbody>
              <tr v-for="payment in payments" class="bg-white border-b" :class="{
                'bg-red-50': payment.type === 1,
                'bg-green-50': payment.type === 2,
              }">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  <span class="font-semibold">{{ moment(payment.created_at).utc(0).format('DD.MM.YYYY') }}</span>
                  <br>
                  в {{ moment(payment.created_at).utc(0).format('HH:mm') }}
                </th>
                <td class="py-4 px-6">
                  <span v-if="payment.type === 1" class="font-semibold">Списание со счёта</span>
                  <span v-else class="font-semibold">Пополнение счёта</span>
                  <br>{{ payment.comment }}
                </td>
                <td class="py-4 px-6">
                  <span class="font-semibold">Сумма</span><br>
                  <span v-if="payment.type === 1">-{{ $filters.currencyFormat(payment.sum) }}</span>
                  <span v-else>+{{ $filters.currencyFormat(payment.sum) }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div v-else class="text-gray-500 text-center">
          У Вас пока не было никаких операций.
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";
import axios from "axios"
import NProgress from "nprogress";

export default {
  data() {
    return {
      payments: [],
      allTasks: [],
    }
  },
  computed: {
    activeTasks() {
      let activeTasks = []

      for(let i = 0; i < this.allTasks.length; i++) {
        if (this.allTasks[i].status === 0 && !this.allTasks[i].deleted_at || this.allTasks[i].status === 1 && !this.allTasks[i].deleted_at) {
          activeTasks.push(this.allTasks[i])
        }
      }

      return activeTasks
    },
    lastWriteOff() {
      for (let i = 0; i < this.payments.length; i++) {
        if (this.payments[i].type === 1) {
          return this.payments[i]
        }
      }

      return null
    },
    activeTasksSum() {
      let sum = 0

      for(let t = 0; t < this.activeTasks.length; t++) {
        let taskPrice = 0
        for(let o = 0; o < this.activeTasks[t].options.length; o++) {
          taskPrice += this.activeTasks[t].options[o].price
        }

        sum += taskPrice * this.activeTasks[t].amount
      }

      return sum
    },
    ...mapState(useUserStore, ['user', 'token']),
  },
  mounted() {
    document.title = 'Финансы'
    NProgress.done()

    this.getPayments()
    if(this.user.type === 2) this.getTasks()
  },
  methods: {
    getPayments() {
      axios.get(import.meta.env.VITE_API_URL + 'payments/', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.payments = res.data.data
        }
      })
    },
    getTasks() {
      axios.get(import.meta.env.VITE_API_URL + 'tasks/user/' + this.user.id + '?page=1&per_page=1000', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.allTasks = res.data.data
        }
      })
    }
  }
}
</script>