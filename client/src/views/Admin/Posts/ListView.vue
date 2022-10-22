<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Управление публикациями</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <AdminMenu />
        <div class="flex">
          <RouterLink :to="{ name: 'ap-posts-create' }" class="bg-green-200 mb-3 text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-green-300 transition duration-300 grid place-content-center">
            Добавить
          </RouterLink>
        </div>
        <div class="flex">
          <RouterLink :to="{ name: 'ap-posts-categories' }" class="bg-blue-200 mb-3 text-blue-700 w-full font-medium py-2 px-2 rounded-lg shadow hover:bg-blue-300 transition duration-300 grid place-content-center">
            Управление категориями
          </RouterLink>
        </div>
        <div class="overflow-x-auto relative shadow-md sm:rounded-lg">
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" class="py-3 px-6">
                ID
              </th>
              <th scope="col" class="py-3 px-6">
                Автор
              </th>
              <th scope="col" class="py-3 px-6 text-center">
                Изображение
              </th>
              <th scope="col" class="py-3 px-6">
                Заголовок
              </th>
              <th scope="col" class="py-3 px-6">
                Категории
              </th>
              <th scope="col" class="py-3 px-6">
                Действия
              </th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="post in posts" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
              <td class="py-4 px-6">
                {{ post.id }}
              </td>
              <td class="py-4 px-6">
                {{ post.author.first_name }} {{ post.author.last_name }}
              </td>
              <td class="py-4 px-6">
                <a :href="mainUrl + post.main_image" target="_blank">
                  <img :src="mainUrl + post.main_image" class="h-14 w-14 mx-auto rounded shadow-lg">
                </a>
              </td>
              <td v-if="!post.deleted_at" class="py-4 px-6 font-medium text-slate-700">
                {{ post.title }}
              </td>
              <td v-else class="py-4 px-6 font-medium text-slate-400">
                {{ post.title }}<br>(архив)
              </td>
              <td class="py-4 px-6 font-medium text-slate-700">
                <span class="ml-1" v-for="category in post.categories">
                  {{ category.title }}<br>
                </span>
              </td>
              <td class="py-4 px-6">
                <RouterLink :to="{ name: 'ap-posts-create', query: { edit_id: post.id } }" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Редактировать</RouterLink><br>
                <button v-if="!post.deleted_at" @click="deletePost(post.id)" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</button>
                <button v-else @click="deletePost(post.id)" class="font-medium text-yellow-600 dark:text-yellow-500 hover:underline">Восстановить</button>
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
      posts: []
    }
  },
  computed: {
    mainUrl() {
      return import.meta.env.VITE_URL
    },
    ...mapState(useUserStore, ['token'])
  },
  mounted() {
    document.title = 'Управление публикациями'

    this.getPosts()
  },
  methods: {
    getPosts() {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'posts/', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.posts = res.data.data

          NProgress.done()
        }
      })
    },
    deletePost(postId) {
      NProgress.start()

      axios.delete(import.meta.env.VITE_API_URL + 'posts/' + postId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getPosts()
      })
    }
  }
}
</script>