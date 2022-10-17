<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
import {moment} from '@/moment'
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Управление пользователями</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <AdminMenu />
        <button @click="toggleEditModal(null)" class="bg-green-200 mb-3 text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-green-300 transition duration-300">
          Добавить
        </button>
        <div class="overflow-x-auto relative shadow-md sm:rounded-lg">
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" class="py-3 px-6">
                ID
              </th>
              <th scope="col" class="py-3 px-6">
                Имя
              </th>
              <th scope="col" class="py-3 px-6">
                Тип
              </th>
              <th scope="col" class="py-3 px-6">
                Онлайн
              </th>
              <th scope="col" class="py-3 px-6">
                Баланс
              </th>
              <th scope="col" class="py-3 px-6">
                Действия
              </th>
            </tr>
            </thead>
            <tbody>
              <tr v-for="user in users" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                <td class="py-4 px-6">
                  {{ user.id }}
                </td>
                <th v-if="!user.deleted_at" scope="row" class="flex items-center py-4 px-6 text-gray-900 whitespace-nowrap dark:text-white">
                  <img class="w-10 h-10 rounded-full" src="@/assets/img/user.png" alt="Jese image">
                  <div class="pl-3">
                    <div class="text-base font-semibold">{{ user.username }}</div>
                    <div class="font-normal text-gray-500">{{ user.email }}</div>
                  </div>
                </th>
                <th v-else scope="row" class="flex items-center py-4 px-6 text-gray-400 whitespace-nowrap dark:text-white">
                  <img class="w-10 h-10 rounded-full" src="@/assets/img/user.png" alt="Jese image">
                  <div class="pl-3">
                    <div class="text-base font-semibold">{{ user.username }}</div>
                    <div class="font-normal text-gray-500">{{ user.email }} (архив)</div>
                  </div>
                </th>
                <td v-if="user.type === 1" class="py-4 px-6">
                  Исполнитель
                </td>
                <td v-if="user.type === 2" class="py-4 px-6">
                  Заказчик
                </td>
                <td v-if="user.type === 3" class="py-4 px-6">
                  Администратор
                </td>
                <td class="py-4 px-6">
                  <div v-if="checkIsOnline(user.last_online)" class="flex items-center">
                    <div class="h-2.5 w-2.5 rounded-full bg-green-400 mr-2"></div> Онлайн
                  </div>
                  <div v-else class="flex items-center">
                    {{ moment(user.last_online).utcOffset(+6, true).fromNow() }}
                  </div>
                </td>
                <td class="py-4 px-6">
                  {{ $filters.currencyFormat(user.balance) }}
                </td>
                <td class="py-4 px-6">
                  <button @click="toggleEditModal(user)" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Редактировать</button><br>
                  <button v-if="!user.deleted_at" @click="deleteUser(user.id)" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</button>
                  <button v-else @click="deleteUser(user.id)" class="font-medium text-yellow-600 dark:text-yellow-500 hover:underline">Восстановить</button>
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
                  <h3 v-if="editUserId !== 0" class="text-lg font-medium leading-6 text-gray-900" id="modal-title">Редактирование пользователя</h3>
                  <h3 v-else class="text-lg font-medium leading-6 text-gray-900" id="modal-title">Добавление пользователя</h3>
                  <div class="mt-2">
                    <div class="mt-6 mb-6">
                      <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Имя</label>
                      <input type="text" @input="onEditFormChange($event, 'editUsername')"  @change="onEditFormChange($event, 'editUsername')" name="username" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" :value="editUsername" placeholder="Имя" required>
                    </div>
                    <div class="mb-6">
                      <label for="countries" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400">Тип</label>
                      <select id="countries" @change="onEditFormChange($event, 'editType')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                        <option value="1" :selected="editType === 1">Исполнитель</option>
                        <option value="2" :selected="editType === 2">Заказчик</option>
                        <option value="3" :selected="editType === 3">Администратор</option>
                      </select>
                    </div>
                    <div class="mb-6">
                      <label for="balance" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Баланс</label>
                      <input @input="onEditFormChange($event, 'editBalance')" @change="onEditFormChange($event, 'editBalance')" type="number" name="balance" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" :value="editBalance" placeholder="Баланс" required>
                    </div>
                    <div class="mb-6">
                      <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">E-Mail</label>
                      <input @input="onEditFormChange($event, 'editEmail')" @change="onEditFormChange($event, 'editEmail')" type="email" name="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" :value="editEmail" placeholder="E-Mail" required>
                    </div>
                    <div class="mb-6">
                      <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Пароль</label>
                      <input @input="onEditFormChange($event, 'editPassword')" @change="onEditFormChange($event, 'editPassword')" type="password" name="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Пароль" required>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
              <button v-if="editUserId !== 0" @click="saveUser" type="button" class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm">Сохранить</button>
              <button v-else @click="storeUser" type="button" class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm">Добавить</button>
              <button @click="toggleEditModal(false)" type="button" class="mt-3 inline-flex w-full justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-base font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">Отменить</button>
            </div>
          </div>
        </div>
      </div>
    </div>
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
      users: [],
      openEditModal: false,
      editUser: null,
      editUserId: 0,
      editUsername: '',
      editType: 0,
      editBalance: 0,
      editEmail: '',
      editPassword: '',
    }
  },
  mounted() {
    document.title = 'Управление пользователями'

    this.getUsers()
  },
  computed: {
    ...mapState(useUserStore, ['user', 'token'])
  },
  methods: {
    onEditFormChange(e, field) {
      this[field] = e.target.value
    },
    toggleEditModal(user = null) {
      this.editUserId = 0
      this.editUsername = ''
      this.editType = 0
      this.editBalance = 0
      this.editEmail = ''
      this.editPassword = ''
      if(user) {
        this.editUserId = user.id
        this.editUsername = user.username
        this.editType = user.type
        this.editBalance = user.balance
        this.editEmail = user.email
      }
      this.editUser = user
      this.openEditModal = !this.openEditModal
    },
    getUsers() {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'users/', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.users = res.data.data

        }

        NProgress.done();
      })
    },
    checkIsOnline(onlineDate) {
      let onlineDateObj = moment(onlineDate).utcOffset(+6, true)
      let currentDateObj = moment()

      if (onlineDateObj.diff(currentDateObj, 'minutes') > -15) {
        return true
      }

      return false
    },
    deleteUser(userId) {
      NProgress.start()

      axios.delete(import.meta.env.VITE_API_URL + 'users/' + userId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getUsers()
      })
    },
    storeUser() {
      NProgress.start()

      axios.post(import.meta.env.VITE_API_URL + 'users/', {
        username: this.editUsername,
        type: parseInt(this.editType),
        balance: parseFloat(this.editBalance),
        email: this.editEmail,
        password: this.editPassword,
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(response => {
        this.toggleEditModal(false);

        this.getUsers()
      })
    },
    saveUser() {
      NProgress.start()

      let data = {}
      if (this.editUser.username !== this.editUsername) data.username = this.editUsername
      if (this.editUser.type !== parseInt(this.editType)) data.type = parseInt(this.editType)
      if (this.editUser.balance !== parseFloat(this.editBalance)) data.balance = parseFloat(this.editBalance)
      if (this.editUser.email !== this.editEmail) data.email = this.editEmail
      if (this.editPassword !== '') data.password = this.editPassword

      axios.put(import.meta.env.VITE_API_URL + 'users/' + this.editUserId, data, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(response => {
        this.toggleEditModal(false);

        this.getUsers()
      })
    },
  }
}
</script>