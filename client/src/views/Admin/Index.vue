<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Панель управления</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <AdminMenu />
        <div class="overflow-x-auto relative mx-3 md:mx-0">
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <tbody>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Количество пользователей
                </th>
                <td class="py-4 px-6">
                  {{ statistics.users_count }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Количество исполнителей
                </th>
                <td class="py-4 px-6">
                  {{ statistics.performers_count }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Количество заказчиков
                </th>
                <td class="py-4 px-6">
                  {{ statistics.customers_count }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Средств на балансе пользователей
                </th>
                <td class="py-4 px-6">
                  {{ $filters.currencyFormat(statistics.users_balance) }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">

                </th>
                <td class="py-4 px-6">

                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Активных задач на бирже
                </th>
                <td class="py-4 px-6">
                  {{ statistics.active_tasks_count }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Активных заказов
                </th>
                <td class="py-4 px-6">
                  {{ statistics.active_orders_count }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Выполненных заказов
                </th>
                <td class="py-4 px-6">
                  {{ statistics.completed_orders_count }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">

                </th>
                <td class="py-4 px-6">

                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Выполнено заказов на сумму в общем
                </th>
                <td class="py-4 px-6">
                  {{ $filters.currencyFormat(statistics.completed_orders_price) }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Выполнено заказов на сумму за сегодня
                </th>
                <td class="py-4 px-6">
                  {{ $filters.currencyFormat(statistics.completed_orders_today_price) }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                  Выполнено заказов на сумму за этот месяц
                </th>
                <td class="py-4 px-6">
                  {{ $filters.currencyFormat(statistics.completed_orders_current_month_price) }}
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
import NProgress from "nprogress";
import axios from "axios"
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";

export default {
  data() {
    return {
      statistics: []
    }
  },
  computed: {
    ...mapState(useUserStore, ['token'])
  },
  mounted() {
    document.title = 'Панель управления'

    this.getStatistics()
  },
  methods: {
    getStatistics() {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'admin/statistics', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.statistics = res.data.data

        }

        NProgress.done();
      })
    }
  }
}
</script>