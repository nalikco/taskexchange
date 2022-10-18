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
                Задача
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
              <td class="py-4 px-6 font-medium text-slate-700">
                #{{ order.task.id }}<br>
                ({{ getTaskOptionStructured(order.task).main.title }})
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
                <button v-if="order.status < 2" @click="deleteOrder(order.id)" class="font-medium text-red-600 dark:text-red-500 hover:underline">Отменить</button>
                <button v-else-if="order.status === 3" @click="deleteOrder(order.id)" class="font-medium text-yellow-600 dark:text-yellow-500 hover:underline">Восстановить</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>
    <div v-if="openEditModal" class="relative z-10" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>

      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
          <div class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
            <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <div class="sm:flex sm:items-start">
                <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                  <h3 v-if="editId !== 0" class="text-lg font-medium leading-6 text-gray-900" id="modal-title">Редактирование опции</h3>
                  <h3 v-else class="text-lg font-medium leading-6 text-gray-900" id="modal-title">Добавление опции</h3>
                  <div class="mt-2">
                    <div class="mt-6 mb-6">
                      <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Наименование</label>
                      <input type="text" @input="onEditFormChange($event, 'editTitle')" @change="onEditFormChange($event, 'editTitle')" name="username" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" :value="editTitle" placeholder="Наименование" required>
                    </div>
                    <div class="mb-6">
                      <label for="countries" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400">Родитель (категория)</label>
                      <select id="countries" @change="onEditFormChange($event, 'editParentId')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                        <option value="0" :selected="editParentId === 0">нет</option>
                        <option v-for="category in showCategories" :value="category.id" :selected="editParentId === category.id">{{ category.title }}</option>
                      </select>
                    </div>
                    <div class="mb-6">
                      <label for="balance" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Цена</label>
                      <input @input="onEditFormChange($event, 'editPrice')" @change="onEditFormChange($event, 'editPrice')" type="number" name="balance" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" :value="editPrice" placeholder="Цена" required>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
              <button v-if="editId !== 0"  @click="saveOption" type="button" class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm">Сохранить</button>
              <button v-else  @click="storeOption" type="button" class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm">Добавить</button>
              <button @click="toggleEditModal(false)" type="button" class="mt-3 inline-flex w-full justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-base font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">Отменить</button>
            </div>
          </div>
        </div>
      </div>
    </div>
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
      editOrder: null,
      editId: 0,
      editTaskId: 0,
      editPerformerId: 0,
      editStatus: 0,
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
    onEditFormChange(e, field) {
      this[field] = e.target.value
    },
    toggleEditModal(order = null) {
      this.editId = 0
      this.editTaskId = ''
      this.editPerformerId = 0
      this.editStatus = 0
      if(order) {
        this.editId = order.id
        this.editParentId = order.parent_id
        this.editTitle = order.title
        this.editPrice = order.price
      }
      this.editOrder = order
      this.openEditModal = !this.openEditModal
    },
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
      NProgress.start()

      axios.delete(import.meta.env.VITE_API_URL + 'orders/' + taskId, {}, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getOrders()
      })
    }
  }
}
</script>