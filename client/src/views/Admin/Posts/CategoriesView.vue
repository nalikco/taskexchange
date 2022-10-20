<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Управление категориями публикаций</h1>
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
                Название
              </th>
              <th scope="col" class="py-3 px-6">
                Действия
              </th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="category in categories" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
              <td class="py-4 px-6">
                {{ category.id }}
              </td>
              <td v-if="!category.deleted_at" class="py-4 px-6 font-medium text-slate-700">
                {{ category.title }}
              </td>
              <td v-else class="py-4 px-6 font-medium text-slate-400">
                {{ category.title }}<br>(архив)
              </td>
              <td class="py-4 px-6">
                <button @click="toggleEditModal(category)" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Редактировать</button><br>
                <button v-if="!category.deleted_at" @click="deleteCategory(category.id)" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</button>
                <button v-else @click="deleteCategory(category.id)" class="font-medium text-yellow-600 dark:text-yellow-500 hover:underline">Восстановить</button>
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
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
              <button v-if="editId !== 0"  @click="saveCategory" type="button" class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm">Сохранить</button>
              <button v-else  @click="storeCategory" type="button" class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 sm:ml-3 sm:w-auto sm:text-sm">Добавить</button>
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
      categories: [],
      openEditModal: false,
      editCategory: null,
      editId: 0,
      editTitle: '',
    }
  },
  computed: {
    ...mapState(useUserStore, ['token'])
  },
  mounted() {
    document.title = 'Управление категориями публикаций'

    this.getCategories()
  },
  methods: {
    onEditFormChange(e, field) {
      this[field] = e.target.value
    },
    toggleEditModal(category = null) {
      this.editId = 0
      this.editTitle = ''
      if(category) {
        this.editId = category.id
        this.editTitle = category.title
      }
      this.editCategory = category
      this.openEditModal = !this.openEditModal
    },
    getCategories() {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'posts/categories/', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.categories = res.data.data

          NProgress.done()
        }
      })
    },
    deleteCategory(categoryId) {
      NProgress.start()

      axios.delete(import.meta.env.VITE_API_URL + 'posts/categories/' + categoryId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getCategories()
      })
    },
    storeCategory() {
      NProgress.start()

      axios.post(import.meta.env.VITE_API_URL + 'posts/categories/', {
        title: this.editTitle
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(response => {
        this.toggleEditModal(false);

        this.getCategories()
      })
    },
    saveCategory() {
      NProgress.start()

      axios.put(import.meta.env.VITE_API_URL + 'posts/categories/' + this.editId, {
        title: this.editTitle
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(response => {
        this.toggleEditModal(false);

        this.getCategories()
      })
    },
  }
}
</script>