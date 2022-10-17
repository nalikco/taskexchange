<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Управление заказами</h1>
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
                Исполнитель
              </th>
              <th scope="col" class="py-3 px-6">
                Заказчик
              </th>
              <th scope="col" class="py-3 px-6">
                Статус
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
            <tr v-for="order in orders" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
              <td class="py-4 px-6">
                {{ order.id }}
              </td>
              <td class="py-4 px-6 font-medium text-slate-700">
                {{ order.offer.performer.username }}
              </td>
              <td class="py-4 px-6 font-medium text-slate-700">
                {{ order.task.customer.username }}
              </td>
              <td v-if="order.deleted_at" class="py-4 px-6 font-medium text-red-500">
                Удалено
              </td>
              <td v-else-if="order.status === 0" class="py-4 px-6 font-medium text-blue-500">
                Активно
              </td>
              <td v-else-if="order.status === 1" class="py-4 px-6 font-medium text-yellow-500">
                Сдано на проверку
              </td>
              <td v-else-if="order.status === 2" class="py-4 px-6 font-medium text-green-500">
                Выполнено
              </td>
              <td v-else-if="order.status === 3" class="py-4 px-6 font-medium text-red-500">
                Отменено
              </td>
              <td class="py-4 px-6 text-slate-700">
                {{ moment(order.created_at).utc(0).format('DD.MM.YYYY в HH:mm') }}
              </td>
              <td class="py-4 px-6">
<!--                <button v-if="!order.deleted_at" @click="deleteOrder(order.id)" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</button>-->
<!--                <button v-else @click="deleteOrder(order.id)" class="font-medium text-yellow-600 dark:text-yellow-500 hover:underline">Восстановить</button>-->
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
      orders: [],
    }
  },
  computed: {
    ...mapState(useUserStore, ['user', 'token'])
  },
  mounted() {
    document.title = 'Управление заказами'

    this.getOrders()
  },
  methods: {
    getOrders() {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'orders/', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.orders = res.data.data
        }

        NProgress.done();
      })
    },
    deleteOrder(taskId) {

    }
  }
}
</script>